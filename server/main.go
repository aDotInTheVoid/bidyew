package main

import (
	"bidyew/card"
	"encoding/json"
	"log"
	"net/http"
	"time"

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

type GetCardResp struct {
	Name string
}

func getCard(ws *websocket.Conn) string {
	msg := Message{"getCard", nil}
	bin, err := json.Marshal(msg)
	must(err)
	must(ws.WriteMessage(websocket.TextMessage, bin))

	var resp GetCardResp
	must(ws.ReadJSON(&resp))
	return resp.Name
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting Handle")

	conn, err := upgrade.Upgrade(w, r, nil)
	must(err)

	hand := card.DealHand()

	for {

		msg := Message{"clearCenter", nil}
		bin, err := json.Marshal(msg)
		must(err)
		must(conn.WriteMessage(websocket.TextMessage, bin))

		time.Sleep(time.Second)

		msg = Message{"setDeck", hand}
		bin, err = json.Marshal(msg)
		must(err)
		must(conn.WriteMessage(websocket.TextMessage, bin))

		time.Sleep(time.Second)

		for _, name := range []string{"setLeft", "setRight", "setUp", "setDown"} {
			card := card.DealHand()[0]

			msg = Message{name, card}
			bin, err := json.Marshal(msg)
			must(err)
			must(conn.WriteMessage(websocket.TextMessage, bin))

			log.Print(getCard(conn))

			time.Sleep(time.Second)

		}
	}
}

func main() {
	log.Println("Starting Server")

	http.HandleFunc("/ws", handleRequest)
	http.ListenAndServe(":9843", nil)
}
