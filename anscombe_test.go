package main

import (
	"github.com/montanaflynn/stats"
	"testing"
)

func Test_LinRegCoef(t *testing.T) {

	test1, _ := stats.LinReg(
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

	got := LinRegCoef(test1)
	want := 0.5001

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}