// utils/groq_client.go
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

type FunctionCall struct {
	Function string                 `json:"function"`
	Parameters map[string]interface{} `json:"parameters"`
}

type GroqRequest struct {
	Model    string       `json:"model"`
	Messages []GroqMessage `json:"messages"`
}

type GroqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type GroqClient struct {
	APIKey string
	Model  string
}

func NewGroqClient(apiKey, model string) *GroqClient {
	return &GroqClient{
		APIKey: apiKey,
		Model:  model,
	}
}

func (c *GroqClient) ChatCompletion(messages []GroqMessage) (string, error) {
	requestBody := GroqRequest{
		Model:    c.Model,
		Messages: messages,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response GroqResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from Groq API")
	}

	return response.Choices[0].Message.Content, nil
}