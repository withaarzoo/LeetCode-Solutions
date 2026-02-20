# 761. Special Binary String

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

# Problem Title

1. Special Binary String

---

## Problem Summary

A special binary string has two properties:

1. The number of 0's is equal to the number of 1's.
2. For every prefix of the string, number of 1's is greater than or equal to number of 0's.

We are allowed to swap two consecutive special substrings.

Our goal is to return the lexicographically largest string possible after performing any number of valid swaps.

---

## Constraints

* 1 <= s.length <= 50
* s[i] is either '0' or '1'
* s is guaranteed to be a special binary string

---

## Intuition

When I read the problem, I noticed something interesting.

This behaves exactly like balanced parentheses.

* '1' behaves like '('
* '0' behaves like ')'

A valid special substring always looks like:

1 + (another special string) + 0

So I thought instead of trying many swaps, I can:

1. Break the string into smallest valid special substrings.
2. Recursively make each one largest.
3. Sort them in descending order.
4. Join them back together.

Sorting works because swapping consecutive special substrings is allowed.

---

## Approach

Step 1: Traverse the string with a counter.

* Increase counter for '1'
* Decrease counter for '0'

Step 2: Whenever counter becomes 0, we found one complete special substring.

Step 3: Remove outer 1 and 0.
Recursively solve the inner part.

Step 4: Store each processed substring.

Step 5: Sort substrings in descending lexicographical order.

Step 6: Concatenate all parts and return the result.

---

## Data Structures Used

* Vector / List / Array to store substrings
* Recursion stack
* Sorting function

---

## Operations & Behavior Summary

* Scan string once to split into valid blocks
* Recursive processing of inner substring
* Sorting of independent special substrings
* Concatenation to form final answer

---

## Complexity

Time Complexity: O(n^2 log n)

* We recursively process substrings.
* Sorting happens at each level.
* n is length of string.

Space Complexity: O(n)

* Recursion stack
* Storage of substrings

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string makeLargestSpecial(string s) {
        vector<string> parts;
        int count = 0, start = 0;

        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '1') count++;
            else count--;

            if (count == 0) {
                string inner = makeLargestSpecial(s.substr(start + 1, i - start - 1));
                parts.push_back("1" + inner + "0");
                start = i + 1;
            }
        }

        sort(parts.begin(), parts.end(), greater<string>());

        string result;
        for (auto &p : parts) result += p;
        return result;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public String makeLargestSpecial(String s) {
        List<String> parts = new ArrayList<>();
        int count = 0, start = 0;

        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '1') count++;
            else count--;

            if (count == 0) {
                String inner = makeLargestSpecial(s.substring(start + 1, i));
                parts.add("1" + inner + "0");
                start = i + 1;
            }
        }

        Collections.sort(parts, Collections.reverseOrder());

        StringBuilder result = new StringBuilder();
        for (String p : parts) result.append(p);
        return result.toString();
    }
}
```

---

### JavaScript

```javascript
var makeLargestSpecial = function(s) {
    let parts = [];
    let count = 0, start = 0;

    for (let i = 0; i < s.length; i++) {
        if (s[i] === '1') count++;
        else count--;

        if (count === 0) {
            let inner = makeLargestSpecial(s.substring(start + 1, i));
            parts.push("1" + inner + "0");
            start = i + 1;
        }
    }

    parts.sort((a, b) => b.localeCompare(a));
    return parts.join('');
};
```

---

### Python3

```python
class Solution:
    def makeLargestSpecial(self, s: str) -> str:
        parts = []
        count = 0
        start = 0

        for i, ch in enumerate(s):
            if ch == '1':
                count += 1
            else:
                count -= 1

            if count == 0:
                inner = self.makeLargestSpecial(s[start+1:i])
                parts.append("1" + inner + "0")
                start = i + 1

        parts.sort(reverse=True)
        return "".join(parts)
```

---

### Go

```go
import "sort"

func makeLargestSpecial(s string) string {
    var parts []string
    count := 0
    start := 0

    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            count++
        } else {
            count--
        }

        if count == 0 {
            inner := makeLargestSpecial(s[start+1 : i])
            parts = append(parts, "1"+inner+"0")
            start = i + 1
        }
    }

    sort.Slice(parts, func(i, j int) bool {
        return parts[i] > parts[j]
    })

    result := ""
    for _, p := range parts {
        result += p
    }

    return result
}
```

---

## Step-by-step Detailed Explanation

1. Initialize count = 0 and start = 0.

2. Traverse the string:

   * If character is '1', increment count.
   * If character is '0', decrement count.

3. Whenever count becomes 0:

   * We found a valid special substring.
   * Extract inner part using substring.
   * Recursively solve inner.
   * Wrap it with '1' and '0'.

4. Store all processed substrings.

5. Sort them in descending order.

6. Concatenate and return.

This ensures lexicographically largest result.

---

## Examples

Example 1:

Input: "11011000"
Output: "11100100"

Example 2:

Input: "10"
Output: "10"

---

## How to use / Run locally

C++

Compile:

g++ solution.cpp -o solution

Run:

./solution

Java

javac Solution.java
java Solution

Python

python solution.py

Go

go run solution.go

---

## Notes & Optimizations

* Think of the string as balanced parentheses.
* Recursion simplifies inner optimization.
* Sorting simulates optimal swapping.
* Constraint is small so recursion is safe.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
