package main

import (
	"github.com/wibbe/glh/gfx"
)

const (
	FrameTick float64 = 1.0 / 20.0
)

const (
	NetworkPort = 6798
)

type Game struct {
	server *Server
	player *Client
	ctx    gfx.Context
}

func NewGame(serverAddress string, ctx gfx.Context) *Game {
	game := &Game{nil, nil, ctx}

	if serverAddress == "localhost" {
		game.server = NewServer("")
	}

	game.player = NewClient(serverAddress)
	return game
}

func (g *Game) Tick(dt float64) {

}

func (g *Game) Draw() {
	g.ctx.Clear(true, true, true)

	if *editMode {

	} else {

	}
}
