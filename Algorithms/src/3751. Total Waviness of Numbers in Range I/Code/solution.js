/**
 * @param {number} num1
 * @param {number} num2
 * @return {number}
 */
var totalWaviness = function (num1, num2) {
  let answer = 0;

  // Check every number in the range
  for (let num = num1; num <= num2; num++) {
    const s = String(num);

    // Numbers with fewer than 3 digits have no peaks or valleys
    if (s.length < 3) {
      continue;
    }

    // Check every middle digit
    for (let i = 1; i < s.length - 1; i++) {
      // Peak condition
      if (s[i] > s[i - 1] && s[i] > s[i + 1]) {
        answer++;
      }
      // Valley condition
      else if (s[i] < s[i - 1] && s[i] < s[i + 1]) {
        answer++;
      }
    }
  }

  return answer;
};
