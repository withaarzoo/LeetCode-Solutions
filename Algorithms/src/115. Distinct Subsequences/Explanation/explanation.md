# 115. Distinct Subsequences

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

The problem gives me two strings, `s` and `t`.

I need to find how many different subsequences of `s` are exactly equal to `t`.

A subsequence keeps the original order of characters, but I can skip any number of characters.

For example:

```text
s = "rabbbit"
t = "rabbit"
```

There are 3 different ways to choose characters from `s` to form `t`, so the answer is `3`.

The important point is that I am counting different ways of selecting characters, not just different resulting strings.

This problem is a classic Dynamic Programming problem and is commonly used to understand subsequence counting, 1D DP optimization, and string DP.

## Constraints

| Constraint                        | Description                                                                     |
| --------------------------------- | ------------------------------------------------------------------------------- |
| `1 <= s.length, t.length <= 1000` | Both strings can contain up to 1000 characters                                  |
| Characters                        | `s` and `t` contain only English letters                                        |
| Answer                            | The test cases are generated so that the answer fits in a 32-bit signed integer |

## Intuition

I first thought about what happens when I process the characters of `s` one by one.

Suppose the current character of `s` matches the character I need in `t`. I have two choices:

1. Skip the current character.
2. Use the current character to build `t`.

That gives me this simple idea:

```text
                    current character
                           |
                  s[i] == t[j]
                     /          \
                   skip         use
                    |             |
              keep old count   add previous count
```

If the characters do not match, I cannot use the current character for that position in `t`.

I then realized that I do not need to remember every subsequence. I only need to remember how many ways I can form every prefix of `t`.

For example:

```text
t = "bag"

Prefix:
""     "b"     "ba"     "bag"
```

I can store the number of ways for these prefixes in a DP array.

The empty string is always possible by choosing nothing:

```text
dp[0] = 1
```

Then, whenever the current character matches, I add the number of ways that could form the previous prefix.

This leads to the transition:

```text
dp[j] = dp[j] + dp[j - 1]
```

I process `j` from right to left so that `dp[j - 1]` still contains the value from before processing the current character.

## Approach

I use dynamic programming with a one-dimensional array.

First, I create:

```text
dp[0 ... m]
```

where `m` is the length of `t`.

The meaning of `dp[j]` is:

```text
dp[j] = number of ways to form the first j characters of t
        using the characters of s processed so far
```

Initially:

```text
target prefix:   ""    b    ba    bag
index:            0    1     2      3
dp:               1    0     0      0
```

The value `dp[0]` is `1` because there is exactly one way to form an empty string: select nothing.

Now I process every character of `s`.

If:

```text
s[i] == t[j - 1]
```

then I can use `s[i]` to complete a subsequence that already forms the first `j - 1` characters of `t`.

So I add:

```text
dp[j - 1]
```

to:

```text
dp[j]
```

The transition is:

```text
dp[j] += dp[j - 1]
```

For example, while processing `"babgbag"` for target `"bag"`:

```text
target:       ""    b    ba    bag
index:         0    1     2      3

initial:       1    0     0      0
after 'b':     1    1     0      0
after 'a':     1    1     1      0
after 'b':     1    2     1      0
...
final:         1    3     4      5
```

The final `dp[3]` is `5`, which means there are 5 ways to form `"bag"`.

The important optimization is updating the DP array from right to left:

```text
m -> m-1 -> m-2 -> ... -> 1
```

If I went from left to right, `dp[j - 1]` could already contain information from the current character. That could make one character of `s` get used more than once.

Going right to left prevents that.

## Data Structures Used

### 1. One-dimensional DP array

I use an array of size `m + 1`.

```text
dp[j]
```

stores the number of ways to form the first `j` characters of `t`.

I use a 1D array instead of a 2D DP table because the current state only needs the previous prefix count.

This reduces the space complexity from `O(n * m)` to `O(m)`.

### 2. Strings

I directly access characters from `s` and `t`.

No extra string or subsequence list is needed because I only need counts, not the actual subsequences.

## Operations & Behavior Summary

The algorithm works like this:

1. Find the length of `t`.
2. Create a DP array of size `m + 1`.
3. Set `dp[0] = 1`.
4. Process every character of `s`.
5. For each character, scan the target `t` from right to left.
6. If the current characters match, add `dp[j - 1]` to `dp[j]`.
7. If they do not match, leave `dp[j]` unchanged.
8. After processing all characters, return `dp[m]`.

In simple pseudocode:

```text
create dp with m + 1 positions
dp[0] = 1

for every character c in s:
    for j from m down to 1:
        if c == t[j - 1]:
            dp[j] += dp[j - 1]

return dp[m]
```

## Complexity

| Type             | Complexity | Explanation                                              |
| ---------------- | ---------: | -------------------------------------------------------- |
| Time Complexity  | `O(n * m)` | I process every character of `s` and scan the target `t` |
| Space Complexity |     `O(m)` | I only store one DP array of size `m + 1`                |

Here:

* `n` = length of `s`
* `m` = length of `t`

The 1D DP optimization is useful because a normal 2D DP solution would require `O(n * m)` extra space.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numDistinct(string s, string t) {
        int m = t.size();

        // dp[j] stores the number of ways to form the first j characters of t.
        // dp[0] is 1 because there is exactly one way to form an empty string.
        vector<unsigned long long> dp(m + 1, 0);
        dp[0] = 1;

        // Process every character of s one by one.
        for (char c : s) {
            // Go from right to left so dp[j - 1] still represents
            // the previous state and the current character is not reused.
            for (int j = m; j >= 1; --j) {
                // If the current character can provide t[j - 1],
                // add all ways of forming the previous prefix.
                if (c == t[j - 1]) {
                    dp[j] += dp[j - 1];
                }
            }
        }

        // dp[m] contains the number of ways to form all of t.
        return (int)dp[m];
    }
};
```

### Java

```java
class Solution {
    public int numDistinct(String s, String t) {
        int m = t.length();

        // dp[j] stores the number of ways to form the first j characters of t.
        // There is exactly one way to form an empty string.
        long[] dp = new long[m + 1];
        dp[0] = 1;

        // Process every character of s.
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);

            // Traverse backwards so dp[j - 1] is still from the previous state.
            // This prevents using the same character of s more than once.
            for (int j = m; j >= 1; j--) {
                // If the characters match, we can use the current character.
                if (c == t.charAt(j - 1)) {
                    dp[j] += dp[j - 1];
                }
            }
        }

        // The problem guarantees that the answer fits in a 32-bit signed integer.
        return (int) dp[m];
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {string} t
 * @return {number}
 */
var numDistinct = function(s, t) {
    const m = t.length;

    // dp[j] stores the number of ways to form the first j characters of t.
    // dp[0] is 1 because choosing nothing forms the empty string.
    const dp = new Array(m + 1).fill(0);
    dp[0] = 1;

    // Process each character of s.
    for (let i = 0; i < s.length; i++) {
        const c = s[i];

        // Traverse backwards so dp[j - 1] is not updated by the current character.
        // This makes sure one character of s is used only once in a transition.
        for (let j = m; j >= 1; j--) {
            // If the characters match, add the ways of forming the previous prefix.
            if (c === t[j - 1]) {
                dp[j] += dp[j - 1];
            }
        }
    }

    // dp[m] contains the number of distinct subsequences equal to t.
    return dp[m];
};
```

### Python3

```python
class Solution:
    def numDistinct(self, s: str, t: str) -> int:
        m = len(t)

        # dp[j] stores the number of ways to form the first j characters of t.
        # There is exactly one way to form an empty string.
        dp = [0] * (m + 1)
        dp[0] = 1

        # Process every character of s.
        for c in s:
            # Go backwards so dp[j - 1] still belongs to the previous state.
            # This prevents the current character from being used multiple times.
            for j in range(m, 0, -1):
                # If the characters match, use the current character to extend
                # every subsequence that already forms the previous prefix.
                if c == t[j - 1]:
                    dp[j] += dp[j - 1]

        # dp[m] contains the number of ways to form all of t.
        return dp[m]
```

### Go

```go
func numDistinct(s string, t string) int {
 m := len(t)

 // dp[j] stores the number of ways to form the first j characters of t.
 // There is exactly one way to form an empty string.
 dp := make([]uint64, m+1)
 dp[0] = 1

 // Process every character of s.
 for i := 0; i < len(s); i++ {
  c := s[i]

  // Traverse backwards so dp[j-1] still represents the previous state.
  // This prevents the current character from being reused.
  for j := m; j >= 1; j-- {
   // If the characters match, add the ways of forming the previous prefix.
   if c == t[j-1] {
    dp[j] += dp[j-1]
   }
  }
 }

 // The problem guarantees that the answer fits in a 32-bit signed integer.
 return int(dp[m])
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### C++

The C++ solution creates a one-dimensional DP array with `m + 1` positions.

The first position represents the empty target, so it starts with:

```text
dp[0] = 1
```

I then iterate through every character of `s`.

For each character, I iterate through `t` from right to left. The reverse direction is important because I need `dp[j - 1]` to represent the state before the current character was processed.

When `s[i]` and `t[j - 1]` are equal, I add:

```text
dp[j - 1]
```

to:

```text
dp[j]
```

This represents using the current character to extend every subsequence that already formed the previous prefix.

At the end, `dp[m]` contains the answer.

I use an integer-compatible numeric type for the intermediate DP values and return the result as an `int`, since the problem guarantees that the final answer fits in a 32-bit signed integer.

### Java

The Java solution follows the same DP logic.

I create a `long` array because intermediate additions can be larger than the final 32-bit result during the calculation.

The first element is initialized to `1` because there is one way to form an empty target.

I process `s` character by character and update the DP array from right to left.

When the current character matches the required target character, I add the previous prefix count.

The final value at index `m` is converted to `int` because the problem guarantees that the answer fits in a 32-bit signed integer.

### JavaScript

The JavaScript solution uses a normal array for the 1D DP table.

I initialize every position with `0` and set the empty-target state to `1`.

For every character in `s`, I scan `t` backwards.

JavaScript's `Number` type can safely represent the integer values required by the problem because the answer is guaranteed to fit in a 32-bit signed integer.

The important part remains the same:

```text
if current character == target character:
    dp[j] += dp[j - 1]
```

The final value is `dp[m]`.

### Python3

The Python solution uses a list as the DP array.

I initialize it with `m + 1` zeros and set:

```text
dp[0] = 1
```

Python integers can grow automatically, so I do not need to choose a special integer type for the DP values.

For every character in `s`, I iterate through the target positions backwards.

If the characters match, I add the number of ways to form the previous prefix.

The answer is stored in `dp[m]`.

### Go

The Go solution uses a one-dimensional slice for the DP state.

I initialize:

```text
dp[0] = 1
```

and process the characters of `s`.

The target positions are visited from right to left for the same reason as in the other implementations: it prevents the current character from being reused in the same iteration.

The final value is converted to `int` and returned.

The algorithm itself does not change between the five languages. Only the syntax and basic data types are different.

## Examples

### Example 1

Input:

```text
s = "rabbbit"
t = "rabbit"
```

Output:

```text
3
```

There are three different ways to select characters from `s` to form `t`.

The repeated `b` characters create multiple valid choices:

```text
r a b b b i t
    ^ ^ ^
    choose different b positions
```

The DP keeps track of these choices without generating every subsequence explicitly.

### Example 2

Input:

```text
s = "babgbag"
t = "bag"
```

Output:

```text
5
```

The DP states grow as characters are processed:

```text
target:       ""    b    ba    bag

initial:       1    0     0      0
after 'b':     1    1     0      0
after 'a':     1    1     1      0
after 'b':     1    2     1      0
...
final:         1    3     4      5
```

So there are `5` different ways to form `"bag"`.

### Example 3

Input:

```text
s = "abc"
t = "abc"
```

Output:

```text
1
```

There is only one possible subsequence that forms `"abc"`:

```text
a -> b -> c
```

Since every character must be selected, the answer is `1`.

## How to Use / Run Locally

### C++

Save the solution as:

```text
solution.cpp
```

Compile it with:

```text
g++ -std=c++17 solution.cpp -o solution
```

Then run:

```text
./solution
```

On Windows, run:

```text
solution.exe
```

The LeetCode version provides the `Solution` class and calls the method automatically. For standalone execution, I would add a small `main` function with test input.

### Java

Save the solution as:

```text
Solution.java
```

Compile it with:

```text
javac Solution.java
```

Then run:

```text
java Solution
```

For LeetCode, I only need to submit the `Solution` class and the required method.

### JavaScript

Save the solution as:

```text
solution.js
```

Run it with Node.js:

```text
node solution.js
```

On LeetCode, the platform calls the provided function automatically.

### Python3

Save the solution as:

```text
solution.py
```

Run:

```text
python3 solution.py
```

On LeetCode, I only need to provide the `Solution` class and its `numDistinct` method.

### Go

Save the solution as:

```text
solution.go
```

Run it with:

```text
go run solution.go
```

For a standalone Go program, a `main` function is also required. On LeetCode, the platform handles the function call.

## Notes & Optimizations

The biggest optimization in this solution is reducing the DP table from two dimensions to one.

A standard 2D DP approach can define:

```text
dp[i][j]
```

as the number of ways to form the first `j` characters of `t` using the first `i` characters of `s`.

That approach takes:

```text
Time:  O(n * m)
Space: O(n * m)
```

The time cannot be improved for this DP approach because I need to consider the relationship between characters of both strings.

However, I can reduce the space to:

```text
Space: O(m)
```

because each state only needs the previous prefix information.

The right-to-left update is essential for this optimization.

Some important edge cases are:

* If `t` is empty, the answer is `1`.
* If `s` is shorter than `t`, the answer is `0`.
* If `s` and `t` are the same, the answer is `1`.
* Repeated characters must be counted correctly.
* Characters that do not match are simply skipped.

I do not generate the actual subsequences because that would use much more time and memory. I only store their counts, which is exactly what the problem asks for.

## Author

Md Aarzoo Islam — [Code with Aarzoo on Instagram](https://www.instagram.com/codewithaarzoo.in/?utm_source=chatgpt.com)
