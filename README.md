# go-finance [![GoDoc](https://godoc.org/github.com/alpeb/go-finance?status.svg)](https://godoc.org/github.com/alpeb/go-finance/fin) [![Go Report Card](https://goreportcard.com/badge/github.com/alpeb/go-finance)](https://goreportcard.com/report/github.com/alpeb/go-finance)
Go library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.

## List of functions

### Rates

- [EffectiveRate](https://godoc.org/github.com/alpeb/go-finance/fin#EffectiveRate)
- [NominalRate](https://godoc.org/github.com/alpeb/go-finance/fin#NominalRate)

### Cashflow

- [NetPresentValue](https://godoc.org/github.com/alpeb/go-finance/fin#NetPresentValue)
- [InternalRateOfReturn](https://godoc.org/github.com/alpeb/go-finance/fin#InternalRateOfReturn)
- [ModifiedInternalRateOfReturn](https://godoc.org/github.com/alpeb/go-finance/fin#ModifiedInternalRateOfReturn)
- [ScheduledNetPresentValue](https://godoc.org/github.com/alpeb/go-finance/fin#ScheduledNetPresentValue)
- [ScheduledInternalRateOfReturn](https://godoc.org/github.com/alpeb/go-finance/fin#ScheduledInternalRateOfReturn)

### TVM

- [PresentValue](https://godoc.org/github.com/alpeb/go-finance/fin#PresentValue)
- [FutureValue](https://godoc.org/github.com/alpeb/go-finance/fin#FutureValue)
- [Payment](https://godoc.org/github.com/alpeb/go-finance/fin#Payment)
- [Periods](https://godoc.org/github.com/alpeb/go-finance/fin#Periods)
- [Rate](https://godoc.org/github.com/alpeb/go-finance/fin#Rate)
- [InterestPayment](https://godoc.org/github.com/alpeb/go-finance/fin#InterestPayment)
- [PrincipalPayment](https://godoc.org/github.com/alpeb/go-finance/fin#PrincipalPayment)

### Bonds

- [DaysDifference](https://godoc.org/github.com/alpeb/go-finance/fin#DaysDifference)
- [DaysPerYear](https://godoc.org/github.com/alpeb/go-finance/fin#DaysPerYear)
- [TBillEquivalentYield](https://godoc.org/github.com/alpeb/go-finance/fin#TBillEquivalentYield)
- [TBillPrice](https://godoc.org/github.com/alpeb/go-finance/fin#TBillPrice)
- [TBillYield](https://godoc.org/github.com/alpeb/go-finance/fin#TBillYield)
- [DiscountRate](https://godoc.org/github.com/alpeb/go-finance/fin#DiscountRate)
- [PriceDiscount](https://godoc.org/github.com/alpeb/go-finance/fin#PriceDiscount)

### Depreciation

- [DepreciationFixedDeclining](https://godoc.org/github.com/alpeb/go-finance/fin#DepreciationFixedDeclining)
- [DepreciationSYD](https://godoc.org/github.com/alpeb/go-finance/fin#DepreciationSYD)
- [DepreciationStraightLine](https://godoc.org/github.com/alpeb/go-finance/fin#DepreciationStraightLine)

## Docs

Checkout the full [docs](https://godoc.org/github.com/alpeb/go-finance/fin).

Also check this excellent blog post [Learning finance concepts using
go-finance](https://blog.aawadia.dev/2020/11/30/finance-concepts-go-finance/)
for some examples and a general intro.

## License

[Mozilla Public License 2.0](https://github.com/alpeb/go-finance/blob/master/LICENSE)
