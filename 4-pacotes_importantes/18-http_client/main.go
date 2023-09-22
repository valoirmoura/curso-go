package main

import (
	"io"
	"net/http"
	"time"
)

func main() {

	//Timeout: tempo para a resposta, nesse caso é de 1 segundo para responder, se não responder..... errooo
	c := http.Client{Timeout: time.Duration(1) * time.Microsecond}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

}
