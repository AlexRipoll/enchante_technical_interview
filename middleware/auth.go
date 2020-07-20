package middleware

import (
	"github.com/AlexRipoll/enchante_technical_interview/config"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := c.Request.Header.Get("Authorization")
		t = strings.TrimSpace(strings.Replace(t, "Bearer", "", -1))

		jwtService := jwt.NewService(config.Params.Jwt.Key, config.Params.Jwt.Ttl)

		if err := jwtService.ValidateToken(t); err != nil {
			tErr := errors.NewError(http.StatusUnauthorized, "not authorized")
			c.JSON(tErr.Status, tErr)
			c.Abort()
			return
		}

		claims := jwtService.Claims()
		userId := claims["id"].(string)
		userRole := claims["role"].(string)
		c.Set("id", userId)
		c.Set("role", userRole)
		c.Next()
	}
}

func AuthenticateAndCheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		t := c.Request.Header.Get("Authorization")
		t = strings.TrimSpace(strings.Replace(t, "Bearer", "", -1))

		jwtService := jwt.NewService(config.Params.Jwt.Key, config.Params.Jwt.Ttl)

		if err := jwtService.ValidateToken(t); err != nil {
			tErr := errors.NewError(http.StatusUnauthorized, "not authorized")
			c.JSON(tErr.Status, tErr)
			c.Abort()
			return
		}

		claims := jwtService.Claims()
		userId := claims["id"].(string)
		userRole := claims["role"].(string)
		if userId == "" || userRole == "" {
			err := errors.NewUnauthorizedError("missing claim")
			c.JSON(err.Status, err)
			c.Abort()
			return
		}
		sellerId := strings.TrimSpace(c.Param("id"))
		if userId != sellerId {
			err := errors.NewForbiddenAccessError("forbidden access")
			c.JSON(err.Status, err)
			c.Abort()
			return
		}

		if role != userRole {
			apiErr := errors.NewForbiddenAccessError("forbidden access")
			c.JSON(apiErr.Status, apiErr)
			c.Abort()
			return
		}
	}
}