package main

import (
	"bufio"
	"fmt"
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

		runCommand(line)
	}

}

func before() {
	createDirData()
}

func createDirData() {
	if !dirExists("data/") {
		createDir("data/")
	}
}
