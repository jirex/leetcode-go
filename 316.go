package main

func main() {

}

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}

	stack := []byte{}
	inStack := [26]bool{}

	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if count[last] == 0 {
					break
				}

				stack = stack[:len(stack)-1]
				inStack[last] = false
			}

			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}

		count[ch-'a']--
	}

	return string(stack)
}
