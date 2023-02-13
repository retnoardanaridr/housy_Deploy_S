package models

import (
	"time"
)

type Transaction struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	TransactionID int                 `json:"transaction_id"`
	CheckIn       string              `json:"checkin" gorm:"type: varchar(255)"`
	CheckOut      string              `json:"checkout" gorm:"type: varchar(255)"`
	HousyID       int                 `json:"housy_id"`
	Housy         HousyResponse       `json:"housy" form:"housy_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID        int                 `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User          UserProfileResponse `json:"user"`
	Total         int                 `json:"total" gorm:"type: varchar(255)"`
	Status        string              `json:"status" gorm:"type: varchar(255)"`
	CreatedAt     time.Time           `json:"-"`
	UpdatedAt     time.Time           `json:"-"`
}
