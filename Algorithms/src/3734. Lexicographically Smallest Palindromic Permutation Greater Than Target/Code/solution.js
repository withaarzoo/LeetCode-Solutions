/**
 * @param {string} s
 * @param {string} target
 * @return {string}
 */
var lexPalindromicPermutation = function (s, target) {
  // Count every character in s.
  const frequency = Array(26).fill(0);

  for (const ch of s) {
    frequency[ch.charCodeAt(0) - 97]++;
  }

  let middle = ""; // Stores the fixed middle character for odd-length strings.
  let oddCount = 0; // Counts characters with odd frequency.

  for (let c = 0; c < 26; c++) {
    if (frequency[c] % 2 === 1) {
      oddCount++;
      middle = String.fromCharCode(97 + c);
    }
  }

  // A palindrome cannot contain more than one odd-frequency character.
  if (oddCount > 1) {
    return "";
  }

  // Each pair contributes one character to the first half.
  const halfCount = frequency.map((count) => Math.floor(count / 2));
  const k = Math.floor(s.length / 2);
  const targetHalf = target.slice(0, k);

  // Find the smallest multiset permutation greater than or equal to targetHalf.
  function smallestGreaterOrEqual(originalCount, prefix) {
    const count = [...originalCount]; // Copy because this function modifies frequencies.
    let matched = 0;

    // Match the target prefix for as long as the required character exists.
    while (matched < k && count[prefix.charCodeAt(matched) - 97] > 0) {
      count[prefix.charCodeAt(matched) - 97]--;
      matched++;
    }

    // The exact prefix itself is possible.
    if (matched === k) {
      return prefix;
    }

    // Move backward until one position can be increased.
    for (let pos = matched; pos >= 0; pos--) {
      // Restore the character used at this position when backtracking.
      if (pos < matched) {
        count[prefix.charCodeAt(pos) - 97]++;
      }

      const current = prefix.charCodeAt(pos) - 97;

      // Use the smallest available character larger than prefix[pos].
      for (let c = current + 1; c < 26; c++) {
        if (count[c] === 0) continue;

        let result = prefix.slice(0, pos);
        result += String.fromCharCode(97 + c);
        count[c]--;

        // Fill the suffix in ascending order to minimize the result.
        for (let ch = 0; ch < 26; ch++) {
          result += String.fromCharCode(97 + ch).repeat(count[ch]);
        }

        return result;
      }
    }

    return ""; // No permutation can be greater than or equal to the prefix.
  }

  // Build the full palindrome from the selected first half.
  function buildPalindrome(half) {
    const reversed = half.split("").reverse().join(""); // Mirror the first half.
    return half + middle + reversed;
  }

  // Change chars into their next lexicographical permutation.
  function nextPermutation(chars) {
    let pivot = chars.length - 2;

    // Find the rightmost position that can be increased.
    while (pivot >= 0 && chars[pivot] >= chars[pivot + 1]) {
      pivot--;
    }

    // The current arrangement is already the largest one.
    if (pivot < 0) {
      return false;
    }

    let swapPos = chars.length - 1;

    // Find the smallest character larger than the pivot from the suffix.
    while (chars[swapPos] <= chars[pivot]) {
      swapPos--;
    }

    // Swap to make the smallest possible increase.
    [chars[pivot], chars[swapPos]] = [chars[swapPos], chars[pivot]];

    // Reverse the suffix so it becomes as small as possible.
    let left = pivot + 1;
    let right = chars.length - 1;

    while (left < right) {
      [chars[left], chars[right]] = [chars[right], chars[left]];
      left++;
      right--;
    }

    return true;
  }

  // Find the smallest first half that can match or exceed targetHalf.
  let half = smallestGreaterOrEqual(halfCount, targetHalf);

  if (half === "" && k > 0) {
    return "";
  }

  // Build and test the smallest candidate.
  const candidate = buildPalindrome(half);

  if (candidate > target) {
    return candidate;
  }

  // The equal first half was not enough, so move to the next permutation.
  const chars = half.split("");

  if (!nextPermutation(chars)) {
    return "";
  }

  // The next first half gives the smallest possible larger palindrome.
  return buildPalindrome(chars.join(""));
};
