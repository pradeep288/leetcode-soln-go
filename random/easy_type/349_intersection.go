package easy_type

func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]struct{})
	var res []int

	for _, v := range nums1 {
		hashMap[v] = struct{}{}
	}

	for _, v := range nums2 {
		if _, ok := hashMap[v]; ok {
			res = append(res, v)
			delete(hashMap, v)
		}
	}

	return res

}
