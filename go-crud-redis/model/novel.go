package model

import "gorm.io/gorm"

type Novel struct {
	*gorm.Model
	Id          int    `gorm:"type:int;primary_key" json:"id"`
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Author      string `gorm:"type:varchar(50);not null" json:"author"`
	Description string `gorm:"type:varchar(50);not null" json:"description"`
}
