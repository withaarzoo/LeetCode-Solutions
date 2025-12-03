/**
 * @param {number[][]} points
 * @return {number}
 */
var countTrapezoids = function (points) {
  const n = points.length;
  const SHIFT = 3000;

  const encodePair = (a, b) => ((a + SHIFT) << 13) ^ (b + SHIFT);

  const gcd = (a, b) => {
    a = Math.abs(a);
    b = Math.abs(b);
    while (b !== 0) {
      const t = a % b;
      a = b;
      b = t;
    }
    return a;
  };

  // slopeKey -> (lineId -> count)
  const bySlope = new Map();
  // vectorKey -> (lineId -> count)
  const byVector = new Map();

  const addTo = (outer, key, lineId) => {
    if (!outer.has(key)) outer.set(key, new Map());
    const inner = outer.get(key);
    inner.set(lineId, (inner.get(lineId) || 0) + 1);
  };

  for (let i = 0; i < n; ++i) {
    const [x1, y1] = points[i];
    for (let j = i + 1; j < n; ++j) {
      const [x2, y2] = points[j];

      let dx = x2 - x1;
      let dy = y2 - y1;

      if (dx < 0 || (dx === 0 && dy < 0)) {
        dx = -dx;
        dy = -dy;
      }

      const g = gcd(dx, dy);
      const ux = dx / g;
      const uy = dy / g;

      const lineId = ux * y1 - uy * x1;

      const slopeKey = encodePair(ux, uy);
      const vectorKey = encodePair(dx, dy);

      addTo(bySlope, slopeKey, lineId);
      addTo(byVector, vectorKey, lineId);
    }
  }

  function countPairs(map) {
    let res = 0;
    for (const inner of map.values()) {
      let sum = 0;
      let sumSq = 0;
      for (const c of inner.values()) {
        sum += c;
        sumSq += c * c;
      }
      res += (sum * sum - sumSq) / 2;
    }
    return res;
  }

  const withParallel = countPairs(bySlope);
  const parallelogramTwo = countPairs(byVector);

  return withParallel - Math.floor(parallelogramTwo / 2);
};
