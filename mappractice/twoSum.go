package mappractice

func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int)
	for i, num := range nums {
		if _, ok := valueMap[target-num]; ok {
			return []int{valueMap[target-num], i}
		}
		valueMap[num] = i
	}
	return []int{}
}
