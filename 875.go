func minEatingSpeed(piles []int, h int) int {
	i, j := 1, 1000000001

	for i < j {
		mid := i + (j-i)/2
		if findHours(piles, mid) <= h {
			j = mid
		} else {
			i = mid + 1
		}
	}

	return i
}

func findHours(piles []int, k int) int {
	hours := 0
	for _, v := range piles {
		hours += v / k
		if v%k > 0 {
			hours++
		}
	}
	return hours
}
