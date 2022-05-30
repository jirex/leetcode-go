package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[rune]int{}, map[rune]int{}

	for _, c := range s1 {
		need[c]++
	}

	left, right := 0, 0

	valid := 0

	for right < len(s2) {

		c := rune(s2[right])
		right++

		if need[c] > 0 {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := rune(s2[left])
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}

				window[d]--
			}
		}
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
