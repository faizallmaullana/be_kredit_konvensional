package database

import (
	"fmt"
	"net/http"

	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
)

func DeleteAllTable(c *gin.Context) {
	// Get all table names
	var tableNames []string
	models.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Pluck("table_name", &tableNames)

	// Drop all tables
	for _, tableName := range tableNames {
		models.DB.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName))
	}

	c.JSON(http.StatusOK, gin.H{"message": "all tables deleted successfuly"})
}
