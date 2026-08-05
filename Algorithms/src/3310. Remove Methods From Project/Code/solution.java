class Solution {
    public List<Integer> remainingMethods(int n, int k, int[][] invocations) {

        // Build the adjacency list
        List<Integer>[] graph = new ArrayList[n];
        for (int i = 0; i < n; i++) {
            graph[i] = new ArrayList<>();
        }

        for (int[] edge : invocations) {
            graph[edge[0]].add(edge[1]);
        }

        // Marks suspicious methods
        boolean[] vis = new boolean[n];

        // DFS from method k
        dfs(k, graph, vis);

        // If a safe method invokes a suspicious one,
        // removal is not allowed
        for (int[] edge : invocations) {
            int u = edge[0];
            int v = edge[1];

            if (!vis[u] && vis[v]) {
                List<Integer> ans = new ArrayList<>();
                for (int i = 0; i < n; i++) {
                    ans.add(i);
                }
                return ans;
            }
        }

        // Return remaining methods
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (!vis[i]) {
                ans.add(i);
            }
        }

        return ans;
    }

    // DFS to mark all reachable methods
    private void dfs(int u, List<Integer>[] graph, boolean[] vis) {
        vis[u] = true;

        for (int v : graph[u]) {
            if (!vis[v]) {
                dfs(v, graph, vis);
            }
        }
    }
}