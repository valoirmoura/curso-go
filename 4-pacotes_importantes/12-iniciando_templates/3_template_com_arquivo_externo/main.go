package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	t := template.Must(template.New("content.html").ParseFiles("content.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"GO", 40},
		{"Java", 100},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}

}
