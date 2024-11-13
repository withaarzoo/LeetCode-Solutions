import java.util.Arrays;

class Solution {
    public long countFairPairs(int[] nums, int lower, int upper) {
        Arrays.sort(nums);
        long count = 0;
        int n = nums.length;
        
        for (int i = 0; i < n - 1; i++) {
            int minVal = lower - nums[i];
            int maxVal = upper - nums[i];
            
            // Find the lower bound for the range
            int start = lowerBound(nums, minVal, i + 1);
            // Find the upper bound for the range
            int end = upperBound(nums, maxVal, i + 1);
            
            // Add the number of valid pairs for this i
            count += (end - start);
        }
        
        return count;
    }
    
    // Custom lower bound implementation
    private int lowerBound(int[] nums, int target, int start) {
        int low = start, high = nums.length;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (nums[mid] < target) {
                low = mid + 1;
            } else {
                high = mid;
            }
        }
        return low;
    }
    
    // Custom upper bound implementation
    private int upperBound(int[] nums, int target, int start) {
        int low = start, high = nums.length;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (nums[mid] <= target) {
                low = mid + 1;
            } else {
                high = mid;
            }
        }
        return low;
    }
}