/**
 * @param {number[]} arr
 * @return {number}
 */
var minJumps = function (arr) {
  const n = arr.length;

  // No jump needed
  if (n === 1) return 0;

  // Store indices for each value
  const map = new Map();

  for (let i = 0; i < n; i++) {
    if (!map.has(arr[i])) {
      map.set(arr[i], []);
    }

    map.get(arr[i]).push(i);
  }

  // BFS queue
  const queue = [0];

  // Visited array
  const visited = new Array(n).fill(false);

  visited[0] = true;

  let steps = 0;

  while (queue.length > 0) {
    let size = queue.length;

    // Process one BFS level
    while (size--) {
      const idx = queue.shift();

      // Last index reached
      if (idx === n - 1) {
        return steps;
      }

      // Move left
      if (idx - 1 >= 0 && !visited[idx - 1]) {
        visited[idx - 1] = true;
        queue.push(idx - 1);
      }

      // Move right
      if (idx + 1 < n && !visited[idx + 1]) {
        visited[idx + 1] = true;
        queue.push(idx + 1);
      }

      // Jump to same-value indices
      for (const nextIdx of map.get(arr[idx])) {
        if (!visited[nextIdx]) {
          visited[nextIdx] = true;
          queue.push(nextIdx);
        }
      }

      // Clear processed indices
      map.set(arr[idx], []);
    }

    // Next BFS level
    steps++;
  }

  return -1;
};
