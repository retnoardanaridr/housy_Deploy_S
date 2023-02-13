package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection to Database
func DatabaseInit() {

	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PORT = os.Getenv("DB_PORT")

	var err error
	// dsn := "root:@tcp(127.0.0.1:3306)/db_housy_mux?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//deploy
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
