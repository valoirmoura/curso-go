package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second) //O contenxto será cancelado sempre que demorar mais de 1 SEGUNDO ou quando alguem executar a função cancel()
	//ctx, cancel := context.WithCancel(ctx)               //já este tipo apenas cancela o contexto quando for chamado a função cancel()
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
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
