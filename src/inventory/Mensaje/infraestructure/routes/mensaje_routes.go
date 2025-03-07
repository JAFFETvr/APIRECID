package routes

import (
	"github.com/gorilla/mux"
	"gym-system/src/inventory/Mensaje/infraestructure/controllers"
)

func SetupRoutes(router *mux.Router, mensajeController *controllers.MensajeController) {
	router.HandleFunc("/mensaje", mensajeController.RecibirMensaje).Methods("POST")
}
