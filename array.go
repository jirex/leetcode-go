func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		s := numbers[left] + numbers[right]
		if s == target {
			return []int{left + 1, right + 1}
		} else if s > target {
			right--
		} else if s < target {
			left++
		}
	}

	return []int{-1, -1}
}