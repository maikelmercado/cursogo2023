package internal

import (
	"log"
	"net/http"

	"github.com/maikelmercado/cursogo2023/internal/controllers"
)

func Server() {

	// Controllers
	http.HandleFunc("/", controllers.MainController)
	http.HandleFunc("/about", controllers.AboutController)

	// Start Server
	println("Servidor iniciado en el puerto 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor , %e", err)
	}

}
