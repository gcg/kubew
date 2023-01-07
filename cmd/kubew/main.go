package main

import (
	"log"

	"github.com/gcg/kubew/internal/screen"
	"github.com/gcg/kubew/internal/ticker"
)

func main() {
	screen.Init()             // init the screen
	go ticker.RefreshScreen() // init the ticker

	if err := screen.App.Run(); err != nil {
		log.Panicln("Cant run")
	}
}
