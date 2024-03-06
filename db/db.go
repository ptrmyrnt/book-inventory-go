package db

import (
	"book-inventory-go/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}
	if db.Find(&data).RecordNotFound() {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{{
		Title:       "Travelling Around the World",
		Author:      "Ria",
		Description: "Buku tentang travelling",
		Stock:       10,
	}, {
		Title:       "Healing Our Journey",
		Author:      "Putri",
		Description: "Buku tentang penyembuhan",
		Stock:       9,
	}, {
		Title:       "Animalia Kingdom",
		Author:      "King Arthur",
		Description: "Buku seputar kingdom animalia",
		Stock:       8,
	}, {
		Title:       "News Update",
		Author:      "Silalahi",
		Description: "Buku seputar berita terkini",
		Stock:       1,
	}}

	for _, v := range data {
		db.Create(&v)
	}
}
