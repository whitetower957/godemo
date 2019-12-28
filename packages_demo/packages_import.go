package packages_demo

import (
	"math/rand"
	"time"
)

/*
import math package to get random seed
*/
func MathRandom() {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100)
	pacPrint(n)
}
