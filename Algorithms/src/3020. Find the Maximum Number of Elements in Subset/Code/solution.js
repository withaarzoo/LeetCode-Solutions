/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumLength = function (nums) {
  // Store frequency of every number
  const freq = new Map();

  for (const x of nums) {
    freq.set(x, (freq.get(x) || 0) + 1);
  }

  let ans = 1;

  // Handle value 1 separately
  if (freq.has(1)) {
    const cnt = freq.get(1);

    // Only odd count of ones is valid
    ans = Math.max(ans, cnt % 2 ? cnt : cnt - 1);
  }

  // Try every distinct starting value
  for (const [start] of freq) {
    if (start === 1) continue;

    let cur = start;
    let len = 0;

    while (freq.has(cur)) {
      // Use two copies if available
      if (freq.get(cur) >= 2) {
        len += 2;

        // Move to the squared value
        cur = cur * cur;
      } else {
        // Single copy becomes the center
        len++;
        break;
      }
    }

    // No center found
    if (len % 2 === 0) len--;

    ans = Math.max(ans, len);
  }

  return ans;
};
