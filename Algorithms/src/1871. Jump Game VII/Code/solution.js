/**
 * @param {string} s
 * @param {number} minJump
 * @param {number} maxJump
 * @return {boolean}
 */
var canReach = function (s, minJump, maxJump) {
  const n = s.length;

  // Queue for BFS
  const queue = [0];

  // Visited array
  const visited = new Array(n).fill(false);
  visited[0] = true;

  // Pointer for queue traversal
  let front = 0;

  // Farthest processed index
  let far = 0;

  while (front < queue.length) {
    const i = queue[front++];

    // Reached last index
    if (i === n - 1) {
      return true;
    }

    // Valid jump range
    const start = Math.max(i + minJump, far + 1);
    const end = Math.min(i + maxJump, n - 1);

    // Explore possible next positions
    for (let j = start; j <= end; j++) {
      // Only move to positions with '0'
      if (s[j] === "0" && !visited[j]) {
        visited[j] = true;
        queue.push(j);
      }
    }

    // Update farthest processed index
    far = Math.max(far, end);
  }

  return false;
};
