var maximumInvitations = function (favorite) {
  const n = favorite.length;
  const inDegree = new Array(n).fill(0);
  const chainLengths = new Array(n).fill(0);
  const visited = new Array(n).fill(false);

  favorite.forEach((fav) => inDegree[fav]++);

  const queue = [];
  for (let i = 0; i < n; i++) {
    if (inDegree[i] === 0) {
      queue.push(i);
    }
  }

  while (queue.length) {
    const node = queue.shift();
    visited[node] = true;

    const next = favorite[node];
    chainLengths[next] = chainLengths[node] + 1;
    if (--inDegree[next] === 0) {
      queue.push(next);
    }
  }

  let maxCycle = 0,
    totalChains = 0;
  for (let i = 0; i < n; i++) {
    if (!visited[i]) {
      let current = i,
        cycleLength = 0;
      while (!visited[current]) {
        visited[current] = true;
        current = favorite[current];
        cycleLength++;
      }

      if (cycleLength === 2) {
        totalChains += 2 + chainLengths[i] + chainLengths[favorite[i]];
      } else {
        maxCycle = Math.max(maxCycle, cycleLength);
      }
    }
  }

  return Math.max(maxCycle, totalChains);
};
