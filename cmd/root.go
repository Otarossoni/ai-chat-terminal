package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartChat() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bem-vindo ao Chat! Digite 'exit' para encerrar.")

	for {
		fmt.Print("Chat> ")

		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if cmd == "exit" {
			fmt.Println("Encerrando Chat...")
			break
		} else {
			fmt.Println(cmd)
		}
	}
}
