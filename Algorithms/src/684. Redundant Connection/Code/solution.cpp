class Solution
{
public:
    vector<int> findRedundantConnection(vector<vector<int>> &edges)
    {
        int n = edges.size();
        vector<int> parent(n + 1), rank(n + 1, 0);

        // Initialize each node as its own parent
        for (int i = 0; i <= n; i++)
        {
            parent[i] = i;
        }

        // Find function with path compression
        function<int(int)> find = [&](int node)
        {
            if (parent[node] != node)
                parent[node] = find(parent[node]); // Path compression
            return parent[node];
        };

        // Union function by rank
        auto unionSets = [&](int u, int v)
        {
            int rootU = find(u), rootV = find(v);
            if (rootU == rootV)
                return false; // Cycle detected
            if (rank[rootU] > rank[rootV])
                parent[rootV] = rootU;
            else if (rank[rootU] < rank[rootV])
                parent[rootU] = rootV;
            else
            {
                parent[rootV] = rootU;
                rank[rootU]++;
            }
            return true;
        };

        // Process each edge
        for (auto &edge : edges)
        {
            if (!unionSets(edge[0], edge[1]))
                return edge;
        }
        return {};
    }
};
