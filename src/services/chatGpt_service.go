package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	dto "github.com/heitor582/m3ChatGo/src/dto"
	models "github.com/heitor582/m3ChatGo/src/models"
)

func SendMessageToChatGpt(userMessage string, chatRoomId uint64, userId uint64) (models.MessageModel, error) {
	url := "https://api.openai.com/v1/chat/completions"

	bearerToken := os.Getenv("CHAT_GPT_TOKEN")

	newMessage := dto.ChatGptMessage{
		Role: "user",
		Content: userMessage,
	}
 
	jsonPayload, err := json.Marshal(dto.NewChatGptMessageDto(newMessage))

	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}
 
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))

	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}
 
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)
 
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Println(body)

	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}
 
	var response dto.ChatGPTResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}

	if len(response.Choices) == 0 {
		return models.MessageModel{}, errors.New("nothing return")
	}	

	var gptMessage models.MessageModel = models.MessageModel {
		Content: response.Choices[len(response.Choices)-1].Message.Content,
		UserID: userId,
		ChatRoomId: chatRoomId,
		CreatedAt: time.Now(),
	}

	return gptMessage, nil
}