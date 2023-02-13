package routes

import (
	"be/handlers"
	"be/pkg/mysql"
	"be/repositories"

	"github.com/gorilla/mux"
)

func AmenitiesRoutes(r *mux.Router) {
	amenitiesRepository := repositories.RepositoryAmenities(mysql.DB)
	h := handlers.HandlerAmenities(amenitiesRepository)

	r.HandleFunc("/amenities", h.FindAmenities).Methods("GET")
	r.HandleFunc("/amenity", h.GetAmenities).Methods("GET")
}
