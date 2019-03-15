package mappractice

import "strings"

func numJewelsInStones(J string, S string) int {
	result := 0
	jMap := make(map[byte]int)
	for e := range J {
		jMap[J[e]] = 1
	}
	for e := range S {
		if _, ok := jMap[S[e]]; ok {
			result = result + 1
		}
	}
	return result
}

func numJewelsInStonesThoughStringsIndex(J string, S string) int {
	result := 0
	for _, v := range S {
		if strings.IndexRune(J, v) != -1 {
			result += 1
		}
	}
	return result
}
