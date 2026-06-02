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
  // Store minimum answer
  let ans = Number.MAX_SAFE_INTEGER;

  const n = landStartTime.length;
  const m = waterStartTime.length;

  // Try every land-water pair
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < m; j++) {
      // Land -> Water
      const landFinish = landStartTime[i] + landDuration[i];
      const waterStart = Math.max(landFinish, waterStartTime[j]);
      const finish1 = waterStart + waterDuration[j];

      // Water -> Land
      const waterFinish = waterStartTime[j] + waterDuration[j];
      const landStart = Math.max(waterFinish, landStartTime[i]);
      const finish2 = landStart + landDuration[i];

      // Update minimum answer
      ans = Math.min(ans, finish1, finish2);
    }
  }

  return ans;
};
