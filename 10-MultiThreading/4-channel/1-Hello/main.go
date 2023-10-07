package main

import "fmt"

func main() {

	canal := make(chan string)

	//Thread 2
	go func() {
		canal <- "Olá Mundo!!!"
	}()

	//Thread 1
	msg := <-canal
	fmt.Println(msg)

}

//Channel
//Fazer comunicação entre Threads
//Segurança para um Thread saber o momento em que ela pode trabalhar com um determinado dado
