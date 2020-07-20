package rest

import (
	mysqlConfig "github.com/AlexRipoll/enchante_technical_interview/config/mysql"
	"github.com/AlexRipoll/enchante_technical_interview/internal/cart"
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
	cartRepository := mysql.CartRepository(mysqlConfig.Session())
	cartHandler := cart.NewHandler(cart.NewService(cartRepository))

	authAdmin := router.Group("/", middleware.AuthenticateAndCheckRole("admin"))
	authSeller := router.Group("/", middleware.AuthenticateAndCheckRole("seller"))
	authUser := router.Group("/", middleware.AuthenticateAndCheckRole("user"))

	router.POST("/login", userHandler.Login)
	router.POST("/users", userHandler.Register)

	// admin only access endpoints
	authAdmin.POST("/admin/users", userHandler.RegisterUser)
	authAdmin.DELETE("/admin/users/:id", userHandler.Delete)
	authAdmin.GET("/users/:id", userHandler.Search)

	// seller only access endpoints
	authSeller.POST("/users/:id/products", productHandler.Add)
	authSeller.GET("/users/:id/products/:product_id", productHandler.Search)
	authSeller.PUT("/users/:id/products/:product_id", productHandler.Update)
	authSeller.DELETE("/users/:id/products/:product_id", productHandler.Delete)
	authSeller.GET("/users/:id/products", productHandler.SearchAll)

	// user only access endpoints
	authUser.POST("/users/:id/orders", cartHandler.Purchase)

	if err := router.Run(":9000"); err != nil {
		panic(err)
	}
}
