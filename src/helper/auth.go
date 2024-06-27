package helper

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(passw *entity.Users) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passw.Password), 14)
	if err != nil {
		return string(bytes), err
	}

	return string(bytes), nil
}

func ValidPassword(passw *entity.Users, passwfromdb string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwfromdb), []byte(passw.Password))
	if err != nil {
		return err
	}
	return nil
}

func GenerateJwt(user *entity.Users) (string, error) {
	jwtKey := []byte(os.Getenv("JWTKEY"))
	expirationTime := time.Now().Add(48 * time.Hour)

	claims := &entity.Users{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Password:  "***",
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func VerifyJwt(tokenString string) error {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&entity.Users{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		},
	)

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token not valid")
	}

	claims, ok := token.Claims.(*entity.Users)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return errors.New("token is expired")
	}

	return err
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func GetUserInfoJwt(tokenString string) entity.Users {
	var userInfo entity.Users
	jwt.ParseWithClaims(
		tokenString,
		&userInfo,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		},
	)

	return userInfo
}
