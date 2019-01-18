package FractalNoise

import (
	"math"
	"time"
)

func MakeNoise(x float64) float64 {
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
