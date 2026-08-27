# 3720. Lexicographically Smallest Permutation Greater Than Target

## Table of Contents

* <a href="#problem-summary">Problem Summary</a>
* <a href="#constraints">Constraints</a>
* <a href="#intuition">Intuition</a>
* <a href="#approach">Approach</a>
* <a href="#data-structures-used">Data Structures Used</a>
* <a href="#operations--behavior-summary">Operations & Behavior Summary</a>
* <a href="#complexity">Complexity</a>
* <a href="#multi-language-solutions">Multi-language Solutions</a>
* <a href="#step-by-step-detailed-explanation-c-java-javascript-python3-go">Step-by-step Detailed Explanation</a>
* <a href="#examples">Examples</a>
* <a href="#how-to-use--run-locally">How to Use / Run Locally</a>
* <a href="#notes--optimizations">Notes & Optimizations</a>
* <a href="#author">Author</a>

## Problem Summary

This repository contains an optimized solution for LeetCode 3720, Lexicographically Smallest Permutation Greater Than Target.

I am given two strings, `s` and `target`, with the same length. Both strings contain only lowercase English letters.

The task is to rearrange the characters of `s` and find the lexicographically smallest permutation that is strictly greater than `target`.

If no permutation of `s` is greater than `target`, I return an empty string.

The main challenge is finding the answer without generating every possible permutation. Since the number of permutations can grow extremely fast, a brute-force approach is not practical.

This solution uses a greedy approach with character frequency counting to build the smallest valid permutation efficiently.

## Constraints

| Constraint         | Value                          |
| ------------------ | ------------------------------ |
| Length of `s`      | `1 <= s.length <= 300`         |
| Length of `target` | `target.length == s.length`    |
| Characters         | Lowercase English letters only |

Because only lowercase English letters are used, I only need to track `26` character frequencies.

## Intuition

The first thing I noticed is that generating all permutations is unnecessary.

To get the lexicographically smallest answer that is greater than `target`, I want to keep the beginning of the answer equal to `target` for as long as possible.

At each position, my first choice is simple: if the current character of `target` is still available in `s`, I use it.

If I cannot match the current character, I try to place the smallest available character that is greater than it.

Sometimes that is not possible either. In that case, I need to go backward and change one of the previously matched positions.

I always try to change the rightmost possible position first. This is important because keeping more characters at the beginning equal to `target` produces a lexicographically smaller result.

Once I place one character that is strictly greater than the corresponding character in `target`, the answer is already valid. Then I simply place all remaining characters in sorted order to make the suffix as small as possible.

## Approach

I use a frequency array to count all characters in `s`.

Then I follow these steps:

1. Start from the first position.
2. Try to match `target[i]` using an available character from `s`.
3. If it is available, use it and continue to the next position.
4. If it is not available, try to use the smallest available character greater than `target[i]`.
5. If such a character exists, use it and append all remaining characters in sorted order.
6. If no larger character exists, move backward through the positions that were previously matched.
7. Restore the character used at that position back into the frequency array.
8. Try to replace it with the smallest available character that is strictly greater.
9. As soon as this is possible, append all remaining characters in sorted order.
10. If no position can be increased, return an empty string.

This greedy algorithm avoids generating permutations and takes advantage of the fact that the alphabet has only `26` lowercase letters.

## Data Structures Used

### Frequency Array

I use an array of size `26`.

Each index represents one lowercase English letter:

* Index `0` represents `a`
* Index `1` represents `b`
* Index `2` represents `c`
* And so on

The value stored at each index tells me how many copies of that character are still available.

This makes it easy to:

* Check whether a character can be used
* Remove a character after using it
* Restore a character while backtracking
* Find the smallest available character greater than another character
* Build the remaining suffix in sorted order

Because the array always has size `26`, these character searches are effectively constant-time operations.

### Result String

I use a result string to build the final permutation.

The prefix remains equal to `target` until I find the position where the answer becomes strictly greater. After that, I append the remaining characters in alphabetical order.

## Operations & Behavior Summary

The algorithm works like this:

```text
Count the frequency of every character in s

Match target from left to right:
    If target[i] is available:
        Use it
    Otherwise:
        Stop matching

Starting from the current position:
    Try to find the smallest available character > target[i]

If found:
    Keep the earlier prefix unchanged
    Use that larger character
    Append all remaining characters in sorted order
    Return the answer

Otherwise:
    Move one position left
    Restore the previously matched character
    Try again

If no position can be increased:
    Return an empty string
```

The important part is that I try positions from right to left when backtracking. This ensures that the first position where the final answer differs from `target` is as far right as possible.

## Complexity

| Complexity       | Value  | Explanation                                                                                                                                                         |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | `n` is the length of `s` and `target`. I process the strings and build the final answer once. Character searches check at most `26` letters, which is constant.     |
| Space Complexity | `O(n)` | The returned answer and remaining suffix can contain up to `n` characters. The frequency array itself uses only `O(1)` extra space because its size is always `26`. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string lexGreaterPermutation(string s, string target) {
        // Store the frequency of every character available in s.
        vector<int> count(26, 0);
        for (char ch : s) {
            count[ch - 'a']++;
        }

        int n = s.size();
        int matched = 0;

        // Try to keep the answer exactly equal to target for as long as possible.
        while (matched < n && count[target[matched] - 'a'] > 0) {
            // Use target[matched] because matching it keeps the prefix smallest.
            count[target[matched] - 'a']--;
            matched++;
        }

        // If matching failed before reaching the end, first try to increase
        // exactly at the position where matching became impossible.
        int start = (matched < n ? matched : n - 1);

        // Move from right to left because changing a later position gives
        // a lexicographically smaller answer than changing an earlier one.
        for (int i = start; i >= 0; i--) {
            // If this position was previously matched, undo that choice
            // so its character becomes available again.
            if (i < matched) {
                count[target[i] - 'a']++;
            }

            // Find the smallest available character strictly greater than target[i].
            int bigger = -1;
            for (int ch = target[i] - 'a' + 1; ch < 26; ch++) {
                if (count[ch] > 0) {
                    bigger = ch;
                    break;
                }
            }

            // If a larger character exists, build the smallest possible answer.
            if (bigger != -1) {
                // Use the larger character to make the whole string > target.
                count[bigger]--;

                // Keep everything before i equal to target.
                string answer = target.substr(0, i);

                // Place the smallest possible larger character at position i.
                answer += char('a' + bigger);

                // Append all remaining characters in sorted order.
                for (int ch = 0; ch < 26; ch++) {
                    answer.append(count[ch], char('a' + ch));
                }

                return answer;
            }
        }

        // No position can be increased, so no valid permutation exists.
        return "";
    }
};
```

### Java

```java
class Solution {
    public String lexGreaterPermutation(String s, String target) {
        // Store the frequency of every character available in s.
        int[] count = new int[26];
        for (char ch : s.toCharArray()) {
            count[ch - 'a']++;
        }

        int n = s.length();
        int matched = 0;

        // Try to keep the answer exactly equal to target for as long as possible.
        while (matched < n && count[target.charAt(matched) - 'a'] > 0) {
            // Use target[matched] because matching it keeps the prefix smallest.
            count[target.charAt(matched) - 'a']--;
            matched++;
        }

        // Start where matching failed, or at the last position if all matched.
        int start = matched < n ? matched : n - 1;

        // Move backward because changing a later position gives a smaller answer.
        for (int i = start; i >= 0; i--) {
            // Restore this character if it was previously used to match target.
            if (i < matched) {
                count[target.charAt(i) - 'a']++;
            }

            // Find the smallest available character greater than target[i].
            int bigger = -1;
            for (int ch = target.charAt(i) - 'a' + 1; ch < 26; ch++) {
                if (count[ch] > 0) {
                    bigger = ch;
                    break;
                }
            }

            // Build the answer when an increasing character is found.
            if (bigger != -1) {
                // Consume the character used to make the answer strictly greater.
                count[bigger]--;

                // Keep the prefix unchanged.
                StringBuilder answer = new StringBuilder(target.substring(0, i));

                // Add the smallest possible character that is greater here.
                answer.append((char) ('a' + bigger));

                // Add every remaining character in sorted order.
                for (int ch = 0; ch < 26; ch++) {
                    while (count[ch]-- > 0) {
                        answer.append((char) ('a' + ch));
                    }
                }

                return answer.toString();
            }
        }

        // No permutation is strictly greater than target.
        return "";
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
var lexGreaterPermutation = function(s, target) {
    // Store the frequency of every character available in s.
    const count = Array(26).fill(0);
    for (const ch of s) {
        count[ch.charCodeAt(0) - 97]++;
    }

    const n = s.length;
    let matched = 0;

    // Match target from left to right for as long as possible.
    while (
        matched < n &&
        count[target.charCodeAt(matched) - 97] > 0
    ) {
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
```

### Python3

```python
class Solution:
    def lexGreaterPermutation(self, s: str, target: str) -> str:
        # Store the frequency of every character available in s.
        count = [0] * 26
        for ch in s:
            count[ord(ch) - ord('a')] += 1

        n = len(s)
        matched = 0

        # Match target from left to right for as long as possible.
        while matched < n and count[ord(target[matched]) - ord('a')] > 0:
            # Using the same character keeps the current prefix smallest.
            count[ord(target[matched]) - ord('a')] -= 1
            matched += 1

        # Start where matching failed, or at the last position if all matched.
        start = matched if matched < n else n - 1

        # Move backward because increasing a later position gives a smaller answer.
        for i in range(start, -1, -1):
            # Restore the character if it was previously used to match target.
            if i < matched:
                count[ord(target[i]) - ord('a')] += 1

            # Find the smallest available character greater than target[i].
            bigger = -1
            for ch in range(ord(target[i]) - ord('a') + 1, 26):
                if count[ch] > 0:
                    bigger = ch
                    break

            # Build the answer as soon as this position can be increased.
            if bigger != -1:
                # Consume the character used to make the answer greater.
                count[bigger] -= 1

                # Keep the prefix unchanged and place the larger character.
                answer = target[:i] + chr(ord('a') + bigger)

                # Append all remaining characters in sorted order.
                for ch in range(26):
                    answer += chr(ord('a') + ch) * count[ch]

                return answer

        # No permutation is strictly greater than target.
        return ""
```

### Go

```go
func lexGreaterPermutation(s string, target string) string {
 // Store the frequency of every lowercase character available in s.
 count := make([]int, 26)
 for i := 0; i < len(s); i++ {
  count[s[i]-'a']++
 }

 n := len(s)
 matched := 0

 // Match target from left to right for as long as possible.
 for matched < n && count[target[matched]-'a'] > 0 {
  // Using the same character keeps the current prefix smallest.
  count[target[matched]-'a']--
  matched++
 }

 // Start where matching failed, or at the last position if all matched.
 start := n - 1
 if matched < n {
  start = matched
 }

 // Try to increase the answer from right to left.
 for i := start; i >= 0; i-- {
  // Restore this character if it was previously used to match target.
  if i < matched {
   count[target[i]-'a']++
  }

  // Find the smallest available character strictly greater than target[i].
  bigger := -1
  for ch := int(target[i]-'a') + 1; ch < 26; ch++ {
   if count[ch] > 0 {
    bigger = ch
    break
   }
  }

  // Build the answer when an increasing character is found.
  if bigger != -1 {
   // Consume the character that makes the string strictly greater.
   count[bigger]--

   // Keep the prefix before i unchanged.
   answer := make([]byte, 0, n)
   answer = append(answer, target[:i]...)

   // Add the smallest available character greater than target[i].
   answer = append(answer, byte('a'+bigger))

   // Append all remaining characters in sorted order.
   for ch := 0; ch < 26; ch++ {
    for count[ch] > 0 {
     answer = append(answer, byte('a'+ch))
     count[ch]--
    }
   }

   return string(answer)
  }
 }

 // No permutation is strictly greater than target.
 return ""
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The same greedy logic is used in all five language versions. Only the syntax for strings, arrays, and character handling changes.

### Step 1: Count all characters

I first count how many times each character appears in `s`.

For example, if:

```text
s = "leet"
```

the available characters are:

```text
e -> 2
l -> 1
t -> 1
```

I store these counts in an array of size `26`.

This is better than repeatedly searching through the original string because I can immediately check whether a character is still available.

### Step 2: Match the target from left to right

I start at position `0`.

If `target[i]` is available, I use that character and decrease its frequency.

I do this because matching the current character keeps the answer as small as possible.

For example:

```text
s      = "abc"
target = "bba"
```

At position `0`, I can use `b`, so I do that.

Now the remaining characters are:

```text
a, c
```

At position `1`, I need another `b` to keep matching `target`, but no `b` remains.

This means I cannot continue with an equal prefix.

### Step 3: Try to increase the current position

Now I look for the smallest remaining character greater than `target[i]`.

In the previous example:

```text
target[1] = 'b'
remaining = 'a', 'c'
```

The smallest character greater than `b` is `c`.

So I choose:

```text
b c
```

At this point, the answer is already strictly greater than `target`.

The first difference is:

```text
target = b b a
answer = b c ...
           ^
```

Since `c` is greater than `b`, the remaining characters no longer need to follow any restriction.

### Step 4: Sort the remaining suffix

After finding the position where the answer becomes greater, I append all unused characters in alphabetical order.

This gives the smallest possible suffix.

For the same example, the only remaining character is `a`, so the final answer is:

```text
bca
```

If I placed the remaining characters in a different order, I could create a larger valid answer, but not the lexicographically smallest one.

### Step 5: Backtrack when the current position cannot be increased

Sometimes no available character is greater than `target[i]`.

Then changing the current position cannot produce a valid answer.

I move one step backward.

Before trying that earlier position, I restore the character that was previously used to match `target`.

This is necessary because that character is no longer fixed at its old position. It becomes available again as part of the remaining character pool.

I continue moving from right to left until I find a position that can be increased.

### Step 6: Why I move from right to left

Suppose I have two possible places where I can make the answer greater.

Changing a later position is always better because it allows the answer to match `target` for a longer prefix.

For example, keeping:

```text
a a ...
```

and changing a later character is lexicographically smaller than changing the second character to something larger:

```text
a b ...
```

That is why I always try the rightmost matched position first.

### Step 7: Return an empty string when no answer exists

If I reach the first position and still cannot place any larger character, then every permutation of `s` is less than or equal to `target`.

In that case, there is no valid answer.

So I return:

```text
""
```

### Language-specific behavior

The algorithm itself does not change between C++, Java, JavaScript, Python3, and Go.

The main implementation differences are:

* C++ can use `vector<int>` and `string`
* Java can use `int[]` and `StringBuilder`
* JavaScript can use a normal array and string operations
* Python3 can use a list for frequencies and string concatenation or a list-based result
* Go can use an integer slice and a byte slice for efficient string construction

In every version, the frequency array contains exactly `26` values and follows the same greedy and backtracking process.

## Examples

### Example 1

**Input**

```text
s = "abc"
target = "bba"
```

**Output**

```text
"bca"
```

**How it works**

* Use `b` to match the first character.
* Another `b` is not available for the second position.
* The smallest remaining character greater than `b` is `c`.
* The remaining character `a` is appended.
* Final answer: `bca`.

### Example 2

**Input**

```text
s = "leet"
target = "coed"
```

**Output**

```text
"eelt"
```

**How it works**

* At the first position, `c` is not available.
* The smallest available character greater than `c` is `e`.
* Once `e` is chosen, the answer is already greater than `target`.
* The remaining characters are placed in sorted order.
* Final answer: `eelt`.

### Example 3

**Input**

```text
s = "baba"
target = "bbaa"
```

**Output**

```text
""
```

**How it works**

* I try to keep the prefix equal to `target`.
* When matching or increasing becomes impossible, I backtrack.
* No position can be replaced with an available character that makes the result strictly greater.
* Therefore, no valid permutation exists.

## How to Use / Run Locally

Clone this repository and move into the project folder:

```bash
git clone <your-repository-url>
cd <your-repository-folder>
```

### Run C++

Make sure a C++ compiler such as `g++` is installed.

Compile the file:

```bash
g++ solution.cpp -o solution
```

Run it:

```bash
./solution
```

On Windows:

```bash
solution.exe
```

### Run Java

Make sure the JDK is installed.

Compile the file:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

If the solution is being used directly on LeetCode, only the `Solution` class needs to be submitted.

### Run JavaScript

Make sure Node.js is installed.

Run the file with:

```bash
node solution.js
```

### Run Python3

Make sure Python 3 is installed.

Run the file with:

```bash
python3 solution.py
```

On some Windows systems, the command may be:

```bash
python solution.py
```

### Run Go

Make sure Go is installed.

Run the file with:

```bash
go run solution.go
```

## Notes & Optimizations

A brute-force solution would generate all permutations of `s`, which can become extremely expensive. With repeated characters, handling duplicate permutations also adds unnecessary complexity.

The greedy approach works because lexicographical comparison depends on the first position where two strings differ.

I always try to keep the prefix equal to `target`. If that becomes impossible, I try to make the first difference as far right as possible.

At that position, I choose the smallest available character that is strictly greater than the target character.

After that point, sorting the remaining characters gives the smallest possible suffix.

The fixed alphabet size is another useful optimization. Since only lowercase English letters are allowed, searching for a larger character requires checking at most `26` possibilities.

Important edge cases include:

* `s` itself is equal to `target`
* The largest permutation of `s` is not greater than `target`
* The first mismatch happens at position `0`
* A larger character is unavailable at the mismatch position, requiring backtracking
* Multiple copies of the same character exist
* The answer becomes greater only at the final position

This makes the solution efficient, clean, and suitable for competitive programming constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/?utm_source=chatgpt.com)
