package main

import (
	"log"
	"os"

	"github.com/sotanodroid/wezario/internal/wezario"
)

func main() {
	app := wezario.Start()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
