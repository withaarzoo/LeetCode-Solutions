# 2573. Find the String with LCP

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given an `n x n` matrix `lcp`.

`lcp[i][j]` represents the length of the longest common prefix between:

* suffix starting from index `i`
* suffix starting from index `j`

We need to return the lexicographically smallest lowercase English string that exactly matches the given LCP matrix.

If no valid string exists, return an empty string.

---

## Constraints

* `1 <= n == lcp.length == lcp[i].length <= 1000`
* `0 <= lcp[i][j] <= n`
* Only lowercase English letters can be used.

---

## Intuition

I noticed that if `lcp[i][j] > 0`, then the first character of both suffixes must be equal.

That means:

```text
word[i] == word[j]
```

So I can group all indices that must contain the same character.

Then I assign the smallest possible characters from `'a'` to `'z'` to keep the final string lexicographically smallest.

After building the string, I verify whether it really produces the same LCP matrix.

---

## Approach

1. Create a `group` array where `group[i]` stores which character group index `i` belongs to.
2. Traverse from left to right.
3. If an index is not assigned yet:

   * assign a new character group
   * mark all positions `j` where `lcp[i][j] > 0` with the same group
4. If more than `26` groups are needed, return `""`.
5. Build the string using the assigned groups.
6. Recompute the LCP matrix using dynamic programming.
7. Compare the generated LCP matrix with the given matrix.
8. If every value matches, return the string.
9. Otherwise, return `""`.

---

## Data Structures Used

* `group[]`

  * Stores which character group each index belongs to.
* `string / char array`

  * Stores the final generated string.
* `dp[][]`

  * Stores recomputed LCP values for validation.

---

## Operations & Behavior Summary

| Operation           | Purpose                                             |
| ------------------- | --------------------------------------------------- |
| Assign groups       | Decide which positions must have the same character |
| Build answer string | Convert groups into letters                         |
| Recompute LCP       | Validate generated string                           |
| Compare matrices    | Ensure exact correctness                            |

---

## Complexity

* **Time Complexity:** `O(n^2)`

  * We use nested loops for grouping and validation.
* **Space Complexity:** `O(n^2)`

  * We use an extra DP matrix of size `n x n`.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string findTheString(vector<vector<int>>& lcp) {
        int n = lcp.size();

        vector<int> group(n, -1);
        int curGroup = 0;

        for (int i = 0; i < n; i++) {
            if (group[i] == -1) {
                if (curGroup == 26) return "";

                group[i] = curGroup++;

                for (int j = i + 1; j < n; j++) {
                    if (lcp[i][j] > 0) {
                        group[j] = group[i];
                    }
                }
            }
        }

        string ans(n, 'a');
        for (int i = 0; i < n; i++) {
            ans[i] = 'a' + group[i];
        }

        vector<vector<int>> dp(n + 1, vector<int>(n + 1, 0));

        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                if (ans[i] == ans[j]) {
                    dp[i][j] = 1 + dp[i + 1][j + 1];
                }
            }
        }

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (dp[i][j] != lcp[i][j]) {
                    return "";
                }
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public String findTheString(int[][] lcp) {
        int n = lcp.length;

        int[] group = new int[n];
        Arrays.fill(group, -1);

        int curGroup = 0;

        for (int i = 0; i < n; i++) {
            if (group[i] == -1) {
                if (curGroup == 26) return "";

                group[i] = curGroup++;

                for (int j = i + 1; j < n; j++) {
                    if (lcp[i][j] > 0) {
                        group[j] = group[i];
                    }
                }
            }
        }

        char[] ans = new char[n];
        for (int i = 0; i < n; i++) {
            ans[i] = (char) ('a' + group[i]);
        }

        int[][] dp = new int[n + 1][n + 1];

        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                if (ans[i] == ans[j]) {
                    dp[i][j] = 1 + dp[i + 1][j + 1];
                }
            }
        }

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (dp[i][j] != lcp[i][j]) {
                    return "";
                }
            }
        }

        return new String(ans);
    }
}
```

### JavaScript

```javascript
var findTheString = function(lcp) {
    const n = lcp.length;

    const group = new Array(n).fill(-1);
    let curGroup = 0;

    for (let i = 0; i < n; i++) {
        if (group[i] === -1) {
            if (curGroup === 26) return "";

            group[i] = curGroup++;

            for (let j = i + 1; j < n; j++) {
                if (lcp[i][j] > 0) {
                    group[j] = group[i];
                }
            }
        }
    }

    const chars = new Array(n);
    for (let i = 0; i < n; i++) {
        chars[i] = String.fromCharCode(97 + group[i]);
    }

    const ans = chars.join("");

    const dp = Array.from({ length: n + 1 }, () => new Array(n + 1).fill(0));

    for (let i = n - 1; i >= 0; i--) {
        for (let j = n - 1; j >= 0; j--) {
            if (ans[i] === ans[j]) {
                dp[i][j] = 1 + dp[i + 1][j + 1];
            }
        }
    }

    for (let i = 0; i < n; i++) {
        for (let j = 0; j < n; j++) {
            if (dp[i][j] !== lcp[i][j]) {
                return "";
            }
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def findTheString(self, lcp: List[List[int]]) -> str:
        n = len(lcp)

        group = [-1] * n
        cur_group = 0

        for i in range(n):
            if group[i] == -1:
                if cur_group == 26:
                    return ""

                group[i] = cur_group
                cur_group += 1

                for j in range(i + 1, n):
                    if lcp[i][j] > 0:
                        group[j] = group[i]

        ans = [''] * n
        for i in range(n):
            ans[i] = chr(ord('a') + group[i])

        ans = ''.join(ans)

        dp = [[0] * (n + 1) for _ in range(n + 1)]

        for i in range(n - 1, -1, -1):
            for j in range(n - 1, -1, -1):
                if ans[i] == ans[j]:
                    dp[i][j] = 1 + dp[i + 1][j + 1]

        for i in range(n):
            for j in range(n):
                if dp[i][j] != lcp[i][j]:
                    return ""

        return ans
```

### Go

```go
func findTheString(lcp [][]int) string {
    n := len(lcp)

    group := make([]int, n)
    for i := 0; i < n; i++ {
        group[i] = -1
    }

    curGroup := 0

    for i := 0; i < n; i++ {
        if group[i] == -1 {
            if curGroup == 26 {
                return ""
            }

            group[i] = curGroup
            curGroup++

            for j := i + 1; j < n; j++ {
                if lcp[i][j] > 0 {
                    group[j] = group[i]
                }
            }
        }
    }

    ans := make([]byte, n)
    for i := 0; i < n; i++ {
        ans[i] = byte('a' + group[i])
    }

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if ans[i] == ans[j] {
                dp[i][j] = 1 + dp[i+1][j+1]
            }
        }
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if dp[i][j] != lcp[i][j] {
                return ""
            }
        }
    }

    return string(ans)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Create a group array

```text
group[i] = which character group position i belongs to
```

Initially, every index is unassigned.

```text
group = [-1, -1, -1, -1]
```

---

### Step 2: Assign smallest possible characters

When I find an unassigned position, I create a new group.

```text
Index 0 -> group 0 -> 'a'
Index 1 -> group 1 -> 'b'
```

This guarantees the lexicographically smallest answer.

---

### Step 3: Merge positions with same starting character

If:

```text
lcp[i][j] > 0
```

Then:

```text
word[i] == word[j]
```

So both positions must belong to the same group.

---

### Step 4: Build the answer string

If:

```text
group = [0, 1, 0, 1]
```

Then:

```text
answer = "abab"
```

---

### Step 5: Recompute LCP matrix

I use DP from back to front.

If characters match:

```text
dp[i][j] = 1 + dp[i + 1][j + 1]
```

Else:

```text
dp[i][j] = 0
```

---

### Step 6: Compare matrices

If recomputed LCP matrix is exactly equal to given matrix, the answer is valid.

Otherwise:

```text
return ""
```

---

## Examples

### Example 1

```text
Input:
lcp = [[4,0,2,0],[0,3,0,1],[2,0,2,0],[0,1,0,1]]

Output:
"abab"
```

### Example 2

```text
Input:
lcp = [[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,1]]

Output:
"aaaa"
```

### Example 3

```text
Input:
lcp = [[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,3]]

Output:
""
```

---

## How to use / Run locally

### C++

```bash
g++ filename.cpp -o output
./output
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node filename.js
```

### Python3

```bash
python filename.py
```

### Go

```bash
go run filename.go
```

---

## Notes & Optimizations

* Grouping ensures equal characters are assigned correctly.
* Using smallest possible character groups guarantees lexicographically smallest answer.
* Final DP validation is necessary because grouping alone is not enough.
* Maximum distinct groups allowed is `26` because only lowercase English letters are available.
* The solution works efficiently for `n <= 1000`.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
