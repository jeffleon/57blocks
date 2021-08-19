package database

import (
	"github.com/jeffleon/57block-movies/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "root:secret@tcp(db:3306)/movies_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	DB = db
	db.AutoMigrate(&models.Movies{})
	seed(db)
}

func seed(db *gorm.DB) {

	movies := []models.Movies{
		{Genre: "Sci-Fi", Name: "Suicide Squad", Description: "...", Director: "James Gun", Release: "2021", UserID: 1},
		{Genre: "Adventure", Name: "Jungle Cruise", Description: "...", Director: "Jaume Collet-Serra", Release: "2021", UserID: 1},
		{Genre: "Action", Name: "Free Guy", Description: "...", Director: "Matt Lieberman", Release: "2021", UserID: 1},
		{Genre: "Action", Name: "Black Widow", Description: "...", Director: "Jac Schaeffer", Release: "2021", UserID: 2},
		{Genre: "Adventure", Name: "Space Jam", Description: "...", Director: "Malcolm D. Lee", Release: "2021", UserID: 2},
		{Genre: "Action", Name: "Fast & Furious 9", Description: "...", Director: "Justin Lin", Release: "2021", UserID: 2},
		{Genre: "Action", Name: "La guerra del mañana", Description: "...", Director: "Chris McKay", Release: "2021", UserID: 3},
		{Genre: "Adventure", Name: "Luca", Description: "...", Director: "Enrico Casarosa", Release: "2021", UserID: 3},
		{Genre: "Action", Name: "Despierta la furia", Description: "...", Director: "Guy Ritchie", Release: "2021", UserID: 3},
	}
	for _, d := range movies {
		db.Model(&models.Movies{}).Create(&d)
	}

}
