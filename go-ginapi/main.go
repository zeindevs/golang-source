package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
}

func router() *gin.Engine {
	r := gin.Default()
	userRoute := r.Group("/user")
	{
		userRoute.GET("/hello/:name", func(c *gin.Context) {
			user := c.Param("name")
			c.String(200, fmt.Sprintf("hello, %s", user))
		})
		userRoute.POST("/post", func(c *gin.Context) {
			body := Message{}
			if err := c.BindJSON(&body); err != nil {
				// c.AbortWithError(http.StatusBadRequest, err)
				c.JSON(http.StatusBadRequest, map[string]string{"errors": err.Error()})
				return
			}
			fmt.Println(body)
			c.JSON(http.StatusAccepted, &body)
		})
		userRoute.POST("/upload", func(c *gin.Context) {
			file, _ := c.FormFile("file")
			log.Println(file.Filename)

			c.SaveUploadedFile(file, "/tmp/tempfile")
			c.String(http.StatusOK, fmt.Sprintf("'%s', uploaded!", file.Filename))
		})
	}
	return r
}

func main() {
	router().Run()
}
