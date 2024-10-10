package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaaaxd/go-crud/controllers"
	"github.com/jaaaxd/go-crud/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	product := r.Group("/products")
	{
		product.POST("/", controllers.CreateProduct)
		product.GET("/", controllers.GetProducts)
		product.GET("/:id", controllers.GetOneProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

	user := r.Group("/users")
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login)
		user.GET("/", controllers.GetAllUsers)
		user.GET("/:id", controllers.GetUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	r.Run(":3000") 
}
