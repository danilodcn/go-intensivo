package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i+1, name)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("A")
	canal := make(chan string)

	go func() {
		fmt.Println("Começou a Função anônima")
		time.Sleep(time.Millisecond * 200)
		task("B")
		canal <- "Mensagem que veio da funçao anônima"
		fmt.Println("Terminou a Função anônima")
	}()
	fmt.Println("Dentro da principal")
	msg := <-canal
	fmt.Println(msg)
}
