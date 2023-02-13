package usersdto

type UserRespone struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email"`
	Password string `json:"password" form:"password"`
	ListAs   string `json:"listas"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Image    string `json:"image"`
}
