package credits

import (
	"net/http"

	jwt_auth "github.com/faizallmaullana/be_kredit_konvensional/Authentication"
	"github.com/gin-gonic/gin"
)

type InputCredit struct {
	IDCustomer          string `json:"id_customer"`
	NamaCustomer        string `json:"nama_customer"`
	NomorTelponCustomer string `json:"nomor_telpon_customer"`
	AlamatCustomer      string `json:"alamat_customer"`
	HargaModal          int    `json:"harga_modal"`
	HargaJual           int    `json:"harga_jual"`
	Cicilan             int    `json:"cicilan"`
	PeriodeCicilan      string `json:"periode_cicilan"`
	DicicilSetiap       string `json:"cicilan_setiap_hari"`
}

func NewCreditAndCustomer(c *gin.Context) {
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
