package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/faizallmaullana/be_kredit_konvensional/controller/authentication"
	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

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
	r.Use(cors.New(corsConfig))
	fmt.Printf("Port: %s \n", port)

	r.POST("/api/v1/kang_kredit/registration", authentication.Registration)

	r.Run(fmt.Sprintf(":%s", port))
}
