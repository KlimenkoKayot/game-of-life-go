package support

func B2I(flag bool) int {
	var i int
	if flag {
		i = 1
	} else {
		i = 0
	}
	return i
}
