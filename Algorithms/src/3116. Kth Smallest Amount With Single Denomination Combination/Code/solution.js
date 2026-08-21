/**
 * @param {number[]} coins
 * @param {number} k
 * @return {number}
 */
var findKthSmallest = function (coins, k) {
  // Sort coins so I can remove redundant larger denominations.
  coins.sort((a, b) => a - b);

  // Keep only coins that produce some multiples not already covered.
  const useful = [];

  for (const coin of coins) {
    let redundant = false;

    // If a kept coin divides this coin, every multiple of this coin
    // is already included among the multiples of the smaller coin.
    for (const prev of useful) {
      if (coin % prev === 0) {
        redundant = true;
        break;
      }
    }

    // Keep the coin only when it is useful.
    if (!redundant) {
      useful.push(coin);
    }
  }

  const m = useful.length;

  // The kth multiple of the smallest coin gives a safe upper bound.
  let low = 1;
  let high = useful[0] * k;

  const totalMasks = 1 << m;

  // Store the LCM and sign for every non-empty subset.
  const lcms = new Array(totalMasks).fill(1);
  const signs = new Array(totalMasks).fill(1);

  // Helper function for the greatest common divisor.
  const gcd = (a, b) => {
    while (b !== 0) {
      [a, b] = [b, a % b];
    }
    return a;
  };

  // Precompute information for every subset.
  for (let mask = 1; mask < totalMasks; mask++) {
    let currentLCM = 1;
    let bits = 0;

    for (let i = 0; i < m; i++) {
      // Include useful[i] when the corresponding bit is set.
      if ((mask & (1 << i)) !== 0) {
        const g = gcd(currentLCM, useful[i]);

        // Divide before multiplying to calculate the LCM safely.
        currentLCM /= g;

        // Cap values that are larger than the maximum binary search value.
        if (currentLCM > Math.floor(high / useful[i])) {
          currentLCM = high + 1;
          break;
        }

        currentLCM *= useful[i];
        bits++;
      }
    }

    // Save the subset LCM.
    lcms[mask] = currentLCM;

    // Odd subsets are added and even subsets are subtracted.
    signs[mask] = bits % 2 === 1 ? 1 : -1;
  }

  // Count how many valid amounts are less than or equal to x.
  const count = (x) => {
    let result = 0;

    for (let mask = 1; mask < totalMasks; mask++) {
      // Ignore subsets whose LCM cannot divide any value up to x.
      if (lcms[mask] <= x) {
        result += signs[mask] * Math.floor(x / lcms[mask]);
      }
    }

    return result;
  };

  // Binary search for the smallest value whose count is at least k.
  while (low < high) {
    const mid = Math.floor(low + (high - low) / 2);

    // Search left if mid already reaches the kth valid amount.
    if (count(mid) >= k) {
      high = mid;
    } else {
      // Otherwise, search larger values.
      low = mid + 1;
    }
  }

  // low is the kth smallest valid amount.
  return low;
};
