package main

type Server struct {
	clients []Client
}

func NewServer() *Server {
	return &Server{}
}
