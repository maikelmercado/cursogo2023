package controllers

import "net/http"

func MainController(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hola Maik"))

}
