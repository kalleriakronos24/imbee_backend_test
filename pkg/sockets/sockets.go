package sockets

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func RunSocketConnection(socket *socketio.Server) *socketio.Server {

	var server = socket

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		err := s.Close()
		if err != nil {
			return err.Error()
		}
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		fmt.Println("closed", reason)
	})

	go func() {
		err := server.Serve()
		if err != nil {

		}
	}()
	defer func(server *socketio.Server) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	return server
}
