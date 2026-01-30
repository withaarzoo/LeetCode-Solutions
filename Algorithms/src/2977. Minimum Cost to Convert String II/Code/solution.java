class Solution {
    static final long INF = Long.MAX_VALUE;

    public long minimumCost(String source, String target,
            String[] original, String[] changed, int[] cost) {

        Map<String, Integer> id = new HashMap<>();
        Set<Integer> lens = new HashSet<>();

        int sz = 0;
        int m = original.length;
        int n = source.length();

        long[][] dist = new long[201][201];
        for (long[] row : dist)
            Arrays.fill(row, INF);

        for (int i = 0; i < m; i++) {
            if (!id.containsKey(original[i])) {
                id.put(original[i], sz++);
                lens.add(original[i].length());
            }
            if (!id.containsKey(changed[i])) {
                id.put(changed[i], sz++);
            }
            int u = id.get(original[i]);
            int v = id.get(changed[i]);
            dist[u][v] = Math.min(dist[u][v], cost[i]);
        }

        for (int i = 0; i < sz; i++)
            dist[i][i] = 0;

        for (int k = 0; k < sz; k++)
            for (int i = 0; i < sz; i++)
                if (dist[i][k] != INF)
                    for (int j = 0; j < sz; j++)
                        if (dist[k][j] != INF)
                            dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);

        long[] dp = new long[n + 1];
        Arrays.fill(dp, INF);
        dp[0] = 0;

        for (int i = 0; i < n; i++) {
            if (dp[i] == INF)
                continue;

            if (source.charAt(i) == target.charAt(i))
                dp[i + 1] = Math.min(dp[i + 1], dp[i]);

            for (int L : lens) {
                if (i + L > n)
                    continue;

                String s = source.substring(i, i + L);
                String t = target.substring(i, i + L);

                if (id.containsKey(s) && id.containsKey(t)) {
                    long d = dist[id.get(s)][id.get(t)];
                    if (d != INF)
                        dp[i + L] = Math.min(dp[i + L], dp[i] + d);
                }
            }
        }

        return dp[n] == INF ? -1 : dp[n];
    }
}
