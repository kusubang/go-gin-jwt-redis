package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kusubang/auth/internal/service"
)

func TokenAuthMiddleware(svc *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := svc.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
