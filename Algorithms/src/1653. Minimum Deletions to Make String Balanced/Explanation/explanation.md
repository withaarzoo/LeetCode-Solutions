# Problem Title

**1653. Minimum Deletions to Make String Balanced**

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

I am given a string `s` that contains only two characters: `'a'` and `'b'`.

The string is called **balanced** if there is **no `'b'` before any `'a'`**.
In simple words, all `'a'` characters must come **before** all `'b'` characters.

I can delete any number of characters from the string.
My task is to find the **minimum number of deletions** required to make the string balanced.

---

## Constraints

* `1 <= s.length <= 10^5`
* `s[i]` is either `'a'` or `'b'`

---

## Intuition

When I started thinking about the problem, I first tried to understand **what makes the string unbalanced**.

I noticed that the problem happens only when:

* a `'b'` appears **before**
* an `'a'`

So the final string must look like:

```bash
aaaaa....bbbbb
```

Whenever I see an `'a'` after some `'b'`, I must fix it.

At that moment, I have only **two choices**:

1. Delete the current `'a'`
2. Delete all previous `'b'`

I realized that at every step, I should choose the option that needs **fewer deletions**.

This observation helped me design a greedy and efficient solution.

---

## Approach

I solve this problem in **one pass**.

I maintain two variables:

* `countB` → how many `'b'` characters I have seen so far
* `deletions` → minimum deletions needed till the current index

Now I scan the string from left to right:

1. If the character is `'b'`

   * I increase `countB`
   * No issue yet

2. If the character is `'a'`

   * This `'a'` becomes a problem if `countB > 0`
   * I calculate:

     * delete this `'a'` → `deletions + 1`
     * delete all previous `'b'` → `countB`
   * I take the minimum of both

This way, I always keep the string balanced with minimum cost.

---

## Data Structures Used

* No complex data structures
* Only integer variables are used

This keeps the solution:

* Simple
* Fast
* Memory efficient

---

## Operations & Behavior Summary

* Traverse string once
* Count previous `'b'`
* Decide minimum deletion at every `'a'`
* Greedy decision at each step
* Final answer stored in `deletions`

---

## Complexity

**Time Complexity:** `O(n)`

* `n` is the length of the string
* The string is traversed only once

**Space Complexity:** `O(1)`

* No extra space is used
* Only constant variables are maintained

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumDeletions(string s) {
        int countB = 0;
        int deletions = 0;

        for (char ch : s) {
            if (ch == 'b') {
                countB++;
            } else {
                deletions = min(deletions + 1, countB);
            }
        }
        return deletions;
    }
};
```

---

### Java

```java
class Solution {
    public int minimumDeletions(String s) {
        int countB = 0;
        int deletions = 0;

        for (char ch : s.toCharArray()) {
            if (ch == 'b') {
                countB++;
            } else {
                deletions = Math.min(deletions + 1, countB);
            }
        }
        return deletions;
    }
}
```

---

### JavaScript

```javascript
var minimumDeletions = function(s) {
    let countB = 0;
    let deletions = 0;

    for (let ch of s) {
        if (ch === 'b') {
            countB++;
        } else {
            deletions = Math.min(deletions + 1, countB);
        }
    }
    return deletions;
};
```

---

### Python3

```python
class Solution:
    def minimumDeletions(self, s: str) -> int:
        countB = 0
        deletions = 0

        for ch in s:
            if ch == 'b':
                countB += 1
            else:
                deletions = min(deletions + 1, countB)

        return deletions
```

---

### Go

```go
func minimumDeletions(s string) int {
    countB := 0
    deletions := 0

    for _, ch := range s {
        if ch == 'b' {
            countB++
        } else {
            if deletions+1 < countB {
                deletions = deletions + 1
            } else {
                deletions = countB
            }
        }
    }
    return deletions
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. I initialize `countB = 0` and `deletions = 0`
2. I start reading characters from left to right
3. Every time I see `'b'`, I increase `countB`
4. When I see `'a'`:

   * I check how many `'b'` came before
   * I calculate the cost of deleting this `'a'`
   * I calculate the cost of deleting previous `'b'`
   * I choose the minimum option
5. After the loop ends, `deletions` contains the final answer

This logic is same in all languages.

---

## Examples

### Example 1

**Input**

```bash
s = "aababbab"
```

**Output**

```bash
2
```

**Explanation**
I delete two characters to make all `'a'` come before `'b'`.

---

### Example 2

**Input**

```bash
s = "bbaaaaabb"
```

**Output**

```bash
2
```

**Explanation**
The best option is to delete the first two `'b'`.

---

## How to use / Run locally

1. Copy the code for your preferred language
2. Paste it into:

   * LeetCode editor, or
   * Your local compiler / IDE
3. Call `minimumDeletions(s)` with input string
4. Run and get the answer

---

## Notes & Optimizations

* This is an optimal greedy solution
* No DP array needed
* Works efficiently even for large input size
* Very common interview problem
* Clean logic with real-world reasoning

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
