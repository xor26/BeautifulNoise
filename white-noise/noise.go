package white_noise

import "math/rand"

func makeNoise() byte {
	return byte(rand.Intn(255))
}

func FillPixels(pixels []byte) {
	for i := 0; i < len(pixels)/4; i++ {
		noise := makeNoise()
		pixels[4*i] = noise
		pixels[4*i+1] = noise
		pixels[4*i+2] = noise
	}
}
