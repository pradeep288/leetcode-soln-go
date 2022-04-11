package medium_type

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hashMap := make(map[int]int)
	res := 0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			hashMap[nums1[i]+nums2[j]] += 1
		}
	}

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums2); j++ {
			s := (nums3[i] + nums4[j]) * -1
			if _, ok := hashMap[s]; ok {
				res += hashMap[s]
			}
		}
	}
	return res
}
