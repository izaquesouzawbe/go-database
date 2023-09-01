package dir

import (
	"fmt"
	"log"
	"os"
)

func CreateDir(nameDir string) {
	err := os.Mkdir(nameDir, os.ModePerm)
	if err != nil {
		log.Println("Error creating directory: ", err)
	} else {
		log.Println("Directory '" + nameDir + "' created successfully")
	}
}

func DirExists(path string) bool {

	_, err := os.Stat(path)
	if err == nil {
		return true // A pasta existe
	} else if os.IsNotExist(err) {
		return false // A pasta não existe
	} else {
		fmt.Println("Ocorreu um erro ao verificar a pasta:", err)
		return false
	}
}

func ListDir(pasta string) ([]string, error) {

	var diretorios []string

	// Abre o diretório
	diretorio, err := os.Open(pasta)
	if err != nil {
		return nil, err
	}
	defer diretorio.Close()

	// Lê o conteúdo do diretório
	info, err := diretorio.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// Itera sobre os itens do diretório
	for _, item := range info {
		if item.IsDir() {
			// Adiciona o nome do diretório ao slice
			diretorios = append(diretorios, item.Name())
		}
	}

	return diretorios, nil
}
