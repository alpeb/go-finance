package fin

import (
	"math"
	"testing"
	"time"
)

func TestDaysDifference(t *testing.T) {
	date1 := time.Date(2005, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()
	date2 := time.Date(2005, time.September, 1, 0, 0, 0, 0, time.UTC).Unix()
	got := DaysDifference(date1, date2, COUNT_ACTUAL_365)
	if got != 62 {
		t.Errorf("DaysDifference(%d, %d) = %d", got)
	}
}

func TestTBillYield(t *testing.T) {
	settlement := time.Date(2008, time.March, 31, 0, 0, 0, 0, time.UTC).Unix()
	maturity := time.Date(2008, time.June, 1, 0, 0, 0, 0, time.UTC).Unix()
	price := 98.45
	got, _ := TBillYield(settlement, maturity, price)
	if math.Abs(got-0.091417) > PRECISION {
		t.Errorf("TBillYield(%d, %d, %f) = %f", settlement, maturity, price, got)
	}

	if _, err := TBillYield(maturity, settlement, price); err == nil {
		t.Errorf("When the settlement happens after the maturity, an error should be returned")
	}

	settlement = time.Date(2008, time.June, 1, 0, 0, 0, 0, time.UTC).Unix()
	maturity = time.Date(2010, time.March, 31, 0, 0, 0, 0, time.UTC).Unix()
	if _, err := TBillYield(settlement, maturity, price); err == nil {
		t.Errorf("When the maturity is more than one year after the settlement, an error should be returned")
	}
}

func TestTBillPrice(t *testing.T) {
	settlement := time.Date(2008, time.March, 31, 0, 0, 0, 0, time.UTC).Unix()
	maturity := time.Date(2008, time.June, 1, 0, 0, 0, 0, time.UTC).Unix()
	discount := 0.09
	got, _ := TBillPrice(settlement, maturity, discount)
	if math.Abs(got-98.45) > PRECISION {
		t.Errorf("TBillPrice(%d, %d, %f) = %f", settlement, maturity, discount, got)
	}

	if _, err := TBillPrice(maturity, settlement, discount); err == nil {
		t.Errorf("When the settlement happens after the maturity, an error should be returned")
	}

	settlement = time.Date(2008, time.June, 1, 0, 0, 0, 0, time.UTC).Unix()
	maturity = time.Date(2010, time.March, 31, 0, 0, 0, 0, time.UTC).Unix()
	if _, err := TBillPrice(settlement, maturity, discount); err == nil {
		t.Errorf("When the maturity is more than one year after the settlement, an error should be returned")
	}
}

func TestTBillEquivalentYield(t *testing.T) {
	var tests = []struct {
		settlement int64
		maturity   int64
		discount   float64
		want       float64
	}{
		{time.Date(2008, time.March, 31, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2008, time.June, 1, 0, 0, 0, 0, time.UTC).Unix(), 0.0914, 0.094151},
		{time.Date(2008, time.March, 20, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2008, time.December, 1, 0, 0, 0, 0, time.UTC).Unix(), 0.0914, 0.097079},
		{time.Date(1993, time.March, 31, 0, 0, 0, 0, time.UTC).Unix(), time.Date(1993, time.December, 15, 0, 0, 0, 0, time.UTC).Unix(), 0.0914, 0.097145},
		{time.Date(1999, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(), 0.0914, 0.099379},
		{time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2001, time.January, 10, 0, 0, 0, 0, time.UTC).Unix(), 0.0914, 0.09994538},
	}

	for _, test := range tests {
		if got, _ := TBillEquivalentYield(test.settlement, test.maturity, test.discount); math.Abs(test.want-got) > PRECISION {
			t.Errorf("TBillEquivalentYield(%d, %d, %f) = %f", test.settlement, test.maturity, test.discount, got)
		}
	}

	if _, err := TBillEquivalentYield(tests[0].maturity, tests[0].settlement, tests[0].discount); err == nil {
		t.Errorf("When the settlement happens after the maturity, an error should be returned")
	}

	settlement := time.Date(2008, time.June, 1, 0, 0, 0, 0, time.UTC).Unix()
	maturity := time.Date(2010, time.March, 31, 0, 0, 0, 0, time.UTC).Unix()
	if _, err := TBillEquivalentYield(settlement, maturity, 0.0914); err == nil {
		t.Errorf("When the maturity is more than one year after the settlement, an error should be returned")
	}
}

func TestDiscountRate(t *testing.T) {
	var tests = []struct {
		settlement int64
		maturity   int64
		price      float64
		redemption float64
		basis      int
		want       float64
	}{
		{time.Date(2007, time.January, 25, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2007, time.June, 15, 0, 0, 0, 0, time.UTC).Unix(), 97.975, 100, COUNT_NASD, 0.052071},
		{time.Date(2007, time.January, 25, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2007, time.June, 15, 0, 0, 0, 0, time.UTC).Unix(), 97.975, 100, COUNT_ACTUAL_ACTUAL, 0.052420},
		{time.Date(2007, time.January, 25, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2007, time.June, 15, 0, 0, 0, 0, time.UTC).Unix(), 97.975, 100, COUNT_ACTUAL_360, 0.051702},
		{time.Date(2007, time.January, 25, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2007, time.June, 15, 0, 0, 0, 0, time.UTC).Unix(), 97.975, 100, COUNT_ACTUAL_365, 0.052420},
		{time.Date(2007, time.January, 25, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2007, time.June, 15, 0, 0, 0, 0, time.UTC).Unix(), 97.975, 100, COUNT_EUROPEAN, 0.052071},
	}

	for _, test := range tests {
		if got := DiscountRate(test.settlement, test.maturity, test.price, test.redemption, test.basis); math.Abs(test.want-got) > PRECISION {
			t.Errorf("DiscountRate(%d, %d, %f, %f, %d) = %f", test.settlement, test.maturity, test.price, test.redemption, test.basis, got)
		}
	}
}

func TestPriceDiscount(t *testing.T) {
	var tests = []struct {
		settlement int64
		maturity   int64
		discount   float64
		redemption float64
		basis      int
		want       float64
	}{
		{time.Date(2008, time.February, 16, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2008, time.March, 1, 0, 0, 0, 0, time.UTC).Unix(), 0.0525, 100, COUNT_NASD, 99.795833},
		{time.Date(2008, time.February, 16, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2008, time.March, 1, 0, 0, 0, 0, time.UTC).Unix(), 0.0525, 100, COUNT_ACTUAL_ACTUAL, 99.799180},
		{time.Date(2008, time.February, 16, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2008, time.March, 1, 0, 0, 0, 0, time.UTC).Unix(), 0.0525, 100, COUNT_ACTUAL_360, 99.795833},
		{time.Date(2008, time.February, 16, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2008, time.March, 1, 0, 0, 0, 0, time.UTC).Unix(), 0.0525, 100, COUNT_ACTUAL_365, 99.798630},
		{time.Date(2008, time.February, 16, 0, 0, 0, 0, time.UTC).Unix(), time.Date(2008, time.March, 1, 0, 0, 0, 0, time.UTC).Unix(), 0.0525, 100, COUNT_EUROPEAN, 99.795833},
	}

	for _, test := range tests {
		if got := PriceDiscount(test.settlement, test.maturity, test.discount, test.redemption, test.basis); math.Abs(test.want-got) > PRECISION {
			t.Errorf("PriceDiscount(%d, %d, %f, %f, %d) = %f", test.settlement, test.maturity, test.discount, test.redemption, test.basis, got)
		}
	}
}
