# Rotate String (LeetCode 796) – String Rotation Check Using Concatenation

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

This problem is about checking whether one string can be transformed into another using rotation.

You are given two strings `s` and `goal`. You need to determine if `s` can become `goal` after performing any number of shifts.

A shift means taking the first character of the string and moving it to the end.

For example:

* "abcde" → "bcdea" after one shift

Your task is to return `true` if `goal` can be obtained this way, otherwise return `false`.

This is a classic string rotation problem often asked in coding interviews and competitive programming.

## Constraints

* 1 ≤ s.length, goal.length ≤ 100
* Both strings contain only lowercase English letters

## Intuition

At first, I thought about simulating every rotation manually. But that felt repetitive.

Then I noticed something important.

If I take the string `s` and concatenate it with itself, all possible rotations of `s` will appear as substrings inside it.

Example:

* s = "abcde"
* s + s = "abcdeabcde"

Now every rotation like "bcdea", "cdeab", etc. is already inside this string.

So instead of rotating again and again, I can just check if `goal` exists inside `s + s`.

## Approach

I followed a simple and clean approach:

1. First, I check if both strings have the same length

   * If not, rotation is impossible

2. Then I create a new string by concatenating `s` with itself

   * This helps capture all possible rotations

3. Finally, I check if `goal` is a substring of this new string

   * If yes, return true
   * Otherwise, return false

This approach avoids unnecessary loops and keeps the logic clean.

## Data Structures Used

* String
  I only use basic string operations like concatenation and substring search. No extra data structures are needed.

## Operations & Behavior Summary

Here’s what the algorithm does step by step:

* Compare lengths of both strings
* Create a new string by doubling the original string
* Search for `goal` inside this doubled string
* Return the result based on whether it is found or not

This is efficient and avoids manual rotation simulation.

## Complexity

| Type             | Complexity |
| ---------------- | ---------- |
| Time Complexity  | O(n)       |
| Space Complexity | O(n)       |

Explanation:

* `n` is the length of the string
* Concatenation takes O(n)
* Substring search also takes O(n) in average cases
* Extra space is used for the new string `s + s`

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool rotateString(string s, string goal) {
        // If lengths are different, rotation is impossible
        if (s.length() != goal.length()) return false;

        // Concatenate s with itself
        string doubled = s + s;

        // Check if goal is a substring of doubled string
        return doubled.find(goal) != string::npos;
    }
};
```

### Java

```java
class Solution {
    public boolean rotateString(String s, String goal) {
        // If lengths are different, rotation is impossible
        if (s.length() != goal.length()) return false;

        // Concatenate s with itself
        String doubled = s + s;

        // Check if goal exists in doubled string
        return doubled.contains(goal);
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {string} goal
 * @return {boolean}
 */
var rotateString = function(s, goal) {
    // If lengths are different, rotation is impossible
    if (s.length !== goal.length) return false;

    // Concatenate s with itself
    let doubled = s + s;

    // Check if goal is a substring
    return doubled.includes(goal);
};
```

### Python3

```python
class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        # If lengths are different, rotation is impossible
        if len(s) != len(goal):
            return False

        # Concatenate s with itself
        doubled = s + s

        # Check if goal is inside doubled string
        return goal in doubled
```

### Go

```go
func rotateString(s string, goal string) bool {
    // If lengths are different, rotation is impossible
    if len(s) != len(goal) {
        return false
    }

    // Concatenate s with itself
    doubled := s + s

    // Check if goal exists in doubled string
    return strings.Contains(doubled, goal)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same across all languages.

I start by checking if the lengths of `s` and `goal` are equal.
This is necessary because rotation does not change string length.

If the lengths are different, I immediately return false.

Next, I create a new string by doing `s + s`.
This step is the key idea behind the solution.

Why this works:
Every rotation of `s` is guaranteed to be a substring of `s + s`.

Then I perform a substring check:

* In C++ → `find()`
* In Java → `contains()`
* In JavaScript → `includes()`
* In Python → `in`
* In Go → `strings.Contains()`

If `goal` is found, I return true. Otherwise false.

If I had used a loop to rotate the string manually, it would still work, but it would be less clean and slightly more complex.

## Examples

### Example 1

Input:
s = "abcde"
goal = "cdeab"

Output:
true

Explanation:
After 2 rotations, "abcde" becomes "cdeab"

---

### Example 2

Input:
s = "abcde"
goal = "abced"

Output:
false

Explanation:
No rotation of "abcde" can produce "abced"

---

### Example 3

Input:
s = "aaaa"
goal = "aaaa"

Output:
true

Explanation:
All rotations are the same since all characters are identical

## How to Use / Run Locally

### C++

1. Save file as `solution.cpp`
2. Compile:

   ```
   g++ solution.cpp -o solution
   ```

3. Run:

   ```
   ./solution
   ```

### Java

1. Save file as `Solution.java`
2. Compile:

   ```
   javac Solution.java
   ```

3. Run:

   ```
   java Solution
   ```

### JavaScript

1. Save file as `solution.js`
2. Run:

   ```
   node solution.js
   ```

### Python3

1. Save file as `solution.py`
2. Run:

   ```
   python3 solution.py
   ```

### Go

1. Save file as `solution.go`
2. Run:

   ```
   go run solution.go
   ```

## Notes & Optimizations

* Always check string length first to avoid unnecessary work
* The concatenation trick is a common pattern in string matching problems
* This solution is optimal for this problem
* Alternative approach is brute-force rotation, but it is less clean
* Works well within constraints since maximum length is small

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
