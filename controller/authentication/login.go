package authentication

import (
	"net/http"

	jwt_auth "github.com/faizallmaullana/be_kredit_konvensional/Authentication"
	"github.com/faizallmaullana/be_kredit_konvensional/controller"
	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
)

type InputLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Algorithm
// 1. Get input
// 2. check existing phone number
// 3. password validation
// 4. if all pass, generate tokenauth
// 5. create return data

func LoginResource(c *gin.Context) {
	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check existing phone
	var user models.Users
	if err := models.DB.Where(" phone = ? ", input.Phone).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is invalid"})
		return
	}

	// pasword validation
	if !controller.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is invalid"})
		return
	}

	// generate token auth
	tokenJWT, err := jwt_auth.GenerateToken(user.ID, user.Phone, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	// return data
	c.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"tokenAuth": tokenJWT,
		"user":      user,
	})
}
