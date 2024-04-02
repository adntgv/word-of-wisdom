package routes

import (
	"applicationDesignTest/internal/business/usecases"
	"applicationDesignTest/internal/datasources/repositories/inmem"
	"applicationDesignTest/internal/http/handlers"

	"github.com/go-chi/chi/v5"
)

type ordersRoute struct {
	router  *chi.Mux
	handler *handlers.OrderHandler
}

func NewOrdersRoute(router *chi.Mux) *ordersRoute {
	ordersRepository := inmem.NewOrderRepository()
	roomsRepository := inmem.NewRoomRepository()
	ordersUsecase := usecases.NewOrderUsecase(ordersRepository)
	roomUsecase := usecases.NewRoomUsecase(roomsRepository)
	handler := handlers.NewOrderHandler(ordersUsecase, roomUsecase)
	return &ordersRoute{
		router:  router,
		handler: handler,
	}
}

func (r *ordersRoute) Register() error {
	r.router.Post("/orders", r.handler.CreateOrder)

	return nil
}
