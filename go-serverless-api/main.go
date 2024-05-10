package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  // If no name is provided in the HTTP request body, throw an error
  return ginLambda.ProxyWithContext(ctx, req)
}

func init() {
  // stdout and stderr are sent to AWS CloudWatch Logs
  log.Print("Gin cold start")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
	})

	ginLambda = ginadapter.New(r)
}

func main() {
  lambda.Start(HandleRequest)
}
