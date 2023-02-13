package repositories

import (
	"be/models"

	"gorm.io/gorm"
)

type AmenitiesRepository interface {
	FindAmenities() ([]models.Amenities, error)
	GetAmenities(ID int) (models.Amenities, error)
}

func RepositoryAmenities(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAmenities() ([]models.Amenities, error) {
	var amenities []models.Amenities
	err := r.db.Find(&amenities).Error

	return amenities, err
}

func (r *repository) GetAmenities(ID int) (models.Amenities, error) {
	var amenities models.Amenities
	err := r.db.First(&amenities).Error

	return amenities, err
}
