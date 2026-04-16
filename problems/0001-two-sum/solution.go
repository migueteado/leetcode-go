package p0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		diff := target - nums[i]
		found, exists := m[diff]
		if exists {
			return []int{found, i}
		}
		m[nums[i]] = i
	}

	return []int{}
}