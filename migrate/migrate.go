package main

import (
	"fmt"

	"github.com/jaaaxd/go-crud/initializers"
	"github.com/jaaaxd/go-crud/models"
)


func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Product{}, &models.User{})
	fmt.Println("successfully migrate")

}