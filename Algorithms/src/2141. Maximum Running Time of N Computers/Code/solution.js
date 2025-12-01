/**
 * @param {number} n
 * @param {number[]} batteries
 * @return {number}
 */
var maxRunTime = function (n, batteries) {
  // Sum of all battery capacities
  let total = 0;
  for (const b of batteries) total += b;

  let low = 0;
  let high = Math.floor(total / n); // Maximum possible time per computer

  // Binary search on time
  while (low < high) {
    const mid = Math.floor((low + high + 1) / 2); // upper mid

    let usable = 0;
    for (const b of batteries) {
      // Each battery can contribute at most mid minutes
      usable += Math.min(b, mid);
      if (usable >= mid * n) break;
    }

    if (usable >= mid * n) {
      // mid is feasible
      low = mid;
    } else {
      // mid is too large
      high = mid - 1;
    }
  }

  return low;
};
