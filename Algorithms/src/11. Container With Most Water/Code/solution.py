class Solution:
    def maxArea(self, height: List[int]) -> int:
        left = 0                      # left pointer index
        right = len(height) - 1       # right pointer index
        max_area = 0                  # best area seen

        while left < right:
            width = right - left
            h = min(height[left], height[right])
            area = h * width
            if area > max_area:
                max_area = area

            # move pointer at smaller height inward
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1

        return max_area
