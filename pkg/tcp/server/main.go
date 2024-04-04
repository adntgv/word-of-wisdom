package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"wordOfWisdom/config"
	"wordOfWisdom/internal/datasources/repositories/inmem"
	"wordOfWisdom/internal/tcp/handlers"
	"wordOfWisdom/internal/tcp/middlewares"
	"wordOfWisdom/pkg/challanger"
)

type Server struct {
	wg         sync.WaitGroup
	ln         net.Listener
	shutdown   chan struct{}
	connection chan net.Conn
	authorized chan net.Conn
	handle     func(net.Conn)
	authorize  func(net.Conn) net.Conn
}

func NewServer(cfg *config.Config) (*Server, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf("%v:%v", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalln(err)
	}

	repo := inmem.NewQuoteRepository()
	chal := challanger.NewChallanger(cfg.Difficulty)

	handler := handlers.NewConnectionHandler(repo).Handle
	authorize := middlewares.NewAuthMiddleware(chal).Authorize

	return &Server{
		wg:         sync.WaitGroup{},
		ln:         ln,
		shutdown:   make(chan struct{}),
		connection: make(chan net.Conn),
		authorized: make(chan net.Conn),
		handle:     handler,
		authorize:  authorize,
	}, nil
}

func (server *Server) Run() error {
	var wg = sync.WaitGroup{}

	server.wg.Add(2)
	go server.AcceptConnections()
	go server.HandleConnections()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down the api server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		return fmt.Errorf("failed to shut down api server: %s", err.Error())
	}

	wg.Wait()
	log.Println("api server successfully shutdown")

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	close(s.shutdown)
	s.ln.Close()

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-time.After(time.Second):
		return fmt.Errorf("timedout waitgroup waiting")
	}
}

func (s *Server) AcceptConnections() {
	defer s.wg.Done()
	for {
		select {
		case <-s.shutdown:
			return
		default:
			conn, err := s.ln.Accept()
			if err != nil {
				log.Println(err)
			} else {
				s.connection <- conn
			}
		}
	}
}

func (s *Server) AuthorizeConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		case conn := <-s.connection:
			authorized := s.authorize(conn)
			if authorized != nil {
				s.authorized <- authorized
			} else {
				conn.Close()
			}
		}
	}
}
func (s *Server) HandleConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		case conn := <-s.connection:
			go s.handle(conn)
		}
	}
}
