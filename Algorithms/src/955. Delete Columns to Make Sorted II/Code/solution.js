var minDeletionSize = function (strs) {
  const n = strs.length;
  const m = strs[0].length;

  const sorted = Array(n - 1).fill(false);
  let deletions = 0;

  for (let col = 0; col < m; col++) {
    let needDelete = false;

    for (let row = 0; row < n - 1; row++) {
      if (!sorted[row] && strs[row][col] > strs[row + 1][col]) {
        needDelete = true;
        break;
      }
    }

    if (needDelete) {
      deletions++;
      continue;
    }

    for (let row = 0; row < n - 1; row++) {
      if (!sorted[row] && strs[row][col] < strs[row + 1][col]) {
        sorted[row] = true;
      }
    }
  }

  return deletions;
};
