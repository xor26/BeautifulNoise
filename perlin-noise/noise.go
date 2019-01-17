package WhiteNoise

import "math/rand"


//permutationTable :=


func makeNoise() byte {
	return byte(rand.Intn(255))
}


func FillWithWhiteNoise(pixels []byte, ) {
	for i := 0; i < len(pixels)/4; i++ {
		noise := makeNoise()
		pixels[4*i] = noise
		pixels[4*i+1] = noise
		pixels[4*i+2] = noise
	}
}