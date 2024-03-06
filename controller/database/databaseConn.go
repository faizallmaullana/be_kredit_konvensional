package database

import (
	"net/http"

	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func CloseDatabase(c *gin.Context) {
	models.DB.Close()

	c.JSON(http.StatusOK, gin.H{"message": "database closed successfully"})
}

func OpenTheDatabase(c *gin.Context) {
	models.ConnectToDatabase()

	c.JSON(http.StatusOK, gin.H{"message": "database open successfully"})
}
