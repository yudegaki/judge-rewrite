package controllers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// TODO: secretを環境変数にする
const authSecretKey = "hoge"

func GenerateAuthenticationToken(id uint, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(authSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetClaimsSubject(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(authSecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["sub"].(uint)
		if ok {
			return 0, fmt.Errorf("Error: Can't decode token")
		}
		// Success
		return id, nil
	}
	return 0, fmt.Errorf("Error: Can't decode token")
}

func ValidateUserName(name string) bool {
	MAX_LENGTH := 20
	// USER_NAME_RULE := regexp.MustCompile("^[0-9a-zA-Z]*$")
	// 20文字以下
	if len(name) > MAX_LENGTH {
		return false
	}
	return true
}

func ValidateUserPasword(password string) bool {
	const MAX_LENGTH = 20
	// 20文字以下
	if len(password) > MAX_LENGTH {
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
