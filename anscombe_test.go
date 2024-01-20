package main

import (
	"github.com/montanaflynn/stats"
	"testing"
)

// ***** BEGIN LinRegCoef TESTING *****

func Test_LinRegCoef_1(t *testing.T) {

	//Test 1 - on first dataset
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
		t.Error("Want 0.5001, got", got)
	}
	
}

func Test_LinRegCoef_2(t *testing.T) {

	//Test 2 - on second dataset
	test2, _ := stats.LinReg(
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

	got := LinRegCoef(test2)
	want := 0.5000

	if got != want {
		t.Error("Want 0.5000, got", got)
	}
	
}

func Test_LinRegCoef_3(t *testing.T) {

	//Test 3 - on third dataset
	test3, _ := stats.LinReg(
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

	got := LinRegCoef(test3)
	want := 0.4997

	if got != want {
		t.Error("Want 0.4997, got", got)
	}
	
}

func Test_LinRegCoef_4(t *testing.T) {

	//Test 4 - on fourth dataset
	test4, _ := stats.LinReg(
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

	got := LinRegCoef(test4)
	want := 0.4999

	if got != want {
		t.Error("Want 0.4999, got", got)
	}
	
}

// *** END LinRegCoef TESTING *****
// ********************************
// *** BEGIN roundFloat TESTING ***

func Test_roundFloat (t *testing.T) {

	got := roundFloat(1.234567, 4)
	want := 1.2346

	if got != want {
		t.Error("Want 1.2346, got", got)
	}
}

// ****END FUNCTION TESTING**********
// **********************************
// ****BEGIN FUNCTION BENCHMARKING***


func BenchmarkLinRegCoef_1(b *testing.B) {

	// Benchmarking LinRegCoef on dataset 1
	benchmark1, _ := stats.LinReg(
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
	for i := 0; i < b.N; i++ {
        LinRegCoef(benchmark1)
    }
}

func BenchmarkLinRegCoef_4(b *testing.B) {

	// Benchmarking LinRegCoef on dataset 4 (the only dataset that will need to use the loop section of LinRegCoef)
	benchmark4, _ := stats.LinReg(
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

	for i := 0; i < b.N; i++ {
        LinRegCoef(benchmark4)
    }
}

// ** END LinRegCoef BENCHMARKING ***
// **********************************
// ** BEGIN roundFloat BENCHMARKING *

func BenchmarkRoundFloat(b *testing.B) {
	num := 5.4637589
	for i := 0; i < b.N; i++ {
        roundFloat(num, 4)
    }
}