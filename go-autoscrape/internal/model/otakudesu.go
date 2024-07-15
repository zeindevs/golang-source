package model

import "gorm.io/gorm"

var (
	StatusOngoing  = "ongoing"
	StatusComplete = "ongoing"
)

type Otakudesu struct {
	*gorm.Model
	Title            string              `gorm:"column:title;type:varchar(255);index"`
	Image            string              `gorm:"column:image;type:varchar(255)"`
	Url              string              `gorm:"column:url;type:varchar(255)"`
	Day              string              `gorm:"column:day;type:varchar(60)"`
	Date             string              `gorm:"column:date;type:varchar(60)"`
	Status           string              `gorm:"column:status;type:varchar(60)"`
	Episode          string              `gorm:"column:episode;type:varchar(60)"`
	Rating           string              `gorm:"column:rating;type:varchar(60)"`
	Japanese         string              `gorm:"column:japanese;type:varchar(255)"`
	Producer         string              `gorm:"column:producer;type:varchar(255)"`
	TotalEpisode     string              `gorm:"column:total_episode;type:varchar(255)"`
	Duration         string              `gorm:"column:duration;type:varchar(255)"`
	ReleasedAt       string              `gorm:"column:released_at;type:varchar(255)"`
	Studio           string              `gorm:"column:studio;type:varchar(255)"`
	Genres           string              `gorm:"column:genres;type:varchar(255)"`
	OtakudesuEpisode []*OtakudesuEpisode `gorm:""`
}

func (Otakudesu) TableName() string {
	return "otakudesu"
}

type OtakudesuEpisode struct {
	*gorm.Model
	Title           string            `gorm:"column:title;type:varchar(255)"`
	Episode         string            `gorm:"column:episode;type:varchar(60)"`
	Url             string            `gorm:"column:url;type:varchar(255)"`
	OtakudesuID     uint              `gorm:"column:otakudesu_id"`
	OtakudesuVideos []*OtakudesuVideo `gorm:""`
}

func (OtakudesuEpisode) TableName() string {
	return "otakudesu_episode"
}

type OtakudesuVideo struct {
	*gorm.Model
	Title              string `gorm:"column:title;type:varchar(255)"`
	Url                string `gorm:"column:url;type:varchar(255)"`
	Provider           string `gorm:"column:provider;type:varchar(255)"`
	Quality            string `gorm:"column:quality;type:varchar(255)"`
	OtakudesuEpisodeID uint   `gorm:"column:otakudesu_episode_id"`
}

func (OtakudesuVideo) TableName() string {
	return "otakudesu_video"
}
