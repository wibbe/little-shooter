package main

import (
	"github.com/wibbe/glh/math"
)

type Client struct {
	position   math.Vector2
	velocity   math.Vector2
	yaw, pitch float32
}

func NewClient(address string) *Client {
	return nil
}
