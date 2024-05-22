// internal/router/router.go
package router

import (
	"github.com/tmpmadula/cantina-shop/internal/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rate limiting and other middleware can be added here if needed
	// r.Use(someMiddleware())

	r.GET("/auth/google/login", api.LoginHandler)
	r.GET("/auth/google/callback", api.CallbackHandler)

	// Protected routes
	auth := r.Group("/")
	auth.Use(api.EnsureLoggedIn)
	{
		auth.POST("/dishes", api.CreateDish)
		auth.GET("/dishes/:id", api.GetDish)
		auth.GET("/dishes", api.ListDishes)
		auth.PUT("/dishes/:id", api.UpdateDish)
		auth.DELETE("/dishes/:id", api.DeleteDish)

		auth.POST("/drinks", api.CreateDrink)
		auth.GET("/drinks/:id", api.GetDrink)
		auth.GET("/drinks", api.ListDrinks)
		auth.PUT("/drinks/:id", api.UpdateDrink)
		auth.DELETE("/drinks/:id", api.DeleteDrink)

		auth.GET("/users/:id", api.GetUser)
		auth.PUT("/users/:id", api.UpdateUser)
		auth.DELETE("/users/:id", api.DeleteUser)

		// Dashboard route to view reviews
		auth.GET("/dashboard/reviews", api.ViewReviews)
	}

	return r
}
