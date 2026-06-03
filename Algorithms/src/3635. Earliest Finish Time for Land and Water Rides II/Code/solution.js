/**
 * @param {number[]} landStartTime
 * @param {number[]} landDuration
 * @param {number[]} waterStartTime
 * @param {number[]} waterDuration
 * @return {number}
 */
var earliestFinishTime = function (
  landStartTime,
  landDuration,
  waterStartTime,
  waterDuration,
) {
  // Computes the best answer when category A is taken first
  const solve = (startA, durA, startB, durB) => {
    const rides = [];

    // Store (start, duration)
    for (let i = 0; i < startB.length; i++) {
      rides.push([startB[i], durB[i]]);
    }

    // Sort by start time
    rides.sort((a, b) => a[0] - b[0]);

    const m = rides.length;

    const starts = new Array(m);
    const prefixMinDur = new Array(m);
    const suffixMinFinish = new Array(m);

    for (let i = 0; i < m; i++) {
      starts[i] = rides[i][0];

      if (i === 0) prefixMinDur[i] = rides[i][1];
      else prefixMinDur[i] = Math.min(prefixMinDur[i - 1], rides[i][1]);
    }

    for (let i = m - 1; i >= 0; i--) {
      const finish = rides[i][0] + rides[i][1];

      if (i === m - 1) suffixMinFinish[i] = finish;
      else suffixMinFinish[i] = Math.min(suffixMinFinish[i + 1], finish);
    }

    let ans = Number.MAX_SAFE_INTEGER;

    for (let i = 0; i < startA.length; i++) {
      const finish1 = startA[i] + durA[i];

      let left = 0;
      let right = m;

      // Upper bound: first start > finish1
      while (left < right) {
        const mid = Math.floor((left + right) / 2);

        if (starts[mid] <= finish1) left = mid + 1;
        else right = mid;
      }

      const pos = left;

      if (pos > 0) {
        ans = Math.min(ans, finish1 + prefixMinDur[pos - 1]);
      }

      if (pos < m) {
        ans = Math.min(ans, suffixMinFinish[pos]);
      }
    }

    return ans;
  };

  return Math.min(
    solve(landStartTime, landDuration, waterStartTime, waterDuration),
    solve(waterStartTime, waterDuration, landStartTime, landDuration),
  );
};
