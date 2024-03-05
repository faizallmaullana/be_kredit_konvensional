package authentication

import (
	"math/rand"
	"net/http"
	"time"
	"unicode"

	jwt_auth "github.com/faizallmaullana/be_kredit_konvensional/Authentication"
	"github.com/faizallmaullana/be_kredit_konvensional/controller"
	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputRegistration struct {
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    int    `json:"token"`
}

func Registration(c *gin.Context) {
	var input InputRegistration
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check unique phone number
	var phone string

	if len(input.Phone) > 0 && input.Phone[:2] == "08" {
		phone = input.Phone
	} else if len(input.Phone) > 0 && input.Phone[:4] == "+628" {
		phone = "0" + input.Phone[3:]
	} else if len(input.Phone) > 0 && input.Phone[:3] == "628" {
		phone = "0" + input.Phone[2:]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid phone number",
		})
	}

	var unique_phone models.Users
	if err := models.DB.Where(" phone = ? ", input.Phone).First(&unique_phone).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already used"})
		return
	}

	// check password strength
	var password string
	isValid, err := controller.CheckPasswordStrength(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if isValid {
		password, _ = controller.HashPassword(input.Password)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password strength check failed."})
		return
	}

	// write to the database (users and profiles)
	dbUsers := models.Users{
		ID:       uuid.New().String(),
		Password: password,
		Role:     input.Role,
		Phone:    phone,
		IsActive: true,
	}

	dbProfile := models.Profiles{
		ID:        uuid.New().String(),
		IDUser:    dbUsers.ID,
		Name:      toProperCase(input.Name),
		CreatedAt: time.Now().UTC().Add(7 * time.Hour),
	}

	models.DB.Create(&dbUsers)
	models.DB.Create(&dbProfile)

	// generate authentication token using jwt
	tokenJWT, err := jwt_auth.GenerateToken(dbUsers.ID, dbUsers.Phone, dbUsers.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	// validate the token
	var CekToken models.Tokens
	if err := models.DB.First(&CekToken).Error; err != nil {
		if err.Error() == "record not found" {

			if input.Token != 111111 {
				c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not accepted"})
				return
			}

			// Seed the random number generator
			rand.Seed(time.Now().UnixNano())

			// Generate a random 6-digit number
			min := 100000 // minimum value of a 6-digit number
			max := 999999 // maximum value of a 6-digit number
			randomNum := min + rand.Intn(max-min+1)

			inputToken := models.Tokens{
				ID:    uuid.New().String(),
				Token: randomNum,
			}

			models.DB.Create(&inputToken)

			input.Token = randomNum
			CekToken.Token = randomNum
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
	}

	if input.Token != CekToken.Token {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not accepted"})
		return
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	min := 100000 // minimum value of a 6-digit number
	max := 999999 // maximum value of a 6-digit number
	randomNum := min + rand.Intn(max-min+1)

	CekToken.Token = randomNum

	models.DB.Model(&CekToken).Update(CekToken)

	// return value
	c.JSON(http.StatusCreated, gin.H{
		"message":   "created success",
		"tokenAuth": tokenJWT,
		"user":      dbUsers,
	})
}

func toProperCase(s string) string {
	// Split the string into words
	words := []rune(s)
	var prev rune
	for i, r := range words {
		if !unicode.IsLetter(prev) && unicode.IsLetter(r) {
			words[i] = unicode.ToUpper(r)
		} else {
			words[i] = unicode.ToLower(r)
		}
		prev = r
	}
	return string(words)
}
