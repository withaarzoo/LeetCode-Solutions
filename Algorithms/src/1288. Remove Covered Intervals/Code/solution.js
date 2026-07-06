/**
 * @param {number[][]} intervals
 * @return {number}
 */
var removeCoveredIntervals = function (intervals) {
  // Sort by start in ascending order.
  // For equal starts, place the larger end first
  // so the covering interval is processed before the covered one.
  intervals.sort((a, b) => {
    if (a[0] === b[0]) {
      return b[1] - a[1];
    }
    return a[0] - b[0];
  });

  // Every valid end is greater than 0, so -1 is a safe initial value.
  let maxEnd = -1;

  // This stores how many intervals are not covered.
  let remaining = 0;

  // Check every interval in the sorted order.
  for (const interval of intervals) {
    // A larger end means no earlier interval can cover this one.
    if (interval[1] > maxEnd) {
      remaining++;

      // Remember the farthest ending point seen so far.
      maxEnd = interval[1];
    }
    // Otherwise, the interval is covered, so I skip it.
  }

  // Return the number of intervals that remain.
  return remaining;
};
