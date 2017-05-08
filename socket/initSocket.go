package socket

import (
	"github.com/googollee/go-socket.io"
)

var socket * socketio.Server

	// init socket.io
	// socket, errsock := socketio.NewServer(nil)
	// if errsock != nil {
	// 	fmt.Print(errsock)
	// }

func socketHandler(c *gin.Context) {
	socket.On("connection", func(so socketio.Socket) {
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
	socket.On("error", func(so socketio.Socket, errsock error) {
		fmt.Println("error:", errsock)
	})
}