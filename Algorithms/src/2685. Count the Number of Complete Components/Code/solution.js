/**
 * @param {number} n
 * @param {number[][]} edges
 * @return {number}
 */
var countCompleteComponents = function (n, edges) {
  // Build adjacency list
  const graph = Array.from({ length: n }, () => []);

  for (const [u, v] of edges) {
    graph[u].push(v);
    graph[v].push(u);
  }

  // Track visited nodes
  const visited = new Array(n).fill(false);

  let answer = 0;

  // DFS for one connected component
  function dfs(node, data) {
    // Mark current node
    visited[node] = true;

    // Count vertex
    data.vertices++;

    // Add degree
    data.degreeSum += graph[node].length;

    // Visit neighbors
    for (const next of graph[node]) {
      if (!visited[next]) {
        dfs(next, data);
      }
    }
  }

  // Traverse every component
  for (let i = 0; i < n; i++) {
    if (visited[i]) continue;

    const data = {
      vertices: 0,
      degreeSum: 0,
    };

    dfs(i, data);

    // Every edge was counted twice
    const edgeCount = data.degreeSum / 2;

    // Check completeness
    if (edgeCount === (data.vertices * (data.vertices - 1)) / 2) {
      answer++;
    }
  }

  return answer;
};
