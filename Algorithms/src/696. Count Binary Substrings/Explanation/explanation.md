# 696. Count Binary Substrings

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given a binary string `s`, I need to return the number of non-empty substrings that:

1. Have equal number of `0`s and `1`s.
2. All `0`s are grouped together.
3. All `1`s are grouped together.

Substrings that appear multiple times must be counted multiple times.

---

## Constraints

* `1 <= s.length <= 10^5`
* `s[i]` is either `'0'` or `'1'`

---

## Intuition

When I first looked at the problem, I noticed something important.

Valid substrings always look like this:

* `0011`
* `01`
* `1100`

That means:

* All `0`s must come together.
* All `1`s must come together.
* The counts must be equal.

So I thought instead of checking all substrings, I should count consecutive groups.

Example:

`00110011`

Groups become:

* `00` → length 2
* `11` → length 2
* `00` → length 2
* `11` → length 2

Between every two adjacent groups, the number of valid substrings equals:

`min(previous_group_length, current_group_length)`

Because I can only match pairs up to the smaller group.

---

## Approach

1. Traverse the string once.
2. Count the size of consecutive groups.
3. When character changes:

   * Add `min(prevGroup, currGroup)` to result.
   * Update `prevGroup`.
   * Reset `currGroup` to 1.
4. After loop ends, add last `min(prevGroup, currGroup)`.
5. Return result.

No nested loops.
No substring generation.
Only one linear pass.

---

## Data Structures Used

* Integer variables only:

  * `prevGroup`
  * `currGroup`
  * `result`

No arrays.
No extra storage.

---

## Operations & Behavior Summary

* Count consecutive characters.
* Compare adjacent group sizes.
* Add minimum of adjacent groups to result.
* Continue until string ends.

---

## Complexity

**Time Complexity:** O(n)

* `n` = length of string.
* I traverse the string once.

**Space Complexity:** O(1)

* Only constant variables used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countBinarySubstrings(string s) {
        int prevGroup = 0;
        int currGroup = 1;
        int result = 0;
        
        for (int i = 1; i < s.length(); i++) {
            if (s[i] == s[i - 1]) {
                currGroup++;
            } else {
                result += min(prevGroup, currGroup);
                prevGroup = currGroup;
                currGroup = 1;
            }
        }
        
        result += min(prevGroup, currGroup);
        return result;
    }
};
```

---

### Java

```java
class Solution {
    public int countBinarySubstrings(String s) {
        int prevGroup = 0;
        int currGroup = 1;
        int result = 0;

        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) == s.charAt(i - 1)) {
                currGroup++;
            } else {
                result += Math.min(prevGroup, currGroup);
                prevGroup = currGroup;
                currGroup = 1;
            }
        }

        result += Math.min(prevGroup, currGroup);
        return result;
    }
}
```

---

### JavaScript

```javascript
var countBinarySubstrings = function(s) {
    let prevGroup = 0;
    let currGroup = 1;
    let result = 0;

    for (let i = 1; i < s.length; i++) {
        if (s[i] === s[i - 1]) {
            currGroup++;
        } else {
            result += Math.min(prevGroup, currGroup);
            prevGroup = currGroup;
            currGroup = 1;
        }
    }

    result += Math.min(prevGroup, currGroup);
    return result;
};
```

---

### Python3

```python
class Solution:
    def countBinarySubstrings(self, s: str) -> int:
        prevGroup = 0
        currGroup = 1
        result = 0

        for i in range(1, len(s)):
            if s[i] == s[i - 1]:
                currGroup += 1
            else:
                result += min(prevGroup, currGroup)
                prevGroup = currGroup
                currGroup = 1

        result += min(prevGroup, currGroup)
        return result
```

---

### Go

```go
func countBinarySubstrings(s string) int {
    prevGroup := 0
    currGroup := 1
    result := 0

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            currGroup++
        } else {
            if prevGroup < currGroup {
                result += prevGroup
            } else {
                result += currGroup
            }
            prevGroup = currGroup
            currGroup = 1
        }
    }

    if prevGroup < currGroup {
        result += prevGroup
    } else {
        result += currGroup
    }

    return result
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Initialize:

   * `prevGroup = 0`
   * `currGroup = 1`
   * `result = 0`

2. Start loop from index 1.

3. If current character equals previous:

   * Increase `currGroup`.

4. If character changes:

   * Add `min(prevGroup, currGroup)` to `result`.
   * Update `prevGroup = currGroup`.
   * Reset `currGroup = 1`.

5. After loop ends:

   * Add final `min(prevGroup, currGroup)`.

6. Return `result`.

Same logic works in all languages.

---

## Examples

**Example 1:**

Input:

```bash
00110011
```

Output:

```bash
6
```

---

**Example 2:**

Input:

```bash
10101
```

Output:

```bash
4
```

---

## How to use / Run locally

### C++

```bash
g++ file.cpp -o output
./output
```

### Java

```bash
javac Solution.java
java Solution
```

### Python

```bash
python file.py
```

### JavaScript

```bash
node file.js
```

### Go

```bash
go run file.go
```

---

## Notes & Optimizations

* Avoid brute force substring generation.
* Never use nested loops.
* One pass solution is optimal.
* Works efficiently even for length 10^5.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
