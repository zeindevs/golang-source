package model

import "gorm.io/gorm"

var (
	StatusOngoing  = "ongoing"
	StatusComplete = "ongoing"
)

type Otakudesu struct {
	*gorm.Model
	Title   string `gorm:"column:title;type:varchar(255);index"`
	Image   string `gorm:"column:image;type:varchar(255)"`
	Url     string `gorm:"column:url;type:varchar(255)"`
	Day     string `gorm:"column:day;type:varchar(60)"`
	Date    string `gorm:"column:date;type:varchar(60)"`
	Status  string `gorm:"column:status;type:varchar(60)"`
	Episode string `gorm:"column:episode;type:varchar(60)"`
}

func (Otakudesu) TableName() string {
	return "otakudesu"
}
