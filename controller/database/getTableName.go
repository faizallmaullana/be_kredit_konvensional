package database

import (
	"net/http"

	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
)

func GetTableName(c *gin.Context) {
	// Get all table names
	var tableNames []string
	models.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Pluck("table_name", &tableNames)

	// Print table names
	for _, tableName := range tableNames {
		c.JSON(http.StatusOK, tableName)
	}
}
