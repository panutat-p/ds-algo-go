package searches

func LinearSearch(li []int, target int) int {
	for k, v := range li {
		if v == target {
			return k
		}
	}
	return -1
}
