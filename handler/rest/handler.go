package rest

import (
	mysqlConfig "github.com/AlexRipoll/enchante_technical_interview/config/mysql"
	"github.com/AlexRipoll/enchante_technical_interview/internal/storage/mysql"
	"github.com/AlexRipoll/enchante_technical_interview/internal/user"
	"github.com/gin-gonic/gin"
)

func Handler() {
	router := gin.Default()

	userRepository := mysql.Repository(mysqlConfig.Session())
	userHandler := user.NewHandler(user.NewService(userRepository))

	router.POST("/users", userHandler.Register)

	if err := router.Run(":9000"); err != nil {
		panic(err)
	}
}
