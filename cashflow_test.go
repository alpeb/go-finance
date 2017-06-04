package fin

import (
	"math"
	"testing"
)

func TestNetPresentValue(t *testing.T) {
	var tests = []struct {
		rate   float64
		values []float64
		want   float64
	}{
		{0.1, []float64{-10000, 3000, 4200, 6800}, 1188.443412},
	}

	for _, test := range tests {
		if got := NetPresentValue(test.rate, test.values); math.Abs(test.want-got) > PRECISION {
			t.Errorf("NetPresentValue(%f, %v) = %f", test.rate, test.values, got)
		}
	}
}

func TestInternalRateOfReturn(t *testing.T) {
	var tests = []struct {
		values []float64
		guess  float64
		want   float64
	}{
		{[]float64{-70000, 12000, 15000, 18000, 21000}, 0.1, -0.02124485},
		{[]float64{-70000, 12000, 15000, 18000, 21000, 26000}, 0.1, 0.086630},
		{[]float64{-70000, 12000, 15000}, -0.40, -0.443507},
	}

	for _, test := range tests {
		if got, _ := InternalRateOfReturn(test.values, test.guess); math.Abs(test.want-got) > PRECISION {
			t.Errorf("InternalRateOfReturn(%v, %f) = %f", test.values, test.guess, got)
		}
	}

	if _, err := InternalRateOfReturn([]float64{70000, 12000, 15000, 18000, 21000}, 0.1); err == nil {
		t.Error("If the cash flow doesn't contain at least one positive value and one negative value, it must return an error")
	}
}

func TestModifiedInternalRateOfReturn(t *testing.T) {
	var tests = []struct {
		values       []float64
		financeRate  float64
		reinvestRate float64
		want         float64
	}{
		{[]float64{-120000, 39000, 30000, 21000, 37000, 46000}, 0.10, 0.12, 0.126094},
		{[]float64{-120000, 39000, 30000, 21000}, 0.10, 0.12, -0.048044},
		{[]float64{-120000, 39000, 30000, 21000, 37000, 46000}, 0.10, 0.14, 0.134759},
	}

	for _, test := range tests {
		if got, _ := ModifiedInternalRateOfReturn(test.values, test.financeRate, test.reinvestRate); math.Abs(test.want-got) > PRECISION {
			t.Errorf("ModifiedInternalRateOfReturn(%v, %f, %f) = %f", test.values, test.financeRate, test.reinvestRate, got)
		}
	}

	if _, err := ModifiedInternalRateOfReturn([]float64{70000, 12000, 15000, 18000, 21000}, 0.1, 0.1); err == nil {
		t.Error("If the cash flow doesn't contain at least one positive value and one negative value, it must return an error")
	}
}
