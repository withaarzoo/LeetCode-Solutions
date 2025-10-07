/**
 * JavaScript (binary-search on a sorted dry-day list)
 * Note: splice removal is O(n). Overall worst-case O(n^2) but commonly passes.
 * @param {number[]} rains
 * @return {number[]}
 */
var avoidFlood = function (rains) {
  const n = rains.length;
  const ans = new Array(n).fill(1);
  const last = new Map(); // lake -> last day index
  const dry = []; // sorted list of dry day indices (kept sorted by appending)

  // upperBound: first index with arr[idx] > target
  const upperBound = (arr, target) => {
    let l = 0,
      r = arr.length;
    while (l < r) {
      let m = (l + r) >> 1;
      if (arr[m] <= target) l = m + 1;
      else r = m;
    }
    return l;
  };

  for (let i = 0; i < n; ++i) {
    if (rains[i] > 0) {
      const lake = rains[i];
      ans[i] = -1;
      if (last.has(lake)) {
        const prev = last.get(lake);
        const idx = upperBound(dry, prev); // first dry day > prev
        if (idx === dry.length) return []; // impossible
        const dryDay = dry[idx];
        ans[dryDay] = lake; // use this dry day to dry the lake
        dry.splice(idx, 1); // remove used dry day
      }
      last.set(lake, i);
    } else {
      dry.push(i); // append keeps `dry` sorted because i increases
    }
  }
  return ans;
};
