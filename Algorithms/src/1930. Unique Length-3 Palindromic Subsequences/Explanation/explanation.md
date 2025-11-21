# Unique Length-3 Palindromic Subsequences (LeetCode 1930)

![Problem Screenshot](/mnt/data/31a07173-a7ec-4f65-a16b-bd9348221068.png)

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given a string `s`, return the **number of unique palindromes of length 3** that are subsequences of `s`.

A length-3 palindrome must have the form `a b a` (same character at both ends, any character in the middle). The same palindrome string is counted once even if it can be formed multiple ways.

## Constraints

* `3 <= s.length <= 10^5`
* `s` consists of only lowercase English letters (`'a'` - `'z'`).

## Intuition

I thought: a length-3 palindrome must look like `a b a`. So if I fix the outer letter `a`, I only need to know which distinct letters `b` appear somewhere between two occurrences of `a`. If I take the earliest occurrence of `a` and the latest occurrence of `a`, any character between them can be a middle character for some `a b a` subsequence. So for each letter `a`, I count distinct letters in the substring between its first and last occurrence. Summing across all 26 letters gives the answer.

## Approach

1. Scan `s` once and record the **first** and **last** indices of every letter `'a'..'z'`.
2. For each letter `ch` from `'a'` to `'z'`:

   * If `first[ch] < last[ch]`, then the letter occurs at least twice.
   * Count distinct letters that appear strictly between `first[ch]` and `last[ch]`. Each distinct middle letter `m` yields palindrome `ch m ch`.
3. Sum those counts for all letters and return the sum.

This approach is simple and efficient because the alphabet size is constant (26).

## Data Structures Used

* Fixed-size arrays (length 26) to store:

  * `first` occurrences (int)
  * `last` occurrences (int)
  * `seen` flags when scanning a substring (boolean/bitset)
* The input string `s` itself.

## Operations & Behavior Summary

* One linear scan to populate `first` and `last`.
* For up to 26 letters, scan the substring between `first` and `last` to mark distinct middle letters.
* For each outer letter, count distinct middle letters and add to result.

## Complexity

* **Time Complexity:** `O(26 * n)` = `O(n)`, where `n` is `s.length`. (26 is constant so this is linear.)

  * One pass to compute `first` and `last` indices: `O(n)`.
  * For each of 26 letters, possibly scanning the substring between its first and last — worst-case 26 full scans = `26 * n` but constant factor only.
* **Space Complexity:** `O(1)` (only constant extra memory: arrays of size 26, a few ints/booleans).

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countPalindromicSubsequence(string s) {
        int n = s.size();
        const int A = 26;
        vector<int> first(A, INT_MAX), last(A, -1);
        // record first and last occurrence for every letter
        for (int i = 0; i < n; ++i) {
            int c = s[i] - 'a';
            first[c] = min(first[c], i);
            last[c] = max(last[c], i);
        }

        int ans = 0;
        // for each outer letter, count distinct middle letters between first and last
        for (int c = 0; c < A; ++c) {
            if (first[c] < last[c]) {
                vector<bool> seen(A, false);
                for (int i = first[c] + 1; i < last[c]; ++i) {
                    seen[s[i] - 'a'] = true;
                }
                for (int j = 0; j < A; ++j) if (seen[j]) ++ans;
            }
        }
        return ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public int countPalindromicSubsequence(String s) {
        int n = s.length();
        int A = 26;
        int[] first = new int[A];
        int[] last = new int[A];
        Arrays.fill(first, Integer.MAX_VALUE);
        Arrays.fill(last, -1);

        // record first and last occurrence for every letter
        for (int i = 0; i < n; ++i) {
            int c = s.charAt(i) - 'a';
            first[c] = Math.min(first[c], i);
            last[c] = Math.max(last[c], i);
        }

        int ans = 0;
        // for each outer letter, count distinct middle letters between first and last
        for (int c = 0; c < A; ++c) {
            if (first[c] < last[c]) {
                boolean[] seen = new boolean[A];
                for (int i = first[c] + 1; i < last[c]; ++i) {
                    seen[s.charAt(i) - 'a'] = true;
                }
                for (int j = 0; j < A; ++j) if (seen[j]) ans++;
            }
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {number}
 */
var countPalindromicSubsequence = function(s) {
    const n = s.length;
    const A = 26;
    const first = new Array(A).fill(Number.MAX_SAFE_INTEGER);
    const last = new Array(A).fill(-1);
    // record first and last occurrence
    for (let i = 0; i < n; ++i) {
        const c = s.charCodeAt(i) - 97;
        first[c] = Math.min(first[c], i);
        last[c] = Math.max(last[c], i);
    }

    let ans = 0;
    for (let c = 0; c < A; ++c) {
        if (first[c] < last[c]) {
            const seen = new Array(A).fill(false);
            for (let i = first[c] + 1; i < last[c]; ++i) {
                seen[s.charCodeAt(i) - 97] = true;
            }
            for (let j = 0; j < A; ++j) if (seen[j]) ans++;
        }
    }
    return ans;
};
```

---

### Python3

```python3
class Solution:
    def countPalindromicSubsequence(self, s: str) -> int:
        n = len(s)
        A = 26
        first = [10**9] * A
        last = [-1] * A
        # record first and last occurrence for each letter
        for i, ch in enumerate(s):
            idx = ord(ch) - ord('a')
            if i < first[idx]:
                first[idx] = i
            if i > last[idx]:
                last[idx] = i

        ans = 0
        # for each outer letter, count distinct middle letters between first and last
        for c in range(A):
            if first[c] < last[c]:
                seen = [False] * A
                for i in range(first[c] + 1, last[c]):
                    seen[ord(s[i]) - ord('a')] = True
                ans += sum(1 for x in seen if x)
        return ans
```

---

### Go

```go
package main

func countPalindromicSubsequence(s string) int {
    n := len(s)
    const A = 26
    first := make([]int, A)
    last := make([]int, A)
    for i := 0; i < A; i++ {
        first[i] = 1<<30
        last[i] = -1
    }
    // record first and last occurrence
    for i := 0; i < n; i++ {
        c := int(s[i] - 'a')
        if i < first[c] {
            first[c] = i
        }
        if i > last[c] {
            last[c] = i
        }
    }

    ans := 0
    // for each outer letter, count distinct middle letters between first and last
    for c := 0; c < A; c++ {
        if first[c] < last[c] {
            seen := make([]bool, A)
            for i := first[c] + 1; i < last[c]; i++ {
                seen[int(s[i]-'a')] = true
            }
            for j := 0; j < A; j++ {
                if seen[j] {
                    ans++
                }
            }
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the key parts of the solution line-by-line in a language-agnostic way, then map to the languages above.

### 1. Record first and last occurrence of each letter

* Why: If a letter appears only once, it cannot be the outer letter `a` in `a b a`. If it appears at least twice, the earliest and latest occurrences form a boundary that covers all possible middle characters that can be used with that outer letter.
* Implementation:

  * Initialize `first` array with a large value (INF) and `last` with `-1`.
  * Traverse the string index `i` from `0` to `n-1`.
  * For each character `ch` compute `idx = ch - 'a'`.
  * Update `first[idx] = min(first[idx], i)` and `last[idx] = max(last[idx], i)`.

**C++/Java/Python/JS/Go mapping:** This corresponds to the first `for` loop in each code block where we fill `first` and `last`.

### 2. For each possible outer letter, check feasibility

* Why: Only letters with `first < last` appear at least twice and can form `a _ a`.
* Implementation:

  * Loop over `c` from `0` to `25`.
  * If `first[c] >= last[c]`, skip — no valid pair.

**Mapping:** The `if (first[c] < last[c])` check in codes.

### 3. Count distinct middle letters strictly between `first` and `last`

* Why: For a fixed outer letter `a`, any distinct `b` that appears between the earliest and latest `a` yields a unique palindrome `a b a`. Counting distinct letters avoids duplicates.
* Implementation:

  * Initialize a `seen[26] = {false}`.
  * For `i` from `first[c] + 1` to `last[c] - 1`, mark `seen[indexOf(s[i])] = true`.
  * Count how many `seen[j]` are true and add to `ans`.

**Complexity note:** Counting distinct letters this way is efficient because the alphabet is size 26.

### 4. Return the total `ans`

* Sum over all 26 letters yields the number of unique length-3 palindromic subsequences.

---

## Examples

1. **Input:** `s = "aabca"`
   **Output:** `3`
   **Explanation:** palindromes: `"aba"`, `"aaa"`, `"aca"`.

2. **Input:** `s = "adc"`
   **Output:** `0`
   **Explanation:** no letter appears twice, so none possible.

3. **Input:** `s = "bbcbaba"`
   **Output:** `4`
   **Explanation:** palindromes: `"bbb"`, `"bcb"`, `"bab"`, `"aba"`.

---

## How to use / Run locally

### Python

1. Save the Python solution in a file `solution.py`.
2. Add a small runner:

```python
if __name__ == "__main__":
    s = "bbcbaba"
    print(Solution().countPalindromicSubsequence(s))
```

3. Run:

```bash
python3 solution.py
```

### JavaScript (Node.js)

1. Save code in `solution.js`.
2. Add:

```javascript
console.log(countPalindromicSubsequence("bbcbaba"));
```

3. Run:

```bash
node solution.js
```

### C++

1. Put the class in a file `solution.cpp`.
2. Add a `main()` to instantiate and call with a test string, and compile:

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

### Java

1. Save as `Solution.java` with a `main()` that calls the method.
2. Compile & run:

```bash
javac Solution.java
java Solution
```

### Go

1. Save code in `main.go`, add a `main()` to call the function.
2. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* The current solution is simple and runs in linear time relative to `n` (with a constant factor 26). It's perfectly adequate for `n` up to `10^5`.
* **Micro-optimization:** I could precompute prefix bitmasks of seen letters so that for any range `[l, r]` I can compute the distinct letters in `O(1)` using bit operations: `mask[r] ^ mask[l-1]` (or using `mask[r] & ~mask[l]` pattern carefully). That would convert the inner scans into constant-time checks per letter, giving a strict `O(n + 26)` runtime, but increases implementation complexity slightly.
* **Memory:** Current approach uses constant extra space (arrays of size 26). The bitmask variant uses one integer per prefix (which is still `O(n)` extra memory but often acceptable).
* The problem asks only for subsequences of length 3, making the alphabet-based approach the most natural.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
