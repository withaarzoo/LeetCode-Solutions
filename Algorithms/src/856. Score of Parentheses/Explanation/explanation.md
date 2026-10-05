# 856. Score of Parentheses – LeetCode Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

The Score of Parentheses problem (LeetCode 856) asks you to calculate a special score for a balanced parentheses string.  

You are given a string `s` that contains only the characters `(` and `)` and is already balanced. Your task is to return an integer score based on three simple rules:

- The string `"()"` is worth 1.
- If you place two balanced strings next to each other (AB), their scores are added.
- If you wrap a balanced string A inside another pair of parentheses `(A)`, the score becomes twice the score of A.

The goal is to compute this score efficiently for any valid input string.

## Constraints

- The length of `s` is between 2 and 50 (inclusive).
- `s` contains only the characters `(` and `)`.
- `s` is guaranteed to be a balanced parentheses string.

## Intuition

When I first looked at the Score of Parentheses problem, I noticed that every time a closing parenthesis immediately follows an opening one, that tiny pair contributes a value that depends on how deeply nested it sits.  

The deeper a `"()"` pair is buried inside other parentheses, the higher its contribution becomes (it gets multiplied by 2 for each surrounding layer).  

This observation led me to track only the current nesting depth while scanning the string once. No need to build an explicit stack of scores; the depth itself tells me the power of two I should add whenever I find a leaf pair.

## Approach

I walk through the string from left to right while keeping two integer variables: a running score and a depth counter.

- When I meet an opening parenthesis I simply increase the depth.
- When I meet a closing parenthesis I first decrease the depth.
- If the character right before this closing parenthesis was an opening parenthesis, I have found a `"()"` pair at the current depth. I therefore add `2^depth` to the score.

Because the string is balanced, the depth never goes negative and ends at zero. After one linear pass the score variable already holds the final answer required by the problem rules.

## Data Structures Used

No extra data structures are required.  

Only a few integer variables are used:
- `score` – accumulates the final result.
- `depth` – tracks the current nesting level of parentheses.

This choice keeps both time and space optimal for the given constraints.

## Operations & Behavior Summary

1. Initialize `score = 0` and `depth = 0`.
2. For every character in the string:
   - If the character is `(`, increase `depth` by one.
   - If the character is `)`, decrease `depth` by one, then check the previous character.
     - When the previous character is also `(`, add `1 << depth` (i.e., `2^depth`) to `score`.
3. After the loop finishes, return `score`.

The algorithm never looks ahead more than one character and never stores intermediate results beyond the two integers.

## Complexity

| Type              | Complexity | Explanation                                      |
|-------------------|------------|--------------------------------------------------|
| Time Complexity   | O(n)       | Each of the n characters is examined exactly once. |
| Space Complexity  | O(1)       | Only a constant number of integer variables are used. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int scoreOfParentheses(string s) {
        int score = 0, depth = 0;
        for (int i = 0; i < s.size(); ++i) {
            if (s[i] == '(') {
                ++depth;
            } else {
                --depth;
                if (s[i - 1] == '(') {
                    score += 1 << depth;
                }
            }
        }
        return score;
    }
};
```

### Java
```java
class Solution {
    public int scoreOfParentheses(String s) {
        int score = 0, depth = 0;
        for (int i = 0; i < s.length(); ++i) {
            if (s.charAt(i) == '(') {
                ++depth;
            } else {
                --depth;
                if (s.charAt(i - 1) == '(') {
                    score += 1 << depth;
                }
            }
        }
        return score;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {number}
 */
var scoreOfParentheses = function(s) {
    let score = 0, depth = 0;
    for (let i = 0; i < s.length; ++i) {
        if (s[i] === '(') {
            ++depth;
        } else {
            --depth;
            if (s[i - 1] === '(') {
                score += 1 << depth;
            }
        }
    }
    return score;
};
```

### TypeScript
```typescript
function scoreOfParentheses(s: string): number {
    let score = 0, depth = 0;
    for (let i = 0; i < s.length; ++i) {
        if (s[i] === '(') {
            ++depth;
        } else {
            --depth;
            if (s[i - 1] === '(') {
                score += 1 << depth;
            }
        }
    }
    return score;
};
```

### Python3
```python
class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        score = depth = 0
        for i in range(len(s)):
            if s[i] == '(':
                depth += 1
            else:
                depth -= 1
                if s[i - 1] == '(':
                    score += 1 << depth
        return score
```

### Go
```go
func scoreOfParentheses(s string) int {
    score, depth := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            depth++
        } else {
            depth--
            if s[i-1] == '(' {
                score += 1 << depth
            }
        }
    }
    return score
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax differs.

I start by declaring two integer variables that will live for the whole function.  
`score` begins at zero so that every contribution can be added to it.  
`depth` also starts at zero because we have not yet entered any parentheses.

The main loop runs from the first character to the last.  

On an opening parenthesis the only action needed is to push the nesting level one step deeper. That single increment is enough to remember how many unmatched opens are currently active.

On a closing parenthesis two things happen in order. First the depth is reduced so that it now represents the exact level of the pair that just closed. Then I look at the character that sits immediately before the current index.  

If that previous character is an opening parenthesis, the two consecutive characters form a primitive `"()"` pair. At this moment the depth value is precisely the number of outer layers surrounding this pair, so I add the corresponding power of two to the score. Using a bit-shift (`1 << depth`) is a fast and portable way to compute that power.

If the previous character is not an opening parenthesis, the current closing parenthesis is matching a larger group that already contains nested material. The multiplication by two for those outer groups has already been accounted for by the deeper leaf pairs, so nothing extra is added.

Because the input is guaranteed to be balanced, the depth counter never becomes negative and returns to zero at the end of the string. Consequently the final value stored in `score` is exactly the result required by the three scoring rules of the problem.

Edge cases such as the shortest string `"()"` or a fully nested string like `"((()))"` are handled automatically by the same depth-tracking logic; no special branches are required.

## Examples

**Example 1**  
Input: `s = "()"`  
Output: `1`  

Trace:  
- index 0: `(`, depth becomes 1  
- index 1: `)`, depth becomes 0, previous character is `(`, therefore add `1 << 0` = 1  
Final score = 1

**Example 2**  
Input: `s = "(())"`  
Output: `2`  

Trace:  
- index 0: `(`, depth = 1  
- index 1: `(`, depth = 2  
- index 2: `)`, depth = 1, previous is `(`, add `1 << 1` = 2  
- index 3: `)`, depth = 0, previous is `)`, no addition  
Final score = 2

**Example 3**  
Input: `s = "()()"`  
Output: `2`  

Trace:  
- index 0: `(`, depth = 1  
- index 1: `)`, depth = 0, previous is `(`, add 1  
- index 2: `(`, depth = 1  
- index 3: `)`, depth = 0, previous is `(`, add 1  
Final score = 2

## How to Use / Run Locally

**C++**  
Save the code in a file named `score.cpp`. Compile with  
`g++ -std=c++17 score.cpp -o score`  
Run with  
`./score`

**Java**  
Save the code in a file named `Solution.java`. Compile with  
`javac Solution.java`  
Run with  
`java Solution`

**JavaScript**  
Save the code in a file named `score.js`. Run with  
`node score.js`

**TypeScript**  
Save the code in a file named `score.ts`. First install the TypeScript compiler if needed (`npm install -g typescript`), then compile with  
`tsc score.ts`  
and run the generated JavaScript with  
`node score.js`

**Python3**  
Save the code in a file named `score.py`. Run with  
`python3 score.py`

**Go**  
Save the code in a file named `score.go`. Run with  
`go run score.go`

In each case you can add a small `main` function or test harness that calls `scoreOfParentheses` with sample strings and prints the results.

## Notes & Optimizations

The depth-tracking method uses only constant extra memory, which is optimal for this problem.  

An alternative stack-based approach also works in linear time: push a zero for every opening parenthesis and, on a closing parenthesis, pop the top value, double it (or turn it into 1 if it was zero), then add it to the new top. That version uses O(n) space and is equally correct, but the depth method is lighter.

Because the maximum length is only 50, both approaches easily pass the time limits. The constant-space solution is preferred when memory is a concern or when you want the simplest possible implementation.

The algorithm relies on the guarantee that the string is balanced; if that guarantee were removed, additional validation would be required.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)