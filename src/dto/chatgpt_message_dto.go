package dto

type ChatGptMessageDto struct {
	Model	string `json:"model" validate:"required,string"`
	Temperature float32 `json:"temperature" validate:"required"`
	Messages []ChatGptMessage `json:"messages" validate:"required"`
}

type ResponseFormat struct{
	Type string `json:"type" validate:"required,string"`
}

type ChatGptMessage struct {
	Role	string `json:"role" validate:"required,string"`
	Content string `json:"content" validate:"required,string"`
}

func NewChatGptMessageDto(messages []ChatGptMessage) ChatGptMessageDto {
	return ChatGptMessageDto{
		Model: "gpt-3.5-turbo",
		Temperature: 1.0,
		Messages: messages,
	}
}