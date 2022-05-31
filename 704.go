package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1

	for i <= j {
		mid := i + (j-i)/2
		n := nums[mid]
		if n == target {
			return mid
		} else if n < target {
			i = mid + 1
		} else if n > target {
			j = mid - 1
		}
	}

	return -1
}
