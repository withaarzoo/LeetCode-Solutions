/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
var shortestBeautifulSubstring = function (s, k) {
  let answer = ""; // I store the best beautiful substring found so far.
  let left = 0; // I keep the left boundary of the sliding window.
  let ones = 0; // I count how many '1' characters are inside the window.

  // I expand the window by moving the right pointer through the string.
  for (let right = 0; right < s.length; right++) {
    // I update the count when the newly added character is '1'.
    if (s[right] === "1") {
      ones++;
    }

    // If I have too many ones, I remove characters from the left
    // until the window contains at most k ones again.
    while (ones > k) {
      if (s[left] === "1") {
        ones--;
      }
      left++;
    }

    // I remove unnecessary leading zeros from a valid window because
    // removing them keeps exactly k ones and makes the substring shorter.
    while (ones === k && s[left] === "0") {
      left++;
    }

    // The window is beautiful when it contains exactly k ones.
    if (ones === k) {
      // I extract the shortest valid candidate ending at right.
      const candidate = s.substring(left, right + 1);

      // I update the answer if this candidate is shorter, or if equal
      // lengths require me to choose the lexicographically smaller string.
      if (
        answer === "" ||
        candidate.length < answer.length ||
        (candidate.length === answer.length && candidate < answer)
      ) {
        answer = candidate;
      }
    }
  }

  // I return the best substring, or an empty string if none exists.
  return answer;
};
