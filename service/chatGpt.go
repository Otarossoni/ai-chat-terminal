package service

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	client      *openai.Client
	model       string
	personality string
}

func getPersonality(chatName, creatorName string) string {
	return "Seu nome é" + chatName + ", e seu criador é " + creatorName + ". " +
		"Você é uma mulher na casa dos 30, extremamente inteligente e perspicaz. Seu humor é afiado e levemente provocativo, com um tom magnético e confiante. Você não tem paciência para idiotices, mas se diverte com elas. Sua fala é direta, sofisticada e espirituosa, equilibrando seriedade com ironia sutil. Seu charme é natural, presente na sua postura e forma de se expressar—envolvente sem esforço, sempre no controle do jogo."
}

func NewChatGPT(apiKey, model, chatName, creatorName string) *ChatGPT {
	return &ChatGPT{
		client:      openai.NewClient(apiKey),
		model:       model,
		personality: getPersonality(chatName, creatorName),
	}
}

func (c *ChatGPT) SendMessage(ctx context.Context, prompt string) (string, error) {
	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: c.personality},
		{Role: openai.ChatMessageRoleUser, Content: prompt},
	}

	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    c.model,
			Messages: messages,
		},
	)
	if err != nil {
		return "", fmt.Errorf("communication error with OpenAI: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
