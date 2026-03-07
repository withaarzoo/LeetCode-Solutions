# Problem Title

1. Minimum Number of Flips to Make the Binary String Alternating

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

You are given a binary string `s` consisting only of characters `0` and `1`.

You can perform two types of operations:

Type 1
Remove the character at the start of the string and append it to the end.

Type 2
Pick any character in the string and flip it. If it is `0` it becomes `1`, and if it is `1` it becomes `0`.

Your task is to return the minimum number of **Type 2 (flip)** operations needed so that the string becomes **alternating**.

A binary string is alternating if no two adjacent characters are equal.

Example alternating strings:

010101
101010

---

## Constraints

1 <= s.length <= 100000

s[i] is either '0' or '1'

---

## Intuition

When I first read the problem, I noticed that I can rotate the string using the first operation. This means I am allowed to change the starting position of the string without any cost.

So instead of thinking about only the original string, I realized that **every rotation of the string is also valid**.

If the length of the string is `n`, there are `n` possible rotations.

At the same time, an alternating binary string can only have two possible patterns:

Pattern 1

010101...

Pattern 2

101010...

So my goal became simple:

For every possible rotation of the string, check how many flips are required to convert it into one of these alternating patterns.

Then choose the minimum value.

However generating all rotations directly would take too much time.

To solve this efficiently, I used a trick.

If I concatenate the string with itself

s + s

then every rotation of the original string appears as a substring of length `n` inside this new string.

Using a sliding window of size `n`, I can check all rotations in linear time.

---

## Approach

Step 1

Let `n` be the length of the string.

Step 2

Create a new string

s2 = s + s

This allows us to simulate every possible rotation.

Step 3

Maintain two mismatch counters

`diff1` for pattern starting with 0

`diff2` for pattern starting with 1

Step 4

Use a sliding window of length `n` over the string `s2`.

Step 5

For each index, compare the character with the expected alternating pattern.

Step 6

Increase mismatch counters when characters do not match the expected pattern.

Step 7

When the window size becomes larger than `n`, remove the effect of the character leaving the window.

Step 8

Each time the window size equals `n`, compute

min(diff1, diff2)

and update the answer.

Step 9

Return the minimum value.

---

## Data Structures Used

String

Used to create `s + s` so we can simulate all rotations.

Integer counters

Used to track mismatches with alternating patterns.

Sliding window

Used to efficiently evaluate every rotation.

---

## Operations & Behavior Summary

Rotation operation

Instead of performing actual rotations, we simulate them using `s + s`.

Flip operation

We count how many flips are required to match alternating patterns.

Sliding window

Allows us to evaluate every rotation in linear time.

---

## Complexity

Time Complexity

O(n)

We iterate through the doubled string once.

Space Complexity

O(1)

Only a few integer variables are used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minFlips(string s) {
        int n = s.size();
        string ss = s + s;

        int diff1 = 0;
        int diff2 = 0;
        int ans = INT_MAX;

        for (int i = 0; i < ss.size(); i++) {

            char expected1 = (i % 2) ? '1' : '0';
            char expected2 = (i % 2) ? '0' : '1';

            if (ss[i] != expected1) diff1++;
            if (ss[i] != expected2) diff2++;

            if (i >= n) {
                char prev = ss[i - n];

                char prevExp1 = ((i - n) % 2) ? '1' : '0';
                char prevExp2 = ((i - n) % 2) ? '0' : '1';

                if (prev != prevExp1) diff1--;
                if (prev != prevExp2) diff2--;
            }

            if (i >= n - 1) {
                ans = min(ans, min(diff1, diff2));
            }
        }

        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public int minFlips(String s) {
        int n = s.length();
        String ss = s + s;

        int diff1 = 0;
        int diff2 = 0;
        int ans = Integer.MAX_VALUE;

        for (int i = 0; i < ss.length(); i++) {

            char c = ss.charAt(i);

            char expected1 = (i % 2 == 0) ? '0' : '1';
            char expected2 = (i % 2 == 0) ? '1' : '0';

            if (c != expected1) diff1++;
            if (c != expected2) diff2++;

            if (i >= n) {
                char prev = ss.charAt(i - n);

                char prevExp1 = ((i - n) % 2 == 0) ? '0' : '1';
                char prevExp2 = ((i - n) % 2 == 0) ? '1' : '0';

                if (prev != prevExp1) diff1--;
                if (prev != prevExp2) diff2--;
            }

            if (i >= n - 1) {
                ans = Math.min(ans, Math.min(diff1, diff2));
            }
        }

        return ans;
    }
}
```

---

### JavaScript

```javascript
var minFlips = function(s) {
    const n = s.length;
    const ss = s + s;

    let diff1 = 0;
    let diff2 = 0;
    let ans = Infinity;

    for (let i = 0; i < ss.length; i++) {

        let expected1 = (i % 2 === 0) ? '0' : '1';
        let expected2 = (i % 2 === 0) ? '1' : '0';

        if (ss[i] !== expected1) diff1++;
        if (ss[i] !== expected2) diff2++;

        if (i >= n) {
            let prev = ss[i - n];

            let prevExp1 = ((i - n) % 2 === 0) ? '0' : '1';
            let prevExp2 = ((i - n) % 2 === 0) ? '1' : '0';

            if (prev !== prevExp1) diff1--;
            if (prev !== prevExp2) diff2--;
        }

        if (i >= n - 1) {
            ans = Math.min(ans, diff1, diff2);
        }
    }

    return ans;
};
```

---

### Python3

```python
class Solution:
    def minFlips(self, s: str) -> int:

        n = len(s)
        ss = s + s

        diff1 = 0
        diff2 = 0
        ans = float('inf')

        for i in range(len(ss)):

            expected1 = '0' if i % 2 == 0 else '1'
            expected2 = '1' if i % 2 == 0 else '0'

            if ss[i] != expected1:
                diff1 += 1

            if ss[i] != expected2:
                diff2 += 1

            if i >= n:
                prev = ss[i - n]

                prevExp1 = '0' if (i - n) % 2 == 0 else '1'
                prevExp2 = '1' if (i - n) % 2 == 0 else '0'

                if prev != prevExp1:
                    diff1 -= 1

                if prev != prevExp2:
                    diff2 -= 1

            if i >= n - 1:
                ans = min(ans, diff1, diff2)

        return ans
```

---

### Go

```go
func minFlips(s string) int {

    n := len(s)
    ss := s + s

    diff1 := 0
    diff2 := 0
    ans := 1<<31 - 1

    for i := 0; i < len(ss); i++ {

        expected1 := byte('0')
        expected2 := byte('1')

        if i%2 == 1 {
            expected1 = '1'
            expected2 = '0'
        }

        if ss[i] != expected1 {
            diff1++
        }

        if ss[i] != expected2 {
            diff2++
        }

        if i >= n {
            prev := ss[i-n]

            prevExp1 := byte('0')
            prevExp2 := byte('1')

            if (i-n)%2 == 1 {
                prevExp1 = '1'
                prevExp2 = '0'
            }

            if prev != prevExp1 {
                diff1--
            }

            if prev != prevExp2 {
                diff2--
            }
        }

        if i >= n-1 {
            if diff1 < ans {
                ans = diff1
            }

            if diff2 < ans {
                ans = diff2
            }
        }
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Duplicate the string

We create `s + s` so that every possible rotation appears as a substring.

1. Maintain two mismatch counters

One for pattern `010101` and another for `101010`.

1. Iterate through the doubled string

At each index we check whether the current character matches the expected alternating pattern.

1. Update mismatch counts

If the character does not match the expected pattern, we increase the mismatch counter.

1. Maintain sliding window

If the window size becomes greater than `n`, we remove the contribution of the character leaving the window.

1. Compute answer

Whenever the window size becomes `n`, we calculate

min(diff1, diff2)

1. Track the minimum result

The smallest value across all windows is the final answer.

---

## Examples

Example 1

Input

s = "111000"

Output

2

Explanation

After rotating and flipping minimal characters, the string becomes alternating.

Example 2

Input

s = "010"

Output

0

Explanation

The string is already alternating.

Example 3

Input

s = "1110"

Output

1

---

## How to use / Run locally

Clone the repository

```bash
git clone <repository-url>
```

Navigate to the project folder

```bash
cd project-folder
```

Compile and run (example for C++)

```bash
g++ solution.cpp
./a.out
```

For Python

```bash
python solution.py
```

---

## Notes & Optimizations

Key optimization

Instead of generating every rotation separately, we use `s + s` and a sliding window.

This reduces the complexity from

O(n^2)

to

O(n)

This approach works efficiently even for the maximum constraint of 100000 characters.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
