package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	DB.AutoMigrate(&Image{})
	//DB.DropTableIfExists(&Vehicle{}, &Customer{}, &Tire{})
	//DB.AutoMigrate(&Company{}).AddForeignKey("id", "customers(id)", "CASCADE", "CASCADE")
	//DB.AutoMigrate(&Company{}).AddForeignKey("veh_id", "vehicles(v_id)", "CASCADE", "CASCADE")
	// DB.AutoMigrate(&Company{})
	// DB.AutoMigrate(&Customer{})
	// DB.AutoMigrate(&Vehicle{})
	// DB.AutoMigrate(&Tire{}).AddForeignKey("tire_id", "vehicles(v_id)", "CASCADE", "CASCADE")
	// DB.Model(&Vehicle{}).AddForeignKey("veh_id", "customers(id)", "CASCADE", "CASCADE")

	// DB, err = gorm.Open(Dbdriver, DBURL)

	// dsn := "username:password@tcp(127.0.0.1:3306)/test"

	// database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("Failed to connect to database!")
	// }

	// err = database.AutoMigrate(&Book{})
	// if err != nil {
	// 	return
	// }

	// DB = database
}
