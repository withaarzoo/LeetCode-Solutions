class Solution {
    public boolean[] pathExistenceQueries(int n, int[] nums, int maxDiff, int[][] queries) {
        // component[i] stores which connected component node i belongs to.
        int[] component = new int[n];

        // Start with component 0 for the first node.
        int componentId = 0;

        // Check every gap between two consecutive sorted values.
        for (int i = 1; i < n; i++) {
            // A gap larger than maxDiff separates the graph into two parts.
            if (nums[i] - nums[i - 1] > maxDiff) {
                componentId++;
            }

            // Store the component of the current node.
            component[i] = componentId;
        }

        // Create one answer for every query.
        boolean[] answer = new boolean[queries.length];

        // Two nodes have a path exactly when their component IDs are equal.
        for (int i = 0; i < queries.length; i++) {
            int u = queries[i][0];
            int v = queries[i][1];

            answer[i] = component[u] == component[v];
        }

        return answer;
    }
} 