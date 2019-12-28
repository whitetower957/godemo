package packages_demo

import (
	"fmt"
	"math/rand"
	"time"
)

/*
import math package to get random seed
*/
func MathRandom() {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100)
	fmt.Println(n)
}
