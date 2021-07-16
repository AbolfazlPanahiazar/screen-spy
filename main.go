package main

import (
	"image/png"
	"os"
	"strconv"
	"time"

	"github.com/vova616/screenshot"
)

func main() {
	for i := 0; i < 10; i++ {
		captureScreen()
		time.Sleep(time.Minute)
	}
}

func captureScreen() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".png")
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
