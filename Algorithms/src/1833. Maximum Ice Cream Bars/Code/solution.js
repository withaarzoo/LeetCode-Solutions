/**
 * @param {number[]} costs
 * @param {number} coins
 * @return {number}
 */
var maxIceCream = function (costs, coins) {
  // Maximum possible cost according to constraints
  const MAX_COST = 100000;

  // Frequency array to count occurrences of each cost
  const freq = new Array(MAX_COST + 1).fill(0);

  // Count every cost
  for (const cost of costs) {
    freq[cost]++;
  }

  // Total purchased ice cream bars
  let answer = 0;

  // Buy from cheapest to most expensive
  for (let cost = 1; cost <= MAX_COST; cost++) {
    // Skip if this cost does not exist
    if (freq[cost] === 0) continue;

    // Maximum bars affordable at current cost
    const canBuy = Math.min(freq[cost], Math.floor(coins / cost));

    // Add purchased bars
    answer += canBuy;

    // Remove spent coins
    coins -= canBuy * cost;
  }

  return answer;
};
