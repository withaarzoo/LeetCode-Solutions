# Maximum Nesting Depth of the Parentheses - LeetCode 1614 Solution

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

You are given a valid parentheses string `s`. The string can contain digits from 0 to 9 and the operators `+`, `-`, `*`, `/` along with opening and closing parentheses.

Your task is to find the maximum nesting depth of the parentheses in the string. Nesting depth means the largest number of parentheses that are open at the same time.

For example, in the expression `(1+(2*3)+((8)/4))+1`, the digit 8 sits inside three levels of parentheses, so the answer is 3.

The problem guarantees that the input is always a valid parentheses string (VPS), so you never have to worry about unbalanced brackets.

## Constraints

- 1 <= s.length <= 100
- s consists of digits 0-9 and the characters `+`, `-`, `*`, `/`, `(`, and `)`
- It is guaranteed that the parentheses expression `s` is a VPS

## Intuition

When I first looked at the problem, I realized that only the parentheses actually matter for calculating depth. Digits and operators can be ignored completely.

The natural way to track nesting is to keep a running count of how many opening parentheses I have seen that are still unmatched. Every time I meet an opening parenthesis the count goes up. Every time I meet a closing parenthesis the count goes down.

The highest value this count ever reaches during the scan is exactly the maximum nesting depth I am looking for.

This approach works because the string is already guaranteed to be valid, so the count will never go negative and will always end at zero.

## Approach

I walk through the string from left to right while maintaining two simple integer variables.

One variable keeps the current depth. The other remembers the maximum depth seen so far.

Whenever the current character is `(`, I increase the current depth by one and then compare it with the maximum. If the new depth is larger, I update the maximum.

Whenever the current character is `)`, I simply decrease the current depth by one.

All other characters are skipped.

At the end of the loop the maximum depth variable holds the answer.

This single-pass method is both simple and efficient.

## Data Structures Used

No advanced data structures are required.

I only use two integer variables:
- `depth` – tracks how many parentheses are currently open
- `maxDepth` – stores the highest value `depth` has reached

A stack could also solve the problem, but it would use extra memory that is unnecessary here. The two counters are enough.

## Operations & Behavior Summary

1. Initialize `depth = 0` and `maxDepth = 0`.
2. For every character in the string:
   - If the character is `(`, increase `depth` by 1 and update `maxDepth` if needed.
   - If the character is `)`, decrease `depth` by 1.
   - Otherwise do nothing.
3. After processing the whole string, return `maxDepth`.

This is essentially a linear scan that updates a running maximum on the fly.

## Complexity

| Metric            | Value | Explanation |
|-------------------|-------|-------------|
| Time Complexity   | O(n)  | We examine each of the n characters in the string exactly once. |
| Space Complexity  | O(1)  | Only a constant amount of extra memory is used for the two integer counters. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int maxDepth(string s) {
        int depth = 0, maxDepth = 0;
        for (char c : s) {
            if (c == '(') {
                depth++;
                if (depth > maxDepth) maxDepth = depth;
            } else if (c == ')') {
                depth--;
            }
        }
        return maxDepth;
    }
};
```

### Java
```java
class Solution {
    public int maxDepth(String s) {
        int depth = 0, maxDepth = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                depth++;
                if (depth > maxDepth) maxDepth = depth;
            } else if (c == ')') {
                depth--;
            }
        }
        return maxDepth;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {number}
 */
var maxDepth = function(s) {
    let depth = 0, maxDepth = 0;
    for (let c of s) {
        if (c === '(') {
            depth++;
            if (depth > maxDepth) maxDepth = depth;
        } else if (c === ')') {
            depth--;
        }
    }
    return maxDepth;
};
```

### TypeScript
```typescript
function maxDepth(s: string): number {
    let depth = 0, maxDepth = 0;
    for (let c of s) {
        if (c === '(') {
            depth++;
            if (depth > maxDepth) maxDepth = depth;
        } else if (c === ')') {
            depth--;
        }
    }
    return maxDepth;
};
```

### Python3
```python
class Solution:
    def maxDepth(self, s: str) -> int:
        depth = 0
        max_depth = 0
        for c in s:
            if c == '(':
                depth += 1
                if depth > max_depth:
                    max_depth = depth
            elif c == ')':
                depth -= 1
        return max_depth
```

### Go
```go
func maxDepth(s string) int {
    depth, maxDepth := 0, 0
    for _, c := range s {
        if c == '(' {
            depth++
            if depth > maxDepth {
                maxDepth = depth
            }
        } else if c == ')' {
            depth--
        }
    }
    return maxDepth
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages, so the explanation applies to every implementation.

I start by declaring two integer variables and setting both of them to zero.  
`depth` will tell me the current nesting level.  
`maxDepth` will remember the deepest level I have seen.

I then iterate over every character of the input string.

When I encounter an opening parenthesis, I first increase `depth`. Right after the increase I check whether this new value is greater than `maxDepth`. If it is, I copy the value into `maxDepth`. Updating the maximum immediately after the increment guarantees I never miss a peak.

When I encounter a closing parenthesis, I decrease `depth`. I do not touch `maxDepth` here because a closing parenthesis can only reduce the current depth, never increase it.

Any other character (a digit or an operator) is simply ignored. It does not affect the nesting count.

Because the problem guarantees a valid parentheses string, `depth` will never become negative and will finish at zero. This removes the need for any extra validation checks.

At the end of the loop I simply return `maxDepth`. That value is the maximum nesting depth of the parentheses.

The same sequence of operations is performed in C++, Java, JavaScript, TypeScript, Python 3, and Go. The only differences are the language-specific ways of writing a loop and comparing characters.

## Examples

**Example 1**  
Input: `s = "(1+(2*3)+((8)/4))+1"`  
Output: `3`  

Trace:  
- First `(` → depth becomes 1  
- Second `(` → depth becomes 2  
- Third `(` → depth becomes 3 (maxDepth is now 3)  
- Later closing parentheses bring the depth back down  
The deepest point was 3, which matches the position of the digit 8.

**Example 2**  
Input: `s = "(1)+((2))+(((3)))"`  
Output: `3`  

Trace:  
- The third group starts with three consecutive opening parentheses, pushing depth to 3.  
- maxDepth is updated to 3 and never goes higher.

**Example 3**  
Input: `s = "()(())((()()))"`  
Output: `3`  

Trace:  
- The last group of parentheses reaches a depth of 3 before closing.  
- All earlier groups stay at depth 1 or 2.

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `main.cpp`.  
2. Compile with `g++ -std=c++17 main.cpp -o main`.  
3. Run with `./main`.  
4. You can hard-code a test string inside the `main` function or read from standard input.

**Java**  
1. Save the code in a file named `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run with `java Solution`.  
4. Add a `main` method that creates an instance of `Solution` and calls `maxDepth` with a test string.

**JavaScript**  
1. Save the code in a file named `maxDepth.js`.  
2. Run with `node maxDepth.js`.  
3. Call the function with a test string and print the result using `console.log`.

**TypeScript**  
1. Save the code in a file named `maxDepth.ts`.  
2. Compile with `tsc maxDepth.ts`.  
3. Run the generated JavaScript file with `node maxDepth.js`.

**Python 3**  
1. Save the code in a file named `solution.py`.  
2. Run with `python3 solution.py`.  
3. Inside an `if __name__ == "__main__":` block, create a `Solution` object and print the result of `maxDepth` on a test string.

**Go**  
1. Save the code in a file named `main.go`.  
2. Run with `go run main.go`.  
3. Add a `main` function that calls `maxDepth` with a test string and prints the returned value.

## Notes & Optimizations

The two-counter approach already achieves the best possible time and space complexity for this problem.

A stack-based solution is also correct: push on `(`, pop on `)`, and keep track of the maximum stack size. It works, but it uses O(n) extra space in the worst case, which is unnecessary given the constraints.

Because the string length is at most 100, even a less optimal solution would pass, yet the constant-space linear scan is still the cleanest choice.

Edge cases worth remembering:
- A string that contains no parentheses at all returns 0.
- A string that is just a single pair `()` returns 1.
- Deeply nested expressions such as `((((1))))` correctly report depth 4.

The algorithm never needs to handle invalid input because the problem statement guarantees a valid parentheses string.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)