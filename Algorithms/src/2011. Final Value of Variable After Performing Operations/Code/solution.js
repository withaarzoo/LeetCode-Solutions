/**
 * @param {string[]} operations
 * @return {number}
 */
var finalValueAfterOperations = function (operations) {
  let X = 0; // start from 0
  for (let op of operations) {
    // if op includes '+' => increment, else decrement
    if (op.indexOf("+") !== -1) X++;
    else X--;
  }
  return X;
};
