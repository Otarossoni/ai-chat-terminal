package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Otarossoni/ai-chat-terminal/config"
	"github.com/Otarossoni/ai-chat-terminal/service"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
)

func StartChat() {
	config.Load()

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)
	chatName := "Nora"
	creatorName := "Otarossoni"

	chatGpt := service.NewChatGPT(config.Get().OpenAiKey, openai.GPT4o, chatName, creatorName)

	myFigure := figure.NewFigure(chatName, "doom", true)
	myFigure.Print()

	fmt.Println(color.HiBlackString("\nBem-vindo ao chat " + chatName + "! Digite 'exit' para encerrar."))

	for {
		fmt.Print(color.GreenString("\nUser> "))

		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if cmd == "exit" {
			aiResponse, errAiResponse := chatGpt.SendMessage(ctx, "Estou de saída, até mais!")
			if errAiResponse != nil {
				fmt.Println("error: ", errAiResponse)
			}

			fmt.Println(color.BlueString(chatName+">"), aiResponse)
			break
		}

		aiResponse, errAiResponse := chatGpt.SendMessage(ctx, cmd)
		if errAiResponse != nil {
			fmt.Println("error: ", errAiResponse)
		}

		fmt.Println(color.BlueString(chatName+">"), aiResponse)
	}
}
