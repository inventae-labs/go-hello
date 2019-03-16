package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direto")

	go f("gorutine")

	// Closure
	go func(msg string) {
		fmt.Println(msg)
	}("indo")

	var input string
	fmt.Scan(&input)
	fmt.Println("pronto")

	// Canais transmitir string
	mensagens := make(chan string)

	go func() { mensagens <- "ping" }()

	// A sintaxe <-canal recebe um valor de um canal. Aqui nós vamos
	// receber a mensagem "ping" que enviamos acima e imprimi-lo.
	msg := <-mensagens
	fmt.Println(msg)

	// Buffer e Sincronização de canal
	msgBuffer := make(chan string, 2)

	go func() { msgBuffer <- "buferizado" }()
	go func() { msgBuffer <- "canal" }()

	for i := 0; i < 2; i++ {
		// Bloqueia até executar uma vez
		fmt.Println(<-msgBuffer)
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "um"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "dois"
	}()

	// Select
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recebido", msg1)
		case msg2 := <-c2:
			fmt.Println("recebido", msg2)
		case <-time.After(time.Second * 5):
			fmt.Printf("Saiu!")
		}
	}

	// Direções de Canais
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "mensagem passada")
	pong(pings, pongs)

	fmt.Println("aguardando receber msg", <-pongs)

	// Timeouts
	// http://goporexemplo.golangbr.org/timeouts.html
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
