var minNumberOfSeconds = function (mountainHeight, workerTimes) {
  const can = (time) => {
    let totalHeight = 0n;

    for (let t of workerTimes) {
      let left = 0n;
      let right = BigInt(mountainHeight);

      while (left <= right) {
        let mid = (left + right) / 2n;

        let required = BigInt(t) * ((mid * (mid + 1n)) / 2n);

        if (required <= time) {
          left = mid + 1n;
        } else {
          right = mid - 1n;
        }
      }

      totalHeight += right;

      if (totalHeight >= BigInt(mountainHeight)) return true;
    }

    return false;
  };

  let left = 1n;
  let right = 10n ** 18n;
  let ans = right;

  while (left <= right) {
    let mid = (left + right) / 2n;

    if (can(mid)) {
      ans = mid;
      right = mid - 1n;
    } else {
      left = mid + 1n;
    }
  }

  return Number(ans);
};
