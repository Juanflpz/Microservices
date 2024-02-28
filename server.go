package main

import (
	"encoding/json"
	"fmt"
	"html/template" //no
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserCredentials estructura para decodificar el JSON enviado en la solicitud POST
type UserCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/login", loginHandler) //al ingresar a la dirección / en el servidor ejecutará el index
	fmt.Println("El servidor está a la escucha en el servidor 80")
	http.ListenAndServe(":80", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //se valida si es una solicitud post
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el JSON del cuerpo de la solicitud
	var credentials UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar si se proporcionaron usuario y clave
	if credentials.User == "" || credentials.Password == "" {
		http.Error(w, "Usuario y clave son obligatorios", http.StatusBadRequest)
		return
	}

	// Generar el token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = credentials.User
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["iss"] = "ingesis.uniquindio.edu.co"

	// Firmar el token
	tokenString, err := token.SignedString([]byte("secret")) // Cambia "secret" por tu clave secreta
	if err != nil {
		http.Error(w, "Error al firmar el token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el token JWT
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, tokenString)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}
func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Página no encontrada")
	} else {
		template.Execute(w, nil)
	}
	fmt.Fprintf(w, "Hola mundo")
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Perform authentication logic here (e.g., check against a database).
		// For simplicity, we'll just check if the username and password are both "admin".
		if username == "admin" && password == "admin" {
			// Successful login, redirect to a welcome page.
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
			return
		}

		// Invalid credentials, show the login page with an error message.
		fmt.Fprintf(w, "Invalid credentials. Please try again.")
		return
	}

	// If not a POST request, serve the login page template.
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
