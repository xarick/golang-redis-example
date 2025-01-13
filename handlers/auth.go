package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-redis-example/cache"
	"github.com/xarick/golang-redis-example/db"
	"github.com/xarick/golang-redis-example/models"
	"github.com/xarick/golang-redis-example/utils"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	err = db.RegisterUser(user.FIO, user.Username, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := db.GetUserByName(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	key := "login_attempts:" + input.Username
	attempts, err := cache.RDB.Get(cache.Ctx, key).Int()
	if err == nil && attempts >= 3 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Account locked. Try again later"})
		return
	}

	err = utils.CheckPassword(user.Password, input.Password)
	if err != nil {
		cache.RDB.Incr(cache.Ctx, key)
		cache.RDB.Expire(cache.Ctx, key, 10*time.Minute)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is incorrect"})
		return
	}

	cache.RDB.Del(cache.Ctx, key)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
