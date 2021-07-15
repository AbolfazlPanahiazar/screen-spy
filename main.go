package main

import (
	"image/png"
	"os"

	"github.com/vova616/screenshot"
)

func main() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./ss.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	f.Close()
}

// build command for not showing the cmd on run of exe file
// go build -ldflags -H=windowsgui main.go
