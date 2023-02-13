package models

type Housy struct {
	ID        int    `json:"id"`
	Name      string `json:"name" gorm:"type: varchar(255)"`
	City      City   `json:"city"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Image     string `json:"image" form:"image" gorm:"type: varchar(255)"`
	CityID    int    `json:"city_id" form:"city_id"`
	Address   string `json:"address" gorm:"type: text"`
	Price     int    `json:"price" gorm:"type: varchar(255)"`
	TypeRent  string `json:"typerent" gorm:"type: varchar(255)"`
	// AmenitiesID []int               `json:"amenities_id" form:"amenities_id" gorm:"-"`
	// Amenities   []Amenities         `json:"amenities" gorm:"many2many:housy_amenities;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bedroom     int                 `json:"bedroom"`
	Bathroom    int                 `json:"bathroom"`
	Description string              `json:"description" form:"description" gorm:"type: varchar(255)"`
	UserID      int                 `json:"user_id"`
	User        UserProfileResponse `json:"user"`
}

type HousyResponse struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	City      CityResponse `json:"city"`
	CityID    int          `json:"city_id" form:"city_id"`
	Thumbnail string       `json:"thumbnail"`
	Image     string       `json:"image"`
	Address   string       `json:"address"`
	Price     int          `json:"price"`
	TypeRent  string       `json:"typerent"`
	// AmenitiesID int                 `json:"amenities_id" gorm:"many2many:housy_amenities"`
	// Amenities   []Amenities         `json:"amenities" form:"amenities_id" gorm:"-"`
	Bedroom     int                 `json:"bedroom"`
	Bathroom    int                 `json:"bathroom"`
	Description string              `json:"description"`
	UserID      int                 `json:"user_id"`
	User        UserProfileResponse `json:"user"`
}

type HouseUserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	CityID  int    `json:"city_id"`
	Address string `json:"address"`
	Price   int    `json:"price"`
	// UserID  int    `json:"-"`
}

func (Housy) TableName() string {
	return "housy"
}

func (HousyResponse) TableName() string {
	return "housy"
}

func (HouseUserResponse) TableName() string {
	return "housy"
}
