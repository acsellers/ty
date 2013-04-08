Package `ty` provides utilities for writing type parametric functions with run
time type safety.

This package contains two sub-packages `fun` and `data` which define some
potentially useful functions and abstractions using the type checker in
this package.

This fork adds 23 more functions over the current parent at the time of writing.
All of those functions are listed at the bottom of this README. Some or all of 
the functions may be upstreamed at some point in the future.

## Requirements

Go tip (or 1.1 when it's released) is required. This package will not work
with Go 1.0.x or earlier.

The very foundation of this package only recently became possible with the
addition of 3 new functions in the standard library `reflect` package:
SliceOf, MapOf and ChanOf. In particular, it provides the ability to
dynamically construct types at run time from component types.

Further extensions to this package can be made if similar functions are added
for structs and functions(?).

## Installation

```bash
go get github.com/BurntSushi/ty
go get github.com/BurntSushi/ty/fun
```

or

```bash
go get github.com/BurntSushi/ty
go get github.com/acsellers/ty/fun
```

## Examples

Squaring each integer in a slice:

```go
square := func(x int) int { return x * x }
nums := []int{1, 2, 3, 4, 5}
squares := Map(square, nums).([]int)
```

Reversing any slice:

```go
slice := []string{"a", "b", "c"}
reversed := Reverse(slice).([]string)
```

Sorting any slice:

```go
// Sort a slice of structs with first class functions.
type Album struct {
  Title string
  Year int
}
albums := []Album{
  {"Born to Run", 1975},
  {"WIESS",       1973},
  {"Darkness",    1978},
  {"Greetings",   1973},
}

less := func(a, b Album) bool { return a.Year < b.Year },
sorted := QuickSort(less, albums).([]Album)
```

Parallel map:

```go
// Compute the prime factorization concurrently
// for every integer in [1000, 10000].
primeFactors := func(n int) []int { // compute prime factors }
factors := ParMap(primeFactors, Range(1000, 10001)).([]int)
```

Asynchronous channel without a fixed size buffer:

```go
s, r := AsyncChan(new(chan int))
send, recv := s.(chan<- int), r.(<-chan int)

// Send as much as you want.
for i := 0; i < 100; i++ {
  s <- i
}
close(s)
for i := range recv {
  // do something with `i`
}
```

Shuffle any slice in place:

```go
jumbleMe := []string{"The", "quick", "brown", "fox"}
Shuffle(jumbleMe)
```

Function memoization:

```go
// Memoizing a recursive function like `fibonacci`.
// Write it like normal:
var fib func(n int64) int64
fib = func(n int64) int64 {
  switch n {
  case 0:
    return 0
  case 1:
    return 1
  }
  return fib(n - 1) + fib(n - 2)
}

// And wrap it with `Memo`.
fib = Memo(fib).(func(int64) int64)

// Will keep your CPU busy for a long time
// without memoization.
fmt.Println(fib(80))
```

## Changes from BurntSushi

This version (acsellers) adds other functions to the current library 
that I wanted or I thought would be interesting to write. 

*  Function Results
  *  All(f, xs)
    * Returns true if every element of xs caused f to return true
  *  Any(f, xs)
    * Returns true if any elements of xs caused f to return true
  *  None(f, xs)
    * Returns true if no elements of xs caused f to return true
  *  One(f, xs)
    * Returns true if exectly one element of xs that caused f to return true
  *  Count(f, xs)
    * Returns the number of elements of xs that caused f to return true
* Iterating
  *  Each(f, xs)
    * Runs a function with no return value on each element of xs
  *  CycleEach(f, xs, n)
    * Runs a function on each element of xs in series n times
  *  CycleMap(f, xs, n)
    * Runs a function on each element of xs in series n times and returns the results
* Grouping
  *  Partition(f, xs)
    * Return two arrays, one in which f returned true for elements of xs and the other where f returned false
  *  GroupBy(f, xs)
    * Return a map of return values of f to the input elements of xs
* Take/Drop
  *  Detect(f, xs)
    * Returns the first element of xs that f returns true for, or nil
  *  Take(f, xs)
    * Returns the series of elements of f that didn't cause f to return true starting from the first element
  *  Drop(f, xs)
    * Returns the series of elements including and following the first element of xs that f returned true for
* Combining
  *  Zip(xs, ys)
    * Take two slices and combine them A, B, A, B until the shorter one runs out
  *  Replace(xs, ys)
    * Replace elements of xs with elements of ys until one runs out
* Mathish
  *  MinInt(f, xs)
    * Run f on each element of xs and return the smallest int that was returned
  *  MinFloat(f, xs)
    * Run f on each element of xs and return the smallest float that was returned
  *  MaxInt(f, xs)
    * Run f on each element of xs and return the largest int that was returned
  *  MaxFloat(f, xs)
    * Run f on each element of xs and return the largest float that was returned
  *  MinMaxInt(f, xs)
    * Run f on each element of xs and return the smallest and largest ints that were returned
  *  MinMaxFloat(f, xs)
    * Run f on each element of xs and return the smallest and largest floats that were returned
  *  SumInt(f, xs)
    * Run f on each element of xs and sum the return values into a single int
  *  SumFloat(f, xs)
    * Run f on each element of xs and sum the return values into a single float
