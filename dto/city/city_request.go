package citydto

type CityRequest struct {
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
