package repositories

import (
	"be/models"

	"gorm.io/gorm"
)

type HousyRepository interface {
	FindHousy() ([]models.Housy, error)
	FindCityById(CityID int) (models.City, error)
	FindAmenitiesByID(AmenitiesID []int) ([]models.Amenities, error)
	GetHouseUser(ID int) (models.Housy, error)
	GetHouseID(ID int) (models.Housy, error)
	AddHouse(house models.Housy) (models.Housy, error)
	UpdateHouse(house models.Housy) (models.Housy, error)
	DeleteHouse(house models.Housy) (models.Housy, error)

	//filter
	GetHouseSinglePath(key string) ([]models.Housy, error)
}

func RepositoryHousy(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHousy() ([]models.Housy, error) {
	var house []models.Housy
	err := r.db.Preload("City").Find(&house).Error

	return house, err
}

func (r *repository) FindCityById(CityID int) (models.City, error) {
	var city models.City
	err := r.db.Find(&city, CityID).Error

	return city, err
}

func (r *repository) FindAmenitiesByID(AmenitiesID []int) ([]models.Amenities, error) {
	var amenities []models.Amenities
	err := r.db.Find(&amenities, AmenitiesID).Error

	return amenities, err
}

func (r *repository) GetHouseID(ID int) (models.Housy, error) {
	var houseId models.Housy

	err := r.db.Preload("User").Preload("City").First(&houseId, ID).Error
	return houseId, err
}

func (r *repository) GetHouseUser(ID int) (models.Housy, error) {
	var housy models.Housy
	err := r.db.Preload("User").First(&housy, "id = ?", ID).Error

	return housy, err
}

func (r *repository) AddHouse(house models.Housy) (models.Housy, error) {
	err := r.db.Preload("User").Preload("City").Create(&house).Error

	return house, err
}

func (r *repository) UpdateHouse(house models.Housy) (models.Housy, error) {
	err := r.db.Save(&house).Error

	return house, err
}

func (r *repository) DeleteHouse(house models.Housy) (models.Housy, error) {
	err := r.db.Delete(&house).Error

	return house, err
}

// filter
func (r *repository) GetHouseSinglePath(key string) ([]models.Housy, error) {
	var housy []models.Housy
	err := r.db.Preload("City").Find(&housy).Error

	return housy, err
}
