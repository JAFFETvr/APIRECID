package routes

import (
   "github.com/gorilla/mux"
"gym-system/src/inventory/Mensaje/infraestructure/controllers"
)

// SetupRoutes configura las rutas de la API.
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/mensaje", controllers.RecibirMensaje).Methods("POST")
}
