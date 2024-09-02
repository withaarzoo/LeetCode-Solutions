/**
 * @param {number[]} chalk - An array where each element represents the amount of chalk a student will use.
 * @param {number} k - The total amount of chalk initially available.
 * @return {number} - The index of the student who will replace the chalk.
 */
var chalkReplacer = function (chalk, k) {
  // Step 1: Calculate the total chalk required for one complete round of all students.
  // We sum up all the elements in the chalk array to get the total chalk usage for one round.
  let totalChalk = chalk.reduce((acc, c) => acc + c, 0);

  // Step 2: Use the modulo operation to reduce k.
  // The idea here is to find out how much chalk will be left after completing several full rounds.
  // Since completing a full round does not affect the outcome, we only need to consider the remainder.
  k %= totalChalk;

  // Step 3: Iterate through the chalk array to determine which student will be unable to continue.
  for (let i = 0; i < chalk.length; i++) {
    // If the remaining chalk (k) is less than the current student's chalk requirement,
    // that student will be the one to replace the chalk, so we return their index.
    if (k < chalk[i]) {
      return i;
    }
    // Otherwise, subtract the current student's chalk usage from k and continue to the next student.
    k -= chalk[i];
  }

  // Step 4: Safety return - this line should technically never be reached
  // because the problem guarantees that there is a student who will run out of chalk.
  return -1;
};
