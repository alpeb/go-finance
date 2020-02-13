package fin

import (
	"math"
	"testing"
	"time"
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
		if got := NetPresentValue(test.rate, test.values); math.Abs(test.want-got) > Precision {
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
		if got, _ := InternalRateOfReturn(test.values, test.guess); math.Abs(test.want-got) > Precision {
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
		if got, _ := ModifiedInternalRateOfReturn(test.values, test.financeRate, test.reinvestRate); math.Abs(test.want-got) > Precision {
			t.Errorf("ModifiedInternalRateOfReturn(%v, %f, %f) = %f", test.values, test.financeRate, test.reinvestRate, got)
		}
	}

	if _, err := ModifiedInternalRateOfReturn([]float64{70000, 12000, 15000, 18000, 21000}, 0.1, 0.1); err == nil {
		t.Error("If the cash flow doesn't contain at least one positive value and one negative value, it must return an error")
	}
}

func TestScheduledNetPresentValue(t *testing.T) {
	var tests = []struct {
		rate   float64
		values []float64
		dates  []time.Time
		want   float64
	}{
		{
			0.09,
			[]float64{
				-10000,
				2750,
				4250,
				3250,
				2750,
			},
			[]time.Time{
				time.Date(2008, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
				time.Date(2008, time.Month(3), 1, 0, 0, 0, 0, time.UTC),
				time.Date(2008, time.Month(10), 30, 0, 0, 0, 0, time.UTC),
				time.Date(2009, time.Month(2), 15, 0, 0, 0, 0, time.UTC),
				time.Date(2009, time.Month(4), 1, 0, 0, 0, 0, time.UTC),
			},
			2086.647602,
		},
		{
			0.2,
			[]float64{
				-2000,
				1000,
				1000,
				1000,
			},
			[]time.Time{
				time.Date(2020, time.Month(2), 12, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.Month(4), 20, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.Month(5), 20, 0, 0, 0, 0, time.UTC),
			},
			900.5182206,
		},
	}

	for _, test := range tests {
		if got, _ := ScheduledNetPresentValue(test.rate, test.values, test.dates); math.Abs(test.want-got) > Precision {
			t.Errorf("ScheduledNetPresentValue(%f, %v, %v) = %f", test.rate, test.values, test.dates, got)
		}
	}

	if _, err := ScheduledNetPresentValue(0.1, []float64{-10000}, []time.Time{}); err == nil {
		t.Error("If values and dates have different lengths, it must return an error")
	}
}

func TestScheduledInternalRateOfReturn(t *testing.T) {
	var tests = []struct {
		values []float64
		dates  []time.Time
		guess  float64
		want   float64
	}{
		{
			[]float64{
				-10000,
				2750,
				4250,
				3250,
				2750,
			},
			[]time.Time{
				time.Date(2008, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
				time.Date(2008, time.Month(3), 1, 0, 0, 0, 0, time.UTC),
				time.Date(2008, time.Month(10), 30, 0, 0, 0, 0, time.UTC),
				time.Date(2009, time.Month(2), 15, 0, 0, 0, 0, time.UTC),
				time.Date(2009, time.Month(4), 1, 0, 0, 0, 0, time.UTC),
			},
			0.1,
			0.373362535,
		},
		{
			[]float64{
				-2000,
				1000,
				1000,
				1000,
			},
			[]time.Time{
				time.Date(2020, time.Month(2), 12, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.Month(4), 20, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.Month(5), 20, 0, 0, 0, 0, time.UTC),
			},
			0.1,
			8.493343973,
		},
	}

	for _, test := range tests {
		if got, _ := ScheduledInternalRateOfReturn(test.values, test.dates, test.guess); math.Abs(test.want-got) > Precision {
			t.Errorf("ScheduledInternalRateOfReturn(%v, %v, %f) = %f", test.values, test.dates, test.guess, got)
		}
	}

	if _, err := ScheduledInternalRateOfReturn([]float64{10000, 2750}, []time.Time{time.Now(), time.Now()}, 0.1); err == nil {
		t.Error("If the cash flow doesn't contain at least one positive value and one negative value, it must return an error")
	}

	if _, err := ScheduledInternalRateOfReturn([]float64{-10000, 2750}, []time.Time{}, 0.1); err == nil {
		t.Error("If values and dates have different lengths, it must return an error")
	}
}
