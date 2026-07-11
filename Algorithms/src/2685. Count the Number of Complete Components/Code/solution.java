class Solution {

    // DFS to visit one connected component
    private void dfs(int node, List<Integer>[] graph, boolean[] visited,
            int[] info) {

        // Mark node as visited
        visited[node] = true;

        // Count one more vertex
        info[0]++;

        // Add current node's degree
        info[1] += graph[node].size();

        // Visit all neighbors
        for (int next : graph[node]) {
            if (!visited[next]) {
                dfs(next, graph, visited, info);
            }
        }
    }

    public int countCompleteComponents(int n, int[][] edges) {

        // Build adjacency list
        List<Integer>[] graph = new ArrayList[n];

        for (int i = 0; i < n; i++) {
            graph[i] = new ArrayList<>();
        }

        for (int[] edge : edges) {
            graph[edge[0]].add(edge[1]);
            graph[edge[1]].add(edge[0]);
        }

        boolean[] visited = new boolean[n];

        int answer = 0;

        // Process every component
        for (int i = 0; i < n; i++) {

            if (visited[i])
                continue;

            // info[0] = vertices
            // info[1] = total degree
            int[] info = new int[2];

            dfs(i, graph, visited, info);

            int vertices = info[0];
            int edgeCount = info[1] / 2;

            // Check whether this component is complete
            if (edgeCount == vertices * (vertices - 1) / 2) {
                answer++;
            }
        }

        return answer;
    }
}