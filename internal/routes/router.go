package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/kusubang/auth/internal/common"
	"github.com/kusubang/auth/internal/handler"
	"github.com/kusubang/auth/internal/logger"
	"github.com/kusubang/auth/internal/middleware"
	"github.com/kusubang/auth/internal/service"
	"github.com/sirupsen/logrus"
)

var authService *service.AuthService
var client *redis.Client

var authLogger *logrus.Entry
var taskLogger *logrus.Entry

func init() {
	client = common.GetRedisClient()
	authService = service.NewAuthService(client)

	logger := logger.GetLogger()

	authLogger = logger.WithFields(logrus.Fields{
		"component": "handler",
		"category":  "auth",
	})
	taskLogger = logger.WithFields(logrus.Fields{
		"component": "handler",
		"category":  "task",
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	authServer := handler.NewAuthServer(authService, authLogger)
	taskServer := handler.NewTaskServer(authService, taskLogger)

	router.Static("/public", "./public")

	router.GET("/task/", taskServer.GetAllTasksHandler)
	router.GET("/task/:id", taskServer.GetTaskHandler)
	router.POST("/task/", taskServer.CreateTaskHandler)

	router.POST("/login/", authServer.LoginHandler)
	router.POST("/logout/", middleware.TokenAuthMiddleware(authService), authServer.LogoutHandler)
	router.POST("/token/refresh", authServer.RefreshHandler)

	return router
}
