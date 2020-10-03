package containerwithmostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := (right - left) * min(height[left], height[right])

	for l, r := left, right; l < r-1; {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

		cAr := (r - l) * min(height[l], height[r])
		if cAr > area {
			area = cAr
			left, right = l, r
		}
	}

	return area
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
