type Solution struct {
	pre []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (this *Solution) PickIndex() int {
	x := rand.Intn(this.pre[len(this.pre)-1]) + 1
	return sort.SearchInts(this.pre, x)
}

func Left_bound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			left = mid + 1
		} else if nums[mid] < target {
			right = mid
		}
	}

	return left
}
