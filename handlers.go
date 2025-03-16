package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

// Signup Handler
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Hash the password before storing it
		hashedPassword, err := HashPassword(password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		// Insert user into the database
		_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		renderTemplate(w, "templates/signup.html")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Retrieve the stored hashed password
		var storedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		// Compare hashed passwords
		if CheckPasswordHash(password, storedPassword) {
			http.Redirect(w, r, "/welcome", http.StatusFound)
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	} else {
		renderTemplate(w, "templates/login.html")
	}
}

// Welcome page
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/welcome.html")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/create.html")
}

// Helper function to render templates
func renderTemplate(w http.ResponseWriter, filename string) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		log.Println("Template error:", err)
		return
	}
	tmpl.Execute(w, nil)
}
