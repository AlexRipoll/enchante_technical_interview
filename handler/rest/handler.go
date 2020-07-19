package rest

import (
	mysqlConfig "github.com/AlexRipoll/enchante_technical_interview/config/mysql"
	"github.com/AlexRipoll/enchante_technical_interview/internal/product"
	"github.com/AlexRipoll/enchante_technical_interview/internal/storage/mysql"
	"github.com/AlexRipoll/enchante_technical_interview/internal/user"
	"github.com/AlexRipoll/enchante_technical_interview/middleware"
	"github.com/gin-gonic/gin"
)

func Handler() {
	router := gin.Default()

	userRepository := mysql.UserRepository(mysqlConfig.Session())
	userHandler := user.NewHandler(user.NewService(userRepository))
	productRepository := mysql.ProductRepository(mysqlConfig.Session())
	productHandler := product.NewHandler(product.NewService(productRepository))

	auth := router.Group("/", middleware.Authenticate())
	authSeller := router.Group("/", middleware.AuthenticateAndCheckRole("seller"))

	auth.POST("/admin/users", userHandler.RegisterUser)
	auth.DELETE("/admin/users/:id", userHandler.Delete)

	router.POST("/login", userHandler.Login)
	router.POST("/users", userHandler.Register)
	router.GET("/users/:id", userHandler.Search)

	authSeller.POST("/users/:id/products", productHandler.Add)
	authSeller.GET("/users/:id/products/:product_id", productHandler.Search)
	authSeller.PUT("/users/:id/products/:product_id", productHandler.Update)
	authSeller.DELETE("/users/:id/products/:product_id", productHandler.Delete)
	authSeller.GET("/users/:id/products", productHandler.SearchAll)

	if err := router.Run(":9000"); err != nil {
		panic(err)
	}
}
