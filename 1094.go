func carPooling(trips [][]int, capacity int) bool {
	pots := make([]int, 1001)

	for _, p := range trips {
		pots[p[1]] += p[0]
		pots[p[2]] -= p[0]
	}

	if pots[0] > capacity {
		return false
	}
	for i := 1; i < 1001; i++ {
		pots[i] += pots[i-1]

		if pots[i] > capacity {
			return false
		}
	}

	return true
}
