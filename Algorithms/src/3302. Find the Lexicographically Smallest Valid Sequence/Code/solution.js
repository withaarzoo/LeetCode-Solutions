/**
 * @param {string} word1
 * @param {string} word2
 * @return {number[]}
 */
var validSequence = function (word1, word2) {
  const n = word1.length; // I store the length of word1 for the two scans.
  const m = word2.length; // I store the length of word2 because the answer has m indices.

  const last = new Array(m).fill(-1); // I store the right-to-left matching position for each suffix.

  let i = n - 1; // I start from the end of word1.
  let j = m - 1; // I start matching from the end of word2.

  // I build a right-to-left matching of word2 inside word1.
  // This tells me where the remaining suffix can be placed.
  while (i >= 0 && j >= 0) {
    // If the characters match, I can use this index for word2[j].
    if (word1[i] === word2[j]) {
      last[j] = i; // I remember this position for the suffix.
      --j; // I now need to match the previous character.
    }

    --i; // I continue searching toward the beginning of word1.
  }

  const ans = []; // I store the final lexicographically smallest sequence.
  let canSkip = true; // I have not used the one allowed mismatch yet.
  j = 0; // I start matching word2 from its first character.

  // I scan from left to right so the first valid index I choose is always the smallest.
  for (i = 0; i < n && j < m; ++i) {
    // If the characters already match, I take this earliest index.
    if (word1[i] === word2[j]) {
      ans.push(i); // I add the current index to the answer.
      ++j; // I move to the next character of word2.
    }
    // Otherwise, I can use the mismatch if the remaining suffix is still possible.
    else if (canSkip && (j === m - 1 || i < last[j + 1])) {
      canSkip = false; // I use the only allowed mismatch.
      ans.push(i); // I choose this earliest possible index.
      ++j; // I move to the next character of word2.
    }
  }

  // If I matched all characters, the constructed sequence is valid.
  if (j === m) {
    return ans;
  }

  // If some character could not be matched, no valid sequence exists.
  return [];
};
