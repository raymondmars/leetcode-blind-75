package twosum

// problem solving ideas

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := target - nums[i]
		if _, ok := table[n]; ok {
			return []int{table[n], i}
		}
		table[nums[i]] = i
	}
	return []int{}
}
