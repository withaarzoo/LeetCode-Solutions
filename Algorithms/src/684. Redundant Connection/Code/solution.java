class Solution {
    public int[] findRedundantConnection(int[][] edges) {
        int n = edges.length;
        int[] parent = new int[n + 1];
        int[] rank = new int[n + 1];

        for (int i = 0; i <= n; i++) {
            parent[i] = i;
            rank[i] = 0;
        }

        // Find function with path compression
        int find(int node) {
            if (parent[node] != node) 
                parent[node] = find(parent[node]);
            return parent[node];
        }

        // Union function by rank
        boolean unionSets(int u, int v) {
            int rootU = find(u);
            int rootV = find(v);
            if (rootU == rootV) return false;
            if (rank[rootU] > rank[rootV]) 
                parent[rootV] = rootU;
            else if (rank[rootU] < rank[rootV]) 
                parent[rootU] = rootV;
            else {
                parent[rootV] = rootU;
                rank[rootU]++;
            }
            return true;
        }

        // Process each edge
        for (int[] edge : edges) {
            if (!unionSets(edge[0], edge[1])) return edge;
        }
        return new int[0];
    }
}
