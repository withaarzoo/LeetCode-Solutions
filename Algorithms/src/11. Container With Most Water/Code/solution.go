func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := 0

    for left < right {
        width := right - left
        h := height[left]
        if height[right] < h {
            h = height[right]
        }
        area := h * width
        if area > maxArea {
            maxArea = area
        }

        // move the pointer that has the smaller height
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}
