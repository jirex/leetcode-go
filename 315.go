package main

func main() {

}

func countSmaller(nums []int) []int {
	indexs := make([]int, len(nums))
	for i := range indexs {
		indexs[i] = i
	}

	res := make([]int, len(nums))
	mergeSort(nums, indexs, res)
	return res
}

func mergeSort(nums []int, indexs []int, res []int) {
	if len(indexs) == 1 {
		return
	}

	mid := (len(indexs) - 1) / 2

	left := make([]int, mid+1)
	copy(left, indexs[:mid+1])
	right := make([]int, len(indexs)-mid-1)
	copy(right, indexs[mid+1:])

	mergeSort(nums, left, res)
	mergeSort(nums, right, res)

	idx := 0
	i, j := 0, 0
	for i < len(left) || j < len(right) {
		if j == len(right) {
			indexs[idx] = left[i]
			i++
			idx++
			continue
		}

		if i == len(left) {
			indexs[idx] = right[j]
			j++
			idx++
			continue
		}

		if nums[left[i]] > nums[right[j]] {
			res[left[i]] += len(right) - j
			indexs[idx] = left[i]
			i++
			idx++
			continue
		}

		indexs[idx] = right[j]
		j++
		idx++
	}
	return
}
