package handler

import (
	"log"
	"net/http"
	"text/template"

	"groupie-tracker/pkg"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusText(http.StatusNotFound), 404)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println("error")
		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}

	data := pkg.DisplayCards("https://groupietrackers.herokuapp.com/api/artists")
	if data == nil {

		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	err = tmp.Execute(w, data)
	if err != nil {
		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
}
