package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app_go/app/models"
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

//cookieを取得する関数
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
