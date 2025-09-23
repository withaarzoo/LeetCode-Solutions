# Problem Title

**165. Compare Version Numbers**

---

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

Given two version strings, `version1` and `version2`, compare them. A version string consists of numeric revisions separated by dots `.`. Each revision's value is the integer interpretation of the substring (leading zeros ignored). Compare corresponding revisions from left-to-right. If one version has fewer revisions, treat the missing revisions as `0`.

Return:

* `-1` if `version1 < version2`
* `1` if `version1 > version2`
* `0` otherwise

---

## Constraints

* `1 <= version1.length, version2.length <= 500`
* `version1` and `version2` only contain digits and `.`
* `version1` and `version2` are valid version numbers
* Each revision fits into a 32-bit integer (so no huge integers per revision)

---

## Intuition

I thought about how to compare versions like numbers written with dots — each part between dots is an integer. So I can move left-to-right, parse each revision into an integer (ignoring leading zeros automatically when parsing), and compare. If one side runs out of parts, I treat missing parts as `0`. To keep memory low, I parse in-place using two pointers rather than creating arrays from `split`.

---

## Approach

1. Keep two indices/pointers `i` and `j` for `version1` and `version2` respectively.
2. At each iteration, parse the next revision integer from each string by reading digits until a `.` or end-of-string.
3. Compare the parsed integers:

   * If `num1 > num2` return `1`.
   * If `num1 < num2` return `-1`.
   * Else continue to the next revision.
4. If both strings finish with no differences, return `0`.

This approach reads each character at most once and uses only O(1) extra space.

---

## Data Structures Used

* Primitive integer variables (pointers and accumulators).
* No arrays or extra containers in the two-pointer method.

---

## Operations & Behavior Summary

* Parsing a revision: loop until `.` or end, accumulate `num = num * 10 + (digit)`.
* Skipping delimiter: if current char is `.`, advance pointer by one.
* Comparison: compare parsed integers immediately to return as soon as a difference appears.

---

## Complexity

* **Time Complexity:** `O(n + m)` where `n = len(version1)` and `m = len(version2)`. Each character is processed once.
* **Space Complexity:** `O(1)` — only a few integer variables and pointers are used.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int compareVersion(string version1, string version2) {
        int i = 0, j = 0;
        int n = version1.size(), m = version2.size();

        while (i < n || j < m) {
            long num1 = 0, num2 = 0;
            // parse next number from version1
            while (i < n && version1[i] != '.') {
                num1 = num1 * 10 + (version1[i] - '0');
                ++i;
            }
            if (i < n && version1[i] == '.') ++i; // skip dot

            // parse next number from version2
            while (j < m && version2[j] != '.') {
                num2 = num2 * 10 + (version2[j] - '0');
                ++j;
            }
            if (j < m && version2[j] == '.') ++j; // skip dot

            if (num1 < num2) return -1;
            if (num1 > num2) return 1;
        }
        return 0;
    }
};
```

### Java

```java
class Solution {
    public int compareVersion(String version1, String version2) {
        int i = 0, j = 0;
        int n = version1.length(), m = version2.length();

        while (i < n || j < m) {
            long num1 = 0, num2 = 0;
            // parse next revision from version1
            while (i < n && version1.charAt(i) != '.') {
                num1 = num1 * 10 + (version1.charAt(i) - '0');
                i++;
            }
            if (i < n && version1.charAt(i) == '.') i++;

            // parse next revision from version2
            while (j < m && version2.charAt(j) != '.') {
                num2 = num2 * 10 + (version2.charAt(j) - '0');
                j++;
            }
            if (j < m && version2.charAt(j) == '.') j++;

            if (num1 < num2) return -1;
            if (num1 > num2) return 1;
        }
        return 0;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} version1
 * @param {string} version2
 * @return {number}
 */
var compareVersion = function(version1, version2) {
    let i = 0, j = 0;
    const n = version1.length, m = version2.length;

    while (i < n || j < m) {
        let num1 = 0, num2 = 0;
        // parse next number from version1
        while (i < n && version1[i] !== '.') {
            num1 = num1 * 10 + (version1.charCodeAt(i) - 48); // '0' -> 48
            i++;
        }
        if (i < n && version1[i] === '.') i++;

        // parse next number from version2
        while (j < m && version2[j] !== '.') {
            num2 = num2 * 10 + (version2.charCodeAt(j) - 48);
            j++;
        }
        if (j < m && version2[j] === '.') j++;

        if (num1 < num2) return -1;
        if (num1 > num2) return 1;
    }
    return 0;
};
```

### Python3

```python3
class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        i, j = 0, 0
        n, m = len(version1), len(version2)

        while i < n or j < m:
            num1 = 0
            num2 = 0
            # parse next revision from version1
            while i < n and version1[i] != '.':
                num1 = num1 * 10 + (ord(version1[i]) - ord('0'))
                i += 1
            if i < n and version1[i] == '.':
                i += 1

            # parse next revision from version2
            while j < m and version2[j] != '.':
                num2 = num2 * 10 + (ord(version2[j]) - ord('0'))
                j += 1
            if j < m and version2[j] == '.':
                j += 1

            if num1 < num2:
                return -1
            if num1 > num2:
                return 1

        return 0
```

### Go

```go
package main

func compareVersion(version1 string, version2 string) int {
    i, j := 0, 0
    n, m := len(version1), len(version2)

    for i < n || j < m {
        var num1, num2 int
        // parse next revision from version1
        for i < n && version1[i] != '.' {
            num1 = num1*10 + int(version1[i]-'0')
            i++
        }
        if i < n && version1[i] == '.' {
            i++
        }

        // parse next revision from version2
        for j < m && version2[j] != '.' {
            num2 = num2*10 + int(version2[j]-'0')
            j++
        }
        if j < m && version2[j] == '.' {
            j++
        }

        if num1 < num2 {
            return -1
        }
        if num1 > num2 {
            return 1
        }
    }
    return 0
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the same algorithm used in all language implementations. The code differs syntactically, but the logic is identical.

1. **Initialize pointers**: `i = 0` for `version1`, `j = 0` for `version2`.
2. **Loop while either string has characters**: `while (i < len1 || j < len2)` — ensures missing revisions act like `0`.
3. **Parse a revision from `version1`**:

   * `num1 = 0`
   * While `i` is inside the string and current char is not `.`, update `num1 = num1 * 10 + digit`.
   * After loop, if the current char is `.`, advance `i` to skip it.
4. **Parse a revision from `version2`** similarly to get `num2`.
5. **Compare `num1` and `num2`**:

   * If `num1 < num2` return `-1`.
   * If `num1 > num2` return `1`.
   * If equal, continue to parse the next revision.
6. **No difference found**: return `0`.

**Important implementation notes:**

* Using `long` for accumulation in languages like C++/Java is safe in case of moderate values; constraint says each revision fits in 32-bit but using a larger type is defensive.
* The parsing loops naturally ignore leading zeros (e.g., "01" => `1`).
* We never allocate intermediate arrays for `split`, keeping memory constant.

---

## Examples

1. Input: `version1 = "1.2"`, `version2 = "1.10"`
   Output: `-1`
   Explanation: `2 < 10` so `version1 < version2`.

2. Input: `version1 = "1.01"`, `version2 = "1.001"`
   Output: `0`
   Explanation: Leading zeros ignored: both revisions `01` and `001` represent `1`.

3. Input: `version1 = "1.0"`, `version2 = "1.0.0"`
   Output: `0`
   Explanation: Missing revisions treated as zero.

---

## How to use / Run locally

### C++ (g++ / clang++)

1. Save the class into a file `solution.cpp` and add a `main()` test harness which calls `Solution::compareVersion` with test inputs (LeetCode platform does not require `main`).
2. Compile and run:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

1. Use the `Solution` class in `Solution.java`. To test locally, either embed a `public static void main` that calls the method, or use an online Java runner.
2. Compile and run:

```bash
javac Solution.java
java Solution
```

### JavaScript (Node.js)

1. Put the `compareVersion` function in `compareVersion.js` and export it or create a small test harness in the same file.
2. Run:

```bash
node compareVersion.js
```

### Python3

1. Put the `Solution` class or a simple wrapper in `compare_version.py`.
2. Run the file with Python:

```bash
python3 compare_version.py
```

### Go

1. Put `compareVersion` in a `main.go` with a `main()` that calls it with test inputs.
2. Build and run:

```bash
go run main.go
```

---

## Notes & Optimizations

* There is an alternate simpler implementation using `split('.')` on both strings, then comparing integer conversions of each substring. That is easier to read but allocates arrays and substrings, which costs extra memory and (slightly) extra time.
* The two-pointer method implemented here is memory friendly (O(1) extra space) and linear-time.
* For languages without automatic big integers, watch out for extremely large revisions; the problem guarantees that each revision fits in 32-bit int.
* This solution returns as soon as it finds a difference — no need to parse the rest.

---

## Author

Created by **[Md. Aarzoo Islam](https://bento.me/withaarzoo)**. Feel free to open issues or pull requests for corrections or improvements.
