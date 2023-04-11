package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"yudegaki.github.com/rewrite-judge/internal/controllers"
	"yudegaki.github.com/rewrite-judge/internal/external"
	"yudegaki.github.com/rewrite-judge/internal/repositories"
)

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const expire = time.Hour * 24

func SignUp(c *gin.Context) {
	var user LoginUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: Invalid request"})
		return
	}

	// validation
	// TODO: error型を使用したvalidationにする
	if controllers.ValidateUserName(user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: Invalid UserName"})
	}
	if controllers.ValidateUserPasword(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: Invalid Password"})
	}

	encryptedPassword, err := controllers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error: Password hashing failed"})
		return
	}

	var newUser repositories.User = repositories.User{
		Name:              user.Username,
		EncryptedPassword: encryptedPassword,
	}
	if err := external.DB.Create(&newUser).Error; err != nil {
		// error
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error: Failed to create user"})
		return
	}

	jwtString, err := controllers.GenerateAuthenticationToken(newUser.ID, expire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error: Failed to generate Token"})
		return
	}
	c.SetCookie("jwt", jwtString, int(expire.Seconds()), "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    newUser,
	})
}

func SignIn(c *gin.Context) {
	var req LoginUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: Invalid request"})
		return
	}

	var user repositories.User
	if err := external.DB.Where("name = ?", req.Username).First(&user); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: Invalid UserName or Password"})
		return
	}

	// Validate Password
	if err := controllers.CheckPasswordHash(req.Password, user.EncryptedPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error: Invalid  UserName or Password"})
		return
	}

	jwtString, err := controllers.GenerateAuthenticationToken(user.ID, expire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error: Failed to generate Token"})
		return
	}
	c.SetCookie("jwt", jwtString, int(expire.Seconds()), "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{})
}

func SignOut(c *gin.Context) {
	// TODO Domainを変える
	c.SetCookie("jwt", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "SignOut Success"})
}

// ログインチェック
func isAuthenticated() gin.HandlerFunc {
	// TODO: secretを環境変数にする
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("hoge"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
