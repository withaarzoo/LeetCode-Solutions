class Solution {
    public int shortestPath(int n, int[][] edges, String labels, int k) {
        List<int[]>[] g = new ArrayList[n];
        for (int i = 0; i < n; i++) g[i] = new ArrayList<>();
        for (int[] e : edges) {
            g[e[0]].add(new int[]{e[1], e[2]});
        }
        int INF = 1000000000;
        int[][] dist = new int[n][k + 1];
        for (int i = 0; i < n; i++) Arrays.fill(dist[i], INF);
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        dist[0][1] = 0;
        pq.offer(new int[]{0, 0, 1});
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int d = cur[0], u = cur[1], cnt = cur[2];
            if (d > dist[u][cnt]) continue;
            if (u == n - 1) return d;
            for (int[] edge : g[u]) {
                int v = edge[0], w = edge[1];
                int ncnt = labels.charAt(u) == labels.charAt(v) ? cnt + 1 : 1;
                if (ncnt > k) continue;
                if (d + w < dist[v][ncnt]) {
                    dist[v][ncnt] = d + w;
                    pq.offer(new int[]{dist[v][ncnt], v, ncnt});
                }
            }
        }
        return -1;
    }
}