package fun

import "testing"

func TestReplace(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{5, 4, 3}
	c := Replace(a, b).([]int)
	assertDeep(t, c, []int{5, 4, 3, 4, 5})

	d := Replace(b, a).([]int)
	assertDeep(t, d, []int{1, 2, 3})
}
