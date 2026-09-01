/**
 * @param {string[]} classroom
 * @param {number} energy
 * @return {number}
 */
var minMoves = function (classroom, energy) {
  const m = classroom.length;
  const n = classroom[0].length;

  // id[r][c] stores the bit assigned to a litter cell.
  const id = Array.from({ length: m }, () => Array(n).fill(-1));

  let k = 0;
  let sr = 0;
  let sc = 0;

  // Find S and assign a unique bit to every L cell.
  for (let r = 0; r < m; r++) {
    for (let c = 0; c < n; c++) {
      if (classroom[r][c] === "S") {
        sr = r;
        sc = c;
      } else if (classroom[r][c] === "L") {
        id[r][c] = k++;
      }
    }
  }

  // If there is no litter, nothing needs to be moved.
  if (k === 0) return 0;

  // When all k bits are 1, every litter cell has been collected.
  const totalMask = (1 << k) - 1;

  // best[r][c][mask] stores the maximum energy seen for this state.
  const best = Array.from({ length: m }, () =>
    Array.from({ length: n }, () => new Int16Array(1 << k).fill(-1)),
  );

  // Each queue item is [row, column, mask, energy, moves].
  const queue = [[sr, sc, 0, energy, 0]];

  // The head index avoids O(n) array shifting during BFS.
  let head = 0;

  // Start with full energy and no collected litter.
  best[sr][sc][0] = energy;

  // Four possible movement directions.
  const dr = [-1, 1, 0, 0];
  const dc = [0, 0, -1, 1];

  while (head < queue.length) {
    // Read the next BFS state without removing earlier elements.
    const [r, c, mask, e, moves] = queue[head++];

    // Try all four neighboring cells.
    for (let d = 0; d < 4; d++) {
      const nr = r + dr[d];
      const nc = c + dc[d];

      // Ignore cells outside the classroom.
      if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;

      // Obstacles cannot be entered.
      if (classroom[nr][nc] === "X") continue;

      // Every move consumes one unit of energy.
      let ne = e - 1;

      // The move is impossible if energy becomes negative.
      if (ne < 0) continue;

      let nmask = mask;

      // Entering R restores energy to its full capacity.
      if (classroom[nr][nc] === "R") {
        ne = energy;
      }

      // Set the bit corresponding to the collected litter.
      if (classroom[nr][nc] === "L") {
        nmask |= 1 << id[nr][nc];
      }

      // All litter is collected, so this is the shortest answer.
      if (nmask === totalMask) {
        return moves + 1;
      }

      // Reaching the same state with less or equal energy is dominated.
      if (ne <= best[nr][nc][nmask]) continue;

      // Store the improved energy value.
      best[nr][nc][nmask] = ne;

      // Add the improved state to BFS.
      queue.push([nr, nc, nmask, ne, moves + 1]);
    }
  }

  // No valid path was able to collect every litter item.
  return -1;
};
