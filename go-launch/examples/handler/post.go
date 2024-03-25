package handler

import (
	"github.com/zeindevs/launch"
	"github.com/zeindevs/launch/db"
	"github.com/zeindevs/launch/examples/data"
)

// 1. Domain Data -> Post
// 2. Request Data -> createPost -> CreatePostParams

// func GetPosts(ctx *launch.Ctx) error {
// 	posts, err := data.FindAllPosts(10)
// 	if err != nil {
// 		return nil
// 	}
// 	return ctx.View("post/list.html", posts)
// }

func CreatePost(c *launch.Ctx) error {
	// params := c.RequestParams()

	post, err := db.Where[data.Post]().Limit(10).Asc("id").Run()
	if err != nil {
		return err
	}
	return c.JSON(post)
}

func GetPost(c *launch.Ctx) error {
	id := c.Param("id")
	// cannot stop this logic
	post, err := data.GetPostsById(id)
	// err -> bad id
	if err != nil {
		return err
	}
	// post, err := db.Where[data.Post]("id", id).Limit(1).Run()
	// if err != nil {
	// 	return err
	// }
	return c.JSON(post)
}
