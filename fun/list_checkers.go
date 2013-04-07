package fun

import "github.com/BurntSushi/ty"

// All has a parametric type:
//
//  func All(f func(A) bool, xs []A) bool
//
// All returns whether all invocations of f on elements of xs
// return true
func All(f, xs interface{}) bool {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		if !call1(vp, vxs.Index(i)).Bool() {
			return false
		}
	}
	return true
}

// Any has a parametric type:
//
//  func Any(f func(A) bool, xs []A) bool
//
// Any returns whether any invocations of f on elements of xs
// return true
func Any(f, xs interface{}) bool {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		if call1(vp, vxs.Index(i)).Bool() {
			return true
		}
	}
	return false
}
