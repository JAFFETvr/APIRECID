package entity

// Mensaje representa el formato del mensaje que recibirá la API.
type Mensaje struct {
	ID       string `json:"id"`
	Contenido string `json:"contenido"`
}
