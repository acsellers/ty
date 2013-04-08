package fun

import (
	"github.com/BurntSushi/ty"
	"reflect"
)

// CycleEach has a parametric type
//
//  func CycleEach(f func(A), xs []A, n int)
//
// CycleEach calls each element of xs with f in order n times
func CycleEach(f, xs interface{}, n int) {
	chk := ty.Check(
		new(func(func(ty.A), []ty.A, int)),
		f, xs, n)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for t := 0; t < n; t++ {
		for i := 0; i < xsLen; i++ {
			call(vp, vxs.Index(i))
		}
	}
}

// CycleMap has a parametric type
//
//  func CycleMap(f func(A) B, xs []A, n int) []B
//
// CycleMap runs Map n times against xs with f and returns the result
func CycleMap(f, xs interface{}, n int) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) ty.B, []ty.A, int) []ty.B),
		f, xs, n)
	vp, vxs, tys := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, xsLen*n, xsLen*n)
	for t := 0; t < n; t++ {
		for i := 0; i < xsLen; i++ {
			vy := call1(vp, vxs.Index(i))
			vys.Index(t*xsLen + i).Set(vy)
		}
	}
	return vys.Interface()
}
