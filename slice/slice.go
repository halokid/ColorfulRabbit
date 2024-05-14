package slice

func InSlice[T string | int](item T, sl []T) bool {
	for _, itemx := range sl {
		if item == itemx {
			return true
		}
	}
	return false
}