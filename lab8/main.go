package main

import (
	"html/template"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) { // Главная страница
	outputPage(w, "templates/index.html", nil)
}

func watchPage(w http.ResponseWriter, r *http.Request) { // Страницы просмотра
	outputPage(w, "templates/watch.html", nil)
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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/watch", watchPage)
	http.HandleFunc("/buy", buyPage)

	http.ListenAndServe(":8000", nil)
}
