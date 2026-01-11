func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	largestRectangleArea := func(h []int) int {
		stack := []int{}
		h = append(h, 0)
		area := 0

		for i := 0; i < len(h); i++ {
			for len(stack) > 0 && h[stack[len(stack)-1]] > h[i] {
				height := h[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				width := i
				if len(stack) > 0 {
					width = i - stack[len(stack)-1] - 1
				}
				area = max(area, height*width)
			}
			stack = append(stack, i)
		}
		return area
	}

	for _, row := range matrix {
		for j := 0; j < cols; j++ {
			if row[j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
