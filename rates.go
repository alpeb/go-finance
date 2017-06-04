package fin

import (
	"errors"
	"math"
)

const PRECISION = 1E-6

// EffectiveRate returns the effective interest rate given the nominal rate and the number of compounding payments per year.
// Excel equivalent: EFFECT
func EffectiveRate(nominal float64, numPeriods int) (float64, error) {
	if numPeriods < 0 {
		return 0, errors.New("numPeriods must be strictly positive")
	}
	return math.Pow(1+nominal/float64(numPeriods), float64(numPeriods)) - 1, nil
}

// NominalRate returns the nominal interest rate given the effective rate and the number of compounding payments per year.
// Excel equivalent: NOMINAL
func NominalRate(effectiveRate float64, numPeriods int) (float64, error) {
	if numPeriods < 0 {
		return 0, errors.New("Number of compounding payments per year must be strictly positive")
	}
	return float64(numPeriods) * (math.Pow(effectiveRate+1, 1/float64(numPeriods)) - 1), nil
}
