package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func FileExists(filename string) bool {
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

func AppendLineToFile(filename string, lines []string) error {

	content := strings.Join(lines, "\n")

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Erro ao gravar no arquivo:", err)
		return nil
	}

	return nil
}

func CreateFile(nameFile string, returnValue bool) *os.File {

	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return nil
	}

	if !returnValue {
		defer file.Close()
	}

	return file

}

func ReadLines(filename string) ([]string, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func LoadFileJSON(caminho string, dados interface{}) error {
	arquivo, err := os.Open(caminho)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	byteValue, err := ioutil.ReadAll(arquivo)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, dados)
	if err != nil {
		return err
	}

	return nil
}
