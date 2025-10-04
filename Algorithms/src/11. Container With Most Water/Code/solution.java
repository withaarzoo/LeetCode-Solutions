class Solution {
    public int maxArea(int[] height) {
        int left = 0;                     // left pointer
        int right = height.length - 1;    // right pointer
        int maxArea = 0;                  // best area so far

        while (left < right) {
            int width = right - left;
            int h = Math.min(height[left], height[right]);
            int area = h * width;
            if (area > maxArea) maxArea = area;

            // move pointer at the smaller height
            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
        }
        return maxArea;
    }
}
