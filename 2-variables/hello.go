package main

import "fmt"
import "reflect"

func main() {
	var nome string = "Fabio"
	var esposa = "Priscila"
	filho := "Pedro"

	var version float32 = 1.1
	altura := 1.71
	var idade int16 = 30

	fmt.Println("Ola Sr.", nome, "sua idade é", idade)
	fmt.Println("Casado com", esposa, "tem um filho chamado", filho)
	fmt.Println("Altura:", altura, "(", reflect.TypeOf(altura), ")")

	fmt.Println("Versao do programa:", version)
}
