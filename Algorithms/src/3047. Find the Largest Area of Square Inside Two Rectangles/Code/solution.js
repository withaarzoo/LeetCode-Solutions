var largestSquareArea = function (bottomLeft, topRight) {
  const n = bottomLeft.length;
  let ans = 0;

  for (let i = 0; i < n; i++) {
    for (let j = i + 1; j < n; j++) {
      const left = Math.max(bottomLeft[i][0], bottomLeft[j][0]);
      const right = Math.min(topRight[i][0], topRight[j][0]);
      const bottom = Math.max(bottomLeft[i][1], bottomLeft[j][1]);
      const top = Math.min(topRight[i][1], topRight[j][1]);

      if (right > left && top > bottom) {
        const side = Math.min(right - left, top - bottom);
        ans = Math.max(ans, side * side);
      }
    }
  }
  return ans;
};
