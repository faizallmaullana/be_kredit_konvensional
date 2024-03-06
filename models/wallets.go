package models

import "time"

// LIST OF TABLES
// - wallets
// - other_transaction

type Wallets struct {
	ID     string `json:"id" gorm:"primary_key"`
	Amount int    `json:"amount"`

	// foreign key
	IDUser string `json:"id_user"`
	User   Users  `json:"user" gorm:"references:IDUser"`
}

type OtherTransaction struct {
	ID          string `json:"id" gorm:"primary_key"`
	Transaction string `json:"transaction"`
	Amount      int    `json:"amount"`
	IsDebit     string `json:"is_debit"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
