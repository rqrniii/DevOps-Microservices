package service

import (
	"context"

	"github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/llm"
)

type AIService struct {
	llm *llm.Client
}

func NewAIService(llmClient *llm.Client) *AIService {
	return &AIService{llm: llmClient}
}

func (s *AIService) GenerateTask(ctx context.Context, userInput string) (string, error) {
	prompt := llm.TaskPrompt(userInput)
	return s.llm.Generate(ctx, prompt)
}
