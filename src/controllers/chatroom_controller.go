package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	dto "github.com/heitor582/m3ChatGo/src/dto"
	"github.com/heitor582/m3ChatGo/src/services"
)

func GetAllChatRooms(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	c.JSON(services.GetAllChatRooms(uint64(userId)))
	return nil
}

func GetChatRoom(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	chatRoom, err := services.GetChatRoom(id, uint64(userId))
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(chatRoom)
	return nil
}

func NewChatRoom(c *fiber.Ctx) error {
	todoDto := new(dto.CreateChatRoomDto)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	if err := c.BodyParser(todoDto); err != nil {
		c.JSON(fiber.ErrBadRequest)
		return errors.New(err.Error())
	}

	todo, err := services.NewChatRoom(*todoDto, uint64(userId))
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(todo)
	return nil
}

func DeleteChatRoom(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	err = services.DeleteChatRoom(id)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	return nil
}