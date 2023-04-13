package pkg

func IsSorted(sl []int) bool {
	for i := 0; i < len(sl)-1; i += 1 {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}

func IsSortedRecursion(sl []int) bool {
	if len(sl) == 0 {
		return true
	}
	if len(sl) == 1 {
		return true
	}

	if sl[0] > sl[1] {
		return false
	} else {
		return IsSortedRecursion(sl[1:])
	}
}
