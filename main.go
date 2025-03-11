package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize database connection
func init() {

	env := godotenv.Load(".env")
	if env != nil {
		log.Fatalf("Error laoding .env fileS")
	}

	sqlps := os.Getenv("SQL_CRED")

	var err error
	// Replace with your MySQL connection details
	db, err = sql.Open("mysql", sqlps)
	if err != nil {
		log.Fatal(err)
	}
}

// HashPassword hashes a given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a plaintext password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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
		tmpl, err := template.New("signup").Parse(`
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Signup</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Arial', sans-serif;
        }

        body {
            background-color: #58CC02; /* Duolingo green */
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
        }

        .container {
            background: white;
            padding: 2rem;
            border-radius: 15px;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
            text-align: center;
            max-width: 350px;
            width: 100%;
        }

        h2 {
            color: #333;
            font-size: 24px;
            margin-bottom: 1rem;
        }

        form {
            display: flex;
            flex-direction: column;
        }

        label {
            text-align: left;
            font-size: 14px;
            font-weight: bold;
            margin-bottom: 5px;
            color: #333;
        }

        input {
            width: 100%;
            padding: 10px;
            margin-bottom: 15px;
            border: 2px solid #ddd;
            border-radius: 8px;
            font-size: 16px;
        }

        input:focus {
            border-color: #58CC02;
            outline: none;
        }

        .btn {
            background-color: #58CC02;
            color: white;
            border: none;
            padding: 12px;
            font-size: 16px;
            border-radius: 8px;
            cursor: pointer;
            transition: 0.3s;
        }

        .btn:hover {
            background-color: #46a800;
        }

        a {
            display: block;
            margin-top: 10px;
            text-decoration: none;
            color: #58CC02;
            font-weight: bold;
        }

        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>

    <div class="container">
        <h2>Signup</h2>
        <form method="POST">
            <label for="username">Username:</label>
            <input type="text" name="username" required>

            <label for="password">Password:</label>
            <input type="password" name="password" required>

            <input type="submit" value="Signup" class="btn">
        </form>
        <a href="/login">Already have an account? Login here</a>
    </div>

</body>
</html>

		`)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	}
}

// Login Handler
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
		tmpl, err := template.New("login").Parse(`
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Login</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Arial', sans-serif;
        }

        body {
            background-color: #58CC02; /* Duolingo green */
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
        }

        .container {
            background: white;
            padding: 2rem;
            border-radius: 15px;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
            text-align: center;
            max-width: 350px;
            width: 100%;
        }

        h2 {
            color: #333;
            font-size: 24px;
            margin-bottom: 1rem;
        }

        form {
            display: flex;
            flex-direction: column;
        }

        label {
            text-align: left;
            font-size: 14px;
            font-weight: bold;
            margin-bottom: 5px;
            color: #333;
        }

        input {
            width: 100%;
            padding: 10px;
            margin-bottom: 15px;
            border: 2px solid #ddd;
            border-radius: 8px;
            font-size: 16px;
        }

        input:focus {
            border-color: #58CC02;
            outline: none;
        }

        .btn {
            background-color: #58CC02;
            color: white;
            border: none;
            padding: 12px;
            font-size: 16px;
            border-radius: 8px;
            cursor: pointer;
            transition: 0.3s;
        }

        .btn:hover {
            background-color: #46a800;
        }

        a {
            display: block;
            margin-top: 10px;
            text-decoration: none;
            color: #58CC02;
            font-weight: bold;
        }

        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>

    <div class="container">
        <h2>Login</h2>
        <form method="POST">
            <label for="username">Username:</label>
            <input type="text" name="username" required>

            <label for="password">Password:</label>
            <input type="password" name="password" required>

            <input type="submit" value="Login" class="btn">
        </form>
        <a href="/signup">Don't have an account? Signup here</a>
    </div>

</body>
</html>
		`)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	}
}

// Welcome page after successful login
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the dashboard!</h1>")
}

func main() {
	// Handle routes
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	// Start the server
	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
