package fun

import "testing"

func TestCycleEach(t *testing.T) {
	ary := make([]int, 0)
	i := 0
	f := func(n int) { ary = append(ary, i*n); i++ }

	CycleEach(f, []int{1, 2, 3}, 2)
	assertDeep(t, ary, []int{0, 2, 6, 3, 8, 15})

	ary = make([]int, 0)
	CycleEach(f, []int{}, 4)
	assertDeep(t, ary, []int{})

	ary = make([]int, 0)
	CycleEach(f, []int{1, 2}, 0)
	assertDeep(t, ary, []int{})
}

func TestCycleMap(t *testing.T) {
	square := func(n int) int { return n * n }
	a := []int{1, 2, 3}
	r := CycleMap(square, a, 2).([]int)
	assertDeep(t, r, []int{1, 4, 9, 1, 4, 9})

	r = CycleMap(square, []int{}, 1).([]int)
	assertDeep(t, r, []int{})
}
