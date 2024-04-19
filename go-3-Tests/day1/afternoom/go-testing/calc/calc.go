package calc

import "slices"

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func OrdemCres(slc []int) []int {
	nslc := slices.Clone(slc)
	slices.Sort[[]int](nslc)
	return nslc
}
