package llm

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
	model  string
}

// NewClient initializes the OpenAI client with API key and model
func NewClient() *Client {
	return &Client{
		client: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
		model:  os.Getenv("OPENAI_MODEL"),
	}
}

// Generate sends a prompt to OpenAI and returns the generated text
func (c *Client) Generate(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: c.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.7,
			MaxTokens:   150,
		},
	)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", nil
	}

	return resp.Choices[0].Message.Content, nil
}
