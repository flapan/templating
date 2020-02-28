package main

import (
	"net/http"

	"github.com/flapan/templating/views"
)

var index *views.View
var contact *views.View

func main() {
	index = views.NewView("bootstrap", "views/index.html")
	contact = views.NewView("bootstrap", "views/contact.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Render(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	contact.Render(w, nil)
}

/* type User struct {
	ID    int
	Email string
}

type ViewUserData struct {
	User User
}

type Widget struct {
	Name  string
	Price int
}

type ViewData struct {
	Name    string
	Widgets []Widget
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/functions", functionHandler)
	http.HandleFunc("/range", rangeHandler)
	http.ListenAndServe(":3000", nil)
}

func functionHandler(w http.ResponseWriter, r *http.Request) {
	testTemplate, err := template.New("functions.html").Funcs(template.FuncMap{
		"hasPermission": func(user User, feature string) bool {
			if user.ID == 1 && feature == "feature-b" {
				return true
			}
			return false
		},
	}).ParseFiles("functions.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html")

	err = testTemplate.Execute(w, ViewUserData{
		User: User{
			ID:    1,
			Email: "someone@gsnail.com",
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rangeHandler(w http.ResponseWriter, r *http.Request) {
	testTemplate, err := template.ParseFiles("range.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html")

	err = testTemplate.Execute(w, ViewData{
		Name: "John Smith",
		Widgets: []Widget{
			Widget{"BlueWidget", 12},
			Widget{"Red Widget", 12},
			Widget{"Green Widget", 12},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	testTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html")

	err = testTemplate.Execute(w, ViewData{
		Name: "John Smith",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/
