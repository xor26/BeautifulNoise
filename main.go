package main

import (
	"fmt"
	"github.com/BeautifulNoise/white-noise"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 1200, 800

type color struct {
	r, g, b byte
}

func fillWithWhiteNoise(pixels []byte) {
	for i := 0; i < winWidth*winHeight; i++ {
		noise := WhiteNoise.MakeNoise()
		pixels[4*i] = noise
		pixels[4*i+1] = noise
		pixels[4*i+2] = noise
	}
}

func setPixelColor(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index > len(pixels)-4 {
		return
	}

	pixels[index] = c.r
	pixels[index+1] = c.g
	pixels[index+2] = c.b
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()


	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		pixels := make([]byte, winWidth*winHeight*4)
		fillWithWhiteNoise(pixels)

		//keyState := sdl.GetKeyboardState()


		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(10)

	}
}