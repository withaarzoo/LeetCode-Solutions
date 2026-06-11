class Solution {
    private static final long MOD = 1_000_000_007L;

    // Fast modular exponentiation
    private long modPow(long base, long exp) {
        long result = 1;

        while (exp > 0) {
            // Multiply when current bit is set
            if ((exp & 1) == 1) {
                result = (result * base) % MOD;
            }

            // Square the base
            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    public int assignEdgeWeights(int[][] edges) {
        int n = edges.length + 1;

        // Build adjacency list
        List<Integer>[] graph = new ArrayList[n + 1];

        for (int i = 0; i <= n; i++) {
            graph[i] = new ArrayList<>();
        }

        for (int[] e : edges) {
            int u = e[0];
            int v = e[1];

            graph[u].add(v);
            graph[v].add(u);
        }

        int maxDepth = 0;

        Deque<int[]> stack = new ArrayDeque<>();
        stack.push(new int[] { 1, 0 });

        boolean[] visited = new boolean[n + 1];
        visited[1] = true;

        while (!stack.isEmpty()) {
            int[] cur = stack.pop();

            int node = cur[0];
            int depth = cur[1];

            maxDepth = Math.max(maxDepth, depth);

            for (int next : graph[node]) {
                if (!visited[next]) {
                    visited[next] = true;
                    stack.push(new int[] { next, depth + 1 });
                }
            }
        }

        return (int) modPow(2, maxDepth - 1);
    }
}