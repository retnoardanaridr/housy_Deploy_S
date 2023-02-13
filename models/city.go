package models

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CityResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (City) TableName() string {
	return "city"
}

func (CityResponse) TableName() string {
	return "city"
}
