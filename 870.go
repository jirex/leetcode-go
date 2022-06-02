import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)

	pairs2 := make([]pair, len(nums2))
	for k,v := range nums2 {
		pairs2[k].k = k
		pairs2[k].v = v
	}

	sort.Slice(pairs2, func(i,j int) bool {return pairs2[i].v < pairs2[j].v})

	result := make([]int, len(nums1))

	min, max := 0, len(nums2)-1

	for _, v1 := range nums1 {
		if v1 > pairs2[min].v {
			result[pairs2[min].k] = v1
			min++
		}else{
			result[pairs2[max].k] = v1
			max--
		}
	}

	return result
}

type pair struct {
	k, v int
}
