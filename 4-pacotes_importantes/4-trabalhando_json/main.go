package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {

	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	fmt.Println(string(res))

	//O Encoder não armazena o valor do JSON, se quisermos armazenar utilizamos o Marshal, o Encoder já entrega o valor direto, pode ser para um Webserver, para um arquivo... enfim...
	json.NewEncoder(os.Stdout).Encode(conta)

	jsonPuro := []byte(`{"Numero": 2, "Saldo": 200}`)
	var contaX Conta

	//Aqui pegamos o valor do JSON e convertemos ele para a Struct, no entanto os valores devem ser identicos ao da Structs, os campos ... os nomes... O Binding deve acontecer com a Struct, ou seja, se o Json tiver um campo diferente quebrará
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)

}
