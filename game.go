package main

type Game struct {
	server *Server
	player *Client
}

func NewGame(serverAddress string) *Game {
	game := &Game{nil, nil}

	if serverAddress == "localhost" {
		game.server = NewServer()
	}

	game.player = NewClient(serverAddress)
	return game
}

func (g *Game) Tick(dt float64) {

}
