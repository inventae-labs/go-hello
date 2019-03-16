package main

import "fmt"

func main() {
	fmt.Println("Escolha um numero do menu")
	fmt.Println("1-Monitorar")
	fmt.Println("2-Logs")
	fmt.Println("3-Sair")

	var commando int
	var nomeComando string

	fmt.Scan(&commando)

	// IFs
	if commando == 1 {
		nomeComando = "Monitorar"
	} else if commando == 2 {
		nomeComando = "Logs"
	} else if commando == 3 {
		nomeComando = "Sair"
	} else {
		nomeComando = "Nao reconhecido"
	}

	// Switchs (PS: o break não é obrigatorio, ele só executa uma opcao)
	switch commando {
	case 1:
		nomeComando = "Monitorar"
	case 2:
		nomeComando = "Logs"
	case 3:
		nomeComando = "Sair"
		break
	default:
		nomeComando = "Nao reconhecido"
	}

	fmt.Println("O comando escolhido foi", commando)
	fmt.Println("Comando", nomeComando)

}
