package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0

	window := map[rune]int{}

	res := 0
	for right < len(s) {
		c := rune(s[right])
		window[c]++
		right++

		for window[c] > 1 {
			window[rune(s[left])]--
			if window[rune(s[left])] == 0 {
				delete(window, rune(s[left]))
			}
			left++
		}

		if len(window) > res {
			res = len(window)
		}
	}
	return res

}
