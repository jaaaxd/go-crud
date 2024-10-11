package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaaaxd/go-crud/initializers"
	"github.com/jaaaxd/go-crud/models"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {

	var reqBody models.Product
	
	// Bind incoming JSON to reqBody 
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: Title, Price, and Type are required"})
		return
	}

	// Map reqBody to var product
	product := models.Product{
		Title:          reqBody.Title,
		Subtitle:       reqBody.Subtitle,
		Desc:           reqBody.Desc,
		Price:          reqBody.Price,
		GuruInfo:       reqBody.GuruInfo,
		Type:           reqBody.Type,
		RelatedStock:   reqBody.RelatedStock,
		ExpectedReturn: reqBody.ExpectedReturn,
	}
	
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}


func GetProducts(c *gin.Context) {

	var products []models.Product

	result := initializers.DB.Find(&products)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetOneProduct(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id")) ; if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

	var product models.Product
	
	result := initializers.DB.First(&product, id)

	// Invalid id
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Other errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func UpdateProduct(c *gin.Context) {
	// Check and convert id to int
	id, err := strconv.Atoi(c.Param("id")) ; if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

	var product models.Product
	ref := initializers.DB.First(&product, id)

	// Invalid id
	if errors.Is(ref.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	// Bind incoming JSON to reqBody 
	var reqBody models.Product
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: Title, Price, and Type are required"})
		return
	}

	// Update product
	result := initializers.DB.Model(&product).Updates(models.Product{
		Title:          reqBody.Title,
		Subtitle:       reqBody.Subtitle,
		Desc:           reqBody.Desc,
		Price:          reqBody.Price,
		GuruInfo:       reqBody.GuruInfo,
		Type:           reqBody.Type,
		RelatedStock:   reqBody.RelatedStock,
		ExpectedReturn: reqBody.ExpectedReturn,
	})


	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {

	// Check and convert id to int
	id, err := strconv.Atoi(c.Param("id")) ; if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

	result := initializers.DB.Delete(&models.Product{}, id)

	// Invalid id
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}