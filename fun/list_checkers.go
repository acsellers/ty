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

// Count has a parametric type:
//
//  func Count(f func(A) bool, xs []A) int
//
// Count returns the number of elements of xs for which f
// returns true
func Count(f, xs interface{}) (matches int) {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		if call1(vp, vxs.Index(i)).Bool() {
			matches++
		}
	}
	return
}

// Detect has a parametric type:
//
//  func Detect(f func(A) bool, xs []A) A
//
// Detect returns the first element for which f returns
// true, if none are returned it returns nil
func Detect(f, xs interface{}) interface{} {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		if call1(vp, vxs.Index(i)).Bool() {
			return vxs.Index(i).Interface()
		}
	}
	return nil
}

// None has a parametric type
//
//  func None(f func(A) bool, xs []A) bool
//
// None returns true if none of the elements in xs caused f to return
// true, false otherwise
func None(f, xs interface{}) bool {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	for i := 0; i < xsLen; i++ {
		if call1(vp, vxs.Index(i)).Bool() {
			return false
		}
	}
	return true
}

// One has a parametric type
//
//  func One(f func(A) bool, xs []A) bool
//
// One returns whether exactly one of the elements in xs
// caused f to return true
func One(f, xs interface{}) bool {
	chk := ty.Check(
		new(func(func(ty.A) bool, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	first := false
	for i := 0; i < xsLen; i++ {
		if call1(vp, vxs.Index(i)).Bool() {
			if first {
				return false
			} else {
				first = true
			}
		}
	}
	return first
}
