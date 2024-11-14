/**
 * @param {number} n
 * @param {number[]} quantities
 * @return {number}
 */
var minimizedMaximum = function (n, quantities) {
  const canDistribute = (maxProducts) => {
    let storesNeeded = 0;
    for (const quantity of quantities) {
      storesNeeded += Math.ceil(quantity / maxProducts);
      if (storesNeeded > n) return false;
    }
    return storesNeeded <= n;
  };

  let low = 1;
  let high = Math.max(...quantities);
  let answer = high;

  while (low <= high) {
    const mid = Math.floor((low + high) / 2);
    if (canDistribute(mid)) {
      answer = mid;
      high = mid - 1;
    } else {
      low = mid + 1;
    }
  }

  return answer;
};
