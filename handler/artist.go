package handler

import (
	"net/http"
	"strconv"
	"text/template"

	"groupie-tracker/pkg"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./templates/artist.html")
	if err != nil {

		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 || id > 52 {
		ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	idshka := pkg.TakeArtist(id, "https://groupietrackers.herokuapp.com/api/artists")
	concert := pkg.DatesConcert(id)
	idshka.Relations = concert
	tmp.Execute(w, idshka)
}
