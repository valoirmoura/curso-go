package main

import "fmt"

func main() {

	salarios := map[string]int{"Valoir": 1000, "Edicleia": 2000}
	fmt.Println(salarios["Valoir"])

	//outra forma de criar map
	idades := make(map[string]int)
	idades["Valoir"] = 35

	//Outra forma de criar map
	instrumentos := map[string]string{}

	for _, instr := range instrumentos {
		fmt.Println(instr)
	}

	for _, idade := range idades {
		fmt.Println(idade)
	}
}
