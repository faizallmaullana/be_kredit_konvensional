package models

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ExportData(c *gin.Context) {
	// Fetch data from each table
	var data []interface{}

	// Fetch data from Table1
	var tableData1 []Users
	if err := DB.Find(&tableData1).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from Table1")
		return
	}
	data = append(data, tableData1)

	// Fetch data from Table2
	var tableData2 []Profiles
	if err := DB.Find(&tableData2).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from Table2")
		return
	}
	data = append(data, tableData2)

	var tableData3 []Wallets
	if err := DB.Find(&tableData3).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from Table2")
		return
	}
	data = append(data, tableData3)

	var tableData4 []OtherTransaction
	if err := DB.Find(&tableData4).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from Table2")
		return
	}
	data = append(data, tableData4)

	var tableData5 []Customers
	if err := DB.Find(&tableData5).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from Table2")
		return
	}
	data = append(data, tableData5)

	var tableData6 []Credits
	if err := DB.Find(&tableData6).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from Table2")
		return
	}
	data = append(data, tableData6)

	var tableData7 []Payments
	if err := DB.Find(&tableData7).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch data from Table2")
		return
	}
	data = append(data, tableData7)

	// Serialize data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to serialize data to JSON")
		return
	}

	// Create a temporary file to store the backup data
	file, err := os.CreateTemp("", "backup_data_*.json")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create temporary file")
		return
	}
	defer os.Remove(file.Name()) // Remove the temporary file after serving

	// Write backup data to the temporary file
	_, err = file.Write(jsonData)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to write backup data to file")
		return
	}

	// Set response headers to make the file downloadable
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=backup_data.json")
	c.Header("Content-Type", "application/json")
	c.File(file.Name())
}
