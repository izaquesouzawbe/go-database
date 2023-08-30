package reading

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to go-database...")
	fmt.Println("")

	for {

		fmt.Print(":~# ")

		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}

		//commands.RunCommand(line)
	}
}
