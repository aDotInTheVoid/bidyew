package main

import (
	"bidyew/card"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func must(err error) {
	if err != nil {
		log.Fatal("Fatal:", err)
	}
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

type Message struct {
	Kind string
	Data interface{}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting Handle")

	conn, err := upgrade.Upgrade(w, r, nil)
	must(err)

	hand := card.DealHand()

	msg := Message{"setDeck", hand}

	bmsg, err := json.Marshal(msg)
	log.Printf("%s", string(bmsg))

	must(err)

	must(conn.WriteMessage(websocket.TextMessage, bmsg))
	log.Println("Closing")

}

func main() {
	log.Println("Starting Server")

	http.HandleFunc("/ws", handleRequest)
	http.ListenAndServe(":9843", nil)
}
