class Solution {
public:
    bool canPartitionGrid(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        
        long long total = 0;
        
        // Step 1: Calculate total sum
        for (auto &row : grid) {
            for (int val : row) {
                total += val;
            }
        }
        
        // Step 2: If total is odd → impossible
        if (total % 2 != 0) return false;
        
        long long target = total / 2;
        
        // Step 3: Try horizontal cuts
        long long rowSum = 0;
        for (int i = 0; i < m - 1; i++) { // ensure non-empty bottom part
            for (int j = 0; j < n; j++) {
                rowSum += grid[i][j];
            }
            if (rowSum == target) return true;
        }
        
        // Step 4: Precompute column sums
        vector<long long> colSum(n, 0);
        for (int j = 0; j < n; j++) {
            for (int i = 0; i < m; i++) {
                colSum[j] += grid[i][j];
            }
        }
        
        // Step 5: Try vertical cuts
        long long curr = 0;
        for (int j = 0; j < n - 1; j++) { // ensure non-empty right part
            curr += colSum[j];
            if (curr == target) return true;
        }
        
        return false;
    }
};