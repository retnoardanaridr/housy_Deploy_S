package handlers

import (
	dto "be/dto/result"
	"be/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handlerAmenities struct {
	AmenitiesRepository repositories.AmenitiesRepository
}

func HandlerAmenities(AmenitiesRepository repositories.AmenitiesRepository) *handlerAmenities {
	return &handlerAmenities{AmenitiesRepository}
}

func (h *handlerAmenities) FindAmenities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := h.AmenitiesRepository.FindAmenities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: categories}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAmenities) GetAmenities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	category, err := h.AmenitiesRepository.GetAmenities(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: category}
	json.NewEncoder(w).Encode(response)
}
