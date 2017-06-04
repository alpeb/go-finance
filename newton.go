package fin

import (
	"errors"
	"math"
)

const MAX_ITERATIONS = 30

func newton(guess float64, function func(float64) float64, derivative func(float64) float64, numIt int) (float64, error) {
	x := guess - function(guess)/derivative(guess)
	if math.Abs(x-guess) < PRECISION {
		return x, nil
	} else if numIt >= MAX_ITERATIONS {
		return 0, errors.New("Solution didn't converge")
	} else {
		return newton(x, function, derivative, numIt+1)
	}
}
