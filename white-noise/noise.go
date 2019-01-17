package WhiteNoise

import "math/rand"

func MakeNoise() byte {
	return byte(rand.Intn(255))
}