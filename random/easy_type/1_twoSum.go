package easy_type

func twoSum(nums []int, target int) []int {
	var hashMap map[int]int

	hashMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := hashMap[nums[temp]]; ok {
			return []int{nums[temp], i}
		}
		hashMap[nums[i]] = i
	}
	return []int{}
}
