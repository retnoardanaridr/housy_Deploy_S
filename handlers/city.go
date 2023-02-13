package handlers

import (
	citydto "be/dto/city"
	dto "be/dto/result"
	"be/models"
	"be/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handlerCity struct {
	CityRepository repositories.CityRepository
}

func HandlerCity(cityRepositoy repositories.CityRepository) *handlerCity {
	return &handlerCity{cityRepositoy}
}

func (h *handlerCity) FindCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	city, err := h.CityRepository.FindCities()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: city}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCity) GetCityById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	city, err := h.CityRepository.GetCityById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: city}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCity) CreateCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(citydto.CityRequest)

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusNotFound, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Field := models.City{
		Name: request.Name,
	}

	city, err := h.CityRepository.CreateCity(Field)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: city}
	json.NewEncoder(w).Encode(response)
}
