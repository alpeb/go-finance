package fin

import (
	"errors"
	"math"
)

// These constants are used in the TVM functions (parameter "paymentType"). They determine wheter payments occur at the end or at the beginning of each period:
const (
	PayEnd = iota
	PayBegin
)

// PresentValue returns the Present Value of a cash flow with constant payments and interest rate (annuities).
//
// Excel equivalent: PV
func PresentValue(rate float64, numPeriods int, pmt float64, fv float64, paymentType int) (pv float64, err error) {
	if numPeriods < 0 {
		return 0, errors.New("Number of periods must be positive")
	}
	if paymentType != PayEnd && paymentType != PayBegin {
		return 0, errors.New("Payment type must be PayEnd or PayBegin")
	}
	if rate != 0 {
		pv = (-pmt*(1+rate*float64(paymentType))*((math.Pow(1+rate, float64(numPeriods))-1)/rate) - fv) / math.Pow(1+rate, float64(numPeriods))
	} else {
		pv = -fv - pmt*float64(numPeriods)
	}
	return pv, nil
}

// FutureValue returns the Future Value of a cash flow with constant payments and interest rate (annuities).
//
// Excel equivalent: FV
func FutureValue(rate float64, numPeriods int, pmt float64, pv float64, paymentType int) (fv float64, err error) {
	if numPeriods < 0 {
		return 0, errors.New("Number of periods must be positive")
	}
	if paymentType != PayEnd && paymentType != PayBegin {
		return 0, errors.New("Payment type must be PayEnd or PayBegin")
	}
	if rate != 0 {
		fv = -pv*math.Pow(1+rate, float64(numPeriods)) - pmt*(1+rate*float64(paymentType))*(math.Pow(1+rate, float64(numPeriods))-1)/rate
	} else {
		fv = -pv - pmt*float64(numPeriods)
	}
	return fv, nil
}

// Payment returns the constant payment (annuity) for a cash flow with a constant interest rate.
//
// Excel equivalent: PMT
func Payment(rate float64, numPeriods int, pv float64, fv float64, paymentType int) (pmt float64, err error) {
	if numPeriods < 0 {
		return 0, errors.New("Number of periods must be positive")
	}
	if paymentType != PayEnd && paymentType != PayBegin {
		return 0, errors.New("Payment type must be PayEnd or PayBegin")
	}
	if rate != 0 {
		pmt = (-fv - pv*math.Pow(1+rate, float64(numPeriods))) / (1 + rate*float64(paymentType)) / ((math.Pow(1+rate, float64(numPeriods)) - 1) / rate)
	} else {
		pmt = (-pv - fv) / float64(numPeriods)
	}
	return pmt, nil
}

// Periods returns the number of periods for a cash flow with constant periodic payments (annuities), and interest rate.
//
// Excel equivalent: NPER
func Periods(rate float64, pmt float64, pv float64, fv float64, paymentType int) (numPeriods float64, err error) {
	if paymentType != PayEnd && paymentType != PayBegin {
		return 0, errors.New("Payment type must be PayEnd or PayBegin")
	}
	if rate != 0 {
		if pmt == 0 && pv == 0 {
			return 0, errors.New("Payment and Present Value can't be both zero when the rate is not zero")
		}
		numPeriods = math.Log((pmt*(1+rate*float64(paymentType))/rate-fv)/(pv+pmt*(1+rate*float64(paymentType))/rate)) / math.Log(1+rate)
	} else {
		if pmt == 0 {
			return 0, errors.New("Rate and Payment can't be both zero")
		}
		numPeriods = (-pv - fv) / pmt
	}
	return numPeriods, nil
}

// Rate returns the periodic interest rate for a cash flow with constant periodic payments (annuities).
// Guess is a guess for the rate, used as a starting point for the iterative algorithm.
//
// Excel equivalent: RATE
func Rate(numPeriods int, pmt float64, pv float64, fv float64, paymentType int, guess float64) (float64, error) {
	if paymentType != PayEnd && paymentType != PayBegin {
		return 0, errors.New("Payment type must be PayEnd or PayBegin")
	}
	function := func(rate float64) float64 {
		return f(rate, numPeriods, pmt, pv, fv, paymentType)
	}
	derivative := func(rate float64) float64 {
		return df(rate, numPeriods, pmt, pv, fv, paymentType)
	}
	return newton(guess, function, derivative, 0)
}

func f(rate float64, numPeriods int, pmt float64, pv float64, fv float64, paymentType int) float64 {
	compounded := math.Pow(1+rate, float64(numPeriods))
	return pv*compounded + pmt*(1+rate*float64(paymentType))*((compounded-1)/rate) + fv
}

func df(rate float64, numPeriods int, pmt float64, pv float64, fv float64, paymentType int) float64 {
	compounded1 := math.Pow(1+rate, float64(numPeriods))
	compounded2 := math.Pow(1+rate, float64(numPeriods-1))
	return float64(numPeriods)*pv*compounded2 + pmt*(float64(paymentType)*(compounded1-1)/rate+(1+rate*float64(paymentType))*(float64(numPeriods)*rate*compounded2-compounded1+1)/math.Pow(rate, 2))
}

// InterestPayment returns the interest payment for a given period for a cash flow with constant periodic payments (annuities)
//
// Excel equivalent: IMPT
func InterestPayment(rate float64, period int, numPeriods int, pv float64, fv float64, paymentType int) (float64, error) {
	if paymentType != PayEnd && paymentType != PayBegin {
		return 0, errors.New("Payment type must be PayEnd or PayBegin")
	}
	interest, _, err := interestAndPrincipal(rate, period, numPeriods, pv, fv, paymentType)
	if err != nil {
		return 0, err
	}
	return interest, nil
}

// PrincipalPayment returns the principal payment for a given period for a cash flow with constant periodic payments (annuities)
//
// Excel equivalent: PPMT
func PrincipalPayment(rate float64, period int, numPeriods int, pv float64, fv float64, paymentType int) (float64, error) {
	if paymentType != PayEnd && paymentType != PayBegin {
		return 0, errors.New("Payment type must be PayEnd or PayBegin")
	}
	_, principal, err := interestAndPrincipal(rate, period, numPeriods, pv, fv, paymentType)
	if err != nil {
		return 0, err
	}
	return principal, nil
}

// interestAndPrincipal returns the interest and principal payment for a given period for a cash flow with constant periodic payments (annuities) and interest rate.
func interestAndPrincipal(rate float64, period int, numPeriods int, pv float64, fv float64, paymentType int) (float64, float64, error) {
	pmt, err := Payment(rate, numPeriods, pv, fv, paymentType)
	if err != nil {
		return 0, 0, err
	}
	capital := pv
	var interest, principal float64
	for i := 1; i <= period; i++ {
		// in first period of advanced payments no interests are paid
		if paymentType == 1 && i == 1 {
			interest = 0
		} else {
			interest = -capital * rate
		}
		principal = pmt - interest
		capital += principal
	}
	return interest, principal, nil
}
