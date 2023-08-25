package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func createDir(nameDir string) {
	err := os.Mkdir(nameDir, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar a pasta:", err)
	} else {
		fmt.Println("Pasta criada com sucesso!")
	}
}

func dirExists(path string) bool {

	_, err := os.Stat(path)
	if err == nil {
		return true // A pasta existe
	} else if os.IsNotExist(err) {
		return false // A pasta não existe
	} else {
		// Tratar outros erros
		fmt.Println("Ocorreu um erro ao verificar a pasta:", err)
		return false
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true // O arquivo existe
	} else if os.IsNotExist(err) {
		return false // O arquivo não existe
	} else {
		// Tratar outros erros
		fmt.Println("Ocorreu um erro ao verificar o arquivo:", err)
		return false
	}
}

func appendLineToFile(filename string, line string) error {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	if err != nil {
		return err
	}

	return nil
}

func createFile(nameFile string) *os.File {

	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return nil
	}

	defer file.Close()

	return file

}

func openFile(nameFile string) *os.File {

	file, err := os.Open(nameFile)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return nil
	}

	return file

}

func closeFile(file *os.File) {
	defer file.Close()
}

func readLines(filename string) ([]string, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}
