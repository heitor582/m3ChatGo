package services

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/heitor582/m3ChatGo/src/configuration"
	dto "github.com/heitor582/m3ChatGo/src/dto"
	models "github.com/heitor582/m3ChatGo/src/models"
	"github.com/heitor582/m3ChatGo/src/functions"
)

func Login(loginDto dto.LoginDto) (string, error) {
	db := configuration.DBConn
	var user models.UserModel
	db.Find(&user, "email = ?", loginDto.Email)
	if user.ID == 0 {
		return "", errors.New("user was not found")
	}
	if !functions.CheckPasswordHash(loginDto.Password, user.Password) {
		return "", errors.New("passwords dont match")
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["company_name"] = user.CompanyName
	claims["id"] = user.ID

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}
func RegisterUser(registerDto dto.RegisterDto) error {
	db := configuration.DBConn
	hashedPassword, err := functions.HashPassword(registerDto.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	var user models.UserModel = models.UserModel{
		CompanyName:  registerDto.CompanyName,
		Email: registerDto.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err = db.Create(&user).Error; err != nil {
		return errors.New(err.Error())
	}
	return nil
}