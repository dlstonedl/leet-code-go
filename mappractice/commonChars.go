package mappractice

func commonChars(A []string) []string {
	var result []string

	cMap1 := make(map[string]int)
	for _, c := range A[0] {
		if _, ok := cMap1[string(c)]; ok {
			cMap1[string(c)] += 1
		} else {
			cMap1[string(c)] = 1
		}
	}

	cMap2 := make(map[string]int)
	for i := 1; i < len(A); i++ {
		for _, c := range A[i] {
			if _, ok := cMap2[string(c)]; ok {
				cMap2[string(c)] += 1
			} else {
				cMap2[string(c)] = 1
			}
		}

		for k := range cMap1 {
			if _, ok := cMap2[k]; ok {
				cMap1[k] = min(cMap1[k], cMap2[k])
			} else {
				cMap1[k] = 0
			}
		}
		cMap2 = make(map[string]int)
	}

	for k, v := range cMap1 {
		for j := 0; j < v; j++ {
			result = append(result, k)
		}
	}

	return result
}

func min(i int, j int) int {
	if i >= j {
		return j
	}

	return i
}
