package fun

import (
	"github.com/BurntSushi/ty"
	"reflect"
)

// Replace has a parametric type
//
//  func Replace(xs, ys []A) []A
//
// Replace changes elements of xs with elements of ys until one array is exhausted
func Replace(xs, ys interface{}) interface{} {
	chk := ty.Check(
		new(func([]ty.A, []ty.A) []ty.A),
		xs, ys)
	vxs, vys, tzs := chk.Args[0], chk.Args[1], chk.Returns[0]

	xsLen, ysLen := vxs.Len(), vys.Len()
	vzs := reflect.MakeSlice(tzs, xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		if i < ysLen {
			vzs.Index(i).Set(vys.Index(i))
		} else {
			vzs.Index(i).Set(vxs.Index(i))
		}
	}

	return vzs.Interface()
}
