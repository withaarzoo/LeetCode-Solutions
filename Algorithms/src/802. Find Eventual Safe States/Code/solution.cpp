class Solution
{
public:
    vector<int> eventualSafeNodes(vector<vector<int>> &graph)
    {
        int n = graph.size();
        vector<vector<int>> reversedGraph(n);
        vector<int> inDegree(n, 0);

        // Reverse the graph and calculate in-degree
        for (int i = 0; i < n; ++i)
        {
            for (int neighbor : graph[i])
            {
                reversedGraph[neighbor].push_back(i);
                inDegree[i]++;
            }
        }

        // Find all terminal nodes
        queue<int> q;
        for (int i = 0; i < n; ++i)
        {
            if (inDegree[i] == 0)
                q.push(i);
        }

        // Topological sorting to find safe nodes
        vector<int> safeNodes;
        while (!q.empty())
        {
            int node = q.front();
            q.pop();
            safeNodes.push_back(node);

            for (int neighbor : reversedGraph[node])
            {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0)
                    q.push(neighbor);
            }
        }

        sort(safeNodes.begin(), safeNodes.end());
        return safeNodes;
    }
};
