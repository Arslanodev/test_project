package controllers

import (
	"example/blog-app/app/models"
	"example/blog-app/app/repositories"
	"example/blog-app/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var body struct {
		Username string
		Password string
		Admin    bool
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Check if user exists
	user, _ := c.repo.GetUserByUsername(body.Username)
	if user.ID != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash a password",
		})

		return
	}

	// Create user
	result := c.repo.RegisterUser(models.User{Username: body.Username, Password: string(hash), Admin: body.Admin})

	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a user",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User has been created"})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	// Get email and pass off reg body
	var body struct {
		Username string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	// Look up requested user
	user, _ := c.repo.GetUserByUsername(body.Username)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare sent in pass with saved user pass has
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}
	// Generate a jwt token
	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *UserController) RequireAuth(ctx *gin.Context) {
	// Get the token from request headers
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Authorization header is missing",
		})

		return
	}
	// Decode/validate it
	token, err := auth.Decode(tokenString)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the exp
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token sub
		var user models.User
		user_id := int(claims["sub"].(float64))
		user, _ = c.repo.GetUserById(user_id)
		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		ctx.Set("user", user)

		// Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}

func (c *UserController) IsAdmin(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	user_data, _ := c.repo.GetUserById(int(user.ID))
	if !user_data.Admin {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

	ctx.Next()
}
