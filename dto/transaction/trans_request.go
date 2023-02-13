package transactiondto

type TransactionRequest struct {
	CheckIn  string `json:"checkin" validate:"required"`
	CheckOut string `json:"checkout" validate:"required"`
	HousyID  int    `json:"housy_id" validate:"required"`
	Total    int    `json:"total" validate:"required"`
	Status   string `json:"status" validate:"required"`
}
