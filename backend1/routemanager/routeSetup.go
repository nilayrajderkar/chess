package routemanager

import (
	gamemanager "chess/gamemanager"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}

func SetupRoutes(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", reader)
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to play chess!")
}

func reader(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))

	if err!=nil{
		log.Println(err)
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("Received message:", string(message))
		if (string(message) == "GAME_INIT"){
			err = conn.WriteMessage(messageType, []byte("Starting a new game"))
			gamemanager.GameManager(conn)
		} else{
			err = conn.WriteMessage(messageType, []byte("Waiting for right input to join game..."))
		}

		if err != nil {
			log.Println(err)
			return
		}
	}
}



