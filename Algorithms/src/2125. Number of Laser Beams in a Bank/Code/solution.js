/**
 * @param {string[]} bank
 * @return {number}
 */
var numberOfBeams = function (bank) {
  let ans = 0;
  let prev = 0; // count of '1's in previous non-empty row

  for (const row of bank) {
    let cnt = 0;
    // count devices in this row
    for (let i = 0; i < row.length; ++i) {
      if (row[i] === "1") cnt++;
    }
    if (cnt > 0) {
      ans += prev * cnt;
      prev = cnt;
    }
  }
  return ans;
};
