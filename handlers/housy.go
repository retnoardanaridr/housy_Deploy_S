package handlers

import (
	housydto "be/dto/housy"
	dto "be/dto/result"
	"be/models"
	"be/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerHousy struct {
	HousyRepository repositories.HousyRepository
}

func HandlerHousy(housyRepository repositories.HousyRepository) *handlerHousy {
	return &handlerHousy{housyRepository}
}

func (h *handlerHousy) FindHousy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	house, err := h.HousyRepository.FindHousy()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHousy) GetHouseID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	house, err := h.HousyRepository.GetHouseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHousy) AddHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "housy"})

	if err != nil {
		fmt.Println(err.Error())
	}
	//get token user
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// var amenitiesId []int
	// for _, r := range r.FormValue("amenities_id") {
	// 	if int(r-'0') >= 0 {
	// 		amenitiesId = append(amenitiesId, int(r-'0'))
	// 	}
	// }

	city_id, _ := strconv.Atoi(r.FormValue("city_id"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	beds, _ := strconv.Atoi(r.FormValue("bedroom"))
	bats, _ := strconv.Atoi(r.FormValue("bathroom"))
	request := housydto.HouseRequest{
		Name:     r.FormValue("name"),
		CityID:   city_id,
		Address:  r.FormValue("address"),
		Price:    price,
		TypeRent: r.FormValue("typerent"),
		// AmenitiesID: amenitiesId,
		Bedroom:     beds,
		Bathroom:    bats,
		Description: r.FormValue("description"),
		Thumbnail:   resp.SecureURL,
	}

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	//get all city data by id
	city, _ := h.HousyRepository.FindCityById(city_id)
	// amenities, _ := h.HousyRepository.FindAmenitiesByID(amenitiesId)

	housy := models.Housy{
		Name:     request.Name,
		City:     city,
		Address:  request.Address,
		Price:    request.Price,
		TypeRent: request.TypeRent,
		// Amenities:   amenities,
		Bedroom:     request.Bedroom,
		Bathroom:    request.Bathroom,
		Description: request.Description,
		UserID:      userId,
		Thumbnail:   resp.SecureURL,
	}

	newHousy, err := h.HousyRepository.AddHouse(housy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataHousy, err := h.HousyRepository.GetHouseUser(newHousy.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: dataHousy}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHousy) DeleteHousy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	housy, err := h.HousyRepository.GetHouseID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadGateway, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	delete, err := h.HousyRepository.DeleteHouse(housy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: delete}
	json.NewEncoder(w).Encode(response)
}
