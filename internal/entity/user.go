package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model

	DisplayName string `json:"displayName"`
	PhoneNumber string `json:"phoneNumber"`
}

func (u *User) Table() string {
	return "users"
}
