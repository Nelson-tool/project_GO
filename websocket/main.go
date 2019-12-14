package main

import (
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
	Sender    string `json:"Sender,omitempty"`
	Recipient string `json:"recipient, omitempty"`
	Content   string `json:"content, omitempty"`
}

var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

//we will create 3 goroutines one for managing client , one for reading socket, and one for writing

func  start(manager *ClientManager) {
	
}

func main() {

}
