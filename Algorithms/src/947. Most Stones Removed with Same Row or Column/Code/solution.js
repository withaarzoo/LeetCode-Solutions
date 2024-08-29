var removeStones = function (stones) {
  const n = stones.length; // Get the number of stones.

  // Create an adjacency list to represent the graph, where each stone is a node.
  const adj = Array.from({ length: n }, () => []);

  // Build the graph by connecting stones that share the same row or column.
  for (let i = 0; i < n; i++) {
    for (let j = i + 1; j < n; j++) {
      // If two stones are in the same row or column, they are connected.
      if (stones[i][0] === stones[j][0] || stones[i][1] === stones[j][1]) {
        adj[i].push(j); // Connect stone i to stone j.
        adj[j].push(i); // Connect stone j to stone i (since the graph is undirected).
      }
    }
  }

  const visited = new Set(); // Create a set to track visited nodes (stones).

  // Depth-First Search (DFS) function to explore all connected stones.
  function dfs(node) {
    visited.add(node); // Mark the current stone as visited.
    for (let neighbor of adj[node]) {
      // For each connected stone (neighbor)...
      if (!visited.has(neighbor)) {
        // If the neighbor hasn't been visited yet...
        dfs(neighbor); // Recursively visit that neighbor.
      }
    }
  }

  let numComponents = 0; // Initialize a counter for connected components.

  // Go through each stone to find all connected components.
  for (let i = 0; i < n; i++) {
    if (!visited.has(i)) {
      // If the stone hasn't been visited...
      dfs(i); // Perform a DFS starting from this stone.
      numComponents++; // Increment the number of connected components.
    }
  }

  // The maximum number of stones that can be removed is the total number of stones
  // minus the number of connected components.
  return n - numComponents;
};
