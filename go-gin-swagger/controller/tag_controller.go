package controller

import "github.com/gin-gonic/gin"

type TagsController struct {
	tagService any
}

func NewTagsController(service any) *TagsController {
	return &TagsController{
		tagService: service,
	}
}

// CreateTags   godoc
// @Summary     Create tags
// @Description Save tags data in Db.
// @Param       tags body request.CreateTagsRequest true "Create tags"
// @Produce     application/json
// @Success     200 {object} response.Response{}
// @Router      /tags [post]
func (c *TagsController) Create(ctx *gin.Context) {

}
