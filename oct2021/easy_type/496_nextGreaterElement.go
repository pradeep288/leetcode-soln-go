package easy_type

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	// Create a HashMap with elements of nums1 as key and their index as value
	for i, v := range nums2 {
		hashMap[v] = i
	}

	res := make([]int, 0)

	for _, val := range nums1 {
		start := hashMap[val] + 1
		updated := false
		for i := start; i < len(nums2); i++ {
			if nums2[i] > val {
				res = append(res, nums2[i])
				updated = true
				break
			}
		}
		if !updated {
			res = append(res, -1)
		}
	}

	return res

}
