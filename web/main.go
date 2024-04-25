package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func Oformlenie(w http.ResponseWriter, r *http.Request) {

	/* 	if r.Method != http.MethodPost {
	   		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
	   		return
	   	}

	   	err := r.ParseForm()
	   	if err != nil {
	   		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
	   		return
	   	} */

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Используйте параметризированные запросы, чтобы избежать SQL-инъекций
	insert, err := db.Query(fmt.Sprintf("INSERT INTO `request` (`name`, `phone`, `description`) VALUES ('%s','%s','%s')", title, anons, full_text))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/loading/", http.StatusSeeOther)

}

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("C:/Users/User/web/PAGE_.html"))

	// Выполняем шаблон и передаем его в ответ
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Loading(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("C:/Users/User/web/templess/index/loading.html"))

	// Выполняем шаблон и передаем его в ответ
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Contacts(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("C:/Users/User/web/templess/index/contacts.html"))

	// Выполняем шаблон и передаем его в ответ
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Zayavka(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("C:/Users/User/web/templess/index/index.html"))

	// Выполняем шаблон и передаем его в ответ
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func Info(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("C:/Users/User/web/templess/index/info.html"))

	// Выполняем шаблон и передаем его в ответ
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {

}

func HandleRequests() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", Contacts) // Добавлен маршрут без завершающего слэша
	http.HandleFunc("/zayavka/", Zayavka)
	http.HandleFunc("/oformlenie/", Oformlenie)
	http.HandleFunc("/info/", Info)
	http.HandleFunc("/loading/", Loading)
	err := http.ListenAndServe(":8040", nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func main() {

	HandleRequests()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/* insert, err := db.Query("INSERT INTO `request` (`name`, `phone`, `description`) VALUES ('Miha', '+87878888', 'dasfssfdsfsfd')")
	if err != nil {
		panic(err)
	}
	defer insert.Close()  */

	fmt.Println("Подключено к MySQL")

}
