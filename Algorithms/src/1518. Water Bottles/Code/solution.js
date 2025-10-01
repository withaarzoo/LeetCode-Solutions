/**
 * @param {number} numBottles
 * @param {number} numExchange
 * @return {number}
 */
var numWaterBottles = function (numBottles, numExchange) {
  let total = numBottles;
  let empties = numBottles;

  while (empties >= numExchange) {
    let newFull = Math.floor(empties / numExchange);
    total += newFull;
    empties = newFull + (empties % numExchange);
  }

  return total;
};
