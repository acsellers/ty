package fun

import "testing"

func TestMinInt(t *testing.T) {
	square := func(x int64) int64 { return x * x }
	min := MinInt(square, []int64{1, 2, 3})
	if min != 1 {
		t.Fatal("TestMinInt: 1 should be the lowest")
	}

	min = MinInt(square, []int64{})
	if min != 0 {
		t.Fatal("TestMinInt: 0 should be returned from 0 length array")
	}
}

func TestMaxInt(t *testing.T) {
	square := func(x int64) int64 { return x * x }
	max := MaxInt(square, []int64{1, 2, 3})
	if max != 9 {
		t.Fatal("TestMaxInt: 9 should be the highest")
	}

	max = MaxInt(square, []int64{})
	if max != 0 {
		t.Fatal("TestMaxInt: 0 should be returned from 0 length array")
	}
}

func TestMinMaxInt(t *testing.T) {
	square := func(x int64) int64 { return x * x }
	min, max := MinMaxInt(square, []int64{1, 2, 3})
	if min != 1 {
		t.Fatal("TestMinMaxInt: 1 should be the lowest")
	}
	if max != 9 {
		t.Fatal("TestMinMaxInt: 9 should be the lowest")
	}

	min, max = MinMaxInt(square, []int64{})
	if min != 0 {
		t.Fatal("TestMinMaxInt: 0 should be returned from 0 length array")
	}
	if max != 0 {
		t.Fatal("TestMinMaxInt: 0 should be returned from 0 length array")
	}
}

func TestMinFloat(t *testing.T) {
	square := func(x float64) float64 { return x * x }
	min := MinFloat(square, []float64{1.0, 2.0, 3.0})
	if min != 1.0 {
		t.Fatal("TestMinFloat: 1.0 should be the lowest")
	}

	min = MinFloat(square, []float64{})
	if min != 0.0 {
		t.Fatal("TestMinFloat: 0 should be returned from 0 length array")
	}
}

func TestMaxFloat(t *testing.T) {
	square := func(x float64) float64 { return x * x }
	max := MaxFloat(square, []float64{1.0, 2.0, 3.0})
	if max != 9.0 {
		t.Fatal("TestMaxFloat: 9.0 should be the highest")
	}

	max = MaxFloat(square, []float64{})
	if max != 0 {
		t.Fatal("TestMaxFloat: 0 should be returned from 0 length array")
	}
}

func TestMinMaxFloat(t *testing.T) {
	square := func(x float64) float64 { return x * x }
	min, max := MinMaxFloat(square, []float64{1.0, 2.0, 3.0})
	if min != 1.0 {
		t.Fatal("TestMinMaxFloat: 1.0 should be the lowest")
	}
	if max != 9 {
		t.Fatal("TestMinMaxFloat: 9.0 should be the lowest")
	}

	min, max = MinMaxFloat(square, []float64{})
	if min != 0 {
		t.Fatal("TestMinMaxFloat: 0 should be returned from 0 length array")
	}
	if max != 0 {
		t.Fatal("TestMinMaxFloat: 0 should be returned from 0 length array")
	}
}
