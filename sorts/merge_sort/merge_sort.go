package merge_sort

func MergeSort(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	}

	left, right := Split(sl)
	left = MergeSort(left)
	right = MergeSort(right)
	return Merge(left, right)
}

func Split(sl []int) ([]int, []int) {
	return sl[:len(sl)/2], sl[len(sl)/2:]
}

func Merge(left []int, right []int) []int {
	var sl []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sl = append(sl, left[i])
			i += 1
		} else {
			sl = append(sl, right[j])
			j += 1
		}
	}

	sl = append(sl, left[i:]...)
	sl = append(sl, right[j:]...)

	return sl
}
