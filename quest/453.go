package quest

func minMoves(nums []int) int {
	//min == max
	var minIdx = 0
	var maxIdx = 0

	for idx := range nums {
		if nums[idx] < nums[minIdx] {
			minIdx = idx
		}
		if nums[idx] > nums[maxIdx] {
			maxIdx = idx
		}
	}
	//其实就是选一个最大的数减1，直到全减到和最小的一样
	count := 0
	for idx := range nums {
		if idx == minIdx {
			continue
		}
		count += nums[idx] - nums[minIdx]
	}
	return count
}
