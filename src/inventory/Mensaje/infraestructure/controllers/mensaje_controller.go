package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gym-system/src/inventory/Mensaje/application/useCases"
	"gym-system/src/inventory/Mensaje/domain/entity"
	"gym-system/src/inventory/Mensaje/domain/repository"
)

// DTO que permite recibir tanto "message" como "contenido"
type mensajeDTO struct {
	Message   string `json:"message"`
	Contenido string `json:"contenido"`
}

type MensajeController struct {
	useCase *useCases.ProcesarMensajeUseCase
}

func NewMensajeController(rabbitRepo repository.RabbitMQRepository) *MensajeController {
	return &MensajeController{
		useCase: useCases.NewProcesarMensajeUseCase(rabbitRepo),
	}
}

func (c *MensajeController) RecibirMensaje(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("JSON recibido:", string(body))

	// Deserializamos en un DTO que admita ambas claves
	var dto mensajeDTO
	if err := json.Unmarshal(body, &dto); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Mapear el DTO a la entidad de dominio
	// Si "contenido" está vacío, usamos "message"
	contenido := dto.Contenido
	if contenido == "" {
		contenido = dto.Message
	}

	// Asignar un ID por defecto o generar uno (aquí usamos "default-id")
	mensaje := entity.Mensaje{
		ID:        "default-id",
		Contenido: contenido,
	}

	fmt.Println("Mensaje después de Unmarshal y mapeo:", mensaje)

	if err := c.useCase.Execute(mensaje); err != nil {
		http.Error(w, "Error al procesar el mensaje", http.StatusInternalServerError)
		return
	}

	fmt.Println("Mensaje recibido y procesado correctamente")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mensaje recibido correctamente"))
}
