package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CreateDir(nameDir string) {
	err := os.Mkdir(nameDir, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar a pasta:", err)
	} else {
		fmt.Println("Pasta criada com sucesso!")
	}
}

func DirExists(path string) bool {

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

func AppendLineToFile(filename string, line string) error {

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

func OpenFile(nameFile string) *os.File {

	file, err := os.Open(nameFile)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return nil
	}

	return file

}

func CloseFile(file *os.File) {
	defer file.Close()
}

func ReadLines(filename string) ([]string, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func ListFiles(diretorio string) ([]string, error) {

	var arquivos []string

	lista, err := ioutil.ReadDir(diretorio)
	if err != nil {
		return nil, err
	}

	for _, arquivoInfo := range lista {
		if !arquivoInfo.IsDir() {
			arquivos = append(arquivos, arquivoInfo.Name())
		}
	}

	return arquivos, nil
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
