package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/wibbe/glh/gfx"
	"github.com/wibbe/glh/math"
	"io/ioutil"
	"os"
)

const (
	SectorCount = 1024
	WallCount   = 1024 * 8
	EntityCount = 128
)

const (
	_ = iota
	EntityPlayStart
	EntityLight
)

type Sector struct {
	StartWall uint16
	WallCount uint16
}

type Wall struct {
	Pos                    math.Vector2
	Texture                uint16
	TexRepeatX, TexRepeatY float32
	TexOffsetX, TexOffsetY float32
}

type Entity struct {
	Pos        math.Vector2
	Kind       uint8
	Arg1, Arg2 uint32
}

type Level struct {
	Sectors  [SectorCount]Sector
	Walls    [WallCount]Wall
	Entities [EntityCount]Entity
	Textures []string
}

func (l *Level) Draw(ctx gfx.Context) {

}

func LoadLevel(filename string) (*Level, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading level %s: %s\n", filename, err.Error())
		return nil, err
	}

	var level Level
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err = decoder.Decode(&level)

	return &level, nil
}

func (l *Level) Save(filename string) error {
	return nil
}
