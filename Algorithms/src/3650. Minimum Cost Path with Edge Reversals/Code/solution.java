class Solution {
    public int minCost(int n, int[][] edges) {
        List<int[]>[] graph = new ArrayList[n];
        for (int i = 0; i < n; i++)
            graph[i] = new ArrayList<>();

        for (int[] e : edges) {
            graph[e[0]].add(new int[] { e[1], e[2] });
            graph[e[1]].add(new int[] { e[0], 2 * e[2] });
        }

        long[] dist = new long[n];
        Arrays.fill(dist, Long.MAX_VALUE);
        dist[0] = 0;

        PriorityQueue<long[]> pq = new PriorityQueue<>(Comparator.comparingLong(a -> a[0]));
        pq.offer(new long[] { 0, 0 });

        while (!pq.isEmpty()) {
            long[] cur = pq.poll();
            long cost = cur[0];
            int node = (int) cur[1];

            if (cost > dist[node])
                continue;

            for (int[] nxt : graph[node]) {
                int v = nxt[0];
                int w = nxt[1];
                if (dist[v] > cost + w) {
                    dist[v] = cost + w;
                    pq.offer(new long[] { dist[v], v });
                }
            }
        }

        return dist[n - 1] == Long.MAX_VALUE ? -1 : (int) dist[n - 1];
    }
}
