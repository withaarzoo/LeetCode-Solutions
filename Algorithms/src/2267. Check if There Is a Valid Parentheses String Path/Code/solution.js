var hasValidPath = function(grid) {
    const m = grid.length, n = grid[0].length;
    if ((m + n - 1) % 2 !== 0) return false;
    const maxBal = m + n;
    const dp = Array.from({length: m}, () => Array.from({length: n}, () => Array(maxBal).fill(false)));
    const start = grid[0][0] === '(' ? 1 : -1;
    if (start < 0) return false;
    dp[0][0][start] = true;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            for (let bal = 0; bal < maxBal; bal++) {
                if (!dp[i][j][bal]) continue;
                if (i + 1 < m) {
                    const nb = bal + (grid[i + 1][j] === '(' ? 1 : -1);
                    if (nb >= 0 && nb < maxBal) dp[i + 1][j][nb] = true;
                }
                if (j + 1 < n) {
                    const nb = bal + (grid[i][j + 1] === '(' ? 1 : -1);
                    if (nb >= 0 && nb < maxBal) dp[i][j + 1][nb] = true;
                }
            }
        }
    }
    return dp[m - 1][n - 1][0];
};