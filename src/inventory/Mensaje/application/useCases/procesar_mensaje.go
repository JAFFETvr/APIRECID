package useCases

import (
	"encoding/json"
	"fmt"
	"gym-system/src/inventory/Mensaje/domain/entity"
	"gym-system/src/inventory/Mensaje/domain/repository"
)

type ProcesarMensajeUseCase struct {
	rabbitRepo repository.RabbitMQRepository
}

func NewProcesarMensajeUseCase(rabbitRepo repository.RabbitMQRepository) *ProcesarMensajeUseCase {
	return &ProcesarMensajeUseCase{rabbitRepo: rabbitRepo}
}

func (pm *ProcesarMensajeUseCase) Execute(mensaje entity.Mensaje) error {
	fmt.Printf("Mensaje recibido en useCase: %+v\n", mensaje)

	if mensaje.ID == "" || mensaje.Contenido == "" {
		return fmt.Errorf("Mensaje inválido: ID o Contenido vacío")
	}

	mensajeJSON, err := json.Marshal(mensaje)
	if err != nil {
		return fmt.Errorf("Error al convertir mensaje a JSON: %s", err)
	}

	err = pm.rabbitRepo.PublishMessage(string(mensajeJSON))
	if err != nil {
		return fmt.Errorf("Error al enviar mensaje a la cola: %s", err)
	}

	fmt.Println("Mensaje enviado a RabbitMQ correctamente:", string(mensajeJSON))
	return nil
}
