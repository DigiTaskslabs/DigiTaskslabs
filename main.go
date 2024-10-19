package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"promotion/router/middleware"
	"strconv"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        InitRouter(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			fmt.Println("failed to listen ", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Start Shutting down service...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("Error shutting down server ", zap.Error(err))
	}
	fmt.Println("Server exiting...")
}
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS)
	root := router.Group("/user")
	{
		root.POST("/register", Register)
	}
	return router
}

// Register member
func Register(ctx *gin.Context) {
	var is_Register_successfully bool = false //Simulated registration success
	var task_token string = "YOUR_TASK_TOKEN" //token given after a task is created
	var telegram_id int64 = 0                 //telegram user id

	//Your registered membership code

	if is_Register_successfully {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", task_token).
			SetMultipartFormData(map[string]string{"uid": strconv.FormatInt(telegram_id, 10)}).
			Post(fmt.Sprintf("%s/%s", "https://api.digitasks.cc", "earn/task/check-ref"))
		if err != nil {
			fmt.Println("openrespProof ", zap.Error(err))
			//return err
		}
		fmt.Println(resp)
	}
}
