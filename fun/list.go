package fun

import (
	"reflect"
	"runtime"
	"sync"

	"github.com/BurntSushi/ty"
)

// Map has a parametric type:
//
//	func Map(f func(A) B, xs []A) []B
//
// Map returns the list corresponding to the return value of applying
// `f` to each element in `xs`.
func Map(f, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) ty.B, []ty.A) []ty.B),
		f, xs)
	vf, vxs, tys := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		vy := call1(vf, vxs.Index(i))
		vys.Index(i).Set(vy)
	}
	return vys.Interface()
}

// Filter has a parametric type:
//
//	func Filter(p func(A) bool, xs []A) []A
//
// Filter returns a new list only containing the elements of `xs` that satisfy
// the predicate `p`.
func Filter(p, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A) []ty.A),
		p, xs)
	vp, vxs, tys := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, 0, xsLen)
	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		if call1(vp, vx).Bool() {
			vys = reflect.Append(vys, vx)
		}
	}
	return vys.Interface()
}

// Foldl has a parametric type:
//
//	func Foldl(f func(A, B) B, init B, xs []A) B
//
// Foldl reduces a list of A to a single element B using a left fold with
// an initial value `init`.
func Foldl(f, init, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A, ty.B) ty.B, ty.B, []ty.A) ty.B),
		f, init, xs)
	vf, vinit, vxs, tb := chk.Args[0], chk.Args[1], chk.Args[2], chk.Returns[0]

	xsLen := vxs.Len()
	vb := zeroValue(tb)
	vb.Set(vinit)
	if xsLen == 0 {
		return vb.Interface()
	}

	vb.Set(call1(vf, vxs.Index(0), vb))
	for i := 1; i < xsLen; i++ {
		vb.Set(call1(vf, vxs.Index(i), vb))
	}
	return vb.Interface()
}

// Foldr has a parametric type:
//
//	func Foldr(f func(A, B) B, init B, xs []A) B
//
// Foldr reduces a list of A to a single element B using a right fold with
// an initial value `init`.
func Foldr(f, init, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A, ty.B) ty.B, ty.B, []ty.A) ty.B),
		f, init, xs)
	vf, vinit, vxs, tb := chk.Args[0], chk.Args[1], chk.Args[2], chk.Returns[0]

	xsLen := vxs.Len()
	vb := zeroValue(tb)
	vb.Set(vinit)
	if xsLen == 0 {
		return vb.Interface()
	}

	vb.Set(call1(vf, vxs.Index(xsLen-1), vb))
	for i := xsLen - 2; i >= 0; i-- {
		vb.Set(call1(vf, vxs.Index(i), vb))
	}
	return vb.Interface()
}

// Concat has a parametric type:
//
//	func Concat(xs [][]A) []A
//
// Concat returns a new flattened list by appending all elements of `xs`.
func Concat(xs interface{}) interface{} {
	chk := ty.Check(
		new(func([][]ty.A) []ty.A),
		xs)
	vxs, tflat := chk.Args[0], chk.Returns[0]

	xsLen := vxs.Len()
	vflat := reflect.MakeSlice(tflat, 0, xsLen*3)
	for i := 0; i < xsLen; i++ {
		vflat = reflect.AppendSlice(vflat, vxs.Index(i))
	}
	return vflat.Interface()
}

// Reverse has a parametric type:
//
//	func Reverse(xs []A) []A
//
// Reverse returns a new slice that is the reverse of `xs`.
func Reverse(xs interface{}) interface{} {
	chk := ty.Check(
		new(func([]ty.A) []ty.A),
		xs)
	vxs, tys := chk.Args[0], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		vys.Index(i).Set(vxs.Index(xsLen - 1 - i))
	}
	return vys.Interface()
}

// Copy has a parametric type:
//
//	func Copy(xs []A) []A
//
// Copy returns a copy of `xs` using Go's `copy` operation.
func Copy(xs interface{}) interface{} {
	chk := ty.Check(
		new(func([]ty.A) []ty.A),
		xs)
	vxs, tys := chk.Args[0], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, xsLen, xsLen)
	reflect.Copy(vys, vxs)
	return vys.Interface()
}

// ParMap has a parametric type:
//
//	func ParMap(f func(A) B, xs []A) []B
//
// ParMap is just like Map, except it applies `f` to each element in `xs`
// concurrently using N worker goroutines (where N is the number of CPUs
// available reported by the Go runtime). If you want to control the number
// of goroutines spawned, use `ParMapN`.
//
// It is important that `f` not be a trivial operation, otherwise the overhead
// of executing it concurrently will result in worse performance than using
// a `Map`.
func ParMap(f, xs interface{}) interface{} {
	n := runtime.NumCPU()
	if n < 1 {
		n = 1
	}
	return ParMapN(f, xs, n)
}

// ParMapN has a parametric type:
//
//	func ParMapN(f func(A) B, xs []A, n int) []B
//
// ParMapN is just like Map, except it applies `f` to each element in `xs`
// concurrently using `n` worker goroutines.
//
// It is important that `f` not be a trivial operation, otherwise the overhead
// of executing it concurrently will result in worse performance than using
// a `Map`.
func ParMapN(f, xs interface{}, n int) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) ty.B, []ty.A) []ty.B),
		f, xs)
	vf, vxs, tys := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	ys := reflect.MakeSlice(tys, xsLen, xsLen)

	if n < 1 {
		n = 1
	}
	work := make(chan int, n)
	wg := new(sync.WaitGroup)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for j := range work {
				// Good golly miss molly. Is `reflect.Value.Index`
				// safe to access/set from multiple goroutines?
				// XXX: If not, we'll need an extra wave of allocation to
				// use real slices of `reflect.Value`.
				ys.Index(j).Set(call1(vf, vxs.Index(j)))
			}
			wg.Done()
		}()
	}
	for i := 0; i < xsLen; i++ {
		work <- i
	}
	close(work)
	wg.Wait()
	return ys.Interface()
}

// Range generates a list of integers corresponding to every integer in
// the half-open interval [x, y).
//
// Range will panic if `end < start`.
func Range(start, end int) []int {
	if end < start {
		panic("range must have end greater than or equal to start")
	}
	r := make([]int, end-start)
	for i := start; i < end; i++ {
		r[i-start] = i
	}
	return r
}

// Each has a parametric type:
//
//  func Each(f func(A), xs []A)
//
// Each runs `f` across each element in `xs`.
func Each(f, xs interface{}) {
	chk := ty.Check(
		new(func(func(ty.A), []ty.A)),
		f, xs)
	vf, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		call(vf, vxs.Index(i))
	}
}

// GroupBy has a parametric type
//
//  func GroupBy(f func(A) B, xs []A) map[B][]A
//
// GroupBy creates a map of return value of f to input element of xs
func GroupBy(f, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) ty.B, []ty.A) map[ty.B][]ty.A),
		f, xs)
	vf, vxs, tys := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vym := reflect.MakeMap(tys)
	for i := 0; i < xsLen; i++ {
		vz := call1(vf, vxs.Index(i))
		mi := vym.MapIndex(vz)
		if !mi.IsValid() {
			vym.SetMapIndex(vz, reflect.MakeSlice(vxs.Type(), 0, 1))
			mi = vym.MapIndex(vz)
		}

		vym.SetMapIndex(vz, reflect.Append(mi, vxs.Index(i)))
	}

	return vym.Interface()
}

// Zip has a parametric type
//
//  func Zip(xs , ys []A) []A
//
// Zip puts the arrays xs and ys together interleaved until the shorter one runs out
func Zip(xs, ys interface{}) interface{} {
	chk := ty.Check(
		new(func([]ty.A, []ty.A) []ty.A),
		xs, ys)
	vxs, vys, vzs := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	ysLen := vys.Len()
	xysLen := xsLen
	if xsLen > ysLen {
		xysLen = ysLen
	}

	zs := reflect.MakeSlice(vzs, xysLen*2, xysLen*2)
	for i := 0; i < xysLen; i++ {
		zs.Index(i * 2).Set(vxs.Index(i))
		zs.Index(i*2 + 1).Set(vys.Index(i))
	}

	return zs.Interface()
}

// Partition has a parametric type
//
//  func Partition(f func(A) bool, xs []A) ([]A, []A)
//
// Partition returns the arrays corresonding to whether the result
// of f returned true or false when called with an element of xs
func Partition(f, xs interface{}) (interface{}, interface{}) {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A) ([]ty.A, []ty.A)),
		f, xs)

	vp, vxs, txs, tys := chk.Args[0], chk.Args[1], chk.Returns[0], chk.Returns[1]

	xsLen := vxs.Len()
	rxs := reflect.MakeSlice(txs, 0, xsLen)
	rys := reflect.MakeSlice(tys, 0, xsLen)

	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		if call1(vp, vx).Bool() {
			rxs = reflect.Append(rxs, vx)
		} else {
			rys = reflect.Append(rys, vx)
		}
	}

	return rxs.Interface(), rys.Interface()
}

// Drop has a parametric type:
//
//  func Drop(f func(A) bool, xs []A) []A
//
// Drop calls f on each element of xs until it returns true, then returns
// that element and the remaining elements of xs
func Drop(f, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A) []ty.A),
		f, xs)
	vp, vxs, txs := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(txs, 0, xsLen)

	found := false
	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		if found || call1(vp, vx).Bool() {
			vys = reflect.Append(vys, vx)
			found = true
		}
	}

	return vys.Interface()
}

// Take has a parametric type:
//
//  func Take(f func(A) bool, xs []A) []A
//
// Take runs f on each element of xs, until f returns true when it
// returns all elements up to the element of xs that f returned true for
func Take(f, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A) []ty.A),
		f, xs)
	vp, vxs, txs := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(txs, 0, xsLen)

	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		if call1(vp, vx).Bool() {
			return vys.Interface()
		}
		vys = reflect.Append(vys, vx)
	}

	return vys.Interface()
}
