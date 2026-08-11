/**
 * @param {number[]} nums
 * @return {number}
 */
var missingInteger = function (nums) {
  // I start the sum with the first element because
  // the first element always belongs to the sequential prefix.
  let sum = nums[0];

  // I scan from the second element to find where the sequence breaks.
  for (let i = 1; i < nums.length; i++) {
    // The prefix continues only when the current value
    // is exactly one greater than the previous value.
    if (nums[i] === nums[i - 1] + 1) {
      sum += nums[i];
    } else {
      // The sequential prefix ends at the previous element.
      break;
    }
  }

  // I store all values in a Set so I can check membership quickly.
  const seen = new Set(nums);

  // I begin checking from the sum of the longest sequential prefix.
  let answer = sum;

  // If the candidate exists in the array, I keep increasing it.
  while (seen.has(answer)) {
    answer++;
  }

  // The first missing candidate is the required answer.
  return answer;
};
