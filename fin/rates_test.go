package fin

import (
	"math"
	"testing"
)

func TestEffectiveRate(t *testing.T) {
	var tests = []struct {
		nominal    float64
		numPeriods int
		want       float64
	}{
		{0.141, 2, 0.14597025},
		{0.139, 3, 0.14553980},
		{0.009, 4, 0.00903042},
		{0.4, 2, 0.44},
		{0.141, 1, 0.141},
		{0.1390, 5, 0.14694625},
		{0.0090, 77, 0.00904009},
		{0.40, 13, 0.48285538},
	}

	for _, test := range tests {
		if got, _ := EffectiveRate(test.nominal, test.numPeriods); math.Abs(test.want-got) > Precision {
			t.Errorf("EffectiveRate(%f, %d) = %f", test.nominal, test.numPeriods, got)
		}
	}

	if _, err := EffectiveRate(0.40, -13); err == nil {
		t.Error("A negative number of periods should produce an error")
	}
}

func TestNominalRate(t *testing.T) {
	var tests = []struct {
		effective  float64
		numPeriods int
		want       float64
	}{
		{0.56, 2, 0.497999},
		{0.34, 4, 0.303643},
		{0.7450, 3, 0.611767},
		{0.1245, 88, 0.117417},
		{0.0320, 9, 0.031554},
		{0.2930, 5, 0.263683},
	}

	for _, test := range tests {
		if got, _ := NominalRate(test.effective, test.numPeriods); math.Abs(test.want-got) > Precision {
			t.Errorf("NominalRate(%f, %d) = %f", test.effective, test.numPeriods, got)
		}
	}

	if _, err := NominalRate(0.2930, -5); err == nil {
		t.Error("A negative number of periods should produce an error")
	}
}
