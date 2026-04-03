/**
 * @param {number[]} robots
 * @param {number[]} distance
 * @param {number[]} walls
 * @return {number}
 */
var maxWalls = function (robots, distance, walls) {
  const n = robots.length;

  const arr = [];
  for (let i = 0; i < n; i++) {
    arr.push([robots[i], distance[i]]);
  }

  arr.sort((a, b) => a[0] - b[0]);
  walls.sort((a, b) => a - b);

  // Dummy robot
  arr.push([1e9, 0]);

  const lowerBound = (arr, target) => {
    let left = 0,
      right = arr.length;

    while (left < right) {
      const mid = Math.floor((left + right) / 2);

      if (arr[mid] < target) left = mid + 1;
      else right = mid;
    }

    return left;
  };

  const upperBound = (arr, target) => {
    let left = 0,
      right = arr.length;

    while (left < right) {
      const mid = Math.floor((left + right) / 2);

      if (arr[mid] <= target) left = mid + 1;
      else right = mid;
    }

    return left;
  };

  const countWalls = (left, right) => {
    if (left > right) return 0;

    return upperBound(walls, right) - lowerBound(walls, left);
  };

  const dp = Array.from({ length: n }, () => [0, 0]);

  dp[0][0] = countWalls(arr[0][0] - arr[0][1], arr[0][0]);

  const firstRightEnd =
    n === 1
      ? arr[0][0] + arr[0][1]
      : Math.min(arr[0][0] + arr[0][1], arr[1][0] - 1);

  dp[0][1] = countWalls(arr[0][0], firstRightEnd);

  for (let i = 1; i < n; i++) {
    const pos = arr[i][0];
    const dist = arr[i][1];

    // Shoot right
    const rightEnd = Math.min(pos + dist, arr[i + 1][0] - 1);
    const rightWalls = countWalls(pos, rightEnd);

    dp[i][1] = Math.max(dp[i - 1][0], dp[i - 1][1]) + rightWalls;

    // Shoot left
    const leftStart = Math.max(pos - dist, arr[i - 1][0] + 1);
    const leftWalls = countWalls(leftStart, pos);

    dp[i][0] = dp[i - 1][0] + leftWalls;

    const prevRightEnd = Math.min(arr[i - 1][0] + arr[i - 1][1], pos - 1);

    const overlapStart = leftStart;
    const overlapEnd = Math.min(prevRightEnd, pos - 1);

    const overlapWalls = countWalls(overlapStart, overlapEnd);

    dp[i][0] = Math.max(dp[i][0], dp[i - 1][1] + leftWalls - overlapWalls);
  }

  return Math.max(dp[n - 1][0], dp[n - 1][1]);
};
