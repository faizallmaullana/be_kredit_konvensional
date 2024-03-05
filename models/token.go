package models

type Tokens struct {
	ID    string `json:"id" gorm:"primary_key"`
	Token int    `json:"token"`
}
