package models

import "time"

// LIST OF TABLES
// - users
// - profiles

type Users struct {
	ID       string `json:"id" gorm:"primary_key"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

type Profiles struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	// foreign key
	IDUser string `json:"id_user"`
	User   Users  `json:"user" gorm:"references:IDUser"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
