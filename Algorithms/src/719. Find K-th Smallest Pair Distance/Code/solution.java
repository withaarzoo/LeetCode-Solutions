import java.util.Arrays;

class Solution {
    // This method counts the number of pairs (i, j) such that the difference
    // between nums[j] and nums[i] is less than or equal to 'mid'.
    private int countPairs(int[] nums, int mid) {
        int count = 0; // Initialize the count of pairs to zero.

        // Use two pointers i and j to find the pairs.
        for (int i = 0, j = 0; i < nums.length; ++i) {
            // Move pointer j to the right as long as the difference between nums[j] and
            // nums[i] is less than or equal to 'mid'.
            while (j < nums.length && nums[j] - nums[i] <= mid) {
                ++j; // Increment j to expand the window of pairs.
            }
            // The number of valid pairs with nums[i] as the first element is (j - i - 1).
            count += j - i - 1;
        }
        return count; // Return the total count of valid pairs.
    }

    // This method finds the k-th smallest distance pair.
    public int smallestDistancePair(int[] nums, int k) {
        Arrays.sort(nums); // Sort the array to allow binary search and two-pointer technique.

        int low = 0; // Initialize the lowest possible distance (difference between identical
                     // elements).
        int high = nums[nums.length - 1] - nums[0]; // Initialize the highest possible distance (difference between max
                                                    // and min elements).

        // Use binary search to find the smallest distance such that there are at least
        // 'k' pairs with that distance.
        while (low < high) {
            int mid = (low + high) / 2; // Calculate the middle value of the current search range.

            // Check if there are at least 'k' pairs with a distance less than or equal to
            // 'mid'.
            if (countPairs(nums, mid) >= k) {
                high = mid; // If yes, search in the lower half.
            } else {
                low = mid + 1; // If no, search in the upper half.
            }
        }

        return low; // 'low' will be the smallest distance with at least 'k' pairs.
    }
}
