package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		}

		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"GO", 40},
			{"Java", 100},
			{"Python", 10},
		})
		if err != nil {
			panic(err)
		}

	})

	http.ListenAndServe(":8282", nil)

}
