package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection to Database
func DatabaseInit() {

	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PORT = os.Getenv("DB_PORT")

	var err error
	// dsn := "root:@tcp(127.0.0.1:3306)/db_housy_mux?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//deploy
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_HOST, DB_PORT, DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local, DB_USER, DB_HOST, DB_PORT, DB_NAME"

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
