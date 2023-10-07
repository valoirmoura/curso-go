package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}

//WaitGroup
// 1- Adicionar qtde de tarefas/operações
// 2- Informar que você terminou uma operação
// 3- Esperar até que as operações sejam finalizadas
