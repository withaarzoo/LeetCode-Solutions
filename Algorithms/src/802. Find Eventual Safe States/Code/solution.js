var eventualSafeNodes = function (graph) {
  const n = graph.length;
  const reversedGraph = Array.from({ length: n }, () => []);
  const inDegree = Array(n).fill(0);

  // Reverse the graph and calculate in-degree
  for (let i = 0; i < n; i++) {
    for (const neighbor of graph[i]) {
      reversedGraph[neighbor].push(i);
      inDegree[i]++;
    }
  }

  // Find all terminal nodes
  const queue = [];
  for (let i = 0; i < n; i++) {
    if (inDegree[i] === 0) queue.push(i);
  }

  // Topological sorting to find safe nodes
  const safeNodes = [];
  while (queue.length > 0) {
    const node = queue.shift();
    safeNodes.push(node);

    for (const neighbor of reversedGraph[node]) {
      inDegree[neighbor]--;
      if (inDegree[neighbor] === 0) queue.push(neighbor);
    }
  }

  return safeNodes.sort((a, b) => a - b);
};
