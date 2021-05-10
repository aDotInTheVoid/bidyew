package main

import (
	"bidyew/card"
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
	ty   string
	data interface{}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting Handle")

	conn, err := upgrade.Upgrade(w, r, nil)
	must(err)

	hand := card.DealHand()

	must(conn.WriteJSON(Message{
		"sethand", hand,
	}))

}

func main() {
	log.Println("Starting Server")

	http.HandleFunc("/ws", handleRequest)
	http.ListenAndServe(":9843", nil)
}
