package models

import "time"

// LIST OF TABLES
// - custommers

type Customers struct {
	ID              string `json:"id" gorm:"primary_key"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	RemainingCredit string `json:"remaining_credit"`
	Note            string `json:"note"`

	// foreign key
	IDUser string `json:"id_user"`
	User   Users  `json:"user" gorm:"references:IDUser"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
