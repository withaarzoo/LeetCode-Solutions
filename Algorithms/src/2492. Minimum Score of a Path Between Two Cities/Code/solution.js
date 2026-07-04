/**
 * @param {number} n
 * @param {number[][]} roads
 * @return {number}
 */
var minScore = function (n, roads) {
  // graph[city] stores [neighbor, road distance].
  const graph = Array.from({ length: n + 1 }, () => []);

  // Build the undirected graph by storing each road in both directions.
  for (const [a, b, distance] of roads) {
    graph[a].push([b, distance]);
    graph[b].push([a, distance]);
  }

  // visited prevents the same city from being processed more than once.
  const visited = new Array(n + 1).fill(false);

  // I use an array as a BFS queue and move a pointer instead of shifting.
  // This keeps every queue operation efficient.
  const queue = [1];
  let front = 0;
  visited[1] = true;

  // Start with the largest possible value.
  let answer = Infinity;

  // Visit every city in the connected component containing city 1.
  while (front < queue.length) {
    const city = queue[front++];

    // Check every road connected to the current city.
    for (const [nextCity, distance] of graph[city]) {
      // Every road in this component can affect the final score.
      answer = Math.min(answer, distance);

      // Add the neighboring city only if it has not been visited.
      if (!visited[nextCity]) {
        visited[nextCity] = true;
        queue.push(nextCity);
      }
    }
  }

  // Return the smallest road distance found in the component.
  return answer;
};
