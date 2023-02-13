package models

import "time"

type Amenities struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	CreatedAd time.Time `json:"-"`
	UpdatedAd time.Time `json:"-"`
}

type AmenitiesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (Amenities) TableName() string {
	return "amenities"
}

func (AmenitiesResponse) TableName() string {
	return "amenities"
}
