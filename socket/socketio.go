package socket

import (
	"fmt"
	"github.com/googollee/go-socket.io"
)

func InitSocket() {
	// init socket.io
	socketIO, errsock := socketio.NewServer(nil)
	if errsock != nil {
		fmt.Println(errsock)
	}
	socketIO.On("connection", func(so socketio.Socket) {
		fmt.Println("on connection")
		so.Join("RPi.EIT")
		so.On("runReconstruction", func(msg string) {
			fmt.Println("this reconstruction")
		})
		so.On("raspiConnect", func(msg string) {
			fmt.Println("raspi connected")
		})
		so.On("disconnection", func() {
			fmt.Println("on disconnect")
		})
	})
	socketIO.On("error", func(so socketio.Socket, errsock error) {
		fmt.Println("error:", errsock)
	})
}