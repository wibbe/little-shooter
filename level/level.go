package level

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
	startWall uint16
	wallCount uint16
}

type Wall struct {
	x, y                   float32
	texture                uint16
	texRepeatX, texRepeatY float32
	texOffsetX, texOffsetY float32
}

type Entity struct {
	x, y       float32
	kind       uint8
	arg1, arg2 uint32
}

type Level struct {
	sectors  [SectorCount]Sector
	walls    [WallCount]Wall
	entities [EntityCount]Entity
}
