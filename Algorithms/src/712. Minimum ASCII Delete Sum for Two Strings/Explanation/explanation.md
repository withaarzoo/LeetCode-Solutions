# Minimum ASCII Delete Sum for Two Strings

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

Given two strings `s1` and `s2`, the task is to make both strings equal by deleting characters.
Each deletion costs the **ASCII value** of the deleted character.

My goal is to return the **minimum total ASCII delete sum** required to make both strings equal.

This is a classic **dynamic programming problem on two strings**.

---

## Constraints

* 1 ≤ length of `s1`, `s2` ≤ 1000
* Strings contain only lowercase English letters

---

## Intuition

When I first read the problem, I realized something important.

Instead of thinking about **what to delete**,
I thought about **what I should keep**.

If both strings end up equal, then the remaining characters must form a **common subsequence**.

So the idea is simple:

* Keep matching characters
* Delete everything else with minimum ASCII cost

This immediately points to **Dynamic Programming**, very similar to LCS but with costs.

---

## Approach

1. I define a DP state where I compare substrings of both strings.
2. If characters match, I move forward without deleting anything.
3. If characters do not match, I try deleting:

   * from the first string
   * or from the second string
4. I take the minimum ASCII cost from both options.
5. I optimize space using **1D DP** instead of a full 2D table.

---

## Data Structures Used

* One-dimensional integer array for Dynamic Programming
* No extra data structures are required

---

## Operations & Behavior Summary

* Compare characters at current indices
* Move both pointers if characters match
* Delete one character if they do not match
* Accumulate ASCII values during deletion
* Final answer stored in DP starting position

---

## Complexity

**Time Complexity:**
O(n × m)
Where `n` is the length of `s1` and `m` is the length of `s2`.

**Space Complexity:**
O(m)
Only one DP row is stored at a time.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumDeleteSum(string s1, string s2) {
        int n = s1.size(), m = s2.size();
        vector<int> dp(m + 1, 0);

        for (int j = m - 1; j >= 0; j--) {
            dp[j] = dp[j + 1] + s2[j];
        }

        for (int i = n - 1; i >= 0; i--) {
            int prev = dp[m];
            dp[m] += s1[i];

            for (int j = m - 1; j >= 0; j--) {
                int temp = dp[j];
                if (s1[i] == s2[j]) {
                    dp[j] = prev;
                } else {
                    dp[j] = min(
                        s1[i] + dp[j],
                        s2[j] + dp[j + 1]
                    );
                }
                prev = temp;
            }
        }
        return dp[0];
    }
};
```

---

### Java

```java
class Solution {
    public int minimumDeleteSum(String s1, String s2) {
        int n = s1.length(), m = s2.length();
        int[] dp = new int[m + 1];

        for (int j = m - 1; j >= 0; j--) {
            dp[j] = dp[j + 1] + s2.charAt(j);
        }

        for (int i = n - 1; i >= 0; i--) {
            int prev = dp[m];
            dp[m] += s1.charAt(i);

            for (int j = m - 1; j >= 0; j--) {
                int temp = dp[j];
                if (s1.charAt(i) == s2.charAt(j)) {
                    dp[j] = prev;
                } else {
                    dp[j] = Math.min(
                        s1.charAt(i) + dp[j],
                        s2.charAt(j) + dp[j + 1]
                    );
                }
                prev = temp;
            }
        }
        return dp[0];
    }
}
```

---

### JavaScript

```javascript
var minimumDeleteSum = function(s1, s2) {
    const n = s1.length, m = s2.length;
    let dp = new Array(m + 1).fill(0);

    for (let j = m - 1; j >= 0; j--) {
        dp[j] = dp[j + 1] + s2.charCodeAt(j);
    }

    for (let i = n - 1; i >= 0; i--) {
        let prev = dp[m];
        dp[m] += s1.charCodeAt(i);

        for (let j = m - 1; j >= 0; j--) {
            let temp = dp[j];
            if (s1[i] === s2[j]) {
                dp[j] = prev;
            } else {
                dp[j] = Math.min(
                    s1.charCodeAt(i) + dp[j],
                    s2.charCodeAt(j) + dp[j + 1]
                );
            }
            prev = temp;
        }
    }
    return dp[0];
};
```

---

### Python3

```python
class Solution:
    def minimumDeleteSum(self, s1: str, s2: str) -> int:
        n, m = len(s1), len(s2)
        dp = [0] * (m + 1)

        for j in range(m - 1, -1, -1):
            dp[j] = dp[j + 1] + ord(s2[j])

        for i in range(n - 1, -1, -1):
            prev = dp[m]
            dp[m] += ord(s1[i])

            for j in range(m - 1, -1, -1):
                temp = dp[j]
                if s1[i] == s2[j]:
                    dp[j] = prev
                else:
                    dp[j] = min(
                        ord(s1[i]) + dp[j],
                        ord(s2[j]) + dp[j + 1]
                    )
                prev = temp

        return dp[0]
```

---

### Go

```go
func minimumDeleteSum(s1 string, s2 string) int {
    n, m := len(s1), len(s2)
    dp := make([]int, m+1)

    for j := m - 1; j >= 0; j-- {
        dp[j] = dp[j+1] + int(s2[j])
    }

    for i := n - 1; i >= 0; i-- {
        prev := dp[m]
        dp[m] += int(s1[i])

        for j := m - 1; j >= 0; j-- {
            temp := dp[j]
            if s1[i] == s2[j] {
                dp[j] = prev
            } else {
                a := int(s1[i]) + dp[j]
                b := int(s2[j]) + dp[j+1]
                if a < b {
                    dp[j] = a
                } else {
                    dp[j] = b
                }
            }
            prev = temp
        }
    }
    return dp[0]
}
```

---

## Step-by-step Detailed Explanation

1. Start DP from the end of both strings.
2. If one string ends, delete all remaining characters from the other.
3. If characters match, move both indices.
4. If characters do not match, try deleting from either string.
5. Choose the option with smaller ASCII cost.
6. Use one DP array to reduce memory usage.
7. The final answer is stored at the start of DP.

---

## Examples

Input
`s1 = "sea"`
`s2 = "eat"`

Output
`231`

Explanation
Delete `s` (115) from `s1` and `t` (116) from `s2`.

---

## How to use / Run locally

1. Clone the repository
2. Choose your preferred language file
3. Compile or run using standard compiler/interpreter
4. Call `minimumDeleteSum(s1, s2)` with test inputs

---

## Notes & Optimizations

* This solution avoids full 2D DP to save memory
* Works efficiently for maximum constraints
* Can be extended to show DP table for learning purposes

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
