package data

import (
	"github.com/zeindevs/launch/db"
	"github.com/zeindevs/launch/examples/handler/params"
)

// Business domain
type Post struct {
	ID          int
	Title       string
	Description string
}

func GetPostsById(id any) (Post, error) {
	return db.Where[Post]("id", id).Limit(1).Asc("id").Run()
}

func FindAllPosts(limit int) (Post, error) {
	return db.Where[Post]().Limit(limit).Asc("id").Run()
}

func CreatePost(params params.CreatePost) (Post, error) {
	return db.Where[Post]().Limit(10).Asc("id").Run()
}

func SpecialQuery() (Post, error) {
	return db.Raw[Post]("SELECT * FROM USER").Limit(1).Desc("title").Run()
}
