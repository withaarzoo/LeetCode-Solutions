/**
 * @param {number[]} cost
 * @return {number}
 */
var minimumCost = function (cost) {
  // Sort from highest cost to lowest cost
  cost.sort((a, b) => b - a);

  let total = 0;

  // Skip every third candy
  for (let i = 0; i < cost.length; i++) {
    if (i % 3 === 2) continue;

    total += cost[i];
  }

  return total;
};
