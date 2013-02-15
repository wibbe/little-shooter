package main

const (
	FrameTick float64 = 1.0 / 20.0
)

const (
	NetworkPort = 6798
)

type Game struct {
	server *Server
	player *Client
}

func NewGame(serverAddress string) *Game {
	game := &Game{nil, nil}

	if serverAddress == "localhost" {
		game.server = NewServer("")
	}

	game.player = NewClient(serverAddress)
	return game
}

func (g *Game) Tick(dt float64) {

}
