package models

type User struct {
	ID       int    `json:"userId"`
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`

	RoleID uint `gorm:"default:0;not null" json:"roleId"`
	Role   Role `gorm:"ForeignKey:RoleID" json:"Role"`
}
