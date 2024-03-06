package admin

import (
	"fmt"
	"net/http"

	jwt_auth "github.com/faizallmaullana/be_kredit_konvensional/Authentication"
	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	fmt.Println("test")
	// cek authorization
	tokenString := c.GetHeader("Authorization")
	profileData, err := jwt_auth.JWTClaims(tokenString, "admin")
	if err == nil {
		if profileData["status"] != "Authorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized to access this feature"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token authorization"})
		return
	}

	var token models.Tokens
	if err := models.DB.First(&token).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token.Token,
	})
}
