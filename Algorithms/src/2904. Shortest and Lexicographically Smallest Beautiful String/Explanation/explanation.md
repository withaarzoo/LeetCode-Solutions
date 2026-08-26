# 2904. Shortest and Lexicographically Smallest Beautiful String

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

This repository contains a solution for LeetCode problem 2904, Shortest and Lexicographically Smallest Beautiful String.

I am given a binary string `s` and a positive integer `k`.

A substring is called beautiful when it contains exactly `k` occurrences of the character `'1'`.

The goal is to find the shortest beautiful substring. If more than one shortest substring exists, I need to return the lexicographically smallest one.

If no substring contains exactly `k` ones, I return an empty string.

The main challenge is handling both conditions correctly. First, I need to minimize the length of the substring. Then, when two valid substrings have the same length, I need to choose the lexicographically smaller one.

This problem is a good example of using a sliding window to solve a substring problem efficiently.

## Constraints

| Constraint        | Value                             |
| ----------------- | --------------------------------- |
| Length of `s`     | `1 <= s.length <= 100`            |
| Value of `k`      | `1 <= k <= s.length`              |
| String characters | `s` contains only `'0'` and `'1'` |

## Intuition

My first observation was simple: a beautiful substring only needs exactly `k` ones.

So instead of checking every possible substring and counting its ones from scratch, I can keep track of a moving window and count how many ones are currently inside it.

I also noticed something important about leading zeros.

If a valid substring starts with one or more `'0'` characters, those zeros can be removed without changing the number of ones. Removing them makes the substring shorter.

For example, if `k = 2`, then:

`00101`

is beautiful because it contains exactly two ones. But:

`101`

is also beautiful and shorter.

This means that whenever I find a window containing exactly `k` ones, I should remove all unnecessary leading zeros before comparing it with my answer.

That observation naturally leads to a sliding window approach.

## Approach

I use two pointers called `left` and `right` to represent the current substring.

I also maintain a counter called `ones` to track how many `'1'` characters are inside the current window.

The process works like this:

1. Move `right` through the string one character at a time.
2. If the current character is `'1'`, increase the `ones` counter.
3. If the window contains more than `k` ones, move `left` forward until the window again contains at most `k` ones.
4. When the window contains exactly `k` ones, remove all leading zeros by moving `left` forward.
5. The remaining window is the shortest beautiful substring ending at the current `right` position.
6. Compare this substring with the current answer.
7. Replace the answer if the new substring is shorter.
8. If both substrings have the same length, keep the lexicographically smaller one.

After processing the entire string, the stored answer is the required shortest and lexicographically smallest beautiful substring.

## Data Structures Used

This solution does not need any complex data structures.

### String

I use a string variable to store the best beautiful substring found so far.

This makes it easy to compare candidate substrings by length and lexicographical order.

### Integer Variables

I use a few integer variables:

* `left` stores the starting position of the sliding window.
* `right` stores the ending position of the sliding window.
* `ones` stores the number of `'1'` characters inside the current window.

No arrays, stacks, queues, maps, or sets are required.

## Operations & Behavior Summary

The algorithm follows this flow:

1. Start with an empty answer.
2. Start the sliding window at the beginning of the string.
3. Expand the window by moving `right`.
4. Count every `'1'` that enters the window.
5. If there are more than `k` ones, shrink the window from the left.
6. Once exactly `k` ones remain, remove unnecessary leading zeros.
7. Treat the current window as a candidate answer.
8. Compare it with the best answer found so far.
9. Continue until every position has been used as the right boundary.
10. Return the best substring, or an empty string if no valid substring exists.

The key idea is that for every ending position, I reduce the window to the shortest possible valid substring before comparing it.

## Complexity

| Complexity       | Value   | Explanation                                                                                                                                                                                                    |
| ---------------- | ------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n²)` | The two pointers move through the string efficiently, but creating and lexicographically comparing candidate substrings can take up to `O(n)` time for multiple valid windows. Here, `n` is the length of `s`. |
| Space Complexity | `O(n)`  | The answer string and temporary candidate substring can contain up to `n` characters. Apart from strings, only a few variables are used.                                                                       |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string shortestBeautifulSubstring(string s, int k) {
        string answer = ""; // I store the best beautiful substring found so far.
        int left = 0;       // I keep the left boundary of the sliding window.
        int ones = 0;       // I count how many '1' characters are inside the window.

        // I expand the window by moving the right pointer through the string.
        for (int right = 0; right < s.size(); right++) {
            // I update the count when the newly added character is '1'.
            if (s[right] == '1') {
                ones++;
            }

            // If I have too many ones, I remove characters from the left
            // until the window contains at most k ones again.
            while (ones > k) {
                if (s[left] == '1') {
                    ones--;
                }
                left++;
            }

            // When I have exactly k ones, I remove leading zeros because
            // they only make the substring longer without adding any ones.
            while (ones == k && s[left] == '0') {
                left++;
            }

            // The current window is beautiful when it contains exactly k ones.
            if (ones == k) {
                int length = right - left + 1;

                // I extract the current shortest candidate ending at right.
                string candidate = s.substr(left, length);

                // I replace the answer if no answer exists yet, if this candidate
                // is shorter, or if equal-length candidates need lexicographical comparison.
                if (answer.empty() ||
                    candidate.size() < answer.size() ||
                    (candidate.size() == answer.size() && candidate < answer)) {
                    answer = candidate;
                }
            }
        }

        // I return the best substring, or an empty string if none was found.
        return answer;
    }
};
```

### Java

```java
class Solution {
    public String shortestBeautifulSubstring(String s, int k) {
        String answer = ""; // I store the best beautiful substring found so far.
        int left = 0;       // I keep the left boundary of the sliding window.
        int ones = 0;       // I count how many '1' characters are inside the window.

        // I expand the window by moving the right pointer through the string.
        for (int right = 0; right < s.length(); right++) {
            // I update the count when the newly added character is '1'.
            if (s.charAt(right) == '1') {
                ones++;
            }

            // If I have too many ones, I shrink the window from the left
            // until it contains at most k ones again.
            while (ones > k) {
                if (s.charAt(left) == '1') {
                    ones--;
                }
                left++;
            }

            // I remove leading zeros because they are unnecessary and only
            // make a valid substring longer.
            while (ones == k && s.charAt(left) == '0') {
                left++;
            }

            // The current window is beautiful when it contains exactly k ones.
            if (ones == k) {
                // I create the current candidate substring.
                String candidate = s.substring(left, right + 1);

                // I update the answer when this candidate is shorter, or when
                // equal-length candidates need lexicographical comparison.
                if (answer.isEmpty() ||
                    candidate.length() < answer.length() ||
                    (candidate.length() == answer.length() &&
                     candidate.compareTo(answer) < 0)) {
                    answer = candidate;
                }
            }
        }

        // I return the best result, or an empty string if no valid substring exists.
        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
var shortestBeautifulSubstring = function(s, k) {
    let answer = ""; // I store the best beautiful substring found so far.
    let left = 0;    // I keep the left boundary of the sliding window.
    let ones = 0;    // I count how many '1' characters are inside the window.

    // I expand the window by moving the right pointer through the string.
    for (let right = 0; right < s.length; right++) {
        // I update the count when the newly added character is '1'.
        if (s[right] === '1') {
            ones++;
        }

        // If I have too many ones, I remove characters from the left
        // until the window contains at most k ones again.
        while (ones > k) {
            if (s[left] === '1') {
                ones--;
            }
            left++;
        }

        // I remove unnecessary leading zeros from a valid window because
        // removing them keeps exactly k ones and makes the substring shorter.
        while (ones === k && s[left] === '0') {
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
```

### Python3

```python
class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        answer = ""  # I store the best beautiful substring found so far.
        left = 0     # I keep the left boundary of the sliding window.
        ones = 0     # I count how many '1' characters are inside the window.

        # I expand the window by moving the right pointer through the string.
        for right in range(len(s)):
            # I update the count when the newly added character is '1'.
            if s[right] == '1':
                ones += 1

            # If I have too many ones, I shrink the window from the left
            # until it contains at most k ones again.
            while ones > k:
                if s[left] == '1':
                    ones -= 1
                left += 1

            # I remove leading zeros because they do not affect the number
            # of ones and only make the current valid substring longer.
            while ones == k and s[left] == '0':
                left += 1

            # The current window is beautiful when it contains exactly k ones.
            if ones == k:
                # I create the current shortest candidate ending at right.
                candidate = s[left:right + 1]

                # I update the answer when this candidate is shorter, or when
                # equal-length candidates need lexicographical comparison.
                if (
                    not answer
                    or len(candidate) < len(answer)
                    or (len(candidate) == len(answer) and candidate < answer)
                ):
                    answer = candidate

        # I return the best substring, or an empty string if none was found.
        return answer
```

### Go

```go
func shortestBeautifulSubstring(s string, k int) string {
 answer := "" // I store the best beautiful substring found so far.
 left := 0    // I keep the left boundary of the sliding window.
 ones := 0    // I count how many '1' characters are inside the window.

 // I expand the window by moving the right pointer through the string.
 for right := 0; right < len(s); right++ {
  // I update the count when the newly added character is '1'.
  if s[right] == '1' {
   ones++
  }

  // If I have too many ones, I remove characters from the left
  // until the window contains at most k ones again.
  for ones > k {
   if s[left] == '1' {
    ones--
   }
   left++
  }

  // I remove unnecessary leading zeros because they do not change
  // the number of ones and only make the valid substring longer.
  for ones == k && s[left] == '0' {
   left++
  }

  // The current window is beautiful when it contains exactly k ones.
  if ones == k {
   // I extract the shortest valid candidate ending at right.
   candidate := s[left : right+1]

   // I update the answer if this candidate is shorter, or if equal
   // lengths require lexicographical comparison.
   if answer == "" ||
    len(candidate) < len(answer) ||
    (len(candidate) == len(answer) && candidate < answer) {
    answer = candidate
   }
  }
 }

 // I return the best substring, or an empty string if none exists.
 return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core logic is the same in C++, Java, JavaScript, Python3, and Go. Only the syntax for strings, loops, and substring operations changes.

I begin by storing an empty result. This represents the case where I have not found any beautiful substring yet.

Next, I initialize the `left` pointer at the beginning of the string and set the number of ones inside the current window to zero.

Then I move the `right` pointer from left to right through the string.

Whenever `right` reaches a `'1'`, I increase the count of ones. At this point, the window may contain fewer than `k`, exactly `k`, or more than `k` ones.

If it contains fewer than `k` ones, I cannot use the window yet. I simply continue expanding it.

If it contains more than `k` ones, the window is invalid. I move `left` forward and remove characters from the window until the number of ones is no longer greater than `k`.

When removing a character from the left side, I only decrease the `ones` counter if that character is `'1'`. Removing a `'0'` does not change the number of ones.

Once the window contains exactly `k` ones, I remove leading zeros.

This step is necessary because a leading zero does not help satisfy the requirement. The substring already has exactly `k` ones, so keeping that zero only makes the answer longer.

I keep moving `left` while the character at `left` is `'0'`.

After this, the current window is the shortest possible beautiful substring ending at the current `right` position.

I then compare it with the best answer found so far.

If no answer has been stored yet, I save the current substring.

If the current substring is shorter than the stored answer, I replace the answer.

If both substrings have the same length, I compare them lexicographically.

Lexicographical comparison works like dictionary order. I compare characters from left to right, and the first different character decides which string is smaller.

For binary strings, `'0'` is lexicographically smaller than `'1'`.

The language-specific comparison is slightly different:

* In C++, strings can be compared directly with comparison operators.
* In Java, `compareTo()` is used for lexicographical comparison.
* In JavaScript, strings can be compared directly using `<`.
* In Python3, strings can also be compared directly using `<`.
* In Go, strings support direct lexicographical comparison using `<`.

The actual algorithm and edge case handling remain the same in all five languages.

One important edge case is when the string does not contain enough ones. In that case, the algorithm never finds a valid window, so the answer remains empty.

Another important case is when multiple shortest beautiful substrings have the same length. The lexicographical comparison ensures that the smallest one is returned.

## Examples

### Example 1

**Input:**

```text
s = "100011001", k = 3
```

**Expected Output:**

```text
"11001"
```

The algorithm finds multiple substrings containing exactly three ones.

After checking their lengths, the shortest valid length is `5`.

Among the shortest candidates, `"11001"` is lexicographically smaller than the other possible substring of the same length, so it becomes the answer.

### Example 2

**Input:**

```text
s = "1011", k = 2
```

**Expected Output:**

```text
"11"
```

The substring `"101"` contains exactly two ones, but its length is `3`.

Later, the substring `"11"` is found. It also contains exactly two ones and has length `2`.

Since `"11"` is shorter, it becomes the answer.

### Example 3

**Input:**

```text
s = "000", k = 1
```

**Expected Output:**

```text
""
```

There are no `'1'` characters in the string.

Since no substring can contain exactly one `'1'`, there is no beautiful substring.

The algorithm returns an empty string.

## How to Use / Run Locally

### C++

Make sure a C++ compiler such as `g++` is installed.

Save the solution in a file named `solution.cpp`.

Compile it:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it:

```bash
./solution
```

On Windows:

```bash
solution.exe
```

### Java

Make sure the Java JDK is installed.

Save the code in a file that matches the class name used for testing.

Compile it:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

For LeetCode submissions, only the required `Solution` class needs to be submitted.

### JavaScript

Make sure Node.js is installed.

Save the solution in a file named `solution.js`.

Run it:

```bash
node solution.js
```

For LeetCode, the function can be submitted directly without additional input or output code.

### Python3

Make sure Python 3 is installed.

Save the solution in a file named `solution.py`.

Run it:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

For LeetCode, submit the `Solution` class with the required method.

### Go

Make sure Go is installed.

Save the solution in a file named `solution.go`.

Run it directly:

```bash
go run solution.go
```

Or build an executable:

```bash
go build solution.go
```

Then run the generated executable.

For LeetCode, submit the required function using the expected function signature.

## Notes & Optimizations

The main optimization in this solution is avoiding a full brute-force check of every possible substring.

A brute-force solution would generate many substrings and repeatedly count how many ones each substring contains. That creates unnecessary repeated work.

The sliding window keeps track of the number of ones while moving through the string.

Another useful observation is removing leading zeros from a valid window. This guarantees that whenever I compare a candidate ending at a particular position, I am using the shortest possible beautiful substring for that ending position.

The lexicographical comparison is only needed when two candidates have the same length. Checking length first avoids unnecessary string comparisons.

Since the constraint is only `s.length <= 100`, this solution is already easily fast enough for the problem.

The most important edge cases are:

* No substring contains exactly `k` ones.
* `k = 1`.
* The string contains only zeros.
* The entire string is the only beautiful substring.
* Multiple shortest beautiful substrings exist.
* Leading zeros appear before a valid group of `k` ones.

This approach is a clean combination of the sliding window technique, substring comparison, and lexicographical ordering for binary strings.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
