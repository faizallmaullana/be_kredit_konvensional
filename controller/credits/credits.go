package credits

import (
	"net/http"
	"time"

	jwt_auth "github.com/faizallmaullana/be_kredit_konvensional/Authentication"
	"github.com/faizallmaullana/be_kredit_konvensional/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputCredit struct {
	IDCustomer          string `json:"id_customer"`
	NamaCustomer        string `json:"nama_customer"`
	NomorTelponCustomer string `json:"nomor_telpon_customer"`
	AlamatCustomer      string `json:"alamat_customer"`
	Product             string `json:"product"`
	HargaModal          int    `json:"harga_modal"`
	HargaJual           int    `json:"harga_jual"`
	Cicilan             int    `json:"cicilan"`
	PeriodeCicilan      string `json:"periode_cicilan"`
	DicicilSetiap       string `json:"cicilan_setiap_hari"`
}

func NewCreditAndCustomer(c *gin.Context) {
	// check authorization
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

	// format input should be json
	var input InputCredit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// timeConvertion
	timeNow := time.Now().Add(7 * time.Hour)
	date := timeNow.Truncate(24 * time.Hour)
	formattedDate := date.Format("02")

	// check pay dialy, weekly, or monthly
	var hariCicilan string
	var periode string
	if input.PeriodeCicilan == "Harian" {
		hariCicilan = ""
		periode = "h"
	} else if input.PeriodeCicilan == "Mingguan" {
		hariCicilan = input.DicicilSetiap
		periode = "m"
	} else if input.PeriodeCicilan == "Bulanan" {
		hariCicilan = formattedDate
		periode = "b"
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there is an error on pay checking"})
		return
	}

	// write the data to the database
	idMitra, _ := profileData["id"].(string)

	Customer := models.Customers{
		ID:              uuid.New().String(),
		Name:            input.NamaCustomer,
		Phone:           input.NomorTelponCustomer,
		Address:         input.AlamatCustomer,
		RemainingCredit: input.HargaJual,
		IDUser:          idMitra,
		CreatedAt:       timeNow,
		IsDeleted:       false,
	}

	Credit := models.Credits{
		ID:           uuid.New().String(),
		IDCustomer:   Customer.ID,
		IDUser:       idMitra,
		Product:      input.Product,
		PayEvery:     hariCicilan,
		CapitalPrice: input.HargaModal,
		SellingPrice: input.HargaJual,
		Cicilan:      input.Cicilan,
		Periode:      periode,
		CreatedAt:    timeNow,
		IsDeleted:    false,
	}

	models.DB.Create(&Customer)
	models.DB.Create(&Credit)

	c.JSON(http.StatusCreated, gin.H{
		"message":   "created",
		"id_mitra":  Customer.IDUser,
		"id_credit": Credit.ID,
	})
}

func NewCreditOnly(c *gin.Context) {
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
}
