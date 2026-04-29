package p0287

func findDuplicate(nums []int) int {
	n := len(nums)

	// Defend from invalid inputs
	// This guarantees the algorithm never go out of bounds
	// and that there is at least one repeated number
	for _, value := range nums {
		if value < 1 || value > n-1 {
			return -1
		}
	}

	slow, fast, result := 0, 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	for {
		slow = nums[slow]
		result = nums[result]
		if slow == result {
			return result
		}
	}
}
