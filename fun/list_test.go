package fun

import (
	"testing"
)

func TestMap(t *testing.T) {
	square := func(x int) int { return x * x }
	squares := Map(square, []int{1, 2, 3, 4, 5}).([]int)

	assertDeep(t, squares, []int{1, 4, 9, 16, 25})
	assertDeep(t, []int{}, Map(square, []int{}).([]int))

	strlen := func(s string) int { return len(s) }
	lens := Map(strlen, []string{"abc", "ab", "a"}).([]int)
	assertDeep(t, lens, []int{3, 2, 1})
}

func TestEach(t *testing.T) {
	results := make([]string, 0)
	greet := func(n string) { results = append(results, "Hey "+n) }
	Each(greet, []string{"a", "b", "c"})

	assertDeep(t, results, []string{"Hey a", "Hey b", "Hey c"})

	results = make([]string, 0)
	Each(greet, []string{})
	assertDeep(t, []string{}, results)
}

func TestFilter(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	evens := Filter(even, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).([]int)

	assertDeep(t, evens, []int{2, 4, 6, 8, 10})
	assertDeep(t, []int{}, Filter(even, []int{}).([]int))
}

func TestFoldl(t *testing.T) {
	// Use an operation that isn't associative so that we know we've got
	// the left/right folds done correctly.
	reducer := func(a, b int) int { return b % a }
	v := Foldl(reducer, 7, []int{4, 5, 6}).(int)

	assertDeep(t, v, 3)
	assertDeep(t, 0, Foldl(reducer, 0, []int{}).(int))
}

func TestFoldr(t *testing.T) {
	// Use an operation that isn't associative so that we know we've got
	// the left/right folds done correctly.
	reducer := func(a, b int) int { return b % a }
	v := Foldr(reducer, 7, []int{4, 5, 6}).(int)

	assertDeep(t, v, 1)
	assertDeep(t, 0, Foldr(reducer, 0, []int{}).(int))
}

func TestConcat(t *testing.T) {
	toflat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	flat := Concat(toflat).([]int)

	assertDeep(t, flat, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestReverse(t *testing.T) {
	reversed := Reverse([]int{1, 2, 3, 4, 5}).([]int)

	assertDeep(t, reversed, []int{5, 4, 3, 2, 1})
}

func TestZip(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	zipped := Zip(a, b)
	assertDeep(t, zipped, []int{1, 2, 3, 4, 5, 6})

	c := []int{1, 2, 3}
	d := []int{}
	zipped = Zip(c, d)
	assertDeep(t, zipped, []int{})

	e := []string{"a", "b", "c", "d"}
	f := []string{"1", "2"}
	zipped = Zip(e, f)
	assertDeep(t, zipped, []string{"a", "1", "b", "2"})
}

func TestPartition(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	even := func(i int) bool { return i%2 == 0 }
	evens, odds := Partition(even, a)
	assertDeep(t, evens, []int{2, 4})
	assertDeep(t, odds, []int{1, 3, 5})

	evens, odds := Partition(even, []int{})
	assertDeep(t, evens, []int{})
	assertDeep(t, odds, []int{})
}

func TestCopy(t *testing.T) {
	orig := []int{1, 2, 3, 4, 5}
	copied := Copy(orig).([]int)

	orig[1] = 999

	assertDeep(t, copied, []int{1, 2, 3, 4, 5})
}

func TestPointers(t *testing.T) {
	type temp struct {
		val int
	}
	square := func(t *temp) *temp { return &temp{t.val * t.val} }
	squares := Map(square, []*temp{
		{1}, {2}, {3}, {4}, {5},
	})

	assertDeep(t, squares, []*temp{
		{1}, {4}, {9}, {16}, {25},
	})
}

func BenchmarkMapSquare(b *testing.B) {
	if flagBuiltin {
		benchmarkMapSquareBuiltin(b)
	} else {
		benchmarkMapSquareReflect(b)
	}
}

func benchmarkMapSquareReflect(b *testing.B) {
	b.StopTimer()
	square := func(a int64) int64 { return a * a }
	list := randInt64Slice(1000, 1<<30)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Map(square, list).([]int64)
	}
}

func benchmarkMapSquareBuiltin(b *testing.B) {
	b.StopTimer()
	square := func(a int64) int64 { return a * a }
	list := randInt64Slice(1000, 1<<30)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ret := make([]int64, len(list))
		for i := 0; i < len(list); i++ {
			ret[i] = square(list[i])
		}
	}
}

func BenchmarkMapPrime(b *testing.B) {
	if flagBuiltin {
		benchmarkMapPrimeBuiltin(b)
	} else {
		benchmarkMapPrimeReflect(b)
	}
}

func benchmarkMapPrimeReflect(b *testing.B) {
	b.StopTimer()
	list := randInt64Slice(1000, 1<<30)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Map(primeFactors, list).([][]int64)
	}
}

func benchmarkMapPrimeBuiltin(b *testing.B) {
	b.StopTimer()
	list := randInt64Slice(1000, 1<<30)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ret := make([][]int64, len(list))
		for i := 0; i < len(list); i++ {
			ret[i] = primeFactors(list[i])
		}
	}
}
