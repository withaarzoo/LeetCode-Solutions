class Solution {
    public int[] pathExistenceQueries(int n, int[] nums, int maxDiff, int[][] queries) {
        // order stores original node indices before sorting them by nums value.
        Integer[] order = new Integer[n];

        for (int i = 0; i < n; i++) {
            order[i] = i;
        }

        // Sort node indices according to their nums values.
        Arrays.sort(order, (a, b) -> Integer.compare(nums[a], nums[b]));

        // pos[node] gives the sorted position of an original node.
        int[] pos = new int[n];

        // values stores nums in sorted order.
        int[] values = new int[n];

        for (int i = 0; i < n; i++) {
            values[i] = nums[order[i]];
            pos[order[i]] = i;
        }

        // Find how many binary lifting levels are needed.
        int log = 1;

        while ((1 << log) <= n) {
            log++;
        }

        // jump[p][i] stores the farthest position after at most 2^p greedy jumps.
        int[][] jump = new int[log][n];

        // Find the farthest one-edge reach for every position using two pointers.
        int r = 0;

        for (int i = 0; i < n; i++) {
            // Keep r at or after the current position.
            if (r < i) {
                r = i;
            }

            // Move r while a direct edge from i is still possible.
            while (r + 1 < n && values[r + 1] - values[i] <= maxDiff) {
                r++;
            }

            // Save the farthest one-jump destination.
            jump[0][i] = r;
        }

        // Build all larger binary jumps.
        for (int p = 1; p < log; p++) {
            for (int i = 0; i < n; i++) {
                // Apply the previous jump twice.
                jump[p][i] = jump[p - 1][jump[p - 1][i]];
            }
        }

        // Store the result of every query.
        int[] answer = new int[queries.length];

        for (int q = 0; q < queries.length; q++) {
            // Convert both original nodes to sorted positions.
            int left = pos[queries[q][0]];
            int right = pos[queries[q][1]];

            // Always process the query from the smaller position to the larger one.
            if (left > right) {
                int temp = left;
                left = right;
                right = temp;
            }

            // The distance from a node to itself is zero.
            if (left == right) {
                answer[q] = 0;
                continue;
            }

            int current = left;
            int distance = 0;

            // Use the largest possible groups of jumps first.
            for (int p = log - 1; p >= 0; p--) {
                if (jump[p][current] < right) {
                    // Take 2^p greedy jumps without passing the target.
                    current = jump[p][current];
                    distance += 1 << p;
                }
            }

            // Check whether one last edge reaches the target.
            if (jump[0][current] >= right) {
                answer[q] = distance + 1;
            } else {
                // No forward progress means the nodes are disconnected.
                answer[q] = -1;
            }
        }

        return answer;
    }
}