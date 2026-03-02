var minSwaps = function (grid) {
  const n = grid.length;
  const trailing = new Array(n).fill(0);

  // Count trailing zeros
  for (let i = 0; i < n; i++) {
    let count = 0;
    for (let j = n - 1; j >= 0; j--) {
      if (grid[i][j] === 0) count++;
      else break;
    }
    trailing[i] = count;
  }

  let swaps = 0;

  for (let i = 0; i < n; i++) {
    let required = n - 1 - i;
    let j = i;

    while (j < n && trailing[j] < required) j++;

    if (j === n) return -1;

    while (j > i) {
      [trailing[j], trailing[j - 1]] = [trailing[j - 1], trailing[j]];
      swaps++;
      j--;
    }
  }

  return swaps;
};
