package models

type User struct {
	ID       int    `json:"userId"`
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}
