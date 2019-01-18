package fractal_noise

import (
	"math"
	"time"
)

func FillPixels(pixels []byte) {
	noise := make([]float64, len(pixels)/4)
	min := 9999.0
	max := -99999.0

	for i := 0; i < len(pixels)/4; i++ {
		noise[i] = makeNoise(float64(i))
		if noise[i] > max {
			max = noise[i]
		}

		if noise[i] < min {
			min = noise[i]
		}
	}

	offset := math.Abs(min)
	scale := 255 / (max + offset)
	for i := 0; i < len(pixels)/4; i++ {
		pixel := noise[i]*scale + offset*scale
		pixels[4*i] = byte(pixel)
		pixels[4*i+1] = byte(pixel)
		pixels[4*i+2] = byte(pixel)
	}
}

func makeNoise(x float64) float64 {
	amplitude := .1
	frequency := .01
	y := math.Sin(x * frequency)
	t := float64(time.Now().UnixNano() / 1000)
	y += math.Sin(x*frequency*2.1+t) * 4.5
	y += math.Sin(x*frequency*1.72+t*1.121) * 4.0
	y += math.Sin(x*frequency*2.221+t*0.437) * 5.0
	y += math.Sin(x*frequency*3.1122+t*4.269) * 2.5
	y *= amplitude * 0.06

	return y
}
