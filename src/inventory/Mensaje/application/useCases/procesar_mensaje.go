package useCases

import (
	"fmt"
	"gym-system/src/inventory/Mensaje/domain/entity"
)

// ProcesarMensaje recibe el mensaje y lo maneja.
func ProcesarMensaje(mensaje entity.Mensaje) error {
	// Aquí iría la lógica de procesamiento, por ejemplo, guardarlo en la BD.
	fmt.Printf(" Mensaje recibido: %+v\n", mensaje)
	
	return nil
}
