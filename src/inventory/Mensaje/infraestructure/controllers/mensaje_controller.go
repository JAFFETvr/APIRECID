package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gym-system/src/inventory/Mensaje/application/useCases"
	"gym-system/src/inventory/Mensaje/domain/entity"
)

// RecibirMensaje maneja la solicitud POST para recibir mensajes desde el consumidor.
func RecibirMensaje(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var mensaje entity.Mensaje

	// Leer el cuerpo de la petición
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Convertir JSON a la estructura Mensaje
	if err := json.Unmarshal(body, &mensaje); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Procesar el mensaje
	if err := useCases.ProcesarMensaje(mensaje); err != nil {
		http.Error(w, "Error al procesar el mensaje", http.StatusInternalServerError)
		return
	}

	fmt.Println(" Mensaje recibido y procesado correctamente")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mensaje recibido correctamente"))
}
