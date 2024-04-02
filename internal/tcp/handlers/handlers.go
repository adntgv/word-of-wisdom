package handlers

import (
	"applicationDesignTest/internal/business/domains"
	"applicationDesignTest/internal/http/requests"
	"encoding/json"
	"log"
	"net/http"
)

type OrderHandler struct {
	orderUsecase domains.OrderUsecase
	roomUsecase  domains.RoomUsecase
}

func NewOrderHandler(orderUsecase domains.OrderUsecase, roomUsecase domains.RoomUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
		roomUsecase:  roomUsecase,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrderRequest requests.CreateOrderRequest
	json.NewDecoder(r.Body).Decode(&newOrderRequest)

	if err := newOrderRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room := newOrderRequest.ToRoom()
	from := newOrderRequest.GetFrom()
	to := newOrderRequest.GetTo()

	if err := h.roomUsecase.Book(r.Context(), room, from, to); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Failed to book room \"%v\": %v", room, err)
		return
	}

	newOrder := newOrderRequest.ToOrder()

	if err := h.orderUsecase.Store(r.Context(), newOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to create order:\n%v\n%v", newOrder, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	log.Printf("Order successfully created: %v", newOrder)
}
