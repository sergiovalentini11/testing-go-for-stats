package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"math"
)

func LinRegCoef(data []stats.Coordinate) float64 {
// LinRegCoef takes in a Coordinate struct, which is a set of ordered pairs run through stats.LinReg.
// Then it iterates through the ordered pairs to find 2 consecutive ordered pairs where x1 - x2 != 0,
// and finally calculates and returns the linear regression coefficient. 
var coeff float64 = 0.0
	for i := 0; i < 10; i++ {
		if data[i].X - data[i+1].X == 0 {
			continue
		} else {
			coeff = (data[i].Y - data[i+1].Y) / (data[i].X - data[i+1].X)
			break
		}
	}
	return coeff
}

// found a helpful rounding function as it doesn't appear that exists in Go yet - credit: https://gosamples.dev/round-float/
func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func main() {

	// creating the Coordinate structs for the ordered pair data
	r1, _ := stats.LinReg(
		[]stats.Coordinate{
			{10.0, 8.04},
			{8.0, 6.95},
			{13.0, 7.58},
			{9.0, 8.81},
			{11.0, 8.33},
			{14.0, 9.96},
			{6.0, 7.24},
			{4.0, 4.26},
			{12.0, 10.84},
			{7.0, 4.82},
			{5.0, 5.68},
		},
	)

	r2, _ := stats.LinReg(
		[]stats.Coordinate{
			{10.0, 9.14},
			{8.0, 8.14},
			{13.0, 8.74},
			{9.0, 8.77},
			{11.0, 9.26},
			{14.0, 8.1},
			{6.0, 6.13},
			{4.0, 3.1},
			{12.0, 9.13},
			{7.0, 7.26},
			{5.0, 4.74},
		},
	)

	r3, _ := stats.LinReg(
		[]stats.Coordinate{
			{10.0, 7.46},
			{8.0, 6.77},
			{13.0, 12.74},
			{9.0, 7.11},
			{11.0, 7.81},
			{14.0, 8.84},
			{6.0, 6.08},
			{4.0, 5.39},
			{12.0, 8.15},
			{7.0, 6.42},
			{5.0, 5.73},
		},
	)

	r4, _ := stats.LinReg(
		[]stats.Coordinate{
			{8.0, 6.58},
			{8.0, 5.76},
			{8.0, 7.71},
			{8.0, 8.84},
			{8.0, 8.47},
			{8.0, 7.04},
			{8.0, 5.25},
			{19.0, 12.5},
			{8.0, 5.56},
			{8.0, 7.91},
			{8.0, 6.89},
		},
	)

	// running LinRegCoef to get the regression coefficients for each dataset, and printing them
	r1_coeff := LinRegCoef(r1)
	r1_coeff = roundFloat(r1_coeff, 4)
	fmt.Println(r1_coeff)

	r2_coeff := LinRegCoef(r2)
	r2_coeff = roundFloat(r2_coeff, 4)
	fmt.Println(r2_coeff)

	r3_coeff := LinRegCoef(r3)
	r3_coeff = roundFloat(r3_coeff, 4)
	fmt.Println(r3_coeff)

	r4_coeff := LinRegCoef(r4)
	r4_coeff = roundFloat(r4_coeff, 4)
	fmt.Println(r4_coeff)

}
