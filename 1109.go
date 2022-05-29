func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for _, b := range bookings {
		diff[b[0]-1] += b[2]
		if b[1] < n {
			diff[b[1]] -= b[2]
		}
	}
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff
}
