package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
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

	url := "https://api.openai.com/v1/chat/completions"

	bearerToken := os.Getenv("CHAT_GPT_TOKEN")
 
	jsonPayload, err := json.Marshal(
		dto.ChatGptMessageDto{
			Model: "gpt-3.5-turbo",
			MaxToken: 500,
			Temperature: 1.0,
			Messages: []dto.ChatGptMessage{
				{
					Role: "system",
					Content: "You are a business man with thousand of years making business",
				},
				{
					Role: "user",
					Content: userMessage.Content,
				},
			},
		},
	)

	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}
 
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))

	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}
 
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)
 
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}
 
	var response dto.ChatGPTResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}

	if len(response.Choices) == 0 {
		return []dto.MessageDto{}, errors.New("nothing return")
	}

	var gptUser models.UserModel
	db.Find(&gptUser, "id = ?", 1)
	if gptUser.ID == 0 {
		return []dto.MessageDto{}, errors.New("user was not found")
	}

	var gptMessage models.MessageModel = models.MessageModel {
		Content: response.Choices[len(response.Choices)-1].Message.Content,
		UserID: userId,
		ChatRoomId: messageDto.ChatRoomId,
		CreatedAt: time.Now(),
	}
	err = db.Create(&gptMessage).Error
	if err != nil {
		return []dto.MessageDto{}, errors.New(err.Error())
	}
	
	return []dto.MessageDto{
		{
			Content: userMessage.Content,
			ChatRoomId: userMessage.ChatRoomId,
			CompanyName: user.CompanyName,
		},
		{
			Content: gptMessage.Content,
			ChatRoomId: gptMessage.ChatRoomId,
			CompanyName: gptUser.CompanyName,
		},
	}, nil
}