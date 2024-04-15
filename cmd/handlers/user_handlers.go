package handlers

import (
	"example/blog-app/cmd/data"
	"example/blog-app/cmd/repositories"
	"example/blog-app/cmd/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	var body struct {
		Name     string
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
	user, err := h.repo.GetUserByUsername(body.Username)
	if user.Username == body.Username {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
		return
	}

	if err != nil {
		log.Println("Identical User not found")
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
	result := h.repo.CreateNewUser(data.User{Name: body.Name, Username: body.Username, Password: string(hash), Admin: body.Admin})

	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a user",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User has been created"})
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
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
	user, _ := h.repo.GetUserByUsername(body.Username)
	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare sent in pass with saved user pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}
	// Generate a jwt token
	token, err := utils.GenerateToken(user.Username, user.Password)

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

func (h *UserHandler) RequireAuth(ctx *gin.Context) {
	// Get the token from request headers
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Authorization header is missing",
		})

		return
	}
	// Decode
	tokenString = tokenString[len("Bearer "):]
	token, err := utils.Decode(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
		return
	}

	// Validate
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Find the user with token username
		var user data.User
		username := claims["username"]
		password := claims["password"]

		if username == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, _ = h.repo.GetUserByUsername(username.(string))

		// Check if user exists
		if user.Username != username {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
			return
		}

		// Check password
		if password != user.Password {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
		}

		// Check the exp
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token expired",
			})
		}

		// Attach to req
		ctx.Set("user", user)

		// Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}

func (h *UserHandler) IsAdmin(ctx *gin.Context) {
	user := ctx.MustGet("user").(data.User)
	user_data, _ := h.repo.GetUserByID(int(user.ID))
	if !user_data.Admin {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

	ctx.Next()
}
