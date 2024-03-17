package service

import (
	"encoding/base64"
	"os"
	"testjun/pkg/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
 jwt.StandardClaims
}

func NewAccessToken(lg string) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS512)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["sub"] = lg

	return accessToken.SignedString([]byte(os.Getenv("TOP")))
}

func NewRefreshToken(lg string) (string, error) {
	hash, err := HashToken(lg)
	if err != nil{
		return "", err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS512)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["sub"] = hash

 	return refreshToken.SignedString([]byte(os.Getenv("SECRET")))
}

func ParseAccessToken(accessToken string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("TOP")), nil
	})
	if err != nil {
		return &UserClaims{}, nil
	}

	return parsedAccessToken.Claims.(*UserClaims), nil
}

func ParseRefreshToken(refreshToken string) (*jwt.StandardClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil{
		return &jwt.StandardClaims{}, nil
	}

 return parsedRefreshToken.Claims.(*jwt.StandardClaims), err
}
func CreateToken(lg string, ) (*models.Token, error){
	t, err := NewAccessToken(lg)
	if err != nil {
	 return &models.Token{}, err
 }
	ref, err := NewRefreshToken(lg)
	if err != nil {
		return &models.Token{}, err
	}
	rtoken := base64.StdEncoding.EncodeToString([]byte(ref))
	return &models.Token{
		AccessToken: t,
		RefreshToken: rtoken,
	}, nil
}