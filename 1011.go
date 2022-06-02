package main

func main() {

}

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 1

	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	for left < right {
		mid := left + (right-left)/2

		if f(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func f(w []int, x int) int {
	days := 0
	for i := 0; i < len(w); {
		cap := x
		for i < len(w) {
			if cap < w[i] {
				break
			} else {
				cap -= w[i]
			}
			i++
		}
		days++
	}
	return days
}
