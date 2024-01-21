# Testing Go for Statistics: MSDS 431

This project was created for the Week 3 assignment in MSDS 431, titled Testing Go for Statistics. anscombe.go is a small program that uses the Anscombe Quartet datasets to calculate their estimated linear regression coefficients. It utilizes the stats library/package created by montanaflynn (github.com/montanaflynn/stats), as well as the fmt and math libraries.

## func roundFloat

I was surprised to learn that neither the Go or the math library had a function that could round to precise decimal places, only functions that rounded to the nearest integer. roundFloat is a function that I found at https://gosamples.dev/round-float/ that fills this need until Go or math add a similar function.

## func LinRegCoef

LinRegCoef accepts a stats.Coordinate struct, which is an array of ordered pairs that must first be passed through stats.LinReg. Then, if x1 - x2 != 0, it calculates the linear regression coefficient, rounds to 4 decimal places using roundFloat, and returns the coefficient. If x1 - x2 = 0, then it loops through the ordered pairs until it finds 2 pairs where x1 - x2 != 0, and then proceeds to calculate and round the coefficient.

## Testing - Python/R vs. Go

In my testing, the Go stats library produced identical linear regression coefficients. 

Python results: 

![Python Results](python-results-1.png "Python Results")