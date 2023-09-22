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
	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))

	err = request.Body.Close()
	if err != nil {
		return
	} //todo arquivo aberto no GO precisa ser fechado

}
