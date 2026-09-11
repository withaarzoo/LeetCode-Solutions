/**
 * @param {number[]} digits
 * @return {number}
 */
var totalNumbers = function (digits) {
  // I store how many times each digit appears.
  const freq = Array(10).fill(0);

  // I build the frequency table to handle duplicate digits correctly.
  for (const digit of digits) {
    freq[digit]++;
  }

  // I store the number of distinct valid 3-digit even numbers.
  let answer = 0;

  // I choose the hundreds digit from 1 to 9 because a leading zero is not allowed.
  for (let first = 1; first <= 9; first++) {
    // I choose the tens digit from 0 to 9.
    for (let second = 0; second <= 9; second++) {
      // I choose only even digits for the ones position.
      for (let third = 0; third <= 8; third += 2) {
        // I skip this combination if any required digit is unavailable.
        if (freq[first] === 0 || freq[second] === 0 || freq[third] === 0) {
          continue;
        }

        // Three equal digits require three copies of the same digit.
        if (first === second && second === third && freq[first] < 3) {
          continue;
        }

        // Equal first and second digits require two copies.
        if (first === second && freq[first] < 2) {
          continue;
        }

        // Equal first and third digits require two copies.
        if (first === third && freq[first] < 2) {
          continue;
        }

        // Equal second and third digits require two copies.
        if (second === third && freq[second] < 2) {
          continue;
        }

        // This combination represents one distinct valid number.
        answer++;
      }
    }
  }

  // I return the total number of valid distinct numbers.
  return answer;
};
