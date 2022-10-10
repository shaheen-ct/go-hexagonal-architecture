package entity

import "gorm.io/gorm"

type Meetup struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int64  `json:"-"`
}
