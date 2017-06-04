package fin

import (
	"errors"
	"math"
)

// DepreciationFixedDeclining returns the depreciation of an asset using the fixed-declining balance method
// Excel equivalent: DB
func DepreciationFixedDeclining(cost float64, salvage float64, life int, period int, month int) (float64, error) {
	if cost < 0 || life < 0 {
		return 0, errors.New("cost and life must be absolute positive numbers")
	}
	if period < 1 {
		return 0, errors.New("period must be greater or equal than one")
	}
	rate := 1 - math.Pow((salvage/cost), (1/float64(life)))
	rate = round(rate, 3)
	accDepreciation := 0.0
	var depreciationPeriod float64
	for i := 1; i <= period; i++ {
		if i == 1 {
			depreciationPeriod = cost * rate * float64(month) / 12
		} else if i == (life + 1) {
			depreciationPeriod = (cost - accDepreciation) * rate * (12 - float64(month)) / 12
		} else {
			depreciationPeriod = (cost - accDepreciation) * rate
		}
		accDepreciation += depreciationPeriod
	}
	return depreciationPeriod, nil
}

// DepreciationStraighLine returns the straight-line depreciation of an asset for each period
// Excel equivalent: SLN
func DepreciationStraightLine(cost float64, salvage float64, life int) (float64, error) {
	if cost < 0 || life < 0 {
		return 0, errors.New("cost and life must be absolute positive numbers")
	}
	return ((cost - salvage) / float64(life)), nil
}

// DepreciationSYD returns the depreciation for an asset in a given period using the sum-of-years' digits method
// Excel equivalent: SYD
func DepreciationSYD(cost float64, salvage float64, life int, per int) float64 {
	return ((cost - salvage) * float64(life-per+1) * 2 / float64(life) / float64(life+1))
}

func round(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}
