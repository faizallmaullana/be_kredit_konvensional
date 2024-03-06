package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectToDatabase() {
	connStr := "user='tanjung' password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
	database, err := gorm.Open("postgres", connStr)

	// database, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func Migration(c *gin.Context) {
	DB.AutoMigrate(
		Users{},
		Profiles{},
		Customers{},
		Credits{},
		Payments{},
		Wallets{},
		OtherTransaction{},
		Tokens{},
	)

	c.JSON(http.StatusCreated, gin.H{"message": "CreatedSuccessfull"})
}
