class Solution
{
public:
    // DFS to visit one connected component
    void dfs(int node, vector<vector<int>> &graph, vector<bool> &vis,
             int &vertices, int &degreeSum)
    {

        // Mark current node as visited
        vis[node] = true;

        // Count this vertex
        vertices++;

        // Add its degree
        degreeSum += graph[node].size();

        // Visit all unvisited neighbors
        for (int next : graph[node])
        {
            if (!vis[next])
            {
                dfs(next, graph, vis, vertices, degreeSum);
            }
        }
    }

    int countCompleteComponents(int n, vector<vector<int>> &edges)
    {

        // Build adjacency list
        vector<vector<int>> graph(n);

        for (auto &e : edges)
        {
            graph[e[0]].push_back(e[1]);
            graph[e[1]].push_back(e[0]);
        }

        // Keep track of visited nodes
        vector<bool> vis(n, false);

        int answer = 0;

        // Process every connected component
        for (int i = 0; i < n; i++)
        {

            if (vis[i])
                continue;

            int vertices = 0;
            int degreeSum = 0;

            // Find all nodes in this component
            dfs(i, graph, vis, vertices, degreeSum);

            // Every edge is counted twice
            int edgeCount = degreeSum / 2;

            // Check whether this component is complete
            if (edgeCount == vertices * (vertices - 1) / 2)
            {
                answer++;
            }
        }

        return answer;
    }
};