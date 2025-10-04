class Solution {
public:
    int maxArea(vector<int>& height) {
        int left = 0;                      // pointer starting at left end
        int right = (int)height.size() - 1;// pointer starting at right end
        int maxArea = 0;                  // tracks maximum area found

        while (left < right) {
            // width between two pointers
            int width = right - left;
            // height limited by the shorter line
            int h = min(height[left], height[right]);
            // compute area
            int area = h * width;
            if (area > maxArea) maxArea = area;

            // move the pointer at the shorter line inward
            if (height[left] < height[right]) {
                ++left;
            } else {
                --right;
            }
        }
        return maxArea;
    }
};
