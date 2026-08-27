/**
 * @param {string} s
 * @param {string} target
 * @return {string}
 */
var lexGreaterPermutation = function (s, target) {
  // Store the frequency of every character available in s.
  const count = Array(26).fill(0);
  for (const ch of s) {
    count[ch.charCodeAt(0) - 97]++;
  }

  const n = s.length;
  let matched = 0;

  // Match target from left to right for as long as possible.
  while (matched < n && count[target.charCodeAt(matched) - 97] > 0) {
    // Using the same character keeps the current prefix smallest.
    count[target.charCodeAt(matched) - 97]--;
    matched++;
  }

  // Start at the failed position, or at the last position if all matched.
  const start = matched < n ? matched : n - 1;

  // Try to increase the answer from right to left.
  for (let i = start; i >= 0; i--) {
    // Restore the character previously used at this matched position.
    if (i < matched) {
      count[target.charCodeAt(i) - 97]++;
    }

    // Find the smallest available character strictly greater than target[i].
    let bigger = -1;
    for (let ch = target.charCodeAt(i) - 97 + 1; ch < 26; ch++) {
      if (count[ch] > 0) {
        bigger = ch;
        break;
      }
    }

    // Once this position can be increased, the rest should be sorted.
    if (bigger !== -1) {
      // Consume the character that makes the string strictly greater.
      count[bigger]--;

      // Keep the prefix before i exactly equal to target.
      let answer = target.slice(0, i);

      // Add the smallest available character greater than target[i].
      answer += String.fromCharCode(97 + bigger);

      // Add all remaining characters in ascending order.
      for (let ch = 0; ch < 26; ch++) {
        answer += String.fromCharCode(97 + ch).repeat(count[ch]);
      }

      return answer;
    }
  }

  // No position can be increased.
  return "";
};
