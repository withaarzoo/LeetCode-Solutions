var checkIfPrerequisite = function (numCourses, prerequisites, queries) {
  // Initialize the graph
  const graph = Array.from({ length: numCourses }, () =>
    Array(numCourses).fill(false)
  );

  // Build the direct edges from prerequisites
  for (const [u, v] of prerequisites) {
    graph[u][v] = true;
  }

  // Floyd-Warshall to compute transitive closure
  for (let k = 0; k < numCourses; k++) {
    for (let i = 0; i < numCourses; i++) {
      for (let j = 0; j < numCourses; j++) {
        if (graph[i][k] && graph[k][j]) {
          graph[i][j] = true;
        }
      }
    }
  }

  // Answer the queries
  return queries.map(([u, v]) => graph[u][v]);
};
