package fin

import (
	"errors"
	"math"
)

// NetPresentValue returns the Net Present Value of a cash flow series given a discount rate
//
// Excel equivalent: NPV
func NetPresentValue(rate float64, values []float64) float64 {
	npv := 0.0
	nper := len(values)
	for i := 1; i <= nper; i++ {
		npv += values[i-1] / math.Pow(1+rate, float64(i))
	}
	return npv
}

// InternalRateOfReturn returns the internal rate of return of a cash flow series.
// Guess is a guess for the rate, used as a starting point for the iterative algorithm.
//
// Excel equivalent: IRR
func InternalRateOfReturn(values []float64, guess float64) (float64, error) {
	min, max := minMaxSlice(values)
	if min*max >= 0 {
		return 0, errors.New("The cash flow must contain at least one positive value and one negative value")
	}

	function := func(rate float64) float64 {
		return NetPresentValue(rate, values)
	}
	derivative := func(rate float64) float64 {
		return dNetPresentValue(rate, values)
	}
	return newton(guess, function, derivative, 0)
}

func dNetPresentValue(rate float64, values []float64) float64 {
	dnpv := 0.0
	nper := len(values)
	for i := 1; i <= nper; i++ {
		dnpv += values[i-1] * float64(-i) * math.Pow(1+rate, float64(i-1)) / math.Pow(1+rate, float64(2*i))
	}
	return dnpv
}

// ModifiedInternalRateOfReturn returns the internal rate of return of a cash flow series, considering both financial and reinvestment rates
//
// financeRate is the rate on the money used in the cash flow.
//
// reinvestRate is the rate received when reinvested
//
// Excel equivalent: MIRR
func ModifiedInternalRateOfReturn(values []float64, financeRate float64, reinvestRate float64) (float64, error) {
	min, max := minMaxSlice(values)
	if min*max >= 0 {
		return 0, errors.New("The cash flow must contain at least one positive value and one negative value")
	}
	positiveFlows := make([]float64, 0)
	negativeFlows := make([]float64, 0)
	for _, value := range values {
		if value >= 0 {
			positiveFlows = append(positiveFlows, value)
			negativeFlows = append(negativeFlows, 0)
		} else {
			positiveFlows = append(positiveFlows, 0)
			negativeFlows = append(negativeFlows, value)
		}
	}
	nper := len(values)
	return math.Pow(-NetPresentValue(reinvestRate, positiveFlows)*math.Pow(1+reinvestRate, float64(nper))/NetPresentValue(financeRate, negativeFlows)/(1+financeRate), 1/float64(nper-1)) - 1, nil
}

func minMaxSlice(values []float64) (float64, float64) {
	min := math.MaxFloat64
	max := -min
	for _, value := range values {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return min, max
}
