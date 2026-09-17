class Solution {
public:
    int minSumOfLengths(vector<int>& arr, int target) {
        int n = arr.size();

        // best[i] stores the shortest target-sum sub-array
        // found completely inside arr[0..i].
        vector<int> best(n, INT_MAX);

        int left = 0;       // Left boundary of the sliding window.
        int sum = 0;        // Sum of the current sliding window.
        int answer = INT_MAX;

        for (int right = 0; right < n; ++right) {
            // Extend the window by including arr[right].
            sum += arr[right];

            // Because all values are positive, shrinking from the left
            // is enough whenever the window sum becomes too large.
            while (sum > target) {
                sum -= arr[left];
                ++left;
            }

            // If the current window has exactly the target sum,
            // it is a valid sub-array.
            if (sum == target) {
                int currentLength = right - left + 1;

                // A previous sub-array must end before 'left'.
                // best[left - 1] gives the shortest such sub-array.
                if (left > 0 && best[left - 1] != INT_MAX) {
                    answer = min(answer,
                                 currentLength + best[left - 1]);
                }

                // If this is the first valid sub-array in the prefix,
                // or it is shorter than the previous best, store it.
                best[right] = currentLength;
            }

            // Even when the current position does not end a valid window,
            // keep the best valid sub-array seen anywhere in the prefix.
            if (right > 0) {
                best[right] = min(best[right], best[right - 1]);
            }
        }

        // If two valid non-overlapping sub-arrays were never found,
        // return -1 as required.
        return answer == INT_MAX ? -1 : answer;
    }
};