package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gym-system/src/inventory/Mensaje/infraestructure/database"
	"gym-system/src/inventory/Mensaje/infraestructure/controllers"
	"gym-system/src/inventory/Mensaje/infraestructure/routes"
	"gym-system/src/inventory/Mensaje/domain/repository"
)

func main() {
	fmt.Println("Iniciando API...")

	// Crear el repositorio de RabbitMQ (implementación de la interfaz)
	rmq, err := database.NewRabbitMQ() // Usamos la implementación de RabbitMQ
	if err != nil {
		log.Fatal("No se pudo conectar a RabbitMQ:", err)
	}
	defer rmq.Close()

	// Crear el controlador de mensajes, inyectando el repositorio
	var rabbitRepo repository.RabbitMQRepository = rmq
	mensajeController := controllers.NewMensajeController(rabbitRepo)

	// Configurar rutas
	router := mux.NewRouter()
	routes.SetupRoutes(router, mensajeController)

	// Iniciar el servidor
	port := ":8081"
	fmt.Println("API escuchando en", port)
	log.Fatal(http.ListenAndServe(port, router))
}
