/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {number[]} values
 * @param {number} k
 * @return {number}
 */
var maxKDivisibleComponents = function (n, edges, values, k) {
  // Build adjacency list
  const adj = Array.from({ length: n }, () => []);
  for (const [u, v] of edges) {
    adj[u].push(v);
    adj[v].push(u);
  }

  let ans = 0;

  // To avoid recursion depth issues in JS, I'll use an explicit stack
  const parent = Array(n).fill(-1);
  const order = []; // will hold nodes in post-order

  // Iterative DFS to generate post-order
  const stack = [0];
  parent[0] = -2; // mark root's parent specially

  while (stack.length > 0) {
    const u = stack.pop();
    order.push(u); // we'll process this later in reverse order
    for (const v of adj[u]) {
      if (v === parent[u]) continue;
      parent[v] = u;
      stack.push(v);
    }
  }

  // remainder[i] = subtree sum % k for node i
  const remainder = Array(n).fill(0);

  // Process nodes in reverse order to simulate post-order
  for (let i = order.length - 1; i >= 0; i--) {
    const u = order[i];
    let sum = values[u] % k;
    for (const v of adj[u]) {
      if (v === parent[u]) continue;
      sum = (sum + remainder[v]) % k;
    }

    if (sum % k === 0) {
      ans++;
      remainder[u] = 0; // this subtree becomes its own component
    } else {
      remainder[u] = sum; // pass remainder to parent
    }
  }

  return ans;
};
