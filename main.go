package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Blog struct {
	*gorm.Model
	ID    uint `gorm: "primaryKey"`
	Title string
	Text  string
}

type User struct {
	*gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

var db *gorm.DB

// Blog Handlers
func GetBlogs(c *gin.Context) {
	var blogs []Blog

	db.Find(&blogs)

	c.JSON(200, gin.H{
		"blogs": blogs,
	})
}

func GetBlogById(c *gin.Context) {
	var blog Blog
	id := c.Param("id")

	db.Find(&blog, id)

	c.JSON(http.StatusOK, gin.H{
		"post": blog,
	})
}

func CreateBlog(c *gin.Context) {
	var blog struct {
		Title string
		Text  string
	}

	c.BindJSON(&blog)

	result := db.Create(
		&Blog{Title: blog.Title, Text: blog.Text},
	)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "Blog was created successfully",
	})
}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&Blog{}, id)
	c.Status(200)
}

// User Handlers
func RegisterUser(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Create user
	result := db.Create(&User{Username: body.Username, Password: string(hash)})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User has been created"})
}

func LoginUser(c *gin.Context) {
	// Get email and pass off reg body
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	// Look up requested user
	var user User

	db.First(&user, "username= ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare sent in pass with saved user pass has
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	privateKey := []byte("my-secret-key")
	fmt.Println(privateKey)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	// send it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetUsers(c *gin.Context) {

}

// Database Connection handler
func ConnectToDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("primary.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Blog{}, &User{})
}

// Main execution
func main() {
	ConnectToDatabase()
	router := gin.Default()

	// Blog routes
	router.GET("/blog", GetBlogs)
	router.GET("blog/:id", GetBlogById)
	router.POST("/blog", CreateBlog)
	router.DELETE("/blog/:id", DeleteBlog)

	// User routes
	router.POST("/user/register", RegisterUser)
	router.POST("/user/login", LoginUser)
	router.GET("/user/validate", RequireAuth, Validate)
	router.GET("/user", GetUsers)

	router.Run("localhost:8080")
}
