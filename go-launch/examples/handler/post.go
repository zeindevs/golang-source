package handler

import (
	"github.com/zeindevs/launch"
	"github.com/zeindevs/launch/db"
	"github.com/zeindevs/launch/examples/data"
	"github.com/zeindevs/launch/examples/handler/params"
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

func CreatePost(ctx *launch.PostCtx[params.CreatePost]) error {
	// params := ctx.RequestParams()

	post, err := db.Where[data.Post]().Limit(10).Asc("id").Run()
	if err != nil {
		return err
	}
	return ctx.JSON(post)
}

func GetPost(ctx *launch.Ctx) error {
	id := ctx.Param("id")
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
	return ctx.JSON(post)
}
