# 3734. Lexicographically Smallest Palindromic Permutation Greater Than Target

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

This problem asks us to find the lexicographically smallest palindromic permutation of `s` that is strictly greater than `target`.

We are given two strings, `s` and `target`, with the same length. Both contain only lowercase English letters.

We can rearrange the characters of `s`, but the final string must satisfy two conditions:

1. It must be a palindrome.
2. It must be lexicographically greater than `target`.

Among all valid palindromic permutations, we need to return the smallest one in lexicographical order.

If `s` cannot form any palindrome, or if no palindromic permutation is greater than `target`, we return an empty string.

The main challenge is that the length can be as large as `300`, so generating and checking every permutation would be far too expensive.

## Constraints

| Constraint         | Value                                                                                      |
| ------------------ | ------------------------------------------------------------------------------------------ |
| Length of `s`      | `1 <= s.length <= 300`                                                                     |
| Length of `target` | `target.length == s.length`                                                                |
| Characters         | Lowercase English letters                                                                  |
| Output             | Lexicographically smallest palindromic permutation strictly greater than `target`, or `""` |

## Intuition

The first thing I noticed is that generating every permutation of `s` is unnecessary.

A palindrome is completely determined by its first half. Once the left half is fixed, the right half must be its reverse. If the string length is odd, the middle character is also fixed.

For example, if the first half is:

```text
abc
```

then the palindrome must look like:

```text
abccba
```

If there is a middle character `d`, it becomes:

```text
abcdcba
```

This means I only need to think about arranging the characters that belong to the first half.

My next observation was about lexicographical comparison. If the first half of my palindrome becomes greater than the corresponding first half of `target`, the complete palindrome is automatically greater than `target`.

So the main goal becomes finding the smallest possible permutation of the first-half characters that can produce a palindrome strictly greater than `target`.

I first try to stay equal to the first half of `target` for as long as possible. If I cannot continue, I go backward and increase the rightmost possible character. Then I fill the remaining positions with the smallest available characters.

This gives the smallest possible larger first half. After building the palindrome, I check whether it is strictly greater than `target`.

If the first half is exactly equal but the resulting palindrome is still not greater, I move to the next lexicographical permutation of the first half.

## Approach

1. Count the frequency of every character in `s`.

2. Check whether `s` can form a palindrome.

   * A palindrome can have at most one character with an odd frequency.
   * If more than one character has an odd count, return `""`.

3. Build the frequency array for the first half.

   * Each character contributes `frequency / 2` copies.
   * If one character has an odd frequency, save it as the middle character.

4. Let `k = n / 2`, where `n` is the length of `s`.

5. Take the first `k` characters of `target`.

6. Find the smallest valid permutation of the first-half characters that is lexicographically greater than or equal to this target prefix.

   * Try matching the target prefix from left to right.
   * If matching fails, move backward.
   * Find the rightmost position where a larger available character can be placed.
   * Choose the smallest such character.
   * Fill all remaining positions in sorted order.

7. Build the complete palindrome using the selected first half.

8. If the palindrome is strictly greater than `target`, return it.

9. Otherwise, find the next lexicographical permutation of the first half.

   * If it exists, build and return the new palindrome.
   * If it does not exist, return `""`.

This approach avoids generating all palindromic permutations and directly constructs the answer.

## Data Structures Used

### Frequency Array

A frequency array of size `26` stores how many times each lowercase English letter appears.

I use it because the input contains only lowercase English letters, so a fixed-size array is faster and simpler than a map or hash table.

### First-Half Frequency Array

Another array of size `26` stores how many copies of each character belong in the first half of the palindrome.

For a character appearing `count` times, the first half receives:

```text
count / 2
```

copies.

### String or Character Array

The first half is stored as a string or character array.

A character array is useful when applying the next lexicographical permutation because characters need to be swapped and reversed efficiently.

### Result String

The final palindrome is built by combining:

```text
first half + middle character + reversed first half
```

## Operations & Behavior Summary

The algorithm works in the following order:

1. Count all characters in `s`.

2. Check the number of characters with odd frequencies.

   * More than one odd frequency means no palindrome is possible.

3. Divide every frequency by `2` to get the available characters for the first half.

4. Compare these characters against the first half of `target`.

5. Try to match the target prefix exactly.

6. If an exact match becomes impossible:

   * Move backward.
   * Restore previously used characters when needed.
   * Increase the rightmost possible position.
   * Use the smallest available larger character.

7. Sort the remaining available characters into the smallest possible suffix.

8. Mirror the first half to create the complete palindrome.

9. Compare the candidate with `target`.

10. If it is not strictly greater, generate the next lexicographical permutation of the first half and mirror it again.

11. Return the resulting palindrome or an empty string if no valid answer exists.

## Complexity

| Complexity       | Value  | Explanation                                                                                                                                                                                                          |
| ---------------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | `n` is the length of `s` and `target`. The algorithm counts frequencies, processes the first half, and builds the palindrome. The alphabet size is fixed at 26, so character searches are effectively constant time. |
| Space Complexity | `O(n)` | The first half and final palindrome require linear space. The frequency arrays use only `O(26)`, which is constant extra space.                                                                                      |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    // Builds the complete palindrome from its left half and optional middle character.
    string buildPalindrome(const string& half, char middle) {
        string result = half; // Start with the chosen left half.

        if (middle != 0) {
            result += middle; // Add the fixed middle character for odd-length strings.
        }

        // Mirror the left half in reverse order to complete the palindrome.
        for (int i = (int)half.size() - 1; i >= 0; --i) {
            result += half[i];
        }

        return result;
    }

    // Finds the lexicographically smallest permutation of the multiset
    // that is greater than or equal to targetHalf.
    string smallestGreaterOrEqual(vector<int> count, const string& targetHalf) {
        int k = targetHalf.size(); // Number of characters in the palindrome's first half.
        int matched = 0; // Number of target characters matched exactly so far.

        // Try to match targetHalf from left to right for as long as possible.
        while (matched < k && count[targetHalf[matched] - 'a'] > 0) {
            --count[targetHalf[matched] - 'a']; // Use this exact character.
            ++matched; // Move to the next position.
        }

        // If every position matched, targetHalf itself is a valid permutation.
        if (matched == k) {
            return targetHalf;
        }

        // Backtrack to find the rightmost position that can be increased.
        for (int pos = matched; pos >= 0; --pos) {
            // When moving left, restore the character that was previously matched.
            if (pos < matched) {
                ++count[targetHalf[pos] - 'a'];
            }

            // Choose the smallest available character strictly greater than targetHalf[pos].
            for (int c = targetHalf[pos] - 'a' + 1; c < 26; ++c) {
                if (count[c] == 0) continue; // This character is not available.

                string result = targetHalf.substr(0, pos); // Keep the prefix unchanged.
                result += char('a' + c); // Increase this position by the smallest possible amount.
                --count[c]; // Consume the chosen larger character.

                // Fill the remaining positions in ascending order for the smallest result.
                for (int ch = 0; ch < 26; ++ch) {
                    result.append(count[ch], char('a' + ch));
                }

                return result;
            }
        }

        return ""; // No permutation can be greater than or equal to targetHalf.
    }

    // Returns the next lexicographical permutation of the first half.
    bool nextPermutation(string& half) {
        int n = half.size();
        int pivot = n - 2; // Start by searching for the rightmost increasing position.

        // Find the rightmost position where half[pivot] < half[pivot + 1].
        while (pivot >= 0 && half[pivot] >= half[pivot + 1]) {
            --pivot;
        }

        // The whole sequence is non-increasing, so no larger permutation exists.
        if (pivot < 0) {
            return false;
        }

        int swapPos = n - 1; // Find the smallest larger character from the suffix.

        // Because the suffix is non-increasing, the first valid character from the right is correct.
        while (half[swapPos] <= half[pivot]) {
            --swapPos;
        }

        swap(half[pivot], half[swapPos]); // Increase the pivot position.

        // Reverse the suffix to make it as small as possible.
        reverse(half.begin() + pivot + 1, half.end());

        return true;
    }

    string lexPalindromicPermutation(string s, string target) {
        vector<int> frequency(26, 0); // Count every character in s.

        for (char ch : s) {
            ++frequency[ch - 'a'];
        }

        char middle = 0; // Stores the unique odd-frequency character, if needed.
        int oddCount = 0; // Counts how many characters have odd frequency.

        for (int c = 0; c < 26; ++c) {
            if (frequency[c] % 2 == 1) {
                ++oddCount;
                middle = char('a' + c);
            }
        }

        // A palindrome can have at most one odd-frequency character.
        if (oddCount > 1) {
            return "";
        }

        vector<int> halfCount(26, 0); // Frequency multiset for the first half.

        for (int c = 0; c < 26; ++c) {
            halfCount[c] = frequency[c] / 2; // Only half of each pair goes to the left side.
        }

        int k = s.size() / 2; // Length of the first half.
        string targetHalf = target.substr(0, k); // Only this prefix controls the first comparison.

        // Find the smallest possible first half that is at least targetHalf.
        string half = smallestGreaterOrEqual(halfCount, targetHalf);

        if (half.empty() && k > 0) {
            return ""; // No first-half permutation can reach targetHalf.
        }

        // Build the palindrome corresponding to this smallest valid first half.
        string candidate = buildPalindrome(half, middle);

        // If it is already strictly greater, it is the required smallest answer.
        if (candidate > target) {
            return candidate;
        }

        // Otherwise, move to the next larger first-half permutation.
        if (!nextPermutation(half)) {
            return "";
        }

        // The next first half produces the smallest possible larger palindrome.
        return buildPalindrome(half, middle);
    }
};
```

### Java

```java
class Solution {

    // Builds the complete palindrome from its left half and optional middle character.
    private String buildPalindrome(String half, char middle) {
        StringBuilder result = new StringBuilder(half); // Start with the chosen left half.

        if (middle != 0) {
            result.append(middle); // Add the fixed middle character for odd-length strings.
        }

        // Mirror the left half in reverse order to complete the palindrome.
        for (int i = half.length() - 1; i >= 0; --i) {
            result.append(half.charAt(i));
        }

        return result.toString();
    }

    // Finds the lexicographically smallest permutation of the multiset
    // that is greater than or equal to targetHalf.
    private String smallestGreaterOrEqual(int[] originalCount, String targetHalf) {
        int[] count = originalCount.clone(); // Work on a copy because counts are modified.
        int k = targetHalf.length(); // Number of characters in the first half.
        int matched = 0; // Number of target characters matched exactly so far.

        // Try to match targetHalf from left to right for as long as possible.
        while (matched < k && count[targetHalf.charAt(matched) - 'a'] > 0) {
            count[targetHalf.charAt(matched) - 'a']--; // Use this exact character.
            matched++; // Move to the next position.
        }

        // If every position matched, targetHalf itself is a valid permutation.
        if (matched == k) {
            return targetHalf;
        }

        // Backtrack to find the rightmost position that can be increased.
        for (int pos = matched; pos >= 0; --pos) {
            // Restore a previously matched character when moving left.
            if (pos < matched) {
                count[targetHalf.charAt(pos) - 'a']++;
            }

            // Choose the smallest available character strictly greater than targetHalf[pos].
            for (int c = targetHalf.charAt(pos) - 'a' + 1; c < 26; ++c) {
                if (count[c] == 0) continue; // This character is not available.

                StringBuilder result = new StringBuilder(targetHalf.substring(0, pos));
                result.append((char) ('a' + c)); // Increase this position minimally.
                count[c]--; // Consume the chosen larger character.

                // Fill every remaining position in ascending order.
                for (int ch = 0; ch < 26; ++ch) {
                    while (count[ch]-- > 0) {
                        result.append((char) ('a' + ch));
                    }
                    count[ch] = Math.max(count[ch], 0); // Keep the count non-negative.
                }

                return result.toString();
            }
        }

        return ""; // No valid permutation can reach targetHalf.
    }

    // Returns true and changes half to its next lexicographical permutation.
    private boolean nextPermutation(char[] half) {
        int pivot = half.length - 2; // Search for the rightmost increasing position.

        // Find the rightmost position that can be increased.
        while (pivot >= 0 && half[pivot] >= half[pivot + 1]) {
            pivot--;
        }

        // No larger permutation exists.
        if (pivot < 0) {
            return false;
        }

        int swapPos = half.length - 1; // Search from the end for the next larger character.

        // The first character from the right that is larger gives the smallest increase.
        while (half[swapPos] <= half[pivot]) {
            swapPos--;
        }

        char temp = half[pivot];
        half[pivot] = half[swapPos];
        half[swapPos] = temp;

        // Reverse the suffix so it becomes the smallest possible suffix.
        int left = pivot + 1;
        int right = half.length - 1;

        while (left < right) {
            temp = half[left];
            half[left] = half[right];
            half[right] = temp;
            left++;
            right--;
        }

        return true;
    }

    public String lexPalindromicPermutation(String s, String target) {
        int[] frequency = new int[26]; // Count every character in s.

        for (char ch : s.toCharArray()) {
            frequency[ch - 'a']++;
        }

        char middle = 0; // Stores the unique odd-frequency character.
        int oddCount = 0; // Counts odd frequencies.

        for (int c = 0; c < 26; ++c) {
            if ((frequency[c] & 1) == 1) {
                oddCount++;
                middle = (char) ('a' + c);
            }
        }

        // A palindrome can have at most one odd-frequency character.
        if (oddCount > 1) {
            return "";
        }

        int[] halfCount = new int[26]; // Store characters used in the first half.

        for (int c = 0; c < 26; ++c) {
            halfCount[c] = frequency[c] / 2;
        }

        int k = s.length() / 2; // Length of the first half.
        String targetHalf = target.substring(0, k); // Prefix that controls the first comparison.

        // Find the smallest first-half permutation that is at least targetHalf.
        String halfString = smallestGreaterOrEqual(halfCount, targetHalf);

        if (halfString.isEmpty() && k > 0) {
            return "";
        }

        // Build the corresponding palindrome.
        String candidate = buildPalindrome(halfString, middle);

        // If it is strictly greater, it is already the smallest valid answer.
        if (candidate.compareTo(target) > 0) {
            return candidate;
        }

        char[] half = halfString.toCharArray(); // Convert to an array for in-place permutation.

        // Otherwise, move to the next larger first-half permutation.
        if (!nextPermutation(half)) {
            return "";
        }

        // Build the palindrome from the next possible first half.
        return buildPalindrome(new String(half), middle);
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {string} target
 * @return {string}
 */
var lexPalindromicPermutation = function(s, target) {
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
    const halfCount = frequency.map(count => Math.floor(count / 2));
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
```

### Python3

```python
class Solution:

    def build_palindrome(self, half: str, middle: str) -> str:
        # The right half is forced to be the reverse of the left half.
        return half + middle + half[::-1]

    def smallest_greater_or_equal(self, original_count, target_half: str) -> str:
        # Work on a copy because the search changes character frequencies.
        count = original_count[:]
        k = len(target_half)
        matched = 0

        # Match the target prefix exactly for as long as possible.
        while matched < k and count[ord(target_half[matched]) - ord('a')] > 0:
            count[ord(target_half[matched]) - ord('a')] -= 1
            matched += 1

        # The complete target prefix itself can be formed.
        if matched == k:
            return target_half

        # Move backward and try to increase the rightmost possible position.
        for pos in range(matched, -1, -1):
            # Restore a character when backtracking over a matched position.
            if pos < matched:
                count[ord(target_half[pos]) - ord('a')] += 1

            current = ord(target_half[pos]) - ord('a')

            # Choose the smallest available character strictly larger than target_half[pos].
            for c in range(current + 1, 26):
                if count[c] == 0:
                    continue

                result = target_half[:pos] + chr(ord('a') + c)
                count[c] -= 1

                # Fill every remaining position in ascending order.
                for ch in range(26):
                    result += chr(ord('a') + ch) * count[ch]

                return result

        # No permutation can reach or exceed target_half.
        return ""

    def next_permutation(self, chars) -> bool:
        # Find the rightmost position that can be increased.
        pivot = len(chars) - 2

        while pivot >= 0 and chars[pivot] >= chars[pivot + 1]:
            pivot -= 1

        # The sequence is already the largest lexicographical permutation.
        if pivot < 0:
            return False

        swap_pos = len(chars) - 1

        # Find the smallest character larger than chars[pivot].
        while chars[swap_pos] <= chars[pivot]:
            swap_pos -= 1

        # Swap the pivot with that character.
        chars[pivot], chars[swap_pos] = chars[swap_pos], chars[pivot]

        # Reverse the suffix to make the next permutation as small as possible.
        chars[pivot + 1:] = reversed(chars[pivot + 1:])

        return True

    def lexPalindromicPermutation(self, s: str, target: str) -> str:
        # Count the frequency of every lowercase English letter.
        frequency = [0] * 26

        for ch in s:
            frequency[ord(ch) - ord('a')] += 1

        middle = ""  # Stores the only possible middle character.
        odd_count = 0

        # A valid palindrome can have at most one odd-frequency character.
        for c in range(26):
            if frequency[c] % 2 == 1:
                odd_count += 1
                middle = chr(ord('a') + c)

        if odd_count > 1:
            return ""

        # Only one character from every pair belongs to the first half.
        half_count = [count // 2 for count in frequency]
        k = len(s) // 2
        target_half = target[:k]

        # Find the smallest first-half permutation that is at least target_half.
        half = self.smallest_greater_or_equal(half_count, target_half)

        if not half and k > 0:
            return ""

        # Build the smallest candidate using that first half.
        candidate = self.build_palindrome(half, middle)

        if candidate > target:
            return candidate

        # If equality was not enough, move to the next first-half permutation.
        chars = list(half)

        if not self.next_permutation(chars):
            return ""

        # This is the smallest palindrome with a strictly larger first half.
        return self.build_palindrome("".join(chars), middle)
```

### Go

```go
func lexPalindromicPermutation(s string, target string) string {
 // Count the frequency of every lowercase English letter.
 frequency := make([]int, 26)
 for i := 0; i < len(s); i++ {
  frequency[s[i]-'a']++
 }

 // Find the optional middle character and count odd frequencies.
 middle := byte(0)
 oddCount := 0
 for c := 0; c < 26; c++ {
  if frequency[c]%2 == 1 {
   oddCount++
   middle = byte('a' + c)
  }
 }

 // A palindrome can have at most one odd-frequency character.
 if oddCount > 1 {
  return ""
 }

 // Build the frequency multiset for the first half.
 halfCount := make([]int, 26)
 for c := 0; c < 26; c++ {
  halfCount[c] = frequency[c] / 2
 }

 k := len(s) / 2
 targetHalf := target[:k]

 // Find the smallest first-half permutation greater than or equal to targetHalf.
 smallestGreaterOrEqual := func(originalCount []int, prefix string) string {
  // Copy counts because this search modifies them.
  count := make([]int, 26)
  copy(count, originalCount)

  matched := 0

  // Match the target prefix exactly while the required character exists.
  for matched < k && count[prefix[matched]-'a'] > 0 {
   count[prefix[matched]-'a']--
   matched++
  }

  // The exact target prefix can be formed.
  if matched == k {
   return prefix
  }

  // Move backward until one position can be increased.
  for pos := matched; pos >= 0; pos-- {
   // Restore a character when backtracking over a matched position.
   if pos < matched {
    count[prefix[pos]-'a']++
   }

   // Choose the smallest available character larger than prefix[pos].
   for c := int(prefix[pos]-'a') + 1; c < 26; c++ {
    if count[c] == 0 {
     continue
    }

    // Keep the earlier prefix unchanged.
    result := make([]byte, 0, k)
    result = append(result, prefix[:pos]...)
    result = append(result, byte('a'+c))
    count[c]--

    // Fill the suffix in ascending order to minimize the result.
    for ch := 0; ch < 26; ch++ {
     for times := 0; times < count[ch]; times++ {
      result = append(result, byte('a'+ch))
     }
    }

    return string(result)
   }
  }

  return ""
 }

 // Build the complete palindrome by mirroring the first half.
 buildPalindrome := func(half string) string {
  result := make([]byte, 0, len(s))
  result = append(result, half...)

  if middle != 0 {
   result = append(result, middle)
  }

  // Append the left half in reverse order.
  for i := len(half) - 1; i >= 0; i-- {
   result = append(result, half[i])
  }

  return string(result)
 }

 // Change a byte slice into its next lexicographical permutation.
 nextPermutation := func(chars []byte) bool {
  pivot := len(chars) - 2

  // Find the rightmost position that can be increased.
  for pivot >= 0 && chars[pivot] >= chars[pivot+1] {
   pivot--
  }

  // No larger permutation exists.
  if pivot < 0 {
   return false
  }

  swapPos := len(chars) - 1

  // Find the smallest larger character in the suffix.
  for chars[swapPos] <= chars[pivot] {
   swapPos--
  }

  // Swap to make the smallest possible increase.
  chars[pivot], chars[swapPos] = chars[swapPos], chars[pivot]

  // Reverse the suffix to make it as small as possible.
  left, right := pivot+1, len(chars)-1
  for left < right {
   chars[left], chars[right] = chars[right], chars[left]
   left++
   right--
  }

  return true
 }

 // Find the smallest possible first half that can reach targetHalf.
 half := smallestGreaterOrEqual(halfCount, targetHalf)

 if half == "" && k > 0 {
  return ""
 }

 // Build the smallest candidate from this first half.
 candidate := buildPalindrome(half)

 // If it is already strictly greater, it is the answer.
 if candidate > target {
  return candidate
 }

 // Otherwise, move to the next larger first-half permutation.
 chars := []byte(half)
 if !nextPermutation(chars) {
  return ""
 }

 // The next first half gives the lexicographically smallest valid answer.
 return buildPalindrome(string(chars))
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core logic is the same in C++, Java, JavaScript, Python3, and Go. The syntax changes, but the algorithm does not.

### 1. Count Character Frequencies

I first count how many times every character appears in `s`.

For example:

```text
s = "aacc"
```

gives:

```text
a -> 2
c -> 2
```

Since all frequencies are even, a palindromic permutation is possible.

For:

```text
s = "abc"
```

the frequencies are:

```text
a -> 1
b -> 1
c -> 1
```

There are three odd-frequency characters, so no palindrome can be formed.

This check is important because there is no reason to continue building permutations when the basic palindrome condition already fails.

### 2. Find the Middle Character

A palindrome can contain at most one character with an odd frequency.

If `s` has odd length and exactly one character appears an odd number of times, that character must be placed in the middle.

For example:

```text
s = "aabbc"
```

The frequencies are:

```text
a -> 2
b -> 2
c -> 1
```

So `c` becomes the middle character.

The final palindrome will always have this structure:

```text
left half + c + reversed(left half)
```

The middle character does not need to be permuted.

### 3. Build the Available First-Half Characters

For every pair of equal characters, one copy goes into the first half and the other automatically belongs to the mirrored second half.

For example:

```text
s = "aaaabb"
```

The first half receives:

```text
a -> 2
b -> 1
```

So the first half must be some permutation of:

```text
aab
```

Possible complete palindromes include:

```text
aabbaa
abaaba
baaaab
```

Instead of generating all of them, I search directly for the smallest valid first half.

### 4. Focus on the First Half of `target`

Let:

```text
k = n / 2
```

where `n` is the length of the strings.

I take the first `k` characters of `target`.

This prefix is important because the first half of a palindrome controls the earliest positions where the final string can differ from `target`.

If my first half becomes greater than the target prefix, the complete palindrome is already greater.

If my first half is smaller, the complete palindrome cannot become greater later because the lexicographical order has already been decided.

That is why I search for a first-half permutation that is at least the target prefix.

### 5. Try to Match the Target Prefix

I start from left to right.

At every position, I check whether the same character from `target` is available.

If it is available, I use it and continue.

For example, suppose the available characters are:

```text
a, a, c
```

and the target prefix is:

```text
aab
```

I can match:

```text
a
a
```

but I cannot match `b`.

At this point, continuing with an exact match is impossible.

### 6. Backtrack and Increase the Rightmost Possible Position

When an exact match fails, I move backward.

I want to change the rightmost possible position because changing a later character produces a smaller lexicographical increase.

At each position, I check whether an available character exists that is greater than the target character at that position.

If multiple larger characters are available, I choose the smallest one.

For the example:

```text
available characters = a, a, c
target prefix = aab
```

I cannot match `b`, but I can replace it with `c`.

The result becomes:

```text
aac
```

This is the smallest available permutation greater than:

```text
aab
```

### 7. Fill the Remaining Characters in Sorted Order

After choosing a larger character, the remaining characters should be placed in ascending order.

This is necessary because I want the smallest possible suffix.

For example, if the remaining characters are:

```text
a, c, d
```

the suffix should be:

```text
acd
```

and not:

```text
cad
```

Sorting the unused characters ensures that the complete first half is lexicographically minimal.

### 8. Build the Palindrome

Once the first half is ready, the rest of the palindrome is forced.

For an even-length string:

```text
half + reverse(half)
```

For an odd-length string:

```text
half + middle + reverse(half)
```

For example:

```text
half = "ab"
middle = "c"
```

produces:

```text
abcba
```

### 9. Compare the Candidate with `target`

I compare the complete palindrome with `target`.

If the candidate is strictly greater, it is the answer.

It is already the smallest possible candidate because the first half was constructed using the smallest valid lexicographical choice.

### 10. Handle an Equal First Half

There is one important edge case.

The first half may be exactly equal to the target prefix, but the full palindrome may still be less than or equal to `target`.

In that case, keeping the same first half will not work.

I need the smallest larger permutation of the first half.

This is where the next permutation algorithm is used.

### 11. Find the Next Lexicographical Permutation

The next permutation algorithm finds the smallest arrangement that is strictly larger than the current arrangement.

It works in three main steps:

1. Find the rightmost position that can be increased.
2. Swap it with the smallest suitable larger character from the suffix.
3. Reverse the remaining suffix to make it as small as possible.

This produces the immediate next lexicographical arrangement.

Since the palindrome is controlled by its first half, this next permutation gives the smallest possible larger palindrome.

### Language-Specific Behavior

The algorithm behaves the same in all five implementations.

* C++ can use `vector<int>`, `string`, and character operations.
* Java can use `int[]`, `StringBuilder`, and `char[]`.
* JavaScript can use arrays and strings.
* Python3 can use lists for frequency arrays and character manipulation.
* Go can use integer slices and byte slices.

The main difference is syntax. The palindrome construction, frequency counting, target-prefix comparison, and next permutation logic remain the same.

## Examples

### Example 1

**Input:**

```text
s = "baba"
target = "abba"
```

**Output:**

```text
"baab"
```

**How it works:**

The characters in `s` can form two palindromic permutations:

```text
abba
baab
```

In lexicographical order:

```text
abba < baab
```

Since the answer must be strictly greater than `"abba"`, the result is:

```text
baab
```

### Example 2

**Input:**

```text
s = "baba"
target = "bbaa"
```

**Output:**

```text
""
```

**How it works:**

The possible palindromic permutations are:

```text
abba
baab
```

Both are lexicographically smaller than `"bbaa"`.

So no valid answer exists, and the result is an empty string.

### Example 3

**Input:**

```text
s = "abc"
target = "abb"
```

**Output:**

```text
""
```

**How it works:**

Every character appears exactly once:

```text
a -> 1
b -> 1
c -> 1
```

There are three odd-frequency characters.

A palindrome can have at most one odd-frequency character, so `s` has no palindromic permutation.

The answer is:

```text
""
```

### Example 4

**Input:**

```text
s = "aac"
target = "abb"
```

**Output:**

```text
"aca"
```

**How it works:**

The only possible palindromic permutation is:

```text
aca
```

Comparing:

```text
aca > abb
```

So `"aca"` is the lexicographically smallest valid answer.

## How to Use / Run Locally

Each language solution can be saved in its own source file and compiled or executed locally.

### C++

Save the solution in:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

On Windows, use:

```bash
solution.exe
```

### Java

Save the solution in:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

Make sure Java is installed and available from your terminal.

### JavaScript

Save the solution in:

```text
solution.js
```

Run it using Node.js:

```bash
node solution.js
```

You can check whether Node.js is installed with:

```bash
node --version
```

### Python3

Save the solution in:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

### Go

Save the solution in:

```text
solution.go
```

Run it directly with:

```bash
go run solution.go
```

Or build an executable:

```bash
go build solution.go
```

Then run the generated executable.

## Notes & Optimizations

The most important optimization is avoiding the generation of all permutations.

Even with only a moderate number of characters, the number of permutations can become extremely large. Checking every permutation and testing whether it is a palindrome would be inefficient.

This solution uses the structure of a palindrome to reduce the problem to arranging only half of the characters.

A fixed frequency array of size `26` is used because the input contains only lowercase English letters. This makes character lookup simple and efficient.

The next permutation step is only needed when the smallest first half that is greater than or equal to the target prefix produces a palindrome that is not strictly greater than `target`.

Important edge cases include:

* `s` cannot form any palindrome.
* The only palindromic permutation is smaller than `target`.
* The smallest palindrome with an equal first half is exactly equal to `target`.
* No next permutation of the first half exists.
* The string length is `1`.

This approach is efficient for the given constraints and runs in linear time with respect to the input length, making it suitable for competitive programming and LeetCode submissions.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
