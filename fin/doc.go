/*
Package fin is a library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.

In the doc for most of the functions we state the equivalent Excel function.

The time value of money (TVM) functions simply are solutions for each one of the terms of the following equation:
    pv(1+r)^n + pmt(1+r.type)((1+r)^n - 1)/r) + fv = 0
Solving for r (rate) is not possible analytically, so a solution is provided through the Newton-Raphson algorithm.*/
package fin
