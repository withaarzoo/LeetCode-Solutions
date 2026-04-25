/**
 * @param {number} side
 * @param {number[][]} points
 * @param {number} k
 * @return {number}
 */
var maxDistance = function (side, points, k) {
  const pts = points.map(([x, y]) => {
    let pos;
    if (x === 0) pos = y;
    else if (y === side) pos = side + x;
    else if (x === side) pos = 3 * side - y;
    else pos = 4 * side - x;
    return { pos, x, y };
  });

  pts.sort((a, b) => a.pos - b.pos);

  function getOffset(x, y, d) {
    if (x === 0) {
      if (d <= 2 * side - y) return d;
      if (d <= side + y) return 2 * side + d - 2 * y;
      return -1;
    } else if (y === side) {
      if (d <= 2 * side - x) return d;
      if (d <= side + x) return 2 * side + d - 2 * x;
      return -1;
    } else if (x === side) {
      if (d <= side + y) return d;
      if (d <= 2 * side - y) return d + 2 * y;
      return -1;
    } else {
      if (d <= side + x) return d;
      if (d <= 2 * side - x) return d + 2 * x;
      return -1;
    }
  }

  function lowerBound(arr, l, r, target) {
    while (l < r) {
      const m = (l + r) >> 1;
      if (arr[m] < target) l = m + 1;
      else r = m;
    }
    return l;
  }

  function can(d) {
    const n = pts.length;
    const pos3 = new Array(3 * n);
    for (let i = 0; i < n; i++) {
      pos3[i] = pts[i].pos;
      pos3[i + n] = pts[i].pos + 4 * side;
      pos3[i + 2 * n] = pts[i].pos + 8 * side;
    }

    const nxt = new Array(2 * n).fill(-1);

    for (let i = 0; i < 2 * n; i++) {
      const p = pts[i % n];
      const off = getOffset(p.x, p.y, d);
      if (off < 0) continue;

      const target = pos3[i] + off;
      const hi = Math.min(i + n, 3 * n);
      const j = lowerBound(pos3, i + 1, hi, target);
      if (j < hi) nxt[i] = j;
    }

    for (let start = 0; start < n; start++) {
      let cur = start;
      let cnt = 1;

      while (cnt < k) {
        cur = nxt[cur];
        if (cur === -1 || cur >= start + n) break;
        cnt++;
      }

      if (cnt >= k) {
        const a = pts[start];
        const b = pts[cur % n];
        if (Math.abs(a.x - b.x) + Math.abs(a.y - b.y) >= d) {
          return true;
        }
      }
    }

    return false;
  }

  let lo = 0,
    hi = 2 * side;
  while (lo < hi) {
    const mid = Math.floor((lo + hi + 1) / 2);
    if (can(mid)) lo = mid;
    else hi = mid - 1;
  }

  return lo;
};
