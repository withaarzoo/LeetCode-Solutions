class Solution {
    public int minSumOfLengths(int[] arr, int target) {
        int n = arr.length;

        // best[i] stores the shortest target-sum sub-array
        // found completely inside arr[0..i].
        int[] best = new int[n];

        // Use a large value to represent "no valid sub-array".
        java.util.Arrays.fill(best, Integer.MAX_VALUE);

        int left = 0;              // Left boundary of the sliding window.
        int sum = 0;                // Sum of the current sliding window.
        int answer = Integer.MAX_VALUE;

        for (int right = 0; right < n; right++) {
            // Add the new element to the current window.
            sum += arr[right];

            // Since all values are positive, move left forward
            // until the window sum becomes at most target.
            while (sum > target) {
                sum -= arr[left];
                left++;
            }

            // A valid target-sum sub-array ends at right.
            if (sum == target) {
                int currentLength = right - left + 1;

                // The previous sub-array must finish before left.
                if (left > 0 && best[left - 1] != Integer.MAX_VALUE) {
                    answer = Math.min(
                        answer,
                        currentLength + best[left - 1]
                    );
                }

                // Store this valid sub-array as the current best
                // for the prefix ending at right.
                best[right] = currentLength;
            }

            // Carry forward the best sub-array found in the prefix
            // even if no valid sub-array ends exactly at right.
            if (right > 0) {
                best[right] = Math.min(best[right], best[right - 1]);
            }
        }

        // Return -1 when two non-overlapping valid sub-arrays do not exist.
        return answer == Integer.MAX_VALUE ? -1 : answer;
    }
}