package ws

import (
	"encoding/json"
	"log"
)

type Hub struct {
	Clients map[string]*Client
	Register chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:	make(map[string]*Client),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if oldClient, exists := h.Clients[client.DeviceID]; exists {
				oldClient.Conn.Close()
			}
			h.Clients[client.DeviceID] = client
			log.Printf("Pantalla conectada: %s", client.DeviceID)

		case client := <-h.Unregister:
			if _, exists := h.Clients[client.DeviceID]; exists {
				delete(h.Clients, client.DeviceID)
				close(client.Send)
				log.Printf("Pantalla desconectada: %s", client.DeviceID)
			}
		}
	}
}

func (h *Hub) SendToDevice(deviceID string, data interface{}) bool {
	client, exists := h.Clients[deviceID]
	if !exists {
		log.Printf("Advertencia: Intento de envío a dispositivo no conectado %s", deviceID)
		return false
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error al serializar JSON para Websocket: %v", err)
		return false
	}

	client.Send <- payload
	return true
}
