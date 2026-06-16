# Process String with Special Operations I - LeetCode 3612 Solution

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

LeetCode 3612, **Process String with Special Operations I**, is a simulation and string manipulation problem.

We are given a string `s` that contains lowercase English letters and three special characters:

* `*`
* `#`
* `%`

The goal is to process the string from left to right and build a new result string according to specific rules.

Rules:

* Lowercase letters are appended to the result.
* `*` removes the last character from the result if one exists.
* `#` duplicates the current result and appends it to itself.
* `%` reverses the current result.

After processing every character, we return the final generated string.

This problem is a good example of a **string simulation algorithm**, where the instructions must be executed exactly in the order they appear.

## Constraints

| Constraint            | Value                                    |
| --------------------- | ---------------------------------------- |
| `1 <= s.length <= 20` | Small input size                         |
| Characters            | Lowercase English letters, `*`, `#`, `%` |

## Intuition

My first observation was that every character directly tells us what operation to perform on the current result string.

There is no need for dynamic programming, greedy logic, or advanced data structures.

The problem statement already describes the exact behavior we need to follow. Whenever I see a problem where operations must be applied one by one in order, my first instinct is usually to simulate the process exactly as described.

Since the input size is very small, a straightforward simulation is both simple and efficient.

## Approach

I maintain a string called `result`.

Then I scan the input string from left to right.

For each character:

1. If it is a lowercase letter, append it to `result`.
2. If it is `*`, remove the last character if the result is not empty.
3. If it is `#`, duplicate the entire current string.
4. If it is `%`, reverse the current string.

After all characters have been processed, return the final value of `result`.

This directly matches the rules given in the problem statement.

## Data Structures Used

### String

The only data structure needed is a string.

Why it works:

* Supports appending characters.
* Supports deleting the last character.
* Supports duplication.
* Supports reversing.

Because the constraints are very small, a simple string is sufficient.

## Operations & Behavior Summary

The algorithm behaves like this:

1. Start with an empty result string.
2. Read characters one by one.
3. Apply the corresponding operation.
4. Keep updating the current result.
5. Continue until the end of the input string.
6. Return the final result.

Pseudo workflow:

* Letter → Add
* `*` → Remove Last Character
* `#` → Duplicate Current String
* `%` → Reverse Current String

## Complexity

| Metric           | Complexity | Explanation                                                                                                                   |
| ---------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | O(n × m)   | `n` is the input length and `m` is the current result length. Duplication and reversal may process the entire current string. |
| Space Complexity | O(m)       | Only the generated result string is stored.                                                                                   |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string processStr(string s) {
        // Stores the current result being built
        string result;

        for (char c : s) {
            // Lowercase letter -> append to result
            if (c >= 'a' && c <= 'z') {
                result.push_back(c);
            }
            // Remove last character if it exists
            else if (c == '*') {
                if (!result.empty()) {
                    result.pop_back();
                }
            }
            // Duplicate current result
            else if (c == '#') {
                result += result;
            }
            // Reverse current result
            else if (c == '%') {
                reverse(result.begin(), result.end());
            }
        }

        return result;
    }
};
```

### Java

```java
class Solution {
    public String processStr(String s) {
        // Stores the current result being built
        StringBuilder result = new StringBuilder();

        for (char c : s.toCharArray()) {
            // Lowercase letter -> append to result
            if (c >= 'a' && c <= 'z') {
                result.append(c);
            }
            // Remove last character if it exists
            else if (c == '*') {
                if (result.length() > 0) {
                    result.deleteCharAt(result.length() - 1);
                }
            }
            // Duplicate current result
            else if (c == '#') {
                result.append(result.toString());
            }
            // Reverse current result
            else if (c == '%') {
                result.reverse();
            }
        }

        return result.toString();
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {string}
 */
var processStr = function(s) {
    // Stores the current result being built
    let result = "";

    for (const c of s) {
        // Lowercase letter -> append to result
        if (c >= 'a' && c <= 'z') {
            result += c;
        }
        // Remove last character if it exists
        else if (c === '*') {
            if (result.length > 0) {
                result = result.slice(0, -1);
            }
        }
        // Duplicate current result
        else if (c === '#') {
            result += result;
        }
        // Reverse current result
        else if (c === '%') {
            result = result.split("").reverse().join("");
        }
    }

    return result;
};
```

### Python3

```python
class Solution:
    def processStr(self, s: str) -> str:
        # Stores the current result being built
        result = ""

        for c in s:
            # Lowercase letter -> append to result
            if 'a' <= c <= 'z':
                result += c

            # Remove last character if it exists
            elif c == '*':
                if result:
                    result = result[:-1]

            # Duplicate current result
            elif c == '#':
                result += result

            # Reverse current result
            elif c == '%':
                result = result[::-1]

        return result
```

### Go

```go
func processStr(s string) string {
    // Stores the current result being built
    result := ""

    for _, c := range s {
        // Lowercase letter -> append to result
        if c >= 'a' && c <= 'z' {
            result += string(c)

        // Remove last character if it exists
        } else if c == '*' {
            if len(result) > 0 {
                result = result[:len(result)-1]
            }

        // Duplicate current result
        } else if c == '#' {
            result += result

        // Reverse current result
        } else if c == '%' {
            chars := []rune(result)

            // Two-pointer reversal
            for l, r := 0, len(chars)-1; l < r; l, r = l+1, r-1 {
                chars[l], chars[r] = chars[r], chars[l]
            }

            result = string(chars)
        }
    }

    return result
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across all five languages.

### Step 1: Create the Result String

We begin with an empty string.

This string stores everything generated so far.

Example:

```text
result = ""
```

### Step 2: Process Characters One by One

We iterate through the input string.

The order matters because each operation changes the current state of the result.

### Step 3: Handle Lowercase Letters

Whenever a normal letter appears:

```text
a
b
c
```

it gets appended to the result.

Example:

```text
result = "ab"
character = 'c'

result = "abc"
```

### Step 4: Handle `*`

This operation removes the last character.

Example:

```text
result = "abcd"
```

After processing `*`:

```text
result = "abc"
```

If the result is already empty, nothing happens.

### Step 5: Handle `#`

This operation duplicates the current string.

Example:

```text
result = "abc"
```

After processing `#`:

```text
result = "abcabc"
```

The entire current string gets appended to itself.

### Step 6: Handle `%`

This operation reverses the current string.

Example:

```text
result = "abcde"
```

After processing `%`:

```text
result = "edcba"
```

### Step 7: Return Final Result

Once all characters have been processed, the resulting string is returned.

This final string represents the answer.

## Examples

### Example 1

Input

```text
s = "a#b%*"
```

Output

```text
"ba"
```

Trace

```text
""      -> start
"a"     -> append a
"aa"    -> duplicate
"aab"   -> append b
"baa"   -> reverse
"ba"    -> remove last character
```

Final Answer

```text
"ba"
```

---

### Example 2

Input

```text
s = "z*#"
```

Output

```text
""
```

Trace

```text
""   -> start
"z"  -> append z
""   -> remove z
""   -> duplicate empty string
```

Final Answer

```text
""
```

---

### Example 3

Input

```text
s = "ab#%"
```

Output

```text
"baba"
```

Trace

```text
"a"
"ab"
"abab"
"baba"
```

Final Answer

```text
"baba"
```

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
```

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

### JavaScript

Run

```bash
node solution.js
```

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

### Go

Run

```bash
go run solution.go
```

Build

```bash
go build solution.go
```

## Notes & Optimizations

* This is primarily a simulation problem.
* The simplest solution is also the most practical one.
* Since the maximum input length is only 20, performance is never an issue.
* A direct implementation is easier to understand and less error-prone.
* Always check whether the string is empty before processing `*`.
* Duplication and reversal operate on the entire current string.
* The order of operations must be preserved exactly as given.

Possible alternative approaches exist using arrays or character buffers, but they provide no meaningful advantage for the given constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
