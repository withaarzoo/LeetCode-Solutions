function countServers(grid: number[][]): number {
    let m = grid.length;
    let n = grid[0].length;
    let row = new Array(m).fill(0);
    let col = new Array(n).fill(0);
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === 1) {
                row[i]++;
                col[j]++;
            }
        }
    }
    let ans = 0;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === 1 && (row[i] > 1 || col[j] > 1)) {
                ans++;
            }
        }
    }
    return ans;
};