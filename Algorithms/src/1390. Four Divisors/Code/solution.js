/**
 * @param {number[]} nums
 * @return {number}
 */
var sumFourDivisors = function (nums) {
  let totalSum = 0;

  for (let num of nums) {
    let cnt = 0;
    let sum = 0;

    for (let d = 1; d * d <= num; d++) {
      if (num % d === 0) {
        let other = num / d;

        cnt++;
        sum += d;

        if (other !== d) {
          cnt++;
          sum += other;
        }

        if (cnt > 4) break;
      }
    }

    if (cnt === 4) {
      totalSum += sum;
    }
  }

  return totalSum;
};
