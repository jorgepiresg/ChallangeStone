package main

import (
	"github.com/jorgepiresg/ChallangeStone/config"
	"github.com/jorgepiresg/ChallangeStone/server"
)

func main() {
	cfg := config.New()
	server := server.New(cfg)
	server.Start()
}
