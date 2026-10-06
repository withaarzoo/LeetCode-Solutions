# Minimum Add to Make Parentheses Valid | LeetCode 921 Solution

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

You are given a string that contains only opening and closing parentheses. The string is considered valid only if every opening parenthesis has a matching closing one in the correct order.

Your task is to find the smallest number of parentheses you need to insert so that the whole string becomes valid. You can insert an opening or closing parenthesis at any position.

The function should take the parentheses string as input and return a single integer — the minimum number of additions required.

This is a classic parentheses matching problem that appears frequently in coding interviews and online judges such as LeetCode.

## Constraints

- 1 <= s.length <= 1000
- s[i] is either '(' or ')'

## Intuition

When I first looked at the problem I realized I do not need to build a new string. I only need to count how many parentheses are missing.

Every time I see a closing parenthesis that has no unmatched opening parenthesis waiting for it, I know I must insert one opening parenthesis. At the end of the string, any leftover opening parentheses will each need a closing parenthesis. Adding these two counts together gives the minimum insertions.

This observation leads to a simple linear scan without using extra space for a stack.

## Approach

I keep two integer counters while walking through the string once.

One counter tracks how many opening parentheses are still unmatched.  
The other counter records every time I already know an insertion is required.

- When I meet '(', I increase the unmatched count.
- When I meet ')', I check the unmatched count.  
  - If it is greater than zero, I can pair the closing parenthesis and decrease the count.  
  - If it is zero, there is nothing to pair with, so I record one insertion.

After the loop finishes, any remaining unmatched openings also need insertions. I add that remaining value to the insertion count and return the total.

This greedy one-pass method guarantees the minimum number of additions because each mismatch is counted exactly once and no extra parentheses are ever added.

## Data Structures Used

Only two integer variables are used:

- An integer to store the current number of unmatched opening parentheses.
- An integer to store the number of insertions already required.

No arrays, stacks, or other dynamic structures are needed. This keeps the solution extremely light on memory.

## Operations & Behavior Summary

1. Initialize both counters to zero.
2. For every character in the string:
   - If the character is '(', increase the unmatched-open counter.
   - If the character is ')':
     - If unmatched-open is positive, decrease it (a match is found).
     - Otherwise increase the insertion counter.
3. After processing every character, add the remaining unmatched-open value to the insertion counter.
4. Return the final insertion count.

The algorithm never looks ahead or rewrites the string; it only counts mismatches as they appear.

## Complexity

| Type              | Value | Explanation |
|-------------------|-------|-------------|
| Time Complexity   | O(n)  | n is the length of the input string. Each character is examined exactly once. |
| Space Complexity  | O(1)  | Only two integer variables are used regardless of the size of the input. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minAddToMakeValid(string s) {
        int open = 0, add = 0;
        for (char c : s) {
            if (c == '(') {
                open++;
            } else {
                if (open > 0) {
                    open--;
                } else {
                    add++;
                }
            }
        }
        return add + open;
    }
};
```

### Java
```java
class Solution {
    public int minAddToMakeValid(String s) {
        int open = 0, add = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                open++;
            } else {
                if (open > 0) {
                    open--;
                } else {
                    add++;
                }
            }
        }
        return add + open;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {number}
 */
var minAddToMakeValid = function(s) {
    let open = 0, add = 0;
    for (let c of s) {
        if (c === '(') {
            open++;
        } else {
            if (open > 0) {
                open--;
            } else {
                add++;
            }
        }
    }
    return add + open;
};
```

### TypeScript
```typescript
function minAddToMakeValid(s: string): number {
    let open = 0, add = 0;
    for (const c of s) {
        if (c === '(') {
            open++;
        } else {
            if (open > 0) {
                open--;
            } else {
                add++;
            }
        }
    }
    return add + open;
};
```

### Python3
```python
class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        open = 0
        add = 0
        for c in s:
            if c == '(':
                open += 1
            else:
                if open > 0:
                    open -= 1
                else:
                    add += 1
        return add + open
```

### Go
```go
func minAddToMakeValid(s string) int {
    open, add := 0, 0
    for _, c := range s {
        if c == '(' {
            open++
        } else {
            if open > 0 {
                open--
            } else {
                add++
            }
        }
    }
    return add + open
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language, so the reasoning below applies to all six implementations.

I begin by declaring two variables that start at zero. One remembers how many opening parentheses are still waiting for a partner. The other remembers every time a closing parenthesis appears with no partner available.

I then examine each character from left to right.

When the character is an opening parenthesis I simply increase the waiting count. There is nothing else to do at that moment.

When the character is a closing parenthesis I look at the waiting count.  
If the count is greater than zero, a matching opening parenthesis exists, so I decrease the count by one.  
If the count is already zero, this closing parenthesis cannot be matched with anything that has appeared so far. I therefore increase the insertion counter by one.

After the entire string has been processed, the waiting count may still hold some value. Each of those leftover opening parentheses needs a closing parenthesis inserted later. I add that remaining value to the insertion counter.

The sum is the smallest number of parentheses that must be added to make the string valid.  

Edge cases are handled automatically:  
- A string that starts with many closing parentheses increases the insertion counter for each of them.  
- A string that ends with many opening parentheses adds those leftover openings at the end.  
- An already valid string keeps both counters at zero and returns zero.

Because the same two counters are used in every language, the only differences between the solutions are syntax (variable declaration, loop style, and return statement).

## Examples

**Example 1**  
Input: s = "())"  
Output: 1  

Trace:  
- '(' → unmatched becomes 1  
- ')' → unmatched becomes 0  
- ')' → unmatched is 0, so insertion count becomes 1  
Final answer = 1 + 0 = 1

**Example 2**  
Input: s = "((("  
Output: 3  

Trace:  
- '(' → unmatched becomes 1  
- '(' → unmatched becomes 2  
- '(' → unmatched becomes 3  
Final answer = 0 + 3 = 3

**Example 3**  
Input: s = "()()"  
Output: 0  

Trace:  
- '(' → unmatched becomes 1  
- ')' → unmatched becomes 0  
- '(' → unmatched becomes 1  
- ')' → unmatched becomes 0  
Final answer = 0 + 0 = 0

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `solution.cpp`.  
2. Compile with `g++ solution.cpp -o solution`.  
3. Run with `./solution` (you may need to add a main function that reads input and calls the method).

**Java**  
1. Save the code in a file named `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run with `java Solution` (add a main method for testing).

**JavaScript**  
1. Save the code in a file named `solution.js`.  
2. Run with `node solution.js` (add console.log calls for testing).

**TypeScript**  
1. Save the code in a file named `solution.ts`.  
2. Compile with `tsc solution.ts`.  
3. Run the generated JavaScript with `node solution.js`.

**Python3**  
1. Save the code in a file named `solution.py`.  
2. Run with `python3 solution.py` (add a main block for testing).

**Go**  
1. Save the code in a file named `solution.go`.  
2. Run with `go run solution.go` (add a main function for testing).

In every language you can hard-code a few test strings inside the main/driver code to verify the results quickly.

## Notes & Optimizations

The two-counter approach is already optimal for both time and space.  

A stack-based solution would also work and is easier to understand for beginners, but it uses O(n) extra memory in the worst case. The counter method avoids that cost completely.

Because the string length is at most 1000, even a slower approach would pass, yet the linear scan is the cleanest and most efficient way to solve the problem.

One common mistake is forgetting to add the remaining unmatched openings at the end. Always remember that both missing openings and missing closings contribute to the final answer.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)