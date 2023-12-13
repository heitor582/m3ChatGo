package dto

type ChatGptMessageDto struct {
	Model	string `json:"model" validate:"required,string"`
	MaxToken uint64 `json:"max_tokens" validate:"required,string"`
	Temperature float32 `json:"temperature" validate:"required"`
	Messages []ChatGptMessage `json:"messages" validate:"required"`
}

type ChatGptMessage struct {
	Role	string `json:"role" validate:"required,string"`
	Content string `json:"content" validate:"required,string"`
}