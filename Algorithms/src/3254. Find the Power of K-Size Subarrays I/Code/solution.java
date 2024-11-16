import java.util.*;

class Solution {
    public int[] resultsArray(int[] nums, int k) {
        int n = nums.length;
        int[] result = new int[n - k + 1];

        for (int i = 0; i <= n - k; i++) {
            int[] subarray = Arrays.copyOfRange(nums, i, i + k);

            // Sort the subarray
            int[] sortedSubarray = subarray.clone();
            Arrays.sort(sortedSubarray);

            // Check if elements are consecutive
            boolean isConsecutive = true;
            for (int j = 1; j < k; j++) {
                if (sortedSubarray[j] - sortedSubarray[j - 1] != 1) {
                    isConsecutive = false;
                    break;
                }
            }

            // Add the result based on conditions
            if (isConsecutive && Arrays.equals(subarray, sortedSubarray)) {
                result[i] = sortedSubarray[k - 1]; // Max element
            } else {
                result[i] = -1;
            }
        }

        return result;
    }
}