class Solution {
    // A DP state stores the best score and the selected original indices.
    static class Node {
        long score;      // Total weight of the selected intervals.
        int[] ids;       // Original indices in sorted order.
        boolean valid;   // True when this state is reachable.

        Node() {
            this.valid = false;
        }

        Node(long score, int[] ids) {
            this.score = score;
            this.ids = ids;
            this.valid = true;
        }
    }

    // Returns true if a is better than b.
    private boolean better(Node a, Node b) {
        // Any valid candidate beats an invalid state.
        if (!a.valid) return false;
        if (!b.valid) return true;

        // Maximum score comes first.
        if (a.score != b.score) return a.score > b.score;

        // Equal scores are resolved using lexicographical order.
        int len = Math.min(a.ids.length, b.ids.length);
        for (int i = 0; i < len; i++) {
            if (a.ids[i] != b.ids[i]) {
                return a.ids[i] < b.ids[i];
            }
        }

        // If one is a prefix of the other, the shorter one is smaller.
        return a.ids.length < b.ids.length;
    }

    // Inserts one index into an already sorted array.
    private int[] addSorted(int[] ids, int value) {
        // Create space for the new index.
        int[] result = Arrays.copyOf(ids, ids.length + 1);

        // Put the new value at the end first.
        result[ids.length] = value;

        // Since there are at most four indices, a small sort is enough.
        Arrays.sort(result);

        return result;
    }

    public int[] maximumWeight(List<List<Integer>> intervals) {
        int n = intervals.size(); // Number of intervals.
        final int K = 4;          // At most four intervals can be chosen.

        // Each row contains left, right, weight, and original index.
        long[][] a = new long[n][4];

        // Save the original index before sorting.
        for (int i = 0; i < n; i++) {
            a[i][0] = intervals.get(i).get(0);
            a[i][1] = intervals.get(i).get(1);
            a[i][2] = intervals.get(i).get(2);
            a[i][3] = i;
        }

        // Sort by right endpoint.
        Arrays.sort(a, Comparator.comparingLong(x -> x[1]));

        // Store right endpoints for binary search.
        long[] ends = new long[n];
        for (int i = 0; i < n; i++) {
            ends[i] = a[i][1];
        }

        // dp[k][i] means the best answer using exactly k intervals
        // from the first i sorted intervals.
        Node[][] dp = new Node[K + 1][n + 1];

        // Choosing zero intervals is always possible with score 0.
        for (int i = 0; i <= n; i++) {
            dp[0][i] = new Node(0, new int[0]);
        }

        // Process every sorted interval.
        for (int i = 1; i <= n; i++) {
            long left = a[i - 1][0];  // Current interval's left endpoint.
            long weight = a[i - 1][2]; // Current interval's weight.
            int index = (int) a[i - 1][3]; // Original index.

            // Find the first previous interval whose end >= left.
            // Everything before that position has end < left.
            int p = lowerBound(ends, i - 1, left);

            // Try choosing exactly k intervals.
            for (int k = 1; k <= K; k++) {
                // Start with the option of skipping the current interval.
                dp[k][i] = dp[k][i - 1];

                // Try taking the current interval when the previous state exists.
                if (dp[k - 1][p] != null && dp[k - 1][p].valid) {
                    // Add the current original index to the previous set.
                    int[] ids = addSorted(dp[k - 1][p].ids, index);

                    // Build the candidate score.
                    Node take = new Node(
                        dp[k - 1][p].score + weight,
                        ids
                    );

                    // Keep the better candidate.
                    if (better(take, dp[k][i])) {
                        dp[k][i] = take;
                    }
                }
            }
        }

        // Compare all choices containing at most four intervals.
        Node answer = new Node();

        for (int k = 1; k <= K; k++) {
            if (better(dp[k][n], answer)) {
                answer = dp[k][n];
            }
        }

        // Return the lexicographically smallest indices among maximum-score answers.
        return answer.ids;
    }

    // Returns the first position whose end >= target.
    private int lowerBound(long[] ends, int length, long target) {
        int lo = 0;
        int hi = length;

        // Standard lower-bound binary search.
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;

            // If this end is already >= target, search to the left.
            if (ends[mid] >= target) {
                hi = mid;
            } else {
                // Otherwise this interval is compatible, so move right.
                lo = mid + 1;
            }
        }

        // 'lo' is the count of compatible previous intervals.
        return lo;
    }
}