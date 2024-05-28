package repo

import (
	"context"
	"encoding/json"
	"errors"
	"go-crud-redis/domain"
	"go-crud-redis/model"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type novelRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

// GetNovelById implements domain.NovelRepo.
func (n *novelRepo) GetNovelById(id int) (model.Novel, error) {
	var novels model.Novel
	var ctx = context.Background()

	// first check data is available in redis
	result, err := n.rdb.Get(ctx, "novel"+strconv.Itoa(id)).Result()
	if err != nil && err != redis.Nil {
		return novels, err
	}

	// if data available in redis, decode it from JSON, and return it
	if len(result) > 0 {
		err := json.Unmarshal([]byte(result), &novels)
		return novels, err
	}

	// if data was not available in redis, get it from databse
	err = n.db.Model(model.Novel{}).Select("id", "name", "description", "author").Where("id=?", id).Find(&novels).Error
	if err != nil {
		return novels, err
	}

	// encode that slice into json before saving into redis
	jsonBytes, err := json.Marshal(novels)
	if err != nil {
		return novels, err
	}
	jsonString := string(jsonBytes)

	// set the json-encoded value in redis
	err = n.rdb.Set(ctx, "novel"+strconv.Itoa(id), jsonString, 24*time.Hour).Err()
	if err != nil {
		return novels, err
	}

	return novels, nil
}

// CreateNovel implements domain.NovelRepo.
func (n *novelRepo) CreateNovel(createNovel model.Novel) error {
	if err := n.db.Create(&createNovel).Error; err != nil {
		return errors.New("internal server error: cannot create novel")
	}

	return nil
}

func NewNovelRepo(db *gorm.DB, rdb *redis.Client) domain.NovelRepo {
	return &novelRepo{
		db:  db,
		rdb: rdb,
	}
}
