package models

import "time"

// LIST OF TABLES
// - credits
// - payments

type Credits struct {
	ID           string `json:"id" gorm:"primary_key"`
	Product      string `json:"product"`
	CapitalPrice int    `json:"capital_price"`
	SellingPrice int    `json:"selling_price"`
	Cicilan      int    `json:"cicilan"`
	Periode      string `json:"periode"`
	PayEvery     string `json:"pay_every"`

	// foreign key
	IDCustomer string    `json:"id_customer"`
	Customer   Customers `json:"customer" gorm:"references:IDCustomer"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Payments struct {
	ID string `json:"id" gorm:"primary_key"`

	// foreign key
	IDCredit string  `json:"id_credit"`
	Credit   Credits `json:"credits" gorm:"references:IDCredits"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
