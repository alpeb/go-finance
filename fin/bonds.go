package fin

import (
	"errors"
	"math"
	"time"
)

// These constants are used in the bonds functions (parameter "basis"), for specifying the basis for the type of day count:
const (
	// US(NASD) 30/360
	CountNasd = iota
	// Actual/actual
	CountActualActual
	// Actual/360
	CountActual360
	// Actual/365
	CountActual365
	// European 30/360
	CountEuropean
)

// DaysDifference returns the difference of days between two dates based on a daycount basis.
// Date1 and date2 are UNIX timestamps (seconds).
func DaysDifference(date1 int64, date2 int64, basis int) int {
	y1, mName1, d1 := time.Unix(date1, 0).Date()
	m1 := int(mName1)
	y2, mName2, d2 := time.Unix(date2, 0).Date()
	m2 := int(mName2)
	switch basis {
	case CountNasd:
		if d2 == 31 && (d1 == 30 || d1 == 31) {
			d2 = 30
		}
		if d1 == 31 {
			d1 = 30
		}
		return (y2-y1)*360 + (m2-m1)*30 + d2 - d1
	case CountActualActual, CountActual360, CountActual365:
		return int((date2 - date1) / 86400)
	case CountEuropean:
		return (y2-y1)*360 + (m2-m1)*30 + d2 - d1
	}
	return 0
}

// DaysPerYear returns the number of days in the year based on a daycount basis.
func DaysPerYear(year int, basis int) int {
	switch basis {
	case CountNasd:
		return 360
	case CountActualActual:
		if isLeap(year) {
			return 366
		}
		return 365
	case CountActual360:
		return 360
	case CountActual365:
		return 365
	case CountEuropean:
		return 360
	}
	return 0
}

// TBillYield returns the yield for a treasury bill
//
// settlement is the unix timestamp (seconds) for the settlement date
//
// maturity is the unix timestamp (seconds) for the maturity date
//
// price is the TBill price per $100 face value
//
// Excel equivalent: TBILLYIELD
func TBillYield(settlement int64, maturity int64, price float64) (float64, error) {
	if settlement >= maturity {
		return 0, errors.New("Maturity must happen before settlement!")
	}
	dsm := float64(maturity-settlement) / float64(86400) // transform to days
	if dsm > 360 {
		return 0, errors.New("maturity can't be more than one year after settlement")
	}
	return (100 - price) * 360 / price / dsm, nil
}

// TBillPrice returns the price per $100 face value for a Treasury bill
//
// settlement is the unix timestamp (seconds) for the settlement date
//
// maturity is the unix timestamp (seconds) for the maturity date
//
// discount is the T-Bill discount rate
//
// Excel equivalent: TBILLPRICE
func TBillPrice(settlement int64, maturity int64, discount float64) (float64, error) {
	if settlement >= maturity {
		return 0, errors.New("maturity must happen before settlement!")
	}
	dsm := float64(maturity-settlement) / float64(86400) // transform to days
	if dsm > 360 {
		return 0, errors.New("maturity can't be more than one year after settlement")
	}
	return 100 * (1 - discount*dsm/360), nil
}

// TBillEquivalentYield returns the bond-equivalent yield for a Treasury bill
//
// settlement is the unix timestamp (seconds) for the settlement date
//
// maturity is the unix timestamp (seconds) for the maturity date
//
// discount is the T-Bill discount rate
//
// Excel equivalent: TBILLEQ
func TBillEquivalentYield(settlement int64, maturity int64, discount float64) (float64, error) {
	if settlement >= maturity {
		return 0, errors.New("Maturity must happen before settlement!")
	}
	dsm := float64(DaysDifference(settlement, maturity, CountActual365))
	ySettlement, mNameSettlement, _ := time.Unix(settlement, 0).Date()
	mSettlement := int(mNameSettlement)
	yMaturity, _, _ := time.Unix(maturity, 0).Date()
	if dsm <= 182 {
		// for one half year or less, the bond-equivalent-yield is equivalent to an actual/365 interest rate
		return 365 * discount / (360 - discount*dsm), nil
	} else if dsm == 366 &&
		((mSettlement <= 2 && isLeap(ySettlement)) ||
			(mSettlement > 2 && isLeap(yMaturity))) {
		return 2 * (math.Sqrt(1-discount*366/(discount*366-360)) - 1), nil
	} else if dsm > 365 {
		return 0, errors.New("maturity can't be more than one year after settlement")
	}
	return (-dsm + math.Sqrt(math.Pow(dsm, 2)-(2*dsm-365)*discount*dsm*365/(discount*dsm-360))) / (dsm - 365/2), nil
}

// DiscountRate returns the discount rate for a bond
//
// settlement is the unix timestamp (seconds) for the settlement date
//
// maturity is the unix timestamp (seconds) for the maturity date
//
// price is the bond's price per $100 face value
//
// redemption is the bond's redemption value per $100 face value
//
// Excel equivalent: DISC
func DiscountRate(settlement int64, maturity int64, price float64, redemption float64, basis int) float64 {
	year, _, _ := time.Unix(settlement, 0).Date()
	daysPerYear := DaysPerYear(year, basis)
	dsm := DaysDifference(settlement, maturity, basis)
	return (redemption - price) * float64(daysPerYear) / redemption / float64(dsm)
}

// PriceDiscount returns the price per $100 face value of a discounted bond
//
// settlement is the unix timestamp (seconds) for the settlement date
//
// maturity is the unix timestamp (seconds) for the maturity date
//
// discount is the bond's discount rate
//
// redemption is the bond's redemption value per $100 face value
//
// Excel equivalent: PRICEDISC
func PriceDiscount(settlement int64, maturity int64, discount float64, redemption float64, basis int) float64 {
	year, _, _ := time.Unix(settlement, 0).Date()
	daysPerYear := DaysPerYear(year, basis)
	dsm := DaysDifference(settlement, maturity, basis)
	return redemption - discount*redemption*float64(dsm)/float64(daysPerYear)
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
