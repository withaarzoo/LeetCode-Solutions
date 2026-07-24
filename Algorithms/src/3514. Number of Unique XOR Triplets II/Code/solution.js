/**
 * @param {number[]} nums
 * @return {number}
 */
var uniqueXorTriplets = function (nums) {
  const MAX_XOR = 2048;

  // Store whether a value exists.
  const present = new Array(MAX_XOR).fill(false);
  for (const x of nums) {
    present[x] = true;
  }

  // Initially only XOR = 0 is reachable.
  let dp = new Array(MAX_XOR).fill(false);
  dp[0] = true;

  // Pick exactly 3 values.
  for (let step = 0; step < 3; step++) {
    const next = new Array(MAX_XOR).fill(false);

    // Extend every reachable XOR.
    for (let cur = 0; cur < MAX_XOR; cur++) {
      if (!dp[cur]) continue;

      // Try every existing value.
      for (let v = 0; v < MAX_XOR; v++) {
        if (present[v]) {
          next[cur ^ v] = true;
        }
      }
    }

    dp = next;
  }

  // Count unique XOR values.
  let ans = 0;
  for (const ok of dp) {
    if (ok) ans++;
  }

  return ans;
};
