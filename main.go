package main

import (
	"fmt"
	"log"
	"net/http"

	"gym-system/src/inventory/Mensaje/infraestructure/controllers"
)

func main() {
	fmt.Println(" Iniciando API...")

	// Definir rutas usando el paquete estándar de Go
	http.HandleFunc("/mensaje", controllers.RecibirMensaje)

	// Iniciar el servidor en el puerto 8080
	port := ":8081"
	fmt.Println(" API escuchando en", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
