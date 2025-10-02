/**
 * @param {number} numBottles
 * @param {number} numExchange
 * @return {number}
 */
var maxBottlesDrunk = function (numBottles, numExchange) {
  let full = numBottles; // full bottles I have
  let empty = 0; // empties I have
  let ans = 0; // total drunk
  let curEx = numExchange; // current exchange requirement

  while (full > 0) {
    ans += full; // drink them
    empty += full;
    full = 0;

    // exchange one-by-one because requirement increases each time
    while (empty >= curEx) {
      empty -= curEx;
      full += 1;
      curEx += 1;
    }
  }
  return ans;
};
