# Testing Go for Statistics: MSDS 431

This project was created for the Week 3 assignment in MSDS 431, titled Testing Go for Statistics. anscombe.go is a small program that uses the Anscombe Quartet datasets to calculate their estimated linear regression coefficients. It utilizes the stats library/package created by montanaflynn (github.com/montanaflynn/stats), as well as the fmt and math libraries.

There are 2 custom functions in anscombe.go: LinRegCoef and roundFloat.

## LinRegCoef

LinRegCoef accepts a stats.Coordinate struct, which is an array of ordered pairs that must first be passed through stats.LinReg. 