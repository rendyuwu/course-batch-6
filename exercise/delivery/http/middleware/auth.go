package middleware

import (
	"context"
	"exercise/domain/entity"
	"exercise/domain/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, web.ResponseError{Message: "UNAUTHORIZED"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, web.ResponseError{Message: "UNAUTHORIZED"})
			c.Abort()
			return
		}

		auths := strings.Split(authHeader, " ")
		if len(auths) != 2 {
			c.JSON(http.StatusUnauthorized, web.ResponseError{Message: "UNAUTHORIZED"})
			c.Abort()
			return
		}

		var user entity.User
		data, err := user.DecryptJWT(auths[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, web.ResponseError{Message: "UNAUTHORIZED"})
			c.Abort()
			return
		}

		if data["user_id"] == nil {
			c.JSON(http.StatusUnauthorized, web.ResponseError{Message: "UNAUTHORIZED"})
			c.Abort()
			return
		}

		if data["user_id"].(float64) < 1 {
			c.JSON(http.StatusUnauthorized, web.ResponseError{Message: "UNAUTHORIZED"})
			c.Abort()
			return
		}

		userID := int(data["user_id"].(float64))
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", userID)
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
