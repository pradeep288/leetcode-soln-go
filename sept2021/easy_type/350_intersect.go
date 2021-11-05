package easy_type

func intersect(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	var res []int
	for _, v := range nums1 {
		set[v]++
	}

	for _, v := range nums2 {
		if set[v] > 0 {
			res = append(res, v)
			set[v]--
		}
	}
	return res
}

/*
class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        result = []
        i=j=0
        nums1.sort()
        nums2.sort()
        while i<len(nums1) and j<len(nums2):
            if nums1[i]==nums2[j]:
                result.append(nums1[i])
                i+=1
                j+=1
            elif nums1[i]<nums2[j]:
                i+=1
            elif nums1[i]>nums2[j]:
                j+=1
        return result
*/
