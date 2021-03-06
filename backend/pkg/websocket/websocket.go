package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return conn, err
	}

	return conn, nil
}

/*
// define a reader which will listen for
// new messages being sent to our WebSocket endpoint
func Reader(conn *websocket.Conn) {
	for {
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(msgType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")

		messageType, r, err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			log.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			log.Println(err)
			return
		}
	}
}
*/
