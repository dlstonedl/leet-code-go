package mappractice

func repeatedNTimes(A []int) int {
	vMap := make(map[int]bool)
	for _, a := range A {
		if value, ok := vMap[a]; value && ok {
			return a
		}
		vMap[a] = true
	}
	return -1
}
