package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	hook "github.com/robotn/gohook"
)

func main() {
	//add()
	fmt.Println("Starting app")
	log.SetFlags(0)
	driver.Main(func(s screen.Screen) {
		fmt.Println("loading image")
		image, err := decode("public/images/grid.png")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("image loaded")
		w := widget.NewSheet(widget.NewImage(image, image.Bounds()))
		if err := widget.RunWindow(s, w, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Title:  "TicTacToe!",
				Width:  image.Bounds().Dx(),
				Height: image.Bounds().Dy(),
			},
		}); err != nil {
			log.Fatal(err)
		}

	})

}

func add() {
	fmt.Println("Adding key event listener")
	/*	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
		hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
			fmt.Println("ctrl-shift-q")
			hook.End()
		})*/

	fmt.Println("--- Please left click button---")
	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		if e.Button == 1 {
			fmt.Println("Click event detected!")
		}
	})

	startProcess := hook.Start()
	<-hook.Process(startProcess)

}

/*
func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}
*/

// TODO: scrolling, such as when images are larger than the window.
func decode(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s: %v", filename, err)
	}
	return m, nil
}
