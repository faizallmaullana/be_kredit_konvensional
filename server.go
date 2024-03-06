package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/faizallmaullana/be_kredit_konvensional/controller/admin"
	"github.com/faizallmaullana/be_kredit_konvensional/controller/authentication"
	"github.com/faizallmaullana/be_kredit_konvensional/controller/database"
	"github.com/faizallmaullana/be_kredit_konvensional/cors"
	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var corsConfig = cors.CORSMiddleware()

func main() {
	port := os.Getenv("PORT")

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Use cors middleware
	models.ConnectToDatabase()
	r.Use(corsConfig)
	fmt.Printf("Port: %s \n", port)

	// Authentication
	r.POST("/api/v1/kang_kredit/registration", authentication.Registration)
	r.POST("/api/v1/kang_kredit/login", authentication.LoginResource)

	// admin
	r.GET("/api/v1/kang_kredit/token", admin.GetToken)

	// database
	r.GET("/api/v1/db/tablename", database.GetTableName)
	r.GET("/api/v1/db/close", database.CloseDatabase)
	r.GET("/api/v1/db/open", database.OpenTheDatabase)
	r.DELETE("/api/v1/db/delete", database.DeleteAllTable)
	r.GET("/api/v1/db/export", models.ExportData)
	r.GET("/api/v1/db/migration", models.Migration)
	r.GET("/api/v1/db/transfer", models.TransferData)
	r.GET("/api/v1/db/download", database.DownloadDatabase)

	r.Run(fmt.Sprintf(":%s", port))
}
