package controllers

import "net/http"

func AboutController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
