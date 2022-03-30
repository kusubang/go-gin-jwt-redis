package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/kusubang/auth/internal/common"
	"github.com/kusubang/auth/internal/handler"
	"github.com/kusubang/auth/internal/middleware"
	"github.com/kusubang/auth/internal/service"
)

var authService *service.AuthService
var client *redis.Client

func init() {
	client = common.GetRedisClient()
	authService = service.NewAuthService(client)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	taskServer := handler.NewTaskServer(authService)
	authServer := handler.NewAuthServer(authService)

	router.GET("/task/", taskServer.GetAllTasksHandler)
	router.GET("/task/:id", taskServer.GetTaskHandler)
	router.POST("/task/", taskServer.CreateTaskHandler)

	router.POST("/login/", authServer.LoginHandler)
	router.POST("/logout/", middleware.TokenAuthMiddleware(authService), authServer.LogoutHandler)
	router.POST("/token/refresh", authServer.RefreshHandler)

	return router
}
