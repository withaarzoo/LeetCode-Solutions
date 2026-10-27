/**
 * @param {number} m
 * @param {number} n
 * @return {string[]}
 */
var createGrid = function(m, n) {
    let grid = Array.from({length: m}, () => Array(n).fill('#'));
    for (let j = 0; j < n; ++j) {
        grid[0][j] = '.';
    }
    for (let i = 0; i < m; ++i) {
        grid[i][n - 1] = '.';
    }
    return grid.map(row => row.join(''));
};