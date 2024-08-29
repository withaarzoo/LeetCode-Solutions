import java.util.*;

class Solution {

    // Helper method to perform Depth First Search (DFS)
    // This method explores all the connected stones starting from the current stone
    // (node)
    private void dfs(int node, List<List<Integer>> adj, Set<Integer> visited) {
        // Mark the current stone as visited
        visited.add(node);

        // Explore all the neighbors (connected stones) of the current stone
        for (int neighbor : adj.get(node)) {
            // If the neighbor stone is not visited, perform DFS on it
            if (!visited.contains(neighbor)) {
                dfs(neighbor, adj, visited);
            }
        }
    }

    public int removeStones(int[][] stones) {
        int n = stones.length; // Total number of stones

        // Create an adjacency list to represent the graph
        // Each stone will have a list of stones it is connected to
        List<List<Integer>> adj = new ArrayList<>();

        // Initialize the adjacency list with empty lists for each stone
        for (int i = 0; i < n; i++) {
            adj.add(new ArrayList<>());
        }

        // Build the graph by connecting stones that share the same row or column
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                // Check if stones i and j are in the same row or column
                if (stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1]) {
                    // If they are, add each stone to the other's adjacency list
                    adj.get(i).add(j);
                    adj.get(j).add(i);
                }
            }
        }

        // Set to keep track of visited stones
        Set<Integer> visited = new HashSet<>();
        int numComponents = 0; // Counter for the number of connected components

        // Perform DFS to find all connected components in the graph
        for (int i = 0; i < n; i++) {
            // If the stone has not been visited, it's a new component
            if (!visited.contains(i)) {
                dfs(i, adj, visited); // Perform DFS to visit all stones in this component
                numComponents++; // Increment the number of components
            }
        }

        // The maximum number of stones that can be removed is total stones minus the
        // number of components
        return n - numComponents;
    }
}
