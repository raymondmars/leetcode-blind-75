package twosum

// description refer: https://leetcode.com/problems/two-sum/descriptoin

/* problem solving ideas
1, To fetch each number from the given array.
2, Store the number in a hash table.
3, Use the target to subtract the fetched number,
if the result of the calculation has existed in the hash table,
It proves there are two numbers in the array added together equal to the target.
4, Return both of the indexes of the two numbers in the array.
*/
func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := target - nums[i]
		if index, ok := table[n]; ok {
			return []int{index, i}
		}
		table[nums[i]] = i
	}
	return []int{}
}
