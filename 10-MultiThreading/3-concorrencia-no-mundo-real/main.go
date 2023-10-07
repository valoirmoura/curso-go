package main

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Se utilizarmos o atomic ele bloqueia o espaço da memória para não ser alterado por outra Thread... muito mas muito util....
		atomic.AddUint64(&number, 1)
		w.Write([]byte("Você teve acesso a essa página " + strconv.FormatUint(number, 10) + " vezes!\n"))
	})

	http.ListenAndServe(":3000", nil)
}

//go run -race main.go //ajuda a verificar problemas de concerrencias

//Em concorrência podemos ter um problema quanto acessamos a mesma variável com multiplas Threads...
//Para isso utilizamos um Lock e UnLock, o que é isso... ele bloqueia a memoria durante o uso daquele espaço não deixando outra Thread altera-lo... por exemplo

//m := sync.Mutex{}
//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//	m.Lock()
//	number++
//	m.Unlock()
//	w.Write([]byte("Você teve acesso a essa página " + strconv.FormatUint(number, 10) + " vezes!\n"))
//})
