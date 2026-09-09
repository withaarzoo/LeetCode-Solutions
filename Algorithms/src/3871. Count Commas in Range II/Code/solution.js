/**
 * @param {number} n
 * @return {number}
 */
var countCommas = function (n) {
  // Start at 1000 because every smaller number has zero commas.
  let start = 1000;

  // Numbers in this first group have one comma.
  let commas = 1;

  // Store the total number of commas.
  let answer = 0;

  // Process each group while its starting value is inside the range.
  while (start <= n) {
    // Check before multiplying so I do not needlessly create a value
    // larger than n and keep all useful calculations within the safe range.
    let end;

    if (start > n / 1000) {
      // The current group reaches only up to n.
      end = n;
    } else {
      // Otherwise, the group ends immediately before start * 1000.
      end = start * 1000 - 1;
    }

    // Count how many integers are inside the current group.
    const count = end - start + 1;

    // Every number in this group uses the same number of commas.
    answer += count * commas;

    // Move to the next group of numbers.
    start *= 1000;

    // The next group has one additional comma per number.
    commas++;
  }

  // Return the total number of commas.
  return answer;
};
