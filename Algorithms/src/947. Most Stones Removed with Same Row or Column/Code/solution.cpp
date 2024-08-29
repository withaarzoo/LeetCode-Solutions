#include <vector>
#include <unordered_set>
using namespace std;

class Solution
{
public:
    // Depth-First Search (DFS) function to explore the connected component
    void dfs(int node, vector<vector<int>> &adj, unordered_set<int> &visited)
    {
        // Mark the current node as visited
        visited.insert(node);

        // Traverse all the neighboring nodes of the current node
        for (int neighbor : adj[node])
        {
            // If the neighbor hasn't been visited yet, perform DFS on it
            if (visited.find(neighbor) == visited.end())
            {
                dfs(neighbor, adj, visited);
            }
        }
    }

    int removeStones(vector<vector<int>> &stones)
    {
        int n = stones.size();      // Number of stones
        vector<vector<int>> adj(n); // Adjacency list to represent the graph

        // Build the graph by connecting stones that share the same row or column
        for (int i = 0; i < n; ++i)
        {
            for (int j = i + 1; j < n; ++j)
            {
                // If two stones are in the same row or column, connect them
                if (stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1])
                {
                    adj[i].push_back(j); // Add j to the neighbors of i
                    adj[j].push_back(i); // Add i to the neighbors of j
                }
            }
        }

        unordered_set<int> visited; // To track visited nodes
        int numComponents = 0;      // Number of connected components

        // Find the number of connected components in the graph
        for (int i = 0; i < n; ++i)
        {
            // If the node hasn't been visited yet, it means it's a new component
            if (visited.find(i) == visited.end())
            {
                dfs(i, adj, visited); // Explore this component using DFS
                numComponents++;      // Increment the component count
            }
        }

        // The number of moves required is the total number of stones minus the number of components
        return n - numComponents;
    }
};
