# 1461. Check If a String Contains All Binary Codes of Size K

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

Given a binary string `s` and an integer `k`, return `true` if every possible binary code of length `k` exists as a substring of `s`. Otherwise, return `false`.

In simple words, I need to check whether all possible binary combinations of size `k` appear somewhere inside the string.

For example, if `k = 2`, possible binary codes are:

00, 01, 10, 11

That means total 4 combinations.

---

## Constraints

* 1 <= s.length <= 5 * 10^5
* s[i] is either '0' or '1'
* 1 <= k <= 20

---

## Intuition

When I first saw this problem, I thought like this:

If the binary code length is `k`, then total possible combinations are `2^k`.

So the real question becomes:

Can I find all `2^k` different substrings of length `k` inside `s`?

If the string does not even have enough substrings, then it is impossible.

Total substrings of length k = n - k + 1

If (n - k + 1) < 2^k → answer must be false.

Instead of storing substrings as strings, I can convert each substring into a number using bit manipulation. That makes the solution faster and more memory efficient.

---

## Approach

1. Let n = length of string.
2. If n < k → return false.
3. Compute total = 2^k using (1 << k).
4. If (n - k + 1) < total → return false.
5. Use a sliding window of size k.
6. Convert first k characters into a number.
7. Store seen patterns in a boolean array of size total.
8. Slide window one step at a time:

   * Left shift current value
   * Remove extra left bit using mask
   * Add new bit
9. Count how many unique patterns I find.
10. If count == total → return true.

---

## Data Structures Used

* Boolean array of size 2^k (to track visited binary codes)
* Integer mask to keep last k bits
* Integer variable to maintain sliding window value

---

## Operations & Behavior Summary

* Bit shifting to simulate sliding window
* Masking to remove extra bits
* Mark visited patterns
* Early return if all combinations found

---

## Complexity

**Time Complexity:** O(n)

I scan the string once using sliding window.

**Space Complexity:** O(2^k)

I use a boolean array of size 2^k.

Since k <= 20, maximum size is manageable.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool hasAllCodes(string s, int k) {
        int n = s.size();
        if (n < k) return false;

        int total = 1 << k;
        if (n - k + 1 < total) return false;

        vector<bool> seen(total, false);
        int mask = total - 1;
        int curr = 0;
        int count = 0;

        for (int i = 0; i < k; i++) {
            curr = (curr << 1) | (s[i] - '0');
        }

        if (!seen[curr]) {
            seen[curr] = true;
            count++;
        }

        for (int i = k; i < n; i++) {
            curr = ((curr << 1) & mask) | (s[i] - '0');
            if (!seen[curr]) {
                seen[curr] = true;
                count++;
                if (count == total) return true;
            }
        }

        return count == total;
    }
};
```

### Java

```java
class Solution {
    public boolean hasAllCodes(String s, int k) {
        int n = s.length();
        if (n < k) return false;

        int total = 1 << k;
        if (n - k + 1 < total) return false;

        boolean[] seen = new boolean[total];
        int mask = total - 1;
        int curr = 0;
        int count = 0;

        for (int i = 0; i < k; i++) {
            curr = (curr << 1) | (s.charAt(i) - '0');
        }

        if (!seen[curr]) {
            seen[curr] = true;
            count++;
        }

        for (int i = k; i < n; i++) {
            curr = ((curr << 1) & mask) | (s.charAt(i) - '0');
            if (!seen[curr]) {
                seen[curr] = true;
                count++;
                if (count == total) return true;
            }
        }

        return count == total;
    }
}
```

### JavaScript

```javascript
var hasAllCodes = function(s, k) {
    const n = s.length;
    if (n < k) return false;

    const total = 1 << k;
    if (n - k + 1 < total) return false;

    const seen = new Array(total).fill(false);
    const mask = total - 1;

    let curr = 0;
    let count = 0;

    for (let i = 0; i < k; i++) {
        curr = (curr << 1) | (s[i] - '0');
    }

    if (!seen[curr]) {
        seen[curr] = true;
        count++;
    }

    for (let i = k; i < n; i++) {
        curr = ((curr << 1) & mask) | (s[i] - '0');
        if (!seen[curr]) {
            seen[curr] = true;
            count++;
            if (count === total) return true;
        }
    }

    return count === total;
};
```

### Python3

```python
class Solution:
    def hasAllCodes(self, s: str, k: int) -> bool:
        n = len(s)
        if n < k:
            return False

        total = 1 << k
        if n - k + 1 < total:
            return False

        seen = [False] * total
        mask = total - 1
        curr = 0
        count = 0

        for i in range(k):
            curr = (curr << 1) | int(s[i])

        if not seen[curr]:
            seen[curr] = True
            count += 1

        for i in range(k, n):
            curr = ((curr << 1) & mask) | int(s[i])
            if not seen[curr]:
                seen[curr] = True
                count += 1
                if count == total:
                    return True

        return count == total
```

### Go

```go
func hasAllCodes(s string, k int) bool {
    n := len(s)
    if n < k {
        return false
    }

    total := 1 << k
    if n-k+1 < total {
        return false
    }

    seen := make([]bool, total)
    mask := total - 1

    curr := 0
    count := 0

    for i := 0; i < k; i++ {
        curr = (curr << 1) | int(s[i]-'0')
    }

    if !seen[curr] {
        seen[curr] = true
        count++
    }

    for i := k; i < n; i++ {
        curr = ((curr << 1) & mask) | int(s[i]-'0')
        if !seen[curr] {
            seen[curr] = true
            count++
            if count == total {
                return true
            }
        }
    }

    return count == total
}
```

---

## Step-by-step Detailed Explanation

1. I calculate total possible binary combinations using 1 << k.
2. I check if string even has enough substrings.
3. I create a boolean array to track visited codes.
4. I convert first k characters into a number.
5. I slide the window:

   * Shift left
   * Remove extra bit using mask
   * Add new bit
6. If new pattern not seen before, mark it.
7. If all patterns found, return true early.

---

## Examples

Example 1:

Input: s = "00110110", k = 2
Output: true

Example 2:

Input: s = "0110", k = 1
Output: true

Example 3:

Input: s = "0110", k = 2
Output: false

---

## How to use / Run locally

C++

* Compile using g++
* Run executable

Java

* Compile using javac
* Run using java

Python

* Run using python3 file.py

JavaScript

* Run using node file.js

Go

* Run using go run file.go

---

## Notes & Optimizations

* Early return improves performance.
* Using bit manipulation avoids heavy substring creation.
* Space usage is controlled because k <= 20.
* Sliding window ensures O(n) time.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
