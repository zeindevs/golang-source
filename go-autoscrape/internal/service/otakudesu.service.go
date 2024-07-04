package service

import (
	"context"
	"log"

	"github.com/zeindevs/goautoscrape/internal/model"
	"gorm.io/gorm"
)

type OtakudesuService struct {
	db *gorm.DB
}

func NewOtakudesuService(db *gorm.DB) *OtakudesuService {
	return &OtakudesuService{
		db: db,
	}
}

func (s *OtakudesuService) Save(ctx context.Context, data []*Otakudesu) error {
	created := 0
	updated := 0
	for _, item := range data {
		var exist model.Otakudesu
		resp := s.db.WithContext(ctx).First(&exist, "title = ?", item.Title)
		if resp.Error == nil {
			update := s.db.WithContext(ctx).Where("id = ?", exist.ID).Updates(item.ToModel())
			if update.Error != nil {
				return update.Error
			}
			log.Println("Update", item.Title)
			updated++
		} else {
			create := s.db.WithContext(ctx).Create(item.ToModel())
			if create.Error != nil {
				return create.Error
			}
			log.Println("Create", item.Title)
			created++
		}
	}

	log.Println("Created:", created, "| Updated:", updated)

	return nil
}
