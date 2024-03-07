package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var SourceDB *gorm.DB
var DestinationDB *gorm.DB

func ConnectToSourceDatabase() {
	database, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("Failed to connect to source database!")
	}
	SourceDB = database
}

func ConnectToDestinationDatabase() {
	connStr := "user='tanjung' password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
	database, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to destination database!")
	}
	DestinationDB = database
}

func TransferData(c *gin.Context) {
	// Connect to source and destination databases
	ConnectToSourceDatabase()
	ConnectToDestinationDatabase()

	// Migrate schemas to destination database if needed
	DestinationDB.AutoMigrate(&Users{}, &Profiles{})

	// Transfer data from the User table
	var users []Users
	SourceDB.Find(&users)
	for _, user := range users {
		DestinationDB.Create(&user)
	}

	// Transfer data from the Product table
	var products []Profiles
	SourceDB.Find(&products)
	for _, product := range products {
		DestinationDB.Create(&product)
	}
}
