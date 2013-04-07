package fun

import "github.com/BurntSushi/ty"

// MinInt has a parametric type:
//
//  func MinInt(f func(A) int64, xs []A) int64
//
// MinInt returns the minimum value returned from f, if the list is
// of length 0, it will return 0
func MinInt(f, xs interface{}) int64 {
	chk := ty.Check(
		new(func(func(ty.A) int64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	if xsLen > 0 {
		min := call1(vp, vxs.Index(0)).Int()
		for i := 1; i < xsLen; i++ {
			vx := vxs.Index(i)
			if lmin := call1(vp, vx).Int(); lmin < min {
				min = lmin
			}
		}
		return min
	}

	return 0
}

// MaxInt has a parametric type:
//
//  func MaxInt(f func(A) int64, xs []A) int64
//
// MaxInt returns the maximum value returned from f, if the list is
// of length 0, it will return 0
func MaxInt(f, xs interface{}) int64 {
	chk := ty.Check(
		new(func(func(ty.A) int64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	if xsLen > 0 {
		max := call1(vp, vxs.Index(0)).Int()
		for i := 1; i < xsLen; i++ {
			vx := vxs.Index(i)
			if lmax := call1(vp, vx).Int(); lmax > max {
				max = lmax
			}
		}
		return max
	}

	return 0
}

// MinMaxInt has a parametric type:
//
//  func MinMaxInt(f func(A) int64, xs []A) (int64, int64)
//
// MinMaxInt returns the minimum and maximum values returned from f, if the list is
// of length 0, it will return 0 and 0
func MinMaxInt(f, xs interface{}) (int64, int64) {
	chk := ty.Check(
		new(func(func(ty.A) int64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	if xsLen > 0 {
		min := call1(vp, vxs.Index(0)).Int()
		max := min
		for i := 1; i < xsLen; i++ {
			vx := vxs.Index(i)
			local := call1(vp, vx).Int()
			if local < min {
				min = local
			}
			if local > max {
				max = local
			}
		}
		return min, max
	}

	return 0, 0
}

// MinFloat has a parametric type:
//
//  func MinFloat(f func(A) float64, xs []A) float64
//
// MinFloat returns the minimum value returned from f, if the list is
// of length 0, it will return 0.0
func MinFloat(f, xs interface{}) float64 {
	chk := ty.Check(
		new(func(func(ty.A) float64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	if xsLen > 0 {
		min := call1(vp, vxs.Index(0)).Float()
		for i := 1; i < xsLen; i++ {
			vx := vxs.Index(i)
			local := call1(vp, vx).Float()
			if local < min {
				min = local
			}
		}
		return min
	}

	return 0.0
}

// MaxFloat has a parametric type:
//
//  func MaxFloat(f func(A) float64, xs []A) float64
//
// MaxFloat returns the minimum value returned from f, if the list is
// of length 0, it will return 0.0
func MaxFloat(f, xs interface{}) float64 {
	chk := ty.Check(
		new(func(func(ty.A) float64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	if xsLen > 0 {
		max := call1(vp, vxs.Index(0)).Float()
		for i := 1; i < xsLen; i++ {
			vx := vxs.Index(i)
			local := call1(vp, vx).Float()
			if local > max {
				max = local
			}
		}
		return max
	}
	return 0.0
}

// MinMaxFloat has a parametric type:
//
//  func MinMaxFloat(f func(A) float64, xs []A) (float64, float64)
//
// MinFloat returns the minimum and maximum values returned from f, if the list is
// of length 0, it will return 0.0 and 0.0
func MinMaxFloat(f, xs interface{}) (float64, float64) {
	chk := ty.Check(
		new(func(func(ty.A) float64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	if xsLen > 0 {
		min := call1(vp, vxs.Index(0)).Float()
		max := min
		for i := 1; i < xsLen; i++ {
			vx := vxs.Index(i)
			local := call1(vp, vx).Float()
			if local < min {
				min = local
			}
			if local > max {
				max = local
			}
		}
		return min, max
	}

	return 0.0, 0.0
}

// SumInt has a parametric type:
//
//  func SumInt(f func(A) int64, xs []A) int64
//
// SumInt returns the sum of the values returned from f
func SumInt(f, xs interface{}) int64 {
	chk := ty.Check(
		new(func(func(ty.A) int64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	var sum int64
	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		sum += call1(vp, vx).Int()
	}

	return sum
}

// SumFloat has a parametric type:
//
//  func SumFloat(f func(A) float64, xs []A) float64
//
// SumFloat returns the sum of the values returned from f
func SumFloat(f, xs interface{}) float64 {
	chk := ty.Check(
		new(func(func(ty.A) float64, []ty.A)),
		f, xs)
	vp, vxs := chk.Args[0], chk.Args[1]

	xsLen := vxs.Len()
	var sum float64
	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		sum += call1(vp, vx).Float()
	}

	return sum
}
