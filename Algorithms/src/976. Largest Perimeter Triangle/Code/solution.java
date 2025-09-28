import java.util.Arrays;

class Solution {
    public int largestPerimeter(int[] nums) {
        // Sort array in non-decreasing order
        Arrays.sort(nums);
        int n = nums.length;
        // Iterate from the end to the start checking triples
        for (int i = n - 1; i >= 2; --i) {
            int a = nums[i];       // largest side in the triple
            int b = nums[i - 1];
            int c = nums[i - 2];
            // If two smaller sides sum greater than the largest side -> valid triangle
            if (b + c > a) return a + b + c;
        }
        return 0; // no valid triangle found
    }
}
