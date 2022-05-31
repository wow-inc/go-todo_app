package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app_go/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		fmt.Println(file)
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	fmt.Println(files)

	templates := template.Must(template.ParseFiles(files...))
	fmt.Println(templates)
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
