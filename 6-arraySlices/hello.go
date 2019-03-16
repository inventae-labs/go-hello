package main

import (
	"fmt"
)

func main() {
	// Array
	var nomes [4]string
	nomes[0] = "Fabio"
	nomes[1] = "Priscila"
	nomes[2] = "Pedro"
	nomes[3] = "Joe"

	fmt.Println("Array", nomes)
	fmt.Println("tamanho:", len(nomes))
	fmt.Println("capactidade:", cap(nomes)) //Capacidade Fixa

	// Slices
	var tecnologias = []string{"NodeJS", "Python", "GoLang"}

	fmt.Println("Array", tecnologias)
	fmt.Println("tamanho:", len(tecnologias))
	fmt.Println("capactidade:", cap(tecnologias)) // Iniciamos com 3 itens define cap com 3

	// For normal
	for i := 0; i < len(tecnologias); i++ {
		fmt.Println(tecnologias[i])
	}

	tecnologias = append(tecnologias, "Java")
	fmt.Println("Array", tecnologias)
	fmt.Println("tamanho:", len(tecnologias))
	fmt.Println("capactidade:", cap(tecnologias)) // Go Dobra a capacitade

	// For usando o range
	for i, tecnologia := range tecnologias {
		fmt.Println("Posicao", i, "Tec", tecnologia)
	}
}
