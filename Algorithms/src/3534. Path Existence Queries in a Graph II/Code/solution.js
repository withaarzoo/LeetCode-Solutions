/**
 * @param {number} n
 * @param {number[]} nums
 * @param {number} maxDiff
 * @param {number[][]} queries
 * @return {number[]}
 */
var pathExistenceQueries = function (n, nums, maxDiff, queries) {
  // Store original node indices so I can sort them by nums value.
  const order = Array.from({ length: n }, (_, i) => i);

  // Sort nodes according to their values.
  order.sort((a, b) => nums[a] - nums[b]);

  // pos[node] gives the sorted position of an original node.
  const pos = new Array(n);

  // values stores nums in sorted order.
  const values = new Array(n);

  for (let i = 0; i < n; i++) {
    values[i] = nums[order[i]];
    pos[order[i]] = i;
  }

  // Find the number of levels needed for binary lifting.
  let LOG = 1;

  while (2 ** LOG <= n) {
    LOG++;
  }

  // jump[p][i] is the farthest position after at most 2^p greedy jumps.
  const jump = Array.from({ length: LOG }, () => new Array(n));

  // Use two pointers to find every one-jump destination.
  let r = 0;

  for (let i = 0; i < n; i++) {
    // Keep r at or after i.
    if (r < i) {
      r = i;
    }

    // Extend r while there is still a direct edge from i.
    while (r + 1 < n && values[r + 1] - values[i] <= maxDiff) {
      r++;
    }

    // Save the farthest position reachable in one edge.
    jump[0][i] = r;
  }

  // Build the binary lifting table.
  for (let p = 1; p < LOG; p++) {
    for (let i = 0; i < n; i++) {
      // Apply the previous level twice to double the jump count.
      jump[p][i] = jump[p - 1][jump[p - 1][i]];
    }
  }

  // Store one result for each query.
  const answer = new Array(queries.length);

  for (let q = 0; q < queries.length; q++) {
    // Convert original nodes into sorted positions.
    let left = pos[queries[q][0]];
    let right = pos[queries[q][1]];

    // The graph is undirected, so always move from left to right.
    if (left > right) {
      [left, right] = [right, left];
    }

    // No edge is needed when both nodes are the same.
    if (left === right) {
      answer[q] = 0;
      continue;
    }

    let current = left;
    let distance = 0;

    // Take the largest groups of jumps that still stop before the target.
    for (let p = LOG - 1; p >= 0; p--) {
      if (jump[p][current] < right) {
        // Skip 2^p greedy jumps at once.
        current = jump[p][current];
        distance += 2 ** p;
      }
    }

    // One final jump must reach the target.
    if (jump[0][current] >= right) {
      answer[q] = distance + 1;
    } else {
      // Otherwise the target belongs to another connected component.
      answer[q] = -1;
    }
  }

  return answer;
};
