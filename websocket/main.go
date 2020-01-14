package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
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

var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for regis := range manager.clients {
		if regis != ignore {
			regis.send <- message
		}
	}
}

// the func start is a method we will create 3 channel one for managing client , one for reading socket, and one for writing

func (manager *ClientManager) start() {
	for {
		select {
		case regis := <-manager.register:
			manager.clients[regis] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected"})
			manager.send(jsonMessage, regis)
		case regis := <-manager.unregister:
			if _, clientOk := manager.clients[regis]; clientOk {
				close(regis.send)
				delete(manager.clients, regis)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has disconnected"})
				manager.send(jsonMessage, regis)
			}
		case message := <-manager.broadcast:
			for regis := range manager.clients {
				select {
				case regis.send <- message:
				default:
					close(regis.send)
					delete(manager.clients, regis)
				}
			}

		}

	}
}
func (rw *Client) read() {
	defer func() {
		manager.unregister <- rw
		rw.socket.Close()
	}()
	for {
		_, message, err := rw.socket.ReadMessage()
		if err != nil {
			manager.unregister <- rw
			rw.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: rw.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (rw *Client) write() {
	defer func() {
		rw.socket.Close()
	}()
	for {
		select {
		case message, clientOk := <-rw.send:
			if !clientOk {
				rw.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			rw.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}
	manager.register <- client
	fmt.Println("ok")
	go client.read()
	go client.write()
}
func checkerro(err error) {
	if err != nil {

		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting server...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	err := http.ListenAndServe(":12345", nil)
	checkerro(err)
}
