package services

import (
	"errors"
	"log"
	"time"

	"github.com/heitor582/m3ChatGo/src/configuration"
	dto "github.com/heitor582/m3ChatGo/src/dto"
	models "github.com/heitor582/m3ChatGo/src/models"
)

func NewMessage(messageDto dto.MessageDto, userId uint64) ([]dto.MessageDto, error) {
	db := configuration.DBConn

	var user models.UserModel
	db.Find(&user, "id = ?", userId)
	if user.ID == 0 {
		return []dto.MessageDto{}, errors.New("user was not found")
	}

	var userMessage models.MessageModel = models.MessageModel {
		Content: messageDto.Content,
		UserID: userId,
		ChatRoomId: messageDto.ChatRoomId,
		CreatedAt: time.Now(),
	}
	err := db.Create(&userMessage).Error
	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}

	var gptUser models.UserModel
	db.Find(&gptUser, "id = ?", 1)
	if gptUser.ID == 0 {
		return []dto.MessageDto{}, errors.New("user was not found")
	}

	log.Println("[mandando mensagem] Achou os usuário e irá mandar a mensagem para a ia")
	
	gptMessage, err := SendMessageToChatGpt(userMessage.Content, messageDto.ChatRoomId, userId)
	if err != nil {
		log.Printf("erro ao mandar mensagem pro gpt %s", err)
		return []dto.MessageDto{}, errors.New(err.Error())
	}

	log.Println("[mandou com sucesso] Mandou a mensagem e a ia retornou")

	err = db.Create(&gptMessage).Error
	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}
	
	return []dto.MessageDto{
		{
			Content: userMessage.Content,
			ChatRoomId: userMessage.ChatRoomId,
			CompanyName: user.CompanyName,
			UserID: userId,
			CreatedAt: userMessage.CreatedAt,
		},
		{
			Content: gptMessage.Content,
			ChatRoomId: gptMessage.ChatRoomId,
			CompanyName: gptUser.CompanyName,
			UserID: gptMessage.ID,
			CreatedAt: gptMessage.CreatedAt,
		},
	}, nil
}