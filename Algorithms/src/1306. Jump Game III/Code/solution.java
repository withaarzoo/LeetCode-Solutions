class Solution {
    
    // DFS helper function
    private boolean dfs(int[] arr, int index, boolean[] visited) {

        // Invalid index
        if (index < 0 || index >= arr.length) {
            return false;
        }

        // Skip already visited indexes
        if (visited[index]) {
            return false;
        }

        // Found value 0
        if (arr[index] == 0) {
            return true;
        }

        // Mark current index as visited
        visited[index] = true;

        // Explore both directions
        int forward = index + arr[index];
        int backward = index - arr[index];

        // Return true if any direction reaches 0
        return dfs(arr, forward, visited) || dfs(arr, backward, visited);
    }

    public boolean canReach(int[] arr, int start) {
        
        // Visited array to avoid cycles
        boolean[] visited = new boolean[arr.length];

        // Start DFS
        return dfs(arr, start, visited);
    }
}