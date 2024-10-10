package controllers

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jaaaxd/go-crud/initializers"
	"github.com/jaaaxd/go-crud/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {

	var reqBody models.User
	
	// get reqBody 
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Check if email already exists
	var existingUser models.User
	res := initializers.DB.Where("email = ?", reqBody.Email).First(&existingUser)
	
	if res.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		return
	}

	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Map to var user
	user := models.User{
		Email:        reqBody.Email,
		Password:     string(hashPassword),
		Firstname:    reqBody.Firstname,
		Lastname:     reqBody.Lastname,
		Experience:   reqBody.Experience,
		Type:         reqBody.Type,
		PhoneNumber:  reqBody.PhoneNumber,
		Birthday:     reqBody.Birthday,
	}

	// Create user
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registered successfully"})
}


func Login(c *gin.Context) {

	var reqBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind incoming JSON to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	var user models.User

	// Find user by email
	result := initializers.DB.Where("email = ?", reqBody.Email).First(&user); 
	
	// Invalid email
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password incorrect"})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), 
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "Login successfully"})
}


func GetAllUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id")) ; if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

	var user models.User
	result := initializers.DB.First(&user, id)

	// Invalid id
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Other errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update info other than Email and Password
func UpdateUser(c *gin.Context) {
	// Check and convert id to int
	id, err := strconv.Atoi(c.Param("id")) ; if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

	var user models.User
	ref := initializers.DB.First(&user, id)

	// Invalid id
	if errors.Is(ref.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	// Bind incoming JSON to reqBody 
	var reqBody struct {
		Firstname   string    `json:"firstname" binding:"required"`
		Lastname    string    `json:"lastname" binding:"required"`
		Experience  string    `json:"experience" binding:"required"`
		Type        string    `json:"type" binding:"required"`
		PhoneNumber string    `json:"phone_number" binding:"required"`
		Birthday    time.Time `json:"birthday" binding:"required"`
	
	}
	
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Update user (excluding Email and Password)
	result := initializers.DB.Model(&user).Updates(models.User{
		Firstname:   reqBody.Firstname,
		Lastname:    reqBody.Lastname,
		Experience:  reqBody.Experience,
		Type:        reqBody.Type,
		PhoneNumber: reqBody.PhoneNumber,
		Birthday:    reqBody.Birthday,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {

	// Check and convert id to int
	id, err := strconv.Atoi(c.Param("id")) ; if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

	result := initializers.DB.Delete(&models.User{}, id)

	// Invalid id
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
