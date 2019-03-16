package main

import "fmt"

func main() {
	const version float32 = 1.0

	fmt.Println("Este programa esta na versao:", version)
	fmt.Println("1 - Iniciar")
	fmt.Println("2 - Logs")
	fmt.Println("3 - Sair")

	var command int
	fmt.Scan(&command) //ponteiro esta esperando um receber um int transforma para ->
	// fmt.Scanf("%d", &command) // formatado para um tipo int, com ponteiro
	fmt.Println("O comando escolhido foi", command, "ponteiro", &command)
}
