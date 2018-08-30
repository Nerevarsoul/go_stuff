package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func on_connection(so socketio.Socket) {
	log.Println("on connection")
	so.Join("chat")
}

func on_chat_message(so socketio.Socket, msg string) {
	log.Println(msg)
	so.BroadcastTo("chat", "chat message", msg)
}

func on_disconnect(so socketio.Socket) {
	log.Println("on disconnect")
}

func on_error(so socketio.Socket, err error) {
	log.Println("error:", err)
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", on_connection)
	server.On("chat message", on_chat_message)
	server.On("disconnection", on_disconnect)
	server.On("error", on_error)

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
