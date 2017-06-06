package fin

import (
	"math"
	"testing"
)

func TestPresentValue(t *testing.T) {
	var tests = []struct {
		rate        float64
		numPeriods  int
		pmt         float64
		fv          float64
		paymentType int
		want        float64
	}{
		{0.08, 20, 500, 0, PayEnd, -4909.073704},
		{0.03, 5, 200, 0, PayEnd, -915.941437},
		{0.29, 7, 100, 0, PayEnd, -286.821438},
		{0, 7, 100, 0, PayEnd, -700.000000},
	}

	for _, test := range tests {
		if got, _ := PresentValue(test.rate, test.numPeriods, test.pmt, test.fv, test.paymentType); math.Abs(test.want-got) > Precision {
			t.Errorf("PresentValue(%f, %d, %f, %f, %d) = %f", test.rate, test.numPeriods, test.pmt, test.fv, test.paymentType, got)
		}
	}

	if _, err := PresentValue(0.29, -7, 100, 0, PayEnd); err == nil {
		t.Error("A negative number of periods should produce an error")
	}

	if _, err := PresentValue(0.29, 7, 100, 0, 3); err == nil {
		t.Error("An invalid payment type should return an error")
	}
}

func TestFutureValue(t *testing.T) {
	var tests = []struct {
		rate        float64
		numPeriods  int
		pmt         float64
		pv          float64
		paymentType int
		want        float64
	}{
		{0.08, 20, 500, 0, PayEnd, -22880.982149},
		{0.03, 5, 200, 0, PayEnd, -1061.827162},
		{0.29, 7, 100, 0, PayEnd, -1705.059664},
		{0, 7, 100, 0, PayEnd, -700.000000},
	}

	for _, test := range tests {
		if got, _ := FutureValue(test.rate, test.numPeriods, test.pmt, test.pv, test.paymentType); math.Abs(test.want-got) > Precision {
			t.Errorf("FutureValue(%f, %d, %f, %f, %d) = %f", test.rate, test.numPeriods, test.pmt, test.pv, test.paymentType, got)
		}
	}

	if _, err := FutureValue(0.29, -7, 100, 0, PayEnd); err == nil {
		t.Error("A negative number of periods should produce an error")
	}

	if _, err := FutureValue(0.29, 7, 100, 0, 3); err == nil {
		t.Error("An invalid payment type should return an error")
	}
}

func TestPayment(t *testing.T) {
	var tests = []struct {
		rate        float64
		numPeriods  int
		pv          float64
		fv          float64
		paymentType int
		want        float64
	}{
		{0.08, 20, 355, 0, PayEnd, -36.157534},
		{0.03, 5, 828, 0, PayEnd, -180.797585},
		{0.29, 7, 477, 0, PayEnd, -166.305561},
		{0, 7, 435, 0, PayEnd, -62.142857},
		{0.1 / 12, 3 * 12, 8000, 0, PayEnd, -258.137498},
		{0.1 / 12, 3 * 12, 8000, 0, PayBegin, -256.004130},
	}

	for _, test := range tests {
		if got, _ := Payment(test.rate, test.numPeriods, test.pv, test.fv, test.paymentType); math.Abs(test.want-got) > Precision {
			t.Errorf("Payment(%f, %d, %f, %f, %d) = %f", test.rate, test.numPeriods, test.pv, test.fv, test.paymentType, got)
		}
	}

	if _, err := Payment(0.29, -7, 435, 0, PayEnd); err == nil {
		t.Error("A negative number of periods should produce an error")
	}

	if _, err := Payment(0.29, 7, 100, 0, 3); err == nil {
		t.Error("An invalid payment type should return an error")
	}
}

func TestPeriods(t *testing.T) {
	var tests = []struct {
		rate        float64
		pmt         float64
		pv          float64
		fv          float64
		paymentType int
		want        float64
	}{
		{0.08, -500, 355, 0, PayEnd, 0.759825},
		{0.03, -200, 828, 0, PayEnd, 4.486566},
		{0.45, -5000, 344, 0, PayEnd, 0.084641},
		{0, -100, 435, 0, PayEnd, 4.350000},
	}

	for _, test := range tests {
		if got, _ := Periods(test.rate, test.pmt, test.pv, test.fv, test.paymentType); math.Abs(test.want-got) > Precision {
			t.Errorf("Periods(%f, %f, %f, %f, %d) = %f", test.rate, test.pmt, test.pv, test.fv, test.paymentType, got)
		}
	}

	if _, err := Periods(0.5, 0, 0, 0, PayEnd); err == nil {
		t.Error("Payment and Present Value both zeroes should return an error")
	}

	if _, err := Periods(0, 0, 0, 0, PayEnd); err == nil {
		t.Error("Rate and Payment both zeroes should return an error")
	}

	if _, err := Periods(0.29, 100, 477, 0, 3); err == nil {
		t.Error("An invalid payment type should return an error")
	}
}

func TestRate(t *testing.T) {
	var tests = []struct {
		numPeriods  int
		pmt         float64
		pv          float64
		fv          float64
		paymentType int
		guess       float64
		want        float64
	}{
		{20, -36.157534, 355, 0, PayEnd, 0.1, 0.08000},
		{5, -180.797585, 828, 0, PayEnd, 0.1, 0.03000},
		{2, -295.208163, 344, 0, PayEnd, 0.1, 0.45000},
	}

	for _, test := range tests {
		if got, _ := Rate(test.numPeriods, test.pmt, test.pv, test.fv, test.paymentType, test.guess); math.Abs(test.want-got) > Precision {
			t.Errorf("Rate(%d, %f, %f, %f, %d, %f) = %f", test.numPeriods, test.pmt, test.pv, test.fv, test.paymentType, test.guess, got)
		}
	}

	if _, err := Rate(20, -36.157534, 255, 0, 3, 0.1); err == nil {
		t.Error("An invalid payment type should return an error")
	}
}

func TestInterestPayment(t *testing.T) {
	var tests = []struct {
		rate        float64
		period      int
		numPeriods  int
		pv          float64
		fv          float64
		paymentType int
		want        float64
	}{
		{0.1 / 12, 3, 3 * 12, 8000, 0, PayEnd, -63.462189},
		{0.1 / 12, 13, 3 * 12, 8000, 0, PayEnd, -46.617168},
		{0.1 / 12, 33, 3 * 12, 8000, 0, PayEnd, -8.428265},
		{0.1 / 12, 3, 3 * 12, 8000, 0, PayBegin, -62.937708},
	}

	for _, test := range tests {
		if got, _ := InterestPayment(test.rate, test.period, test.numPeriods, test.pv, test.fv, test.paymentType); math.Abs(test.want-got) > Precision {
			t.Errorf("InterestPayment(%f, %d, %d, %f, %f, %d) = %f", test.rate, test.period, test.numPeriods, test.pv, test.fv, test.paymentType, got)
		}
	}

	if _, err := InterestPayment(0.1/12, 3, 3*12, 8000, 0, 2); err == nil {
		t.Error("An invalid payment type should return an error")
	}
}

func TestPrincipalPayment(t *testing.T) {
	var tests = []struct {
		rate        float64
		period      int
		numPeriods  int
		pv          float64
		fv          float64
		paymentType int
		want        float64
	}{
		{0.1 / 12, 3, 3 * 12, 8000, 0, PayEnd, -194.675308},
		{0.1 / 12, 13, 3 * 12, 8000, 0, PayEnd, -211.5203289},
		{0.1 / 12, 33, 3 * 12, 8000, 0, PayEnd, -249.709231},
		{0.1 / 12, 3, 3 * 12, 8000, 0, PayBegin, -193.066421},
	}

	for _, test := range tests {
		if got, _ := PrincipalPayment(test.rate, test.period, test.numPeriods, test.pv, test.fv, test.paymentType); math.Abs(test.want-got) > Precision {
			t.Errorf("PrincipalPayment(%f, %d, %d, %f, %f, %d) = %f", test.rate, test.period, test.numPeriods, test.pv, test.fv, test.paymentType, got)
		}
	}

	if _, err := PrincipalPayment(0.1/12, 3, 3*12, 8000, 0, 2); err == nil {
		t.Error("An invalid payment type should return an error")
	}
}
