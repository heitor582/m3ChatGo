package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	dto "github.com/heitor582/m3ChatGo/src/dto"
	"github.com/heitor582/m3ChatGo/src/services"
)

func SendMessage(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)

	message := new(dto.MessageDto)
	if err := c.BodyParser(message); err != nil {
		c.JSON(fiber.ErrBadRequest)
		return errors.New(err.Error())
	}

	messages, err := services.NewMessage(*message, uint64(userId))
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(messages)
	return nil
}