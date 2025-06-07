// services/ai_service.go
package services

import (
	"ai-assistant/repositories"
	"ai-assistant/utils"
	"encoding/json"
	"errors"
)

type AIService struct {
	groqClient    *utils.GroqClient
	systemRepo    *repositories.SystemRepository
	memoryRepo    *repositories.MemoryRepository
}

func NewAIService(groqClient *utils.GroqClient, systemRepo *repositories.SystemRepository, memoryRepo *repositories.MemoryRepository) *AIService {
	return &AIService{
		groqClient:    groqClient,
		systemRepo:    systemRepo,
		memoryRepo:    memoryRepo,
	}
}

func (s *AIService) ProcessMessage(sessionID string, message string) (string, error) {
	// 1. Ambil konteks memori dari MongoDB
	memories, _ := s.memoryRepo.GetRecentMemories(sessionID, 5)
	
	// 2. Ambil system prompt dari MariaDB
	systemPrompt, _ := s.systemRepo.GetSetting("ai_system_prompt")
	
	// 3. Bangun messages untuk Groq
	messages := []utils.GroqMessage{
		{Role: "system", Content: systemPrompt},
	}
	
	for _, m := range memories {
		messages = append(messages, utils.GroqMessage{
			Role:    "assistant",
			Content: m.Content,
		})
	}
	
	messages = append(messages, utils.GroqMessage{
		Role:    "user",
		Content: message,
	})
	
	// 4. Panggil Groq API
	response, err := s.groqClient.ChatCompletion(messages)
	if err != nil {
		return "", err
	}
	
	// 5. Cek function calling
	if funcCall := tryParseFunctionCall(response); funcCall != nil {
		result, err := s.ExecuteFunction(funcCall)
		if err != nil {
			return "", err
		}
		
		// Tambahkan hasil fungsi ke messages
		messages = append(messages, utils.GroqMessage{
			Role:    "function",
			Name:    funcCall.Function,
			Content: result,
		})
		
		// Panggil ulang API dengan konteks lengkap
		response, err = s.groqClient.ChatCompletion(messages)
		if err != nil {
			return "", err
		}
	}
	
	// 6. Simpan memori baru
	s.memoryRepo.StoreMemory(sessionID, response)
	
	return response, nil
}

func tryParseFunctionCall(content string) *utils.FunctionCall {
	var fc utils.FunctionCall
	if err := json.Unmarshal([]byte(content), &fc); err == nil {
		if fc.Function != "" && fc.Parameters != nil {
			return &fc
		}
	}
	return nil
}

func (s *AIService) ExecuteFunction(fc *utils.FunctionCall) (string, error) {
	switch fc.Function {
	case "get_weather":
		return s.getWeather(fc.Parameters.City)
	// Tambahkan fungsi lain di sini
	default:
		return "", errors.New("function not implemented")
	}
}

func (s *AIService) getWeather(city string) (string, error) {
	// Implementasi sebenarnya
	weather := map[string]string{
		"temperature": "30°C",
		"condition":   "Sunny",
		"city":        city,
	}
	return json.Marshal(weather)
}