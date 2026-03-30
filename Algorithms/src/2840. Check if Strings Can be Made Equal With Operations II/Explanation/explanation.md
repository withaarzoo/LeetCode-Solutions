# 2840. Check if Strings Can be Made Equal With Operations II

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

You are given two strings `s1` and `s2` of the same length.

You can swap any two characters in a string only if:

* `i < j`
* `(j - i)` is even

This means:

* Characters at even indexes can only move to even indexes.
* Characters at odd indexes can only move to odd indexes.

You need to return `true` if both strings can be made equal after performing any number of such operations.

---

## Constraints

```text
n == s1.length == s2.length
1 <= n <= 10^5
s1 and s2 consist only of lowercase English letters
```

---

## Intuition

I thought about what kind of swaps are actually possible.

Since I can only swap indices whose difference is even:

* Even index characters can only move among even positions.
* Odd index characters can only move among odd positions.

So instead of thinking about all possible swaps, I only need to check:

1. Do both strings have the same characters at even indexes?
2. Do both strings have the same characters at odd indexes?

If both are true, then I can always rearrange them to make the strings equal.

---

## Approach

1. Create two frequency arrays:

   * One for even index characters
   * One for odd index characters
2. Traverse both strings together.
3. For every index:

   * If index is even:

     * Increase count for `s1[i]`
     * Decrease count for `s2[i]`
   * If index is odd:

     * Increase count for `s1[i]`
     * Decrease count for `s2[i]`
4. At the end, if all frequencies become `0`, return `true`.
5. Otherwise, return `false`.

---

## Data Structures Used

* Two integer arrays of size `26`

  * `even[26]`
  * `odd[26]`

These arrays store the frequency difference of characters between `s1` and `s2`.

---

## Operations & Behavior Summary

| Index Type | Allowed Movement              |
| ---------- | ----------------------------- |
| Even Index | Can move only to even indexes |
| Odd Index  | Can move only to odd indexes  |

| What I Compare                   | Why                                                  |
| -------------------------------- | ---------------------------------------------------- |
| Even-index character frequencies | Because even positions only swap with even positions |
| Odd-index character frequencies  | Because odd positions only swap with odd positions   |

---

## Complexity

* Time Complexity: `O(n)`

  * I traverse the strings only once.
  * `n` is the length of the strings.

* Space Complexity: `O(1)`

  * I only use fixed-size arrays of size `26`.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool checkStrings(string s1, string s2) {
        vector<int> even(26, 0), odd(26, 0);

        for (int i = 0; i < s1.size(); i++) {
            if (i % 2 == 0) {
                even[s1[i] - 'a']++;
                even[s2[i] - 'a']--;
            } else {
                odd[s1[i] - 'a']++;
                odd[s2[i] - 'a']--;
            }
        }

        for (int i = 0; i < 26; i++) {
            if (even[i] != 0 || odd[i] != 0) {
                return false;
            }
        }

        return true;
    }
};
```

### Java

```java
class Solution {
    public boolean checkStrings(String s1, String s2) {
        int[] even = new int[26];
        int[] odd = new int[26];

        for (int i = 0; i < s1.length(); i++) {
            if (i % 2 == 0) {
                even[s1.charAt(i) - 'a']++;
                even[s2.charAt(i) - 'a']--;
            } else {
                odd[s1.charAt(i) - 'a']++;
                odd[s2.charAt(i) - 'a']--;
            }
        }

        for (int i = 0; i < 26; i++) {
            if (even[i] != 0 || odd[i] != 0) {
                return false;
            }
        }

        return true;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var checkStrings = function(s1, s2) {
    const even = new Array(26).fill(0);
    const odd = new Array(26).fill(0);

    for (let i = 0; i < s1.length; i++) {
        if (i % 2 === 0) {
            even[s1.charCodeAt(i) - 97]++;
            even[s2.charCodeAt(i) - 97]--;
        } else {
            odd[s1.charCodeAt(i) - 97]++;
            odd[s2.charCodeAt(i) - 97]--;
        }
    }

    for (let i = 0; i < 26; i++) {
        if (even[i] !== 0 || odd[i] !== 0) {
            return false;
        }
    }

    return true;
};
```

### Python3

```python
class Solution:
    def checkStrings(self, s1: str, s2: str) -> bool:
        even = [0] * 26
        odd = [0] * 26

        for i in range(len(s1)):
            if i % 2 == 0:
                even[ord(s1[i]) - ord('a')] += 1
                even[ord(s2[i]) - ord('a')] -= 1
            else:
                odd[ord(s1[i]) - ord('a')] += 1
                odd[ord(s2[i]) - ord('a')] -= 1

        for i in range(26):
            if even[i] != 0 or odd[i] != 0:
                return False

        return True
```

### Go

```go
func checkStrings(s1 string, s2 string) bool {
    even := make([]int, 26)
    odd := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        if i%2 == 0 {
            even[s1[i]-'a']++
            even[s2[i]-'a']--
        } else {
            odd[s1[i]-'a']++
            odd[s2[i]-'a']--
        }
    }

    for i := 0; i < 26; i++ {
        if even[i] != 0 || odd[i] != 0 {
            return false
        }
    }

    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Create frequency arrays

I create two arrays:

* `even[26]`
* `odd[26]`

These arrays store character frequencies separately for even and odd indexes.

```text
even[0] -> count of 'a' at even positions
even[1] -> count of 'b' at even positions
...
```

---

### Step 2: Traverse both strings together

I run a loop from `0` to `n - 1`.

For every index:

* If index is even:

  * Add frequency of `s1[i]`
  * Remove frequency of `s2[i]`

* If index is odd:

  * Add frequency of `s1[i]`
  * Remove frequency of `s2[i]`

This helps me compare both strings in one pass.

---

### Step 3: Check if all counts become zero

At the end:

* If every value in `even[]` is `0`
* And every value in `odd[]` is `0`

Then both strings have identical characters at even and odd positions.

So I return `true`.

Otherwise, I return `false`.

---

## Examples

### Example 1

```text
Input:
s1 = "abcdba"
s2 = "cabdab"

Output:
true
```

Explanation:

```text
Even positions:
s1 -> a, c, b
s2 -> c, b, a

Odd positions:
s1 -> b, d, a
s2 -> a, d, b
```

Both have the same character frequencies.

---

### Example 2

```text
Input:
s1 = "abe"
s2 = "bea"

Output:
false
```

Explanation:

```text
Even positions:
s1 -> a, e
s2 -> b, a
```

Even-position frequencies do not match.

---

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* I do not sort anything.
* Sorting would take `O(n log n)`.
* Frequency counting is faster and works in `O(n)`.
* Since only lowercase English letters are used, array size remains fixed at `26`.
* This makes the solution very memory efficient.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
