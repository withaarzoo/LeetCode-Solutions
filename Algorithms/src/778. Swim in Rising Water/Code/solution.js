/**
 * Min-heap implementation for arrays where element is [time, r, c]
 * comparator compares element[0]
 */
class MinHeap {
    constructor() { this.heap = []; }
    size() { return this.heap.length; }
    peek() { return this.heap[0]; }
    push(val) {
        this.heap.push(val);
        this._bubbleUp(this.heap.length - 1);
    }
    pop() {
        if (this.heap.length === 0) return undefined;
        const top = this.heap[0];
        const last = this.heap.pop();
        if (this.heap.length > 0) {
            this.heap[0] = last;
            this._bubbleDown(0);
        }
        return top;
    }
    _bubbleUp(i) {
        while (i > 0) {
            const p = Math.floor((i - 1) / 2);
            if (this.heap[p][0] <= this.heap[i][0]) break;
            [this.heap[p], this.heap[i]] = [this.heap[i], this.heap[p]];
            i = p;
        }
    }
    _bubbleDown(i) {
        const n = this.heap.length;
        while (true) {
            let smallest = i;
            const l = 2 * i + 1, r = 2 * i + 2;
            if (l < n && this.heap[l][0] < this.heap[smallest][0]) smallest = l;
            if (r < n && this.heap[r][0] < this.heap[smallest][0]) smallest = r;
            if (smallest === i) break;
            [this.heap[smallest], this.heap[i]] = [this.heap[i], this.heap[smallest]];
            i = smallest;
        }
    }
}

/**
 * @param {number[][]} grid
 * @return {number}
 */
var swimInWater = function(grid) {
    const n = grid.length;
    const visited = Array.from({length: n}, () => Array(n).fill(false));
    const heap = new MinHeap();
    heap.push([grid[0][0], 0, 0]);
    const dirs = [[1,0],[-1,0],[0,1],[0,-1]];

    while (heap.size() > 0) {
        const [t, r, c] = heap.pop();
        if (visited[r][c]) continue;
        visited[r][c] = true;
        if (r === n - 1 && c === n - 1) return t;
        for (const [dr, dc] of dirs) {
            const nr = r + dr, nc = c + dc;
            if (nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc]) {
                const nt = Math.max(t, grid[nr][nc]);
                heap.push([nt, nr, nc]);
            }
        }
    }
    return -1;
};
