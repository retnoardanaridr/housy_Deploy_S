package routes

import (
	"be/handlers"
	"be/pkg/middleware"
	"be/pkg/mysql"
	"be/repositories"

	"github.com/gorilla/mux"
)

func HousyRoutes(r *mux.Router) {
	HousyRepository := repositories.RepositoryHousy(mysql.DB)
	h := handlers.HandlerHousy(HousyRepository)

	r.HandleFunc("/housy", h.FindHousy).Methods("GET")
	r.HandleFunc("/house/{id}", h.GetHouseID).Methods("GET")
	r.HandleFunc("/house", middleware.Auth(middleware.UploadThumbnail(h.AddHouse))).Methods("POST")
	// r.HandleFunc("/house/{id}", middleware.Auth(h.UpdateHousy)).Methods("PATCH")
	r.HandleFunc("/house/{id}", middleware.Auth(h.DeleteHousy)).Methods("DELETE")
	// r.HandleFunc("/filter/houses?{sortBy}", (h.FilterSinglePath)).Methods("GET")
}
