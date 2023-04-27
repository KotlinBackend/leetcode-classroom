package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // 創建一個空的 map，用來存儲元素值和它們的索引
	for i, num := range nums {
		another := target - num          // 計算目標值與當前元素值之差
		if index, ok := m[another]; ok { // 如果差值存在於 map 中，則返回這兩個元素的索引
			return []int{index, i}
		}
		m[num] = i // 否則，將當前元素值和它的索引加入 map 中
	}
	return nil // 如果找不到符合要求的元素，則返回空
}
