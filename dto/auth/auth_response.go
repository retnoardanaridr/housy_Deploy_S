package authdto

type LoginResponse struct {
	Username string `json:"username" gorm:"type: varchar(255)"`
	ListAs   string `json:"listas" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
}

type CheckAuthResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	ListAs   string `json:"listas" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
}
