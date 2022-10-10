package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Meetups  []*Meetup `json:"meetups" gorm:"foreignKey:UserID;references:ID"`
}
