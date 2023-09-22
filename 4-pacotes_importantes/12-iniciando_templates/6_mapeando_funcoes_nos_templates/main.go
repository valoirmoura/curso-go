package main

import (
	"html/template" //sempre usar o html/template quando for expor os dados em html pois ele já se blinda melhor com segurança do que o text/template
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		}

		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))
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
