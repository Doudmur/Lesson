package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	CreatedAt string
}

type server struct {
	db *sql.DB
}

func database() server {
	database, _ := sql.Open("sqlite3", "db.db")
	server := server{db: database}
	return server
}

func (s *server) create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fn := r.FormValue("fn")
		ln := r.FormValue("ln")
		em := r.FormValue("email")
		createdAt := time.Now()

		if _, err := s.db.Exec("INSERT INTO users(firstName, lastName, email, createdAt) VALUES ($1, $2, $3, $4);", fn, ln, em, createdAt); err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	t, err := template.ParseFiles("templates/create.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func (s *server) delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if _, err := s.db.Exec("DELETE FROM users WHERE id=$1;", id); err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) update(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("fn") == "" {
		id := r.FormValue("id")
		data := map[string]interface{}{"id": id}
		t, _ := template.ParseFiles("templates/update.html")
		t.Execute(w, data)
		return
	}
	id := r.FormValue("id")
	fn := r.FormValue("fn")
	ln := r.FormValue("ln")
	email := r.FormValue("email")
	if _, err := s.db.Exec("UPDATE users SET firstName=$1, lastName=$2, email=$3 WHERE id=$4;", fn, ln, email, id); err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	res, err := s.db.Query("SELECT * FROM users;") // get all users
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	for res.Next() { // put every user to array
		var user User
		err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err = res.Err(); err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, users)
	if err != nil {
		return
	}
}

func main() {
	s := database()
	defer s.db.Close()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", s.index)
	http.HandleFunc("/create", s.create)
	http.HandleFunc("/delete", s.delete)
	http.HandleFunc("/update", s.update)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
