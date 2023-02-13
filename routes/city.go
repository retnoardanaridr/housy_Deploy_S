package routes

import (
	"be/handlers"
	"be/pkg/mysql"
	"be/repositories"

	"github.com/gorilla/mux"
)

func CityRoutes(r *mux.Router) {
	CityRepository := repositories.ReposityCity(mysql.DB)
	h := handlers.HandlerCity(CityRepository)

	r.HandleFunc("/cities", h.FindCities).Methods("GET")
	r.HandleFunc("/city/{id}", h.GetCityById).Methods("GET")
	r.HandleFunc("/city", h.CreateCity).Methods("POST")
}
