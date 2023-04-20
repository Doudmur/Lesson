package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type server struct {
	db *sql.DB
}

func dbConnect() server {
	db, err := sql.Open("sqlite3", "db.sql")
	fmt.Println("Opening database")
	if err != nil {
		log.Fatal(err)
	}

	s := server{db: db}

	return s
}

func indexPage(w http.ResponseWriter, r *http.Request) { // Главная страница
	outputPage(w, "templates/index.html", nil)
}

func watchPage(w http.ResponseWriter, r *http.Request) { // Страницы просмотра
	outputPage(w, "templates/watch.html", nil)
}

func (s *server) registerPage(w http.ResponseWriter, r *http.Request) { // Страница регистрации
	if r.Method == "POST" {
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		username := r.FormValue("username")
		pass := r.FormValue("pass")

		_, err := s.db.Exec("insert into users(username, full_name, last_name, password, role) values ($1, $2, $3, $4, 'user')", username, fname, lname, pass)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Hello")
		text := username + " user was created successfully!"

		data := map[string]interface{}{"text": text}
		outputPage(w, "templates/register.html", data)

	}
	outputPage(w, "templates/register.html", nil)
}

func buyPage(w http.ResponseWriter, r *http.Request) { //Страница с формой

	data := make(map[string]interface{})

	if r.Method == "POST" { // Если была отправлена форма

		if err := r.ParseForm(); err != nil { // Если есть ошибка, отправляем ее
			data = map[string]interface{}{"error": err}
			outputPage(w, "templates/buy.html", data)
			return
		}

		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		// Отправляем инфу если нет ошибок
		data = map[string]interface{}{"firstName": firstName, "lastName": lastName}
	}

	outputPage(w, "templates/buy.html", data)
}

func outputPage(w http.ResponseWriter, htmlFile string, data interface{}) { // Функция для вывода страниц
	t, _ := template.ParseFiles(htmlFile)
	t.Execute(w, data)
}

func main() {
	s := dbConnect()
	defer s.db.Close()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/watch", watchPage)
	http.HandleFunc("/buy", buyPage)
	http.HandleFunc("/register", s.registerPage)

	http.ListenAndServe(":8080", nil)
}
