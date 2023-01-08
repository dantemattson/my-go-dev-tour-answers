package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func errorFunction(approximation float64, original float64) (ret float64) {
	ret = math.Abs(approximation*approximation - original)
	return
}

func debugOutput(numTimesLooped int, approximation float64, original float64) {
	fmt.Printf("Newton's method looped %d times\nError: %.12f\n", numTimesLooped, errorFunction(approximation, original))
}

func Sqrt(original float64) (approximation float64, e error) {
	if original < 0 {
		e = ErrNegativeSqrt(original)
		return
	}
	maxitr, epsilon := 100, 1.0e-15

	approximation = 1.0
	for i := 0; i < maxitr; i++ {
		approximation -= (approximation*approximation - original) / (2 * approximation)
		if errorFunction(approximation, original) <= epsilon {
			debugOutput(i, approximation, original)
			return
		}
	}
	debugOutput(maxitr, approximation, original)
	return
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(-3))
}
