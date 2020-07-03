package containerwithmostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ar := area(height, left, right)

	for l, r := left, right; l < r-1; {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

		cAr := area(height, l, r)
		if cAr > ar {
			ar = cAr
			left, right = l, r
		}
	}

	return ar
}

func area(height []int, l, r int) int {
	return (r - l) * min(height[l], height[r])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
