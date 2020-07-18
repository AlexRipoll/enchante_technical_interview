package middleware

import (
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

		jwtService := jwt.NewService("secret-Key", 3600)

		if err := jwtService.ValidateToken(t); err != nil {
			tErr := errors.NewError(http.StatusUnauthorized, "not authorized")
			c.JSON(tErr.Status, tErr)
			c.Abort()
			return
		}

		claims := jwtService.Claims()
		userId := claims["id"].(string)
		c.Set("id", userId)
		c.Next()
	}
}
