var minOperations = function (s, k) {
  const n = s.length;
  let zero = 0;

  for (let c of s) if (c === "0") zero++;

  if (zero === 0) return 0;

  if (n === k) {
    if (zero === n) return 1;
    if (zero === 0) return 0;
    return -1;
  }

  const one = n - zero;
  const base = n - k;

  let ans = Infinity;

  // Odd case
  if (k % 2 === zero % 2) {
    let m = Math.max(Math.ceil(zero / k), Math.ceil(one / base));

    if (m % 2 === 0) m++;

    ans = Math.min(ans, m);
  }

  // Even case
  if (zero % 2 === 0) {
    let m = Math.max(Math.ceil(zero / k), Math.ceil(zero / base));

    if (m % 2 === 1) m++;

    ans = Math.min(ans, m);
  }

  return ans === Infinity ? -1 : ans;
};
