package pkg

func IsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}
