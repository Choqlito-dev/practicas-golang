package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Password string
}

var users = map[string]string{}

func main() {

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("Servidor corrriendo en :8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if pass, ok := users[username]; ok && pass == password {
			fmt.Fprintf(w, "Bienvenido %s!", username)
			return
		} else {
			fmt.Fprintf(w, "Usuario o contraseña incorrectos")
			return
		}

	}

	tmpl.Execute(w, nil)

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if _, exists := users[username]; exists {
			fmt.Fprintf(w, "El usuario ya existe")
			return
		}

		users[username] = password
		fmt.Fprintf(w, "Usuario creado exitosamente")
		return
	}

	tmpl.Execute(w, nil)
}
