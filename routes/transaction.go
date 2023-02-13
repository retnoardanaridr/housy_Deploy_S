package routes

import (
	"be/handlers"
	"be/pkg/middleware"
	"be/pkg/mysql"
	"be/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	r.HandleFunc("/transaction", middleware.Auth(h.AddTransaction)).Methods("POST")
	r.HandleFunc("/orders", middleware.Auth(h.FindOrders)).Methods("GET")
	r.HandleFunc("/order", middleware.Auth(h.GetOrder)).Methods("GET")
	r.HandleFunc("/check-out", middleware.Auth(h.PayTransaction)).Methods("PATCH")
	r.HandleFunc("/history", middleware.Auth(h.GetTransactionUser)).Methods("GET")
	r.HandleFunc("/order/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
