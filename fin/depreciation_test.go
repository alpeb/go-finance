package fin

import (
	"math"
	"testing"
)

func TestDepreciationFixedDeclining(t *testing.T) {
	var tests = []struct {
		cost    float64
		salvage float64
		life    int
		period  int
		month   int
		want    float64
	}{
		{1000000, 100000, 6, 1, 7, 186083.333333},
		{1000000, 100000, 6, 2, 7, 259639.416667},
		{1000000, 100000, 6, 3, 7, 176814.442750},
		{1000000, 100000, 6, 4, 7, 120410.635513},
		{1000000, 100000, 6, 5, 7, 81999.642784},
		{1000000, 100000, 6, 6, 7, 55841.756736},
		{1000000, 100000, 6, 7, 7, 15845.098474},
	}

	for _, test := range tests {
		if got, _ := DepreciationFixedDeclining(test.cost, test.salvage, test.life, test.period, test.month); math.Abs(test.want-got) > Precision {
			t.Errorf("DepreciationFixedDeclining(%f, %f, %d, %d, %d) = %f", test.cost, test.salvage, test.life, test.period, test.month, got)
		}
	}
}

func TestDepreciationStraightLine(t *testing.T) {
	var tests = []struct {
		cost    float64
		salvage float64
		life    int
		want    float64
	}{
		{30000, 7500, 10, 2250},
	}

	for _, test := range tests {
		if got, _ := DepreciationStraightLine(test.cost, test.salvage, test.life); math.Abs(test.want-got) > Precision {
			t.Errorf("DepreciationStraightLine(%f, %f, %d) = %f", test.cost, test.salvage, test.life, got)
		}
	}
}

func TestDepreciationSYD(t *testing.T) {
	var tests = []struct {
		cost    float64
		salvage float64
		life    int
		per     int
		want    float64
	}{
		{30000, 7500, 10, 1, 4090.909091},
		{30000, 7500, 10, 10, 409.090909},
	}

	for _, test := range tests {
		if got := DepreciationSYD(test.cost, test.salvage, test.life, test.per); math.Abs(test.want-got) > Precision {
			t.Errorf("DepreciationSYD(%f, %f, %d, %d) = %f", test.cost, test.salvage, test.life, test.per, got)
		}
	}
}
