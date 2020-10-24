package _map

func Equals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		yv, ok := y[k]
		if !ok || xv != yv {
			return false
		}
	}
	return true
}
