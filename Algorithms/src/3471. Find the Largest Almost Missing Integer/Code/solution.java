class Solution {
    public int largestInteger(int[] nums, int k) {
        int n = nums.length;

        // nums[i] is between 0 and 50, so a fixed array can store all frequencies.
        // This keeps the extra space constant.
        int[] freq = new int[51];

        // Count how many times every value occurs in nums.
        for (int x : nums) {
            freq[x]++;
        }

        // When k = 1, each subarray contains one element.
        // So only values occurring exactly once are valid.
        if (k == 1) {
            // Check from the largest possible value to the smallest.
            for (int x = 50; x >= 0; x--) {
                if (freq[x] == 1) {
                    return x;
                }
            }

            // No value occurs exactly once.
            return -1;
        }

        // When k = n, the entire array is the only subarray.
        // Therefore, the largest value in nums is the answer.
        if (k == n) {
            int answer = 0;

            // Find the maximum value in the array.
            for (int x : nums) {
                answer = Math.max(answer, x);
            }

            return answer;
        }

        // For 1 < k < n, only the first and last elements
        // can occur in exactly one subarray of size k.
        int answer = -1;

        // The first value is valid only if it appears once in the whole array.
        if (freq[nums[0]] == 1) {
            answer = Math.max(answer, nums[0]);
        }

        // The last value is valid only if it appears once in the whole array.
        if (freq[nums[n - 1]] == 1) {
            answer = Math.max(answer, nums[n - 1]);
        }

        // Return the largest valid candidate, or -1 if none exists.
        return answer;
    }
}