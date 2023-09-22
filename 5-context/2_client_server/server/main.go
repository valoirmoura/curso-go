package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		//Log imprime no command line (stdout)
		log.Println("Request processada com sucesso")
		//imprime no Browse
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo client")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)

	}
}
