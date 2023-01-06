package handler

import (
	"net/http"
	"text/template"
)

type ErrorBody struct {
	Message string
	Code    int
}

func ErrorHandler(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	// fmt.Println(code)
	errorBody := ErrorBody{Message: message, Code: code}
	html, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err = html.Execute(w, errorBody); err != nil {
		// fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
