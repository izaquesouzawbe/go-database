package main

import (
	"bufio"
	"fmt"
	"go-database/commands"
	"os"
)

func main() {

	before()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to go-database...")
	fmt.Println("")

	for {

		fmt.Print(":~# ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}

		commands.RunCommand(line)
	}

}
