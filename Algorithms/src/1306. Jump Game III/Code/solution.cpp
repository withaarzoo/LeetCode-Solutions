class Solution {
public:
    
    // DFS function to explore reachable indexes
    bool dfs(vector<int>& arr, int index, vector<bool>& visited) {
        
        // If index goes outside array, this path is invalid
        if (index < 0 || index >= arr.size()) {
            return false;
        }

        // If already visited, no need to process again
        if (visited[index]) {
            return false;
        }

        // If current value is 0, answer is found
        if (arr[index] == 0) {
            return true;
        }

        // Mark current index as visited
        visited[index] = true;

        // Move forward
        int forward = index + arr[index];

        // Move backward
        int backward = index - arr[index];

        // Return true if any path reaches value 0
        return dfs(arr, forward, visited) || dfs(arr, backward, visited);
    }

    bool canReach(vector<int>& arr, int start) {
        
        // Track visited indexes
        vector<bool> visited(arr.size(), false);

        // Start DFS from given index
        return dfs(arr, start, visited);
    }
};