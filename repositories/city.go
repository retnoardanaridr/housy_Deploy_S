package repositories

import (
	"be/models"

	"gorm.io/gorm"
)

type CityRepository interface {
	FindCities() ([]models.City, error)
	GetCityById(ID int) (models.City, error)
	CreateCity(city models.City) (models.City, error)
}

func ReposityCity(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCities() ([]models.City, error) {
	var city []models.City
	err := r.db.Find(&city).Error

	return city, err
}

func (r *repository) GetCityById(ID int) (models.City, error) {
	var cityId models.City

	err := r.db.First(&cityId, ID).Error
	return cityId, err
}

func (r *repository) CreateCity(city models.City) (models.City, error) {
	err := r.db.Create(&city).Error

	return city, err
}
