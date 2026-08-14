# 3090. Maximum Length Substring With Two Occurrences

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

LeetCode 3090, **Maximum Length Substring With Two Occurrences**, asks us to find the longest substring where every character appears at most two times.

The input is a string `s` containing only lowercase English letters.

For example, if the string is `bcbbbcba`, a valid substring cannot contain three `b` characters. We need to find the longest continuous part of the string that follows this rule.

The output is a single integer representing the maximum length of such a substring.

This problem is a good example of the **sliding window**, **two pointers**, and **frequency counting** techniques commonly used in competitive programming and DSA interviews.

## Constraints

* `2 <= s.length <= 100`
* `s` consists only of lowercase English letters.
* Every character in the selected substring can appear at most `2` times.

## Intuition

I first noticed that I am not looking for any random set of characters. I need a **substring**, so all selected characters must be next to each other.

A simple approach would be to generate every possible substring and count how many times each character appears. But that repeats a lot of work.

Instead, I can maintain one window over the string.

I keep two pointers, `left` and `right`. The `right` pointer expands the window, while `left` moves forward whenever the window becomes invalid.

The only thing that can make the window invalid is a character appearing more than two times.

So whenever a character reaches a count of `3`, I move `left` forward until that character appears at most two times again.

This gives a clean **sliding window solution with O(n) time complexity**.

## Approach

I use a sliding window with two pointers.

1. Start `left` at the beginning of the string.
2. Move `right` from left to right.
3. Add the current character to a frequency array.
4. If its frequency becomes greater than `2`, move `left` forward.
5. Decrease the frequency of every character that leaves the window.
6. Continue until the current window becomes valid.
7. Calculate the current window length.
8. Store the maximum length found so far.
9. Return the maximum length.

Because both pointers only move forward, each character is processed a limited number of times.

The important idea is that I do not restart the search whenever a window becomes invalid. I simply shrink the existing window and continue.

## Data Structures Used

### Frequency Array

I use an integer array of size `26`.

Each position represents one lowercase English letter:

* Index `0` represents `a`
* Index `1` represents `b`
* Index `2` represents `c`
* ...
* Index `25` represents `z`

The array stores how many times each character appears inside the current sliding window.

Since the input contains only lowercase English letters, an array is simpler and faster than using a hash map.

### Two Pointers

I use two indexes:

* `left` represents the start of the current substring.
* `right` represents the end of the current substring.

Together, they define the current sliding window.

## Operations & Behavior Summary

The algorithm works like this:

1. Create a frequency array with 26 positions.
2. Set `left = 0`.
3. Start moving `right` through the string.
4. Increase the frequency of `s[right]`.
5. Check whether the current character now appears more than two times.
6. If it does, move `left` forward.
7. Decrease the frequency of characters removed from the window.
8. Stop shrinking when the window becomes valid again.
9. Calculate `right - left + 1`.
10. Update the maximum answer.
11. Continue until `right` reaches the end of the string.
12. Return the maximum length.

In simple pseudocode:

```text
start left at 0
answer = 0

for every right from 0 to n - 1:
    add s[right] to the frequency array

    while s[right] appears more than 2 times:
        remove s[left] from the frequency array
        move left forward

    update answer using the current window length

return answer
```

## Complexity

| Complexity       |   Cost | Explanation                                                                             |
| ---------------- | -----: | --------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | `right` moves through the string once, and `left` also moves forward at most `n` times. |
| Space Complexity | `O(1)` | The frequency array always contains exactly 26 positions, regardless of the input size. |

Here, `n` is the length of the input string `s`.

The sliding window is efficient because I never move either pointer backward.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximumLengthSubstring(string s) {
        int freq[26] = {};

        int left = 0;

        int ans = 0;

        for (int right = 0; right < s.size(); right++) {
            freq[s[right] - 'a']++;

            while (freq[s[right] - 'a'] > 2) {
                freq[s[left] - 'a']--;

                left++;
            }

            ans = max(ans, right - left + 1);
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int maximumLengthSubstring(String s) {
        int[] freq = new int[26];

        int left = 0;

        int ans = 0;

        for (int right = 0; right < s.length(); right++) {
            freq[s.charAt(right) - 'a']++;

            while (freq[s.charAt(right) - 'a'] > 2) {
                freq[s.charAt(left) - 'a']--;

                left++;
            }

            ans = Math.max(ans, right - left + 1);
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {number}
 */
var maximumLengthSubstring = function(s) {
    const freq = new Array(26).fill(0);

    let left = 0;

    let ans = 0;

    for (let right = 0; right < s.length; right++) {
        const index = s.charCodeAt(right) - 97;
        freq[index]++;

        while (freq[index] > 2) {
            freq[s.charCodeAt(left) - 97]--;

            left++;
        }

        ans = Math.max(ans, right - left + 1);
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def maximumLengthSubstring(self, s: str) -> int:
        freq = [0] * 26

        left = 0

        ans = 0

        for right in range(len(s)):
            index = ord(s[right]) - ord('a')
            freq[index] += 1

            while freq[index] > 2:
                freq[ord(s[left]) - ord('a')] -= 1

                left += 1

            ans = max(ans, right - left + 1)

        return ans
```

### Go

```go
func maximumLengthSubstring(s string) int {
 freq := make([]int, 26)

 left := 0

 ans := 0

 for right := 0; right < len(s); right++ {
  index := int(s[right] - 'a')
  freq[index]++

  for freq[index] > 2 {
   freq[int(s[left]-'a')]--

   left++
  }

  length := right - left + 1
  if length > ans {
   ans = length
  }
 }

 return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core logic is the same in all five languages. Only the syntax used to create arrays, access characters, and calculate indexes changes.

### Step 1: Create the frequency array

I create an array containing 26 zeros.

This array keeps track of the number of times each lowercase character appears inside the current window.

For example, if the current window contains:

```text
a b a c
```

then the frequency for `a` is `2`, `b` is `1`, and `c` is `1`.

I use an array instead of a map because the problem guarantees lowercase English letters.

### Step 2: Set the left pointer

I start with:

```text
left = 0
```

This means the sliding window initially starts at the first character.

The `right` pointer will expand the window from left to right.

At any point, the current substring is:

```text
s[left ... right]
```

### Step 3: Move the right pointer

For every position from the beginning to the end of the string, I add the current character to the window.

I increase its frequency by `1`.

For example, if the current character is `b`, I increase the frequency stored for `b`.

This keeps the frequency array synchronized with the current window.

### Step 4: Check the frequency limit

The problem allows each character to appear at most twice.

So if the current character reaches a frequency of `3`, the current window is invalid.

For example:

```text
b c b b
```

Here, `b` appears three times.

I cannot keep this entire window because it violates the problem condition.

### Step 5: Shrink from the left

When the window becomes invalid, I start moving `left` forward.

Before moving it, I remove the character at `left` from the frequency array.

Then I increment `left`.

I repeat this until the invalid character appears at most twice.

For example:

```text
b c b b
```

Removing the first `b` gives:

```text
c b b
```

Now `b` appears only twice, so the window is valid again.

### Step 6: Update the answer

Once the window is valid, I calculate its length:

```text
right - left + 1
```

I compare this value with the maximum answer found so far.

If the current window is longer, I update the answer.

### Step 7: Continue scanning

I continue moving `right` until I reach the end of the string.

There is no need to restart the search after shrinking the window.

The current valid window already gives me the best possible substring ending at the current `right` position.

### C++ behavior

In C++, I use a fixed-size integer array with 26 positions.

Character indexes can be calculated using:

```text
s[index] - 'a'
```

This converts a lowercase letter into a number from `0` to `25`.

### Java behavior

In Java, I use an integer array of size `26`.

A character can be converted to an array index by subtracting `'a'`.

The rest of the sliding window logic is exactly the same.

### JavaScript behavior

In JavaScript, I create an array with 26 zero values.

I can convert a character into its alphabet position using `charCodeAt()`.

For example, the character code of `a` is used as the starting point, so subtracting it gives an index between `0` and `25`.

### Python3 behavior

In Python, I use a list containing 26 zeros.

The `ord()` function converts a character into its numeric character code.

Subtracting the code for `a` gives the required array index.

Python's `max()` function can then be used to update the best answer.

### Go behavior

In Go, I use an integer slice with 26 positions.

Since the input contains lowercase English letters, subtracting `'a'` from a character gives the correct alphabet index.

The sliding window works the same way as in the other implementations.

## Examples

### Example 1

**Input:**

```text
s = "bcbbbcba"
```

**Output:**

```text
4
```

The window can contain at most two occurrences of every character.

As soon as `b` appears for the third time, I move the left pointer forward until the window becomes valid again.

One longest valid substring has length `4`.

### Example 2

**Input:**

```text
s = "aaaa"
```

**Output:**

```text
2
```

Only two `a` characters are allowed.

The window grows to:

```text
aa
```

When the third `a` is added, I move `left` forward and remove the first `a`.

The same thing happens for the fourth `a`.

So the longest valid substring has length `2`.

### Example 3

**Input:**

```text
s = "abcabc"
```

**Output:**

```text
6
```

Every character appears exactly twice in the complete string.

Since no character appears more than two times, the entire string is a valid substring.

Therefore, the answer is `6`.

## How to Use / Run Locally

The repository contains solutions for C++, Java, JavaScript, Python3, and Go.

You can copy the solution for your preferred language into a local file and test it with the sample inputs.

### C++

Save the solution as:

```text
solution.cpp
```

Compile it using:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Then run:

```bash
./solution
```

If you are using Windows with MinGW, you can run the generated executable directly.

### Java

Save the solution as:

```text
Solution.java
```

Compile it using:

```bash
javac Solution.java
```

Then run:

```bash
java Solution
```

### JavaScript

Save the solution as:

```text
solution.js
```

Make sure Node.js is installed.

Run:

```bash
node solution.js
```

### Python3

Save the solution as:

```text
solution.py
```

Run:

```bash
python3 solution.py
```

On some Windows installations, you may need:

```bash
python solution.py
```

### Go

Save the solution as:

```text
solution.go
```

Run it directly with:

```bash
go run solution.go
```

You can also build an executable with:

```bash
go build solution.go
```

## Notes & Optimizations

The main edge case is a string containing the same character many times.

For example:

```text
aaaaaa
```

The window can never contain more than two `a` characters, so the answer is `2`.

Another simple case is when every character appears at most twice in the entire string. In that situation, the complete string is already a valid answer.

The fixed-size frequency array is a good choice here because the input contains only lowercase English letters. A hash map would also work, but it is unnecessary for this constraint.

A brute-force solution can check many possible substrings, but it repeats character-counting work. The sliding window avoids that repeated work and gives `O(n)` time.

The two-pointer technique is the key optimization. The `right` pointer only moves forward, and the `left` pointer also only moves forward. This is why the nested `while` loop does not make the overall solution `O(n²)`.

This problem is useful practice for recognizing when a substring problem can be solved with **sliding window**, **two pointers**, and **frequency counting**.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
