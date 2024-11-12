/**
 * @param {number[][]} items
 * @param {number[]} queries
 * @return {number[]}
 */
var maximumBeauty = function (items, queries) {
  items.sort((a, b) => a[0] - b[0]);
  let priceBeauty = [];
  let maxBeauty = 0;

  for (let [price, beauty] of items) {
    maxBeauty = Math.max(maxBeauty, beauty);
    priceBeauty.push([price, maxBeauty]);
  }

  return queries.map((query) => {
    let left = 0,
      right = priceBeauty.length - 1;
    while (left <= right) {
      let mid = left + Math.floor((right - left) / 2);
      if (priceBeauty[mid][0] <= query) left = mid + 1;
      else right = mid - 1;
    }
    return right >= 0 ? priceBeauty[right][1] : 0;
  });
};
