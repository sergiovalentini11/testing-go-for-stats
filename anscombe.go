package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	// "gonum.org/v1/plot"
)

func LinRegCoef(data []stats.Coordinate) {

	for i := 0; i < 10; i++ {
		if data[i].X-data[i+1].X == 0 {
			continue
		} else {
			coeff := (data[i].Y - data[i+1].Y) / (data[i].X - data[i+1].X)
			fmt.Println("")
			fmt.Printf("Regression coefficient = %.4f", coeff)
			fmt.Println("")
			break
		}
	}
}

func main() {

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

	LinRegCoef(r1)
	LinRegCoef(r2)
	LinRegCoef(r3)
	LinRegCoef(r4)

}
