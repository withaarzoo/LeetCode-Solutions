# 3474. Lexicographically Smallest Generated String

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

## Problem Summary

We are given two strings:

* `str1` of length `n`
* `str2` of length `m`

We need to generate a string `word` of length `n + m - 1`.

For every index `i`:

* If `str1[i] == 'T'`, then substring `word[i...i+m-1]` must be exactly equal to `str2`
* If `str1[i] == 'F'`, then substring `word[i...i+m-1]` must not be equal to `str2`

We need to return the lexicographically smallest valid generated string.

If it is impossible to generate such a string, return an empty string.

## Constraints

```text
1 <= n == str1.length <= 10^4
1 <= m == str2.length <= 500
str1 consists only of 'T' and 'F'
str2 consists only of lowercase English letters
```

## Intuition

I thought about solving the problem in two phases.

First, every position where `str1[i] == 'T'` forces a complete copy of `str2` into the answer.

That means those characters are fixed.

Then, after placing all required substrings, I fill the remaining empty positions with `'a'` because I want the lexicographically smallest answer.

After that, I check every `'F'` position.

If a substring accidentally becomes equal to `str2`, I try to change one character in that substring.

To keep the answer as small as possible, I try changing from right to left.

## Approach

1. Let:

   * `n = str1.length`
   * `m = str2.length`
   * Final answer length = `n + m - 1`

2. Create an answer array filled with `'?'`

3. For every index `i` where `str1[i] == 'T'`:

   * Force `ans[i + j] = str2[j]`
   * If there is a conflict, return `""`
   * Mark those positions as fixed

4. Replace all remaining `'?'` with `'a'`

5. For every index `i` where `str1[i] == 'F'`:

   * Check whether substring `ans[i...i+m-1]` equals `str2`
   * If not equal, continue
   * Otherwise, try to modify one unfixed character from right to left
   * If no modification is possible, return `""`

6. Return the final answer string

## Data Structures Used

* Character array / string builder for constructing the answer
* Boolean array `fixed[]` to mark positions fixed by `'T'`

## Operations & Behavior Summary

| Operation                 | Purpose                                   |
| ------------------------- | ----------------------------------------- |
| Fill answer with `?`      | Represents unknown positions              |
| Apply all `T` constraints | Forces required substrings                |
| Fill remaining with `a`   | Makes answer lexicographically smallest   |
| Check all `F` constraints | Ensures forbidden substrings do not match |
| Modify from right to left | Keeps string as small as possible         |

## Complexity

* Time Complexity: `O(n * m)`

  * `n` is length of `str1`
  * `m` is length of `str2`

* Space Complexity: `O(n + m)`

  * Extra space is used for answer array and fixed array

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string generateString(string str1, string str2) {
        int n = str1.size();
        int m = str2.size();
        int len = n + m - 1;

        string ans(len, '?');
        vector<bool> fixed(len, false);

        for (int i = 0; i < n; i++) {
            if (str1[i] == 'T') {
                for (int j = 0; j < m; j++) {
                    int pos = i + j;

                    if (ans[pos] != '?' && ans[pos] != str2[j]) {
                        return "";
                    }

                    ans[pos] = str2[j];
                    fixed[pos] = true;
                }
            }
        }

        for (int i = 0; i < len; i++) {
            if (ans[i] == '?') ans[i] = 'a';
        }

        for (int i = 0; i < n; i++) {
            if (str1[i] == 'F') {
                bool same = true;

                for (int j = 0; j < m; j++) {
                    if (ans[i + j] != str2[j]) {
                        same = false;
                        break;
                    }
                }

                if (!same) continue;

                bool changed = false;

                for (int j = m - 1; j >= 0; j--) {
                    int pos = i + j;

                    if (fixed[pos]) continue;

                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c != ans[pos] && c != str2[j]) {
                            ans[pos] = c;
                            changed = true;
                            break;
                        }
                    }

                    if (changed) break;
                }

                if (!changed) return "";
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public String generateString(String str1, String str2) {
        int n = str1.length();
        int m = str2.length();
        int len = n + m - 1;

        char[] ans = new char[len];
        boolean[] fixed = new boolean[len];

        for (int i = 0; i < len; i++) {
            ans[i] = '?';
        }

        for (int i = 0; i < n; i++) {
            if (str1.charAt(i) == 'T') {
                for (int j = 0; j < m; j++) {
                    int pos = i + j;
                    char ch = str2.charAt(j);

                    if (ans[pos] != '?' && ans[pos] != ch) {
                        return "";
                    }

                    ans[pos] = ch;
                    fixed[pos] = true;
                }
            }
        }

        for (int i = 0; i < len; i++) {
            if (ans[i] == '?') ans[i] = 'a';
        }

        for (int i = 0; i < n; i++) {
            if (str1.charAt(i) == 'F') {
                boolean same = true;

                for (int j = 0; j < m; j++) {
                    if (ans[i + j] != str2.charAt(j)) {
                        same = false;
                        break;
                    }
                }

                if (!same) continue;

                boolean changed = false;

                for (int j = m - 1; j >= 0; j--) {
                    int pos = i + j;

                    if (fixed[pos]) continue;

                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c != ans[pos] && c != str2.charAt(j)) {
                            ans[pos] = c;
                            changed = true;
                            break;
                        }
                    }

                    if (changed) break;
                }

                if (!changed) return "";
            }
        }

        return new String(ans);
    }
}
```

### JavaScript

```javascript
var generateString = function(str1, str2) {
    const n = str1.length;
    const m = str2.length;
    const len = n + m - 1;

    const ans = Array(len).fill('?');
    const fixed = Array(len).fill(false);

    for (let i = 0; i < n; i++) {
        if (str1[i] === 'T') {
            for (let j = 0; j < m; j++) {
                const pos = i + j;

                if (ans[pos] !== '?' && ans[pos] !== str2[j]) {
                    return "";
                }

                ans[pos] = str2[j];
                fixed[pos] = true;
            }
        }
    }

    for (let i = 0; i < len; i++) {
        if (ans[i] === '?') ans[i] = 'a';
    }

    for (let i = 0; i < n; i++) {
        if (str1[i] === 'F') {
            let same = true;

            for (let j = 0; j < m; j++) {
                if (ans[i + j] !== str2[j]) {
                    same = false;
                    break;
                }
            }

            if (!same) continue;

            let changed = false;

            for (let j = m - 1; j >= 0; j--) {
                const pos = i + j;

                if (fixed[pos]) continue;

                for (let c = 97; c <= 122; c++) {
                    const ch = String.fromCharCode(c);

                    if (ch !== ans[pos] && ch !== str2[j]) {
                        ans[pos] = ch;
                        changed = true;
                        break;
                    }
                }

                if (changed) break;
            }

            if (!changed) return "";
        }
    }

    return ans.join('');
};
```

### Python3

```python
class Solution:
    def generateString(self, str1: str, str2: str) -> str:
        n = len(str1)
        m = len(str2)
        length = n + m - 1

        ans = ['?'] * length
        fixed = [False] * length

        for i in range(n):
            if str1[i] == 'T':
                for j in range(m):
                    pos = i + j

                    if ans[pos] != '?' and ans[pos] != str2[j]:
                        return ""

                    ans[pos] = str2[j]
                    fixed[pos] = True

        for i in range(length):
            if ans[i] == '?':
                ans[i] = 'a'

        for i in range(n):
            if str1[i] == 'F':
                same = True

                for j in range(m):
                    if ans[i + j] != str2[j]:
                        same = False
                        break

                if not same:
                    continue

                changed = False

                for j in range(m - 1, -1, -1):
                    pos = i + j

                    if fixed[pos]:
                        continue

                    for c in range(ord('a'), ord('z') + 1):
                        ch = chr(c)

                        if ch != ans[pos] and ch != str2[j]:
                            ans[pos] = ch
                            changed = True
                            break

                    if changed:
                        break

                if not changed:
                    return ""

        return ''.join(ans)
```

### Go

```go
func generateString(str1 string, str2 string) string {
    n := len(str1)
    m := len(str2)
    length := n + m - 1

    ans := make([]byte, length)
    fixed := make([]bool, length)

    for i := 0; i < length; i++ {
        ans[i] = '?'
    }

    for i := 0; i < n; i++ {
        if str1[i] == 'T' {
            for j := 0; j < m; j++ {
                pos := i + j

                if ans[pos] != '?' && ans[pos] != str2[j] {
                    return ""
                }

                ans[pos] = str2[j]
                fixed[pos] = true
            }
        }
    }

    for i := 0; i < length; i++ {
        if ans[i] == '?' {
            ans[i] = 'a'
        }
    }

    for i := 0; i < n; i++ {
        if str1[i] == 'F' {
            same := true

            for j := 0; j < m; j++ {
                if ans[i+j] != str2[j] {
                    same = false
                    break
                }
            }

            if !same {
                continue
            }

            changed := false

            for j := m - 1; j >= 0; j-- {
                pos := i + j

                if fixed[pos] {
                    continue
                }

                for c := byte('a'); c <= byte('z'); c++ {
                    if c != ans[pos] && c != str2[j] {
                        ans[pos] = c
                        changed = true
                        break
                    }
                }

                if changed {
                    break
                }
            }

            if !changed {
                return ""
            }
        }
    }

    return string(ans)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Decide the final length

The generated string length will always be:

```text
n + m - 1
```

So I create an array of that size.

### Step 2: Apply all `T` constraints

Whenever I see `str1[i] == 'T'`, I force the substring starting from index `i` to become equal to `str2`.

Example:

```text
str1 = TFTF
str2 = ab
```

At index `0`, I place:

```text
ab???
```

At index `2`, I place:

```text
abab?
```

If two placements conflict, I return an empty string.

### Step 3: Fill empty positions with `a`

After applying all forced substrings, some positions may still remain empty.

To make the answer lexicographically smallest, I fill those positions with `'a'`.

### Step 4: Process all `F` constraints

For every index where `str1[i] == 'F'`, I check whether the substring equals `str2`.

If it does not match, then it is already valid.

### Step 5: Break invalid matches

If a forbidden substring exactly matches `str2`, then I need to break it.

I try to change one unfixed character from right to left.

Changing later positions helps keep the overall answer smaller lexicographically.

### Step 6: Return impossible if needed

If every character in that substring is fixed and cannot be changed, then no valid answer exists.

So I return:

```text
""
```

## Examples

### Example 1

```text
Input:
str1 = "TFTF"
str2 = "ab"

Output:
"ababa"
```

### Example 2

```text
Input:
str1 = "TFTF"
str2 = "abc"

Output:
""
```

### Example 3

```text
Input:
str1 = "F"
str2 = "d"

Output:
"a"
```

## How to use / Run locally

```bash
g++ solution.cpp -o solution
./solution
```

```bash
javac Solution.java
java Solution
```

```bash
node solution.js
```

```bash
python solution.py
```

```bash
go run solution.go
```

## Notes & Optimizations

* I use a greedy approach to keep the answer lexicographically smallest
* I always fill empty positions with `'a'`
* I only change characters when a forbidden substring becomes equal to `str2`
* I modify from right to left because changing later characters affects lexicographical order less
* Time complexity is efficient enough for the given constraints

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
