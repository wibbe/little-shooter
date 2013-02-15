package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

type Server struct {
	clients      map[uint32]*ServerClient
	nextClientId uint32
	currentLevel string
	clientsMutex sync.Mutex
}

type ServerClient struct {
	id uint32
}

func NewServer(levelToLoad string) *Server {
	server := &Server{clients: make(map[uint32]*ServerClient), nextClientId: 1, currentLevel: levelToLoad}
	go handleServerListen(server)
	return server
}

func handleServerListen(server *Server) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", NetworkPort))
	if err != nil {
		fmt.Fprint(os.Stderr, "Error listening: ", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error accepting client: ", err.Error())
			return
		}

		server.clientsMutex.Lock()
		client := &ServerClient{server.nextClientId}
		server.clients[client.id] = client
		server.nextClientId += 1
		server.clientsMutex.Unlock()

		go handleClientConnection(conn, client)
	}
}

func handleClientConnection(con net.Conn, client *ServerClient) {

}
