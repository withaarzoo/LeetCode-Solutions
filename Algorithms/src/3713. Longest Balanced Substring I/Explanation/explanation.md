# 3713. Longest Balanced Substring I

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

You are given a string `s` consisting of lowercase English letters.

A substring is called **balanced** if all distinct characters inside that substring appear the same number of times.

Your task is to return the length of the longest balanced substring.

---

## Constraints

* 1 <= s.length <= 1000
* `s` contains only lowercase English letters

---

## Intuition

When I first read the problem, I focused on one important line:

All distinct characters in the substring must appear the same number of times.

So I thought like this:

If a substring has:

* 1 distinct character → always balanced
* 2 distinct characters → both must appear equal times
* 3 distinct characters → all three must appear equal times

That means:

length of substring = distinct_characters × frequency_of_each_character

Since n is at most 1000, I realized I can check all substrings using a brute force approach in O(n²), which is acceptable.

The key idea is:
While expanding a substring, keep track of:

* frequency of characters
* number of distinct characters
* maximum frequency

If at any point:

length == distinct × maxFrequency

then the substring is balanced.

---

## Approach

1. Loop through each index as starting point.
2. Create a frequency array of size 26.
3. Keep track of:

   * distinct characters
   * maximum frequency
4. Expand substring from i to j.
5. Update frequency and distinct count.
6. Check if length == distinct × maxFrequency.
7. Update answer if valid.

---

## Data Structures Used

* Frequency array of size 26
* Integer variables for:

  * distinct count
  * maximum frequency
  * result length

No extra complex data structures are used.

---

## Operations & Behavior Summary

For each substring:

* Increase character frequency
* Update distinct count
* Update max frequency
* Check balanced condition
* Update maximum answer

---

## Complexity

**Time Complexity:** O(n²)

* We check all substrings.
* n is length of string.

**Space Complexity:** O(1)

* Only 26-length frequency array used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int longestBalanced(string s) {
        int n = s.size();
        int ans = 0;

        for (int i = 0; i < n; i++) {
            vector<int> freq(26, 0);
            int distinct = 0;
            int maxFreq = 0;

            for (int j = i; j < n; j++) {
                int idx = s[j] - 'a';

                if (freq[idx] == 0)
                    distinct++;

                freq[idx]++;
                maxFreq = max(maxFreq, freq[idx]);

                int length = j - i + 1;

                if (length == distinct * maxFreq)
                    ans = max(ans, length);
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int longestBalanced(String s) {
        int n = s.length();
        int ans = 0;

        for (int i = 0; i < n; i++) {
            int[] freq = new int[26];
            int distinct = 0;
            int maxFreq = 0;

            for (int j = i; j < n; j++) {
                int idx = s.charAt(j) - 'a';

                if (freq[idx] == 0)
                    distinct++;

                freq[idx]++;
                maxFreq = Math.max(maxFreq, freq[idx]);

                int length = j - i + 1;

                if (length == distinct * maxFreq)
                    ans = Math.max(ans, length);
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
var longestBalanced = function(s) {
    const n = s.length;
    let ans = 0;

    for (let i = 0; i < n; i++) {
        let freq = new Array(26).fill(0);
        let distinct = 0;
        let maxFreq = 0;

        for (let j = i; j < n; j++) {
            let idx = s.charCodeAt(j) - 97;

            if (freq[idx] === 0)
                distinct++;

            freq[idx]++;
            maxFreq = Math.max(maxFreq, freq[idx]);

            let length = j - i + 1;

            if (length === distinct * maxFreq)
                ans = Math.max(ans, length);
        }
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def longestBalanced(self, s: str) -> int:
        n = len(s)
        ans = 0

        for i in range(n):
            freq = [0] * 26
            distinct = 0
            maxFreq = 0

            for j in range(i, n):
                idx = ord(s[j]) - ord('a')

                if freq[idx] == 0:
                    distinct += 1

                freq[idx] += 1
                maxFreq = max(maxFreq, freq[idx])

                length = j - i + 1

                if length == distinct * maxFreq:
                    ans = max(ans, length)

        return ans
```

### Go

```go
func longestBalanced(s string) int {
    n := len(s)
    ans := 0

    for i := 0; i < n; i++ {
        freq := make([]int, 26)
        distinct := 0
        maxFreq := 0

        for j := i; j < n; j++ {
            idx := int(s[j] - 'a')

            if freq[idx] == 0 {
                distinct++
            }

            freq[idx]++
            if freq[idx] > maxFreq {
                maxFreq = freq[idx]
            }

            length := j - i + 1

            if length == distinct*maxFreq {
                if length > ans {
                    ans = length
                }
            }
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation

For each starting index:

* Create frequency array
* Set distinct and maxFreq to 0

For each ending index:

* Increase frequency of current character
* If first occurrence, increase distinct
* Update maxFreq
* Compute length
* Check if length == distinct × maxFreq
* Update result

Repeat for all starting positions.

---

## Examples

Input: "abba"
Output: 4
Explanation: a and b both appear 2 times.

Input: "zzabccy"
Output: 4
Explanation: z, a, b, c each appear once in substring "zabc".

Input: "aba"
Output: 2
Explanation: "ab" or "ba".

---

## How to use / Run locally

1. Copy the solution into your local IDE.
2. Compile using your preferred language compiler.
3. Run with test cases.

Example for C++:

```bash
g++ file.cpp
./a.out
```

Example for Python:

```bash
python file.py
```

---

## Notes & Optimizations

* Since n ≤ 1000, O(n²) is efficient enough.
* We avoid re-checking frequencies by expanding substring incrementally.
* Space usage is constant.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
