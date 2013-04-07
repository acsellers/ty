package fun

import "testing"

func TestAll(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	allEven := All(even, []int{1, 2, 3, 4, 5, 6})
	if allEven {
		t.Fatal("TestAll: list are not all even")
	}

	lessThanTen := func(x int) bool { return x < 10 }
	less := All(lessThanTen, []int{1, 2, 3, 4, 5})
	if !less {
		t.Fatal("TestAll: list is all less than 10")
	}
}

func TestAny(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	allEven := Any(even, []int{1, 2, 3, 4, 5, 6})
	assertDeep(t, allEven, true)

	greaterThanTen := func(x int) bool { return x > 10 }
	greater := Any(greaterThanTen, []int{1, 2, 3, 4, 5})
	assertDeep(t, greater, false)
}

func TestCount(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	countEven := Count(even, []int{1, 2, 3, 4, 5, 6})
	assertDeep(t, countEven, 3)

	countEven = Count(even, []int{})
	assertDeep(t, countEven, 0)
}

func TestDetect(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	firstEven := Detect(even, []int{1, 2, 3, 4, 5, 6})
	assertDeep(t, firstEven.(int), 2)

	firstEven = Detect(even, []int{})
	assertDeep(t, firstEven, nil)
}
