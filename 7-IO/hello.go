package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	comando := menu()

	switch comando {
	case 1:
		lerArquivo()
	case 2:
		escreverArquivo("Hello", "Word")
	case 3:
		os.Exit(0) // Exit with success
	default:
		os.Exit(-1) // Exit with error
	}
}

func menu() int {
	fmt.Println("Escolha uma opcao")
	fmt.Println("1-Ler Arquivo")
	fmt.Println("2-Escrever Arquivo")
	fmt.Println("3-Sair")

	var comando int
	fmt.Scan(&comando)

	if comando > 3 || comando < 1 {
		fmt.Println("Comando inválido...")
		return menu()
	}

	return comando
}

func lerArquivo() {
	fmt.Println("Digite o local do arquivo:")

	var fileName string
	fmt.Scan(&fileName)

	if file, err := os.Open(fileName); err != nil {
		fmt.Println("Erro ao ler o arquivo", err)
	} else {
		var lines []string
		cursor := bufio.NewReader(file)
		for {
			line, err := cursor.ReadString('\n')
			line = strings.TrimSpace(line)
			lines = append(lines, line)

			if err == io.EOF {
				break
			}
		}
		file.Close()
		fmt.Println("Conteudo", lines)
	}

}

func escreverArquivo(nums ...string) {
	fmt.Println("Digite o local do arquivo:")

	var fileName string
	fmt.Scan(&fileName)

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo", err)
		return
	}

	toggle := false
	for i := 1; i <= 100; i++ {
		toggle = !toggle

		// https://golang.org/src/time/format.go
		line := time.Now().Format("02/01/2016 15:04:05") + " - i: " + strconv.Itoa(i) + " toggle: " + strconv.FormatBool(toggle) + " - "

		for _, item := range nums {
			line += item
		}

		line += "\n"

		file.WriteString(line)
	}

	file.Close()
}
