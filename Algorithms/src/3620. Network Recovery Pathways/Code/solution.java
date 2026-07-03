class Solution {
    public int findMaxPathScore(int[][] edges, boolean[] online, long k) {
        int n = online.length;

        // Build graph
        ArrayList<int[]>[] graph = new ArrayList[n];
        for (int i = 0; i < n; i++)
            graph[i] = new ArrayList<>();

        int[] indegree = new int[n];

        for (int[] e : edges) {
            graph[e[0]].add(new int[] { e[1], e[2] });
            indegree[e[1]]++;
        }

        // Topological order
        Queue<Integer> q = new ArrayDeque<>();
        for (int i = 0; i < n; i++)
            if (indegree[i] == 0)
                q.offer(i);

        ArrayList<Integer> topo = new ArrayList<>();

        while (!q.isEmpty()) {
            int u = q.poll();
            topo.add(u);

            for (int[] edge : graph[u]) {
                if (--indegree[edge[0]] == 0)
                    q.offer(edge[0]);
            }
        }

        int left = 0, right = 1_000_000_000;
        int ans = -1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            long INF = Long.MAX_VALUE / 4;
            long[] dp = new long[n];
            Arrays.fill(dp, INF);
            dp[0] = 0;

            for (int u : topo) {

                // Skip unreachable nodes
                if (dp[u] == INF)
                    continue;

                // Skip offline intermediate nodes
                if (u != 0 && u != n - 1 && !online[u])
                    continue;

                for (int[] edge : graph[u]) {
                    int v = edge[0];
                    int w = edge[1];

                    // Ignore small edges
                    if (w < mid)
                        continue;

                    // Cannot enter offline intermediate nodes
                    if (v != n - 1 && !online[v])
                        continue;

                    dp[v] = Math.min(dp[v], dp[u] + w);
                }
            }

            if (dp[n - 1] <= k) {
                ans = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return ans;
    }
}