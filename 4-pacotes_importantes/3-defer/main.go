package main

import (
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	//O Defer será executado no final, depois de que tudo for executado, ele serve para garantir que não esqueçamos de fechar....
	defer request.Body.Close()

	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))

}
