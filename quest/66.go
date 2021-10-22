package quest

func plusOne(digits []int) []int {
	cnt := len(digits)
	carry := false
	for i := 1; i <= cnt; i++ {
		if digits[cnt-i] >= 9 {
			carry = true
			digits[cnt-i] = 0
			continue
		} else {
			carry = false
			digits[cnt-i] += 1
			break
		}
	}

	if carry {
		digits = append([]int{1}, digits...)
	}

	return digits
}
