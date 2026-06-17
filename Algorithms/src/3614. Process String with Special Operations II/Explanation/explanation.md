# LeetCode 3614. Process String with Special Operations II

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given a string `s` containing lowercase English letters and special characters `*`, `#`, and `%`, we need to process the string from left to right and determine the character located at index `k` in the final generated string.

The special operations behave as follows:

| Character | Operation                            |
| --------- | ------------------------------------ |
| `a-z`     | Append the character                 |
| `*`       | Remove the last character if present |
| `#`       | Duplicate the current string         |
| `%`       | Reverse the current string           |

The challenge is that the final string can become extremely large, making direct simulation impossible.

Instead of building the actual string, we need an efficient algorithm that can determine the kth character while handling string lengths up to `10^15`.

This problem combines string processing, reverse simulation, indexing, and length tracking techniques commonly used in advanced DSA and competitive programming problems.

---

## Constraints

| Constraint              | Value                            |
| ----------------------- | -------------------------------- |
| `1 <= s.length <= 10^5` | String length                    |
| `0 <= k <= 10^15`       | Target index                     |
| Characters              | Lowercase letters, `*`, `#`, `%` |
| Final processed length  | At most `10^15`                  |

---

## Intuition

My first thought was to simulate the entire string and apply every operation exactly as described.

That quickly breaks down because the `#` operation doubles the current string. After several duplications, the string can grow to an enormous size.

I realized that I never actually need the full string.

The only thing I care about is:

> Which original character eventually ends up at position `k`?

That observation changes the problem completely.

Instead of constructing the final string, I track only the length after each operation. Then I walk backward through the operations and keep translating index `k` back to its previous position until I reach the original character responsible for that index.

---

## Approach

1. Process the string from left to right.
2. Store the length after every operation.
3. Never build the actual string.
4. After processing all operations, check whether `k` is inside the final length.
5. Traverse the operations in reverse order.
6. Undo each operation:

   * Undo duplication using modulo.
   * Undo reversal using mirrored indexing.
   * Handle appended letters.
   * Ignore removed characters because they no longer exist.
7. Continue until the exact character responsible for index `k` is found.
8. Return that character.

This reverse simulation approach avoids creating huge strings and works efficiently even for the largest test cases.

---

## Data Structures Used

### Array / Vector

Used to store the resulting string length after each operation.

Why?

Because while moving backward, I need to know:

* Length before an operation
* Length after an operation

Having these values precomputed makes reversing operations easy.

### Integer Variables

Used to maintain:

* Current processed length
* Target index `k`

All lengths use 64-bit integers because values can reach `10^15`.

---

## Operations & Behavior Summary

The algorithm performs the following stages:

### Stage 1: Track Lengths

Process each character and update the current length.

* Letter → length + 1
* `*` → length - 1 if possible
* `#` → length × 2
* `%` → length unchanged

### Stage 2: Validate k

If:

```text
k >= finalLength
```

then the required index does not exist.

Return:

```text
'.'
```

### Stage 3: Reverse Simulation

Move from right to left.

#### Letter

Check whether the current index points to the newly appended character.

#### Duplicate (`#`)

Map index back into the original half.

#### Reverse (`%`)

Mirror the index.

#### Delete (`*`)

No adjustment is required because surviving characters keep their positions.

### Stage 4: Return Answer

As soon as a matching appended character is found, return it.

---

## Complexity

| Metric           | Complexity | Explanation                            |
| ---------------- | ---------- | -------------------------------------- |
| Time Complexity  | `O(n)`     | One forward pass and one backward pass |
| Space Complexity | `O(n)`     | Stores length after every operation    |

Where:

* `n` = length of input string `s`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    char processStr(string s, long long k) {
        int n = s.size();

        // len[i] = length of the result string after processing s[i]
        vector<long long> len(n);
        long long curLen = 0;

        for (int i = 0; i < n; i++) {
            char c = s[i];

            if (c >= 'a' && c <= 'z') {
                // Appending a letter increases length by 1
                curLen++;
            } else if (c == '*') {
                // Remove last character if it exists
                if (curLen > 0) curLen--;
            } else if (c == '#') {
                // Duplicate the whole string
                curLen *= 2;
            } else { // '%'
                // Reversing does not change length
            }

            len[i] = curLen;
        }

        // k is outside the final string
        if (k >= curLen) return '.';

        // Walk backwards and undo operations
        for (int i = n - 1; i >= 0; i--) {
            char c = s[i];

            long long after = len[i];
            long long before = (i == 0 ? 0 : len[i - 1]);

            if (c >= 'a' && c <= 'z') {
                // This letter was appended at index "before"
                if (k == before) return c;
            } else if (c == '#') {
                // T + T -> map back into the first copy
                if (before > 0) k %= before;
            } else if (c == '%') {
                // Reverse operation
                k = before - 1 - k;
            } else {
                // '*' removed the last character.
                // All remaining indices stay unchanged.
            }
        }

        return '.';
    }
};
```

### Java

```java
class Solution {
    public char processStr(String s, long k) {
        int n = s.length();

        // len[i] = length after processing s[i]
        long[] len = new long[n];
        long curLen = 0;

        for (int i = 0; i < n; i++) {
            char c = s.charAt(i);

            if (c >= 'a' && c <= 'z') {
                // Append character
                curLen++;
            } else if (c == '*') {
                // Remove last character if present
                if (curLen > 0) curLen--;
            } else if (c == '#') {
                // Duplicate string
                curLen *= 2;
            } else { // '%'
                // Length unchanged
            }

            len[i] = curLen;
        }

        // Out of bounds
        if (k >= curLen) return '.';

        // Undo operations from right to left
        for (int i = n - 1; i >= 0; i--) {
            char c = s.charAt(i);

            long before = (i == 0 ? 0 : len[i - 1]);

            if (c >= 'a' && c <= 'z') {
                // Letter was appended at index "before"
                if (k == before) return c;
            } else if (c == '#') {
                // Undo duplication
                if (before > 0) k %= before;
            } else if (c == '%') {
                // Undo reverse
                k = before - 1 - k;
            } else {
                // '*' needs no index adjustment
            }
        }

        return '.';
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {number} k
 * @return {character}
 */
var processStr = function(s, k) {
    const n = s.length;

    // Use BigInt because lengths can reach 1e15
    const len = new Array(n);
    let curLen = 0n;

    for (let i = 0; i < n; i++) {
        const c = s[i];

        if (c >= 'a' && c <= 'z') {
            // Append character
            curLen++;
        } else if (c === '*') {
            // Remove last character if it exists
            if (curLen > 0n) curLen--;
        } else if (c === '#') {
            // Duplicate string
            curLen *= 2n;
        } else {
            // '%' does not change length
        }

        len[i] = curLen;
    }

    let idx = BigInt(k);

    // Out of bounds
    if (idx >= curLen) return '.';

    // Undo operations
    for (let i = n - 1; i >= 0; i--) {
        const c = s[i];
        const before = (i === 0 ? 0n : len[i - 1]);

        if (c >= 'a' && c <= 'z') {
            // Letter was appended at position "before"
            if (idx === before) return c;
        } else if (c === '#') {
            // Undo duplication
            if (before > 0n) idx %= before;
        } else if (c === '%') {
            // Undo reverse
            idx = before - 1n - idx;
        } else {
            // '*' keeps surviving indices unchanged
        }
    }

    return '.';
};
```

### Python3

```python
class Solution:
    def processStr(self, s: str, k: int) -> str:
        n = len(s)

        # lengths[i] = length after processing s[i]
        lengths = [0] * n
        cur_len = 0

        for i, ch in enumerate(s):
            if 'a' <= ch <= 'z':
                # Append a character
                cur_len += 1
            elif ch == '*':
                # Remove last character if present
                if cur_len > 0:
                    cur_len -= 1
            elif ch == '#':
                # Duplicate the whole string
                cur_len *= 2
            else:
                # '%' only reverses, length stays same
                pass

            lengths[i] = cur_len

        # k is outside the final string
        if k >= cur_len:
            return '.'

        # Undo operations from right to left
        for i in range(n - 1, -1, -1):
            ch = s[i]
            before = 0 if i == 0 else lengths[i - 1]

            if 'a' <= ch <= 'z':
                # This letter was appended at index "before"
                if k == before:
                    return ch
            elif ch == '#':
                # Undo T + T
                if before > 0:
                    k %= before
            elif ch == '%':
                # Undo reverse
                k = before - 1 - k
            else:
                # '*' does not change surviving indices
                pass

        return '.'
```

### Go

```go
func processStr(s string, k int64) byte {
 n := len(s)

 // lengths[i] = length after processing s[i]
 lengths := make([]int64, n)
 var curLen int64 = 0

 for i := 0; i < n; i++ {
  c := s[i]

  if c >= 'a' && c <= 'z' {
   // Append a character
   curLen++
  } else if c == '*' {
   // Remove last character if present
   if curLen > 0 {
    curLen--
   }
  } else if c == '#' {
   // Duplicate the whole string
   curLen *= 2
  } else {
   // '%' only reverses, length unchanged
  }

  lengths[i] = curLen
 }

 // k is outside the final string
 if k >= curLen {
  return '.'
 }

 // Undo operations from right to left
 for i := n - 1; i >= 0; i-- {
  c := s[i]

  var before int64
  if i > 0 {
   before = lengths[i-1]
  }

  if c >= 'a' && c <= 'z' {
   // Letter was appended at index "before"
   if k == before {
    return c
   }
  } else if c == '#' {
   // Undo duplication
   if before > 0 {
    k %= before
   }
  } else if c == '%' {
   // Undo reverse
   k = before - 1 - k
  } else {
   // '*' keeps surviving indices unchanged
  }
 }

 return '.'
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across all five languages.

Only syntax changes.

### Step 1: Record Lengths

While scanning the string, I maintain a running length.

For example:

```text
a#b
```

Processing:

```text
a    -> length = 1
#    -> length = 2
b    -> length = 3
```

I store each intermediate length.

This gives me enough information to reverse every operation later.

---

### Step 2: Check Whether k Exists

Suppose the final length is:

```text
8
```

and:

```text
k = 10
```

Index 10 is outside the string.

The answer is immediately:

```text
'.'
```

No further work is needed.

---

### Step 3: Walk Backward

Now the interesting part begins.

Instead of asking:

```text
What is character k?
```

I ask:

```text
Where did position k come from?
```

Every operation is reversed.

---

### Step 4: Undo Appended Letters

Suppose:

```text
ab
```

becomes:

```text
abc
```

Character `c` occupies the last position.

If my current index points exactly to that location, then I have found the answer.

Otherwise, the answer belongs to the previous string.

---

### Step 5: Undo Duplication

Consider:

```text
abc
```

After duplication:

```text
abcabc
```

Indices:

```text
0 1 2 3 4 5
a b c a b c
```

If I want:

```text
index 4
```

it corresponds to:

```text
index 1
```

in the original string.

That is why:

```text
k %= originalLength
```

correctly maps the position back.

---

### Step 6: Undo Reversal

Suppose:

```text
abc
```

becomes:

```text
cba
```

Length is:

```text
3
```

Mapping:

```text
0 -> 2
1 -> 1
2 -> 0
```

The reverse mapping becomes:

```text
k = length - 1 - k
```

This restores the original position.

---

### Step 7: Handle Deletion

Suppose:

```text
abc
```

becomes:

```text
ab
```

Only the last character disappears.

All remaining indices stay exactly where they are.

Therefore no index transformation is needed.

---

### Step 8: Find the Character

Eventually the reverse traversal reaches the exact letter that originally created position `k`.

That letter is returned as the answer.

---

## Examples

### Example 1

Input

```text
s = "a#b%*"
k = 1
```

Output

```text
a
```

Trace

```text
a      -> a
#      -> aa
b      -> aab
%      -> baa
*      -> ba
```

Index 1 contains:

```text
a
```

---

### Example 2

Input

```text
s = "cd%#*#"
k = 3
```

Output

```text
d
```

Trace

```text
cd
dc
dcdc
dcd
dcddcd
```

Index 3 contains:

```text
d
```

---

### Example 3

Input

```text
s = "z*#"
k = 0
```

Output

```text
.
```

Trace

```text
z
(empty)
(empty)
```

No character exists at index 0.

---

## How to Use / Run Locally

### C++

Compile

```bash
g++ main.cpp -O2 -std=c++17
```

Run

```bash
./a.out
```

---

### Java

Compile

```bash
javac Main.java
```

Run

```bash
java Main
```

---

### JavaScript

Run

```bash
node main.js
```

---

### Python3

Run

```bash
python main.py
```

---

### Go

Run

```bash
go run main.go
```

---

## Notes & Optimizations

* Never construct the final string.
* The final length can reach `10^15`.
* Reverse simulation is the key optimization.
* Length tracking makes every operation reversible.
* A naive solution may cause memory overflow.
* This approach easily handles the largest constraints.
* Using 64-bit integers is mandatory.
* The solution is optimal for both time and space within the problem constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
