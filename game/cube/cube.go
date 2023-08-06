package cube

import "math/rand"

func Dice() int {
	return rand.Intn(6)
}
