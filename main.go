package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)


var tmpl = template.Must(template.ParseFiles("template/index.html"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("Visited your website.")

			data := struct{
				Name string
				Result string
			} {}
		if r.Method == http.MethodPost {
			r.ParseForm()
			data.Name = r.FormValue("username")

			num1Str := r.FormValue("num1")
			num2Str := r.FormValue("num2")
			op := r.FormValue("operation")

			if num1Str != "" && num2Str != "" && op != "" {
				num1, err1 := strconv.ParseFloat(num1Str, 64)
				num2, err2 := strconv.ParseFloat(num2Str, 64)
				if err1 == nil && err2 == nil {
					op = strings.TrimSpace(op)
					switch op {
					case "+" :
						data.Result = fmt.Sprintf("%f", num1 + num2)

					case "-" :
						data.Result = fmt.Sprintf("%f", num1- num2)

					case "*":
						data.Result = fmt.Sprintf("%f", num1 * num2)

					case "/":
						if  num2 == 0 {
							data.Result = "Error cannot be divided"
						} else {
							data.Result = fmt.Sprintf("%f", num1/num2)
						}

					default :
						data.Result = "Error."
					} 
				

				} else {
					data.Result = "Error: Invalid operation."
				}
			}
		}


		tmpl.Execute(w, data)
		
	})

	fmt.Println("The server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}