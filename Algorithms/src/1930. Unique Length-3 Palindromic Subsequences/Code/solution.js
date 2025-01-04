/**
 * @param {string} s
 * @return {number}
 */
var countPalindromicSubsequence = function (s) {
  let first = Array(26).fill(-1);
  let last = Array(26).fill(-1);

  for (let i = 0; i < s.length; i++) {
    let index = s.charCodeAt(i) - 97;
    if (first[index] === -1) first[index] = i;
    last[index] = i;
  }

  let result = 0;
  for (let i = 0; i < 26; i++) {
    if (first[i] !== -1 && last[i] > first[i]) {
      let middleChars = new Set();
      for (let j = first[i] + 1; j < last[i]; j++) {
        middleChars.add(s[j]);
      }
      result += middleChars.size;
    }
  }

  return result;
};
