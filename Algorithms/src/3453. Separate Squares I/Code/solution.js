var separateSquares = function (squares) {
  let totalArea = 0;
  let low = 1e18,
    high = -1e18;

  for (let [x, y, l] of squares) {
    totalArea += l * l;
    low = Math.min(low, y);
    high = Math.max(high, y + l);
  }

  for (let i = 0; i < 80; i++) {
    let mid = (low + high) / 2;
    let areaBelow = 0;

    for (let [x, y, l] of squares) {
      if (mid <= y) continue;
      if (mid >= y + l) areaBelow += l * l;
      else areaBelow += l * (mid - y);
    }

    if (areaBelow * 2 < totalArea) low = mid;
    else high = mid;
  }
  return low;
};
