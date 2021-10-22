package quest

//得看看双指针的正确性证明
func maxArea(height []int) int {
	cnt := len(height)
	i := 0
	j := cnt - 1
	max := 0
	for i != j {
		h := 0
		w := j - i
		if height[i] < height[j] {
			h = height[i]
			i += 1
		} else {
			h = height[j]
			j -= 1
		}
		area := h * w
		if area > max {
			max = area
		}
	}
	return max
}
