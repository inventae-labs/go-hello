package main

import "fmt"
import "os"

// Clousure
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	var comando int
	timer := counter()

	timer()
	fmt.Println("usando o endereco", &comando, "para gravar o comando")

	timer()
	intro()

	timer()
	menu()

	timer()
	lerComando(&comando)

	timer()
	commandName, _ := interpretaComando(&comando) // skip return

	timer()
	fmt.Println("Voce escolheu o comando", comando, commandName)

	timer()
	switch comando {
	case 1:
		timer()
		monitorar()
	case 2:
		timer()
		logs()
	case 3:
		timer()
		sair()
	default:
		os.Exit(-1) // Code Error = -1 - Error
	}

	fmt.Println("Ticks", timer())
}

func intro() {
	version := 1.2
	fmt.Println("Ola, amigos")
	fmt.Println("Versao", version)
}

func menu() {
	fmt.Println("Menu:")
	fmt.Println("1-Monitorar")
	fmt.Println("2-Logs")
	fmt.Println("3-Sair")
}

func lerComando(comando *int) {
	fmt.Scan(comando)
}

func interpretaComando(comando *int) (string, bool) {
	name := "Não reconhecido"
	reconheceu := true

	fmt.Println("Lendo o comando do endereco", comando)

	switch *comando {
	case 1:
		name = "Monitorar"
	case 2:
		name = "Logs"
	case 3:
		name = "Sair"
	default:
		reconheceu = false
	}

	return name, reconheceu
}

func monitorar() {
	fmt.Println("Monitorar")
}

func logs() {
	fmt.Println("Logs")
}

func sair() {
	fmt.Println("Saindo do programa...")
	os.Exit(0) // Code Error = 0 - Success
}
