package dto

type ChatGPTResponse struct {
    ID               string   `json:"id"`
    Object           string   `json:"object"`
    Created          int64   `json:"created"`
    Model            string   `json:"model"`
    SystemFingerprint string  `json:"system_fingerprint"`
    Choices          []Choice `json:"choices"`
    Usage            Usage    `json:"usage"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type Choice struct {
    Index         int64   `json:"index"`
    Message       Message `json:"message"`
    FinishReason  string  `json:"finish_reason"`
}

type Usage struct {
    PromptTokens       int64 `json:"prompt_tokens"`
    CompletionTokens   int64 `json:"completion_tokens"`
    TotalTokens        int64 `json:"total_tokens"`
}