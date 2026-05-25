class Solution {
public:
    bool canReach(string s, int minJump, int maxJump) {
        int n = s.size();

        // Queue for BFS traversal
        queue<int> q;

        // Visited array to avoid revisiting indices
        vector<bool> visited(n, false);

        // Start from index 0
        q.push(0);
        visited[0] = true;

        // Keeps track of the farthest processed index
        int far = 0;

        while (!q.empty()) {
            int i = q.front();
            q.pop();

            // If we reached last index, answer is true
            if (i == n - 1)
                return true;

            // Calculate jump range
            int start = max(i + minJump, far + 1);
            int end = min(i + maxJump, n - 1);

            // Explore all valid next positions
            for (int j = start; j <= end; j++) {

                // Only visit positions containing '0'
                if (s[j] == '0' && !visited[j]) {
                    visited[j] = true;
                    q.push(j);
                }
            }

            // Update farthest processed position
            far = max(far, end);
        }

        return false;
    }
};