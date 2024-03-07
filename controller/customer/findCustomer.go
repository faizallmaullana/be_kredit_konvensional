package customer

import (
	"net/http"

	jwt_auth "github.com/faizallmaullana/be_kredit_konvensional/Authentication"
	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
)

// find name of customer
func FindCustomer(c *gin.Context) {
	name := c.Param("name")
	tokenString := c.GetHeader("Authorization")
	profileData, err := jwt_auth.JWTClaims(tokenString, "all")
	if err == nil {
		if profileData["status"] != "Authorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized to access this feature"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token authorization"})
		return
	}

	var customer []models.Customers
	search := models.DB.Where(" id_user = ? ", profileData["id"])

	if err := search.Where(" name LIKE ? ", "%"+name+"%").Find(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Customer Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "data founded",
		"customer": customer,
	})
}
