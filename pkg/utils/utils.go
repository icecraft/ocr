package utils

func MinIntArrWithPos(v []int) (int, int) {
	if v == nil || len(v) == 0 {
		return -1, -1
	}
	var pos int
	n := v[0]
	for i := range v {
		if n > v[i] {
			n = v[i]
			pos = i
		}
	}
	return pos, n
}
