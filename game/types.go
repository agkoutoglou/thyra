package game

import (
	"bytes"
	"net"
)

type Reply struct {
	world  bytes.Buffer
	events string
	intro  bytes.Buffer
}

type ClientRequest struct {
	Client Client
	Cmd    string
}

type LoginRequest struct {
	Username string
	Conn     net.Conn
	Reply    chan bool
}

// Player holds all variables for a character.
type Player struct {
	Nickname string `toml:"nickname"`
	PC
	Area     string `toml:"area"`
	Room     string `toml:"room"`
	Position string `toml:"position"`
}

type Area struct {
	Name  string          `toml:"name"`
	Intro string          `toml:"intro"`
	Rooms map[string]Room `toml:"rooms"`
}

type Room struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	Cubes       []Cube `toml:"cubes"`
}

type Cube struct {
	ID    string `toml:"id"`
	POSX  string `toml:"posx"`
	POSY  string `toml:"posy"`
	Exits []Exit `toml:"exits"`
	Type  string `toml:"type"`
}

type Exit struct {
	ToArea   string `toml:"toarea"`
	ToRoom   string `toml:"toroom"`
	ToCubeID string `toml:"tocubeid"`
}
