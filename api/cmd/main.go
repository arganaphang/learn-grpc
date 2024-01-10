package main

import (
	"context"
	"log"
	"net/http"

	pb "application/user"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	app.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	app.GET("/users", func(ctx *gin.Context) {
		res, err := client.GetUsers(context.Background(), &pb.GetUsersRequest{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed get users",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": res.Message,
			"data":    res.Data,
		})
	})

	app.Run("0.0.0.0:8000")
}
