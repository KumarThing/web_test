package main

import (
	"html/template"
	"net/http"
	"fmt"
	"log"
)


var tmpl = template.Must(template.ParseFiles("template/index.html"))

func main() {
	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("Visited your website.")

			data := struct{
				Name string
			} {}
		if r.Method == http.MethodPost {
			r.ParseForm()
			data.Name = r.FormValue("username")
		}

		tmpl.Execute(w, data)
		
	})

	fmt.Println("The server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}