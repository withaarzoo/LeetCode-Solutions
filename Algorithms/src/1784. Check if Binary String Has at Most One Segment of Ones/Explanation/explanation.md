# 1784. Check if Binary String Has at Most One Segment of Ones

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

Given a binary string `s` without leading zeros, determine whether the string contains **at most one contiguous segment of ones**.

A segment of ones means a continuous block of `1` characters.

If the string contains only one block of `1`s, return `true`. Otherwise return `false`.

Example:

* `111000` → valid (one segment of ones)
* `1001` → invalid (two segments of ones)

---

## Constraints

* `1 <= s.length <= 100`
* `s[i]` is either `'0'` or `'1'`
* `s[0]` is `'1'`

---

## Intuition

When I read the problem, I noticed something important.

The string does not contain leading zeros, which means the string always starts with `1`. That means the first segment of ones starts immediately.

The only way a second segment of ones can appear is if:

1. I encounter a `0`
2. Later another `1` appears again

That pattern looks like `01`.

So if the substring `01` appears anywhere in the string, it means a new segment of ones started after a zero. That means the string has more than one segment of ones.

Therefore the solution becomes very simple:

If the string contains `01`, return `false`.

If it does not contain `01`, return `true`.

---

## Approach

1. Start scanning the binary string.
2. Check if the substring `01` exists.
3. If `01` appears, that means:

   * the first segment of ones ended
   * a zero appeared
   * another `1` started a new segment
4. In that case return `false`.
5. If the substring `01` never appears, all ones belong to a single segment.
6. Return `true`.

---

## Data Structures Used

No special data structures are required.

The algorithm only inspects characters in the given string.

Space usage remains constant.

---

## Operations & Behavior Summary

| Operation       | Purpose                                   |
| --------------- | ----------------------------------------- |
| String Scan     | Iterate through the binary string         |
| Substring Check | Detect pattern `01`                       |
| Boolean Return  | Determine if more than one segment exists |

---

## Complexity

Time Complexity: `O(n)`

`n` is the length of the binary string. The string is scanned once to check whether the pattern `01` appears.

Space Complexity: `O(1)`

No extra data structures are used. Only constant memory is required.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool checkOnesSegment(string s) {
        // If "01" appears, another segment of ones exists
        return s.find("01") == string::npos;
    }
};
```

### Java

```java
class Solution {
    public boolean checkOnesSegment(String s) {
        return !s.contains("01");
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {boolean}
 */
var checkOnesSegment = function(s) {
    return !s.includes("01");
};
```

### Python3

```python
class Solution:
    def checkOnesSegment(self, s: str) -> bool:
        return "01" not in s
```

### Go

```go
func checkOnesSegment(s string) bool {
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' && s[i+1] == '1' {
            return false
        }
    }
    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core idea in every language is identical. We only check whether the pattern `01` exists.

### Step 1

The string starts with `1`, so the first segment of ones begins immediately.

Example:

```bash
111000
```

The first segment is `111`.

### Step 2

If we ever see `01`, it means a new segment of ones starts.

Example:

```bash
1001
```

Breaking it down:

```bash
1 -> first segment starts
0 -> ones segment ended
0 -> still zero
1 -> new ones segment started
```

This means there are two segments of ones.

### Step 3

Languages like C++, Java, JavaScript, and Python allow direct substring checks.

Example in C++:

```cpp
s.find("01")
```

If the result is not found, it returns `string::npos`.

So we check:

```bash
s.find("01") == string::npos
```

If true, only one segment exists.

### Step 4

In Go, we manually scan adjacent characters.

```bash
for i := 0; i < len(s)-1; i++
```

Then check:

```bash
s[i] == '0' && s[i+1] == '1'
```

If that condition is true, return `false`.

Otherwise continue scanning.

### Step 5

If the entire string is scanned and `01` never appears, then the string contains only one segment of ones.

Return `true`.

---

## Examples

Example 1

Input

```bash
s = "1001"
```

Output

```bash
false
```

Explanation

The ones appear in two separate segments.

---

Example 2

Input

```bash
s = "110"
```

Output

```bash
true
```

Explanation

There is only one continuous block of ones.

---

## How to use / Run locally

Example using C++.

1. Create a file `main.cpp`
2. Paste the C++ solution
3. Compile the program

```bash
g++ main.cpp -o solution
```

1. Run the executable

```bash
./solution
```

For Python

```bash
python solution.py
```

For JavaScript

```bash
node solution.js
```

---

## Notes & Optimizations

* The solution works in linear time.
* Only a single scan of the string is needed.
* No additional memory is required.
* Detecting the pattern `01` is the simplest and most efficient way to determine whether multiple segments of ones exist.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
