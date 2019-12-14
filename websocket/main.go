package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

//createstructureforclient manager
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

//client parameters
type Client struct {
	id     string
	name   string
	socket *websocket.Conn
	send   chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient, omitempty"`
	Content   string `json:"content, omitempty"`
}

var manager = ClientManager {
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

//we will create 3 goroutines one for managing client , one for reading socket, and one for writing

func  (manager *ClientManager) start(){
	for {
		select {
		case regis := <-manager.register:
			manager.clients[regis] = true
			jsonMessage,_ := json.Marshal(&Message{Content:"/A new socket has connected"})
			manager.send(jsonMessage, regis)
		case regis := <-manager.unregister:
			if _, client_ok := manager.clients[regis]; client_ok {
				close(regis.send)
				delete(manager.clients, regis)
				jsonMessage, _ := json.Marshal(&Message{Content:"/A new socket has disconnected"})
				manager.send(jsonMessage, regis)
			}
		case regis := <-manager.broadcast:
			for regis := 

		}

	}

	
}

func main() {

}
