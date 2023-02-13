package transactiondto

import (
	"be/models"
)

type TransactionResponse struct {
	ID           int                        `json:"id" gorm:"primary_key:auto_increment"`
	TransationID int                        `json:"transaction_id"`
	CheckIn      string                     `json:"checkin"`
	CheckOut     string                     `json:"checkout"`
	Housy        models.HousyResponse       `json:"housy"`
	UserID       int                        `json:"user_id"`
	User         models.UserProfileResponse `json:"user"`
	Total        int                        `json:"total"`
	Status       string                     `json:"status"`
}
