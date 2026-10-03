# Longest Valid Parentheses - LeetCode 32 Solution

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

You are given a string that contains only two characters: opening parentheses `(` and closing parentheses `)`. Your task is to find the length of the longest valid (well-formed) parentheses substring inside that string.

A valid parentheses substring means every opening bracket has a matching closing bracket in the correct order. You do not need to return the substring itself — only its length.

This is a classic string and stack problem that appears frequently in coding interviews and competitive programming contests. It tests how well you can track matching pairs without scanning the same characters multiple times.

## Constraints

- The length of the string s is between 0 and 3 * 10^4 (inclusive).
- Every character in s is either `(` or `)`.

## Intuition

When I first looked at this problem, I noticed that every valid parentheses sequence must have an equal number of opening and closing brackets. The length of any valid stretch is always even.

My initial instinct was to use a stack to store the positions of unmatched opening brackets. That works, but it uses extra memory. Then I realized I could just keep two simple counters — one for left brackets and one for right brackets — and walk through the string twice. 

Walking only from left to right misses some cases where the string starts with extra closing brackets. Walking from both directions solves that issue completely and keeps the memory usage constant.

## Approach

I use two integer counters called left and right, plus a variable to store the maximum length found so far.

In the first pass I scan the string from left to right.  
- Every time I see `(`, I increase the left counter.  
- Every time I see `)`, I increase the right counter.  
- When the two counters become equal, I know I just finished a valid stretch, so I update the maximum length with twice the value of the right counter.  
- If the right counter becomes larger than the left counter, the current stretch can never become valid, so I reset both counters to zero.

In the second pass I scan the string from right to left and do almost the same thing, except the reset condition is swapped. This second pass catches the valid stretches that the first pass missed.

At the end I simply return the largest length I found.

This approach solves the longest valid parentheses problem in linear time and constant space.

## Data Structures Used

No complex data structures are required.  
I only use a few integer variables for the two counters and the answer.  

This keeps the solution extremely light on memory and makes it easy to implement in any language.

## Operations & Behavior Summary

1. Initialize left = 0, right = 0, maxLen = 0.  
2. Walk from the start of the string to the end.  
   - Count opening and closing brackets.  
   - Update maxLen whenever the counts become equal.  
   - Reset both counts when there are more closing brackets than opening ones.  
3. Reset the counters to zero.  
4. Walk from the end of the string back to the start.  
   - Count the brackets again.  
   - Update maxLen whenever the counts become equal.  
   - Reset both counts when there are more opening brackets than closing ones.  
5. Return the final maxLen value.

## Complexity

| Type              | Value | Explanation                                                                 |
|-------------------|-------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n)  | We scan the entire string twice. Each character is processed in constant time. n is the length of the input string. |
| Space Complexity  | O(1)  | Only a fixed number of integer variables are used. No arrays or stacks whose size depends on n are allocated. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int longestValidParentheses(string s) {
        int left = 0, right = 0, maxLen = 0;
        for (char c : s) {
            if (c == '(') left++;
            else right++;
            if (left == right) maxLen = max(maxLen, 2 * right);
            else if (right > left) left = right = 0;
        }
        left = right = 0;
        for (int i = s.size() - 1; i >= 0; i--) {
            if (s[i] == '(') left++;
            else right++;
            if (left == right) maxLen = max(maxLen, 2 * left);
            else if (left > right) left = right = 0;
        }
        return maxLen;
    }
};
```

### Java
```java
class Solution {
    public int longestValidParentheses(String s) {
        int left = 0, right = 0, maxLen = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') left++;
            else right++;
            if (left == right) maxLen = Math.max(maxLen, 2 * right);
            else if (right > left) left = right = 0;
        }
        left = right = 0;
        for (int i = s.length() - 1; i >= 0; i--) {
            if (s.charAt(i) == '(') left++;
            else right++;
            if (left == right) maxLen = Math.max(maxLen, 2 * left);
            else if (left > right) left = right = 0;
        }
        return maxLen;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {number}
 */
var longestValidParentheses = function(s) {
    let left = 0, right = 0, maxLen = 0;
    for (let c of s) {
        if (c === '(') left++;
        else right++;
        if (left === right) maxLen = Math.max(maxLen, 2 * right);
        else if (right > left) left = right = 0;
    }
    left = right = 0;
    for (let i = s.length - 1; i >= 0; i--) {
        if (s[i] === '(') left++;
        else right++;
        if (left === right) maxLen = Math.max(maxLen, 2 * left);
        else if (left > right) left = right = 0;
    }
    return maxLen;
};
```

### TypeScript
```typescript
function longestValidParentheses(s: string): number {
    let left = 0, right = 0, maxLen = 0;
    for (let c of s) {
        if (c === '(') left++;
        else right++;
        if (left === right) maxLen = Math.max(maxLen, 2 * right);
        else if (right > left) left = right = 0;
    }
    left = right = 0;
    for (let i = s.length - 1; i >= 0; i--) {
        if (s[i] === '(') left++;
        else right++;
        if (left === right) maxLen = Math.max(maxLen, 2 * left);
        else if (left > right) left = right = 0;
    }
    return maxLen;
};
```

### Python3
```python
class Solution:
    def longestValidParentheses(self, s: str) -> int:
        left = right = maxLen = 0
        for c in s:
            if c == '(':
                left += 1
            else:
                right += 1
            if left == right:
                maxLen = max(maxLen, 2 * right)
            elif right > left:
                left = right = 0
        left = right = 0
        for c in reversed(s):
            if c == '(':
                left += 1
            else:
                right += 1
            if left == right:
                maxLen = max(maxLen, 2 * left)
            elif left > right:
                left = right = 0
        return maxLen
```

### Go
```go
func longestValidParentheses(s string) int {
    left, right, maxLen := 0, 0, 0
    for _, c := range s {
        if c == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            if 2*right > maxLen {
                maxLen = 2 * right
            }
        } else if right > left {
            left, right = 0, 0
        }
    }
    left, right = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            if 2*left > maxLen {
                maxLen = 2 * left
            }
        } else if left > right {
            left, right = 0, 0
        }
    }
    return maxLen
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages. Only the syntax changes.

I start by creating three integer variables: left, right, and maxLen. All of them begin at zero.  

In the first loop I examine every character from left to right.  
- If the character is an opening parenthesis I add one to left.  
- If it is a closing parenthesis I add one to right.  
- Right after updating the counters I check two conditions.  
  - When left equals right I have a complete valid stretch. Its length is 2 * right, so I keep the larger of that value and the current maxLen.  
  - When right becomes greater than left the stretch is already broken. I reset both counters so the next characters start a fresh count.  

After the first loop finishes I set left and right back to zero.  

In the second loop I examine every character from right to left.  
The counting rules stay the same, but the reset condition is reversed: I reset when left becomes larger than right. This direction catches the cases where the string begins with extra closing brackets.  

When both loops are finished, maxLen holds the length of the longest valid parentheses substring. I simply return that value.

Edge cases are handled automatically:  
- An empty string returns 0.  
- A string with no matching pairs returns 0.  
- Strings that start or end with unmatched brackets are correctly truncated by the reset logic.

## Examples

**Example 1**  
Input: s = "(()"  
Output: 2  

Trace:  
Forward pass sees two opens and one close. Counters never become equal, so maxLen stays 0.  
Backward pass starts at the last character. It finds a matching pair of length 2 and updates maxLen to 2.

**Example 2**  
Input: s = ")()()"  
Output: 4  

Trace:  
Forward pass finds the stretch "()()" of length 4.  
Backward pass also finds the same stretch. The answer is 4.

**Example 3**  
Input: s = ""  
Output: 0  

Trace:  
Both passes see an empty string, so maxLen remains 0.

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `main.cpp`.  
2. Open a terminal and run:  
   `g++ -std=c++17 main.cpp -o main`  
3. Run the program:  
   `./main`

**Java**  
1. Save the code in a file named `Solution.java`.  
2. Compile:  
   `javac Solution.java`  
3. Run:  
   `java Solution`

**JavaScript**  
1. Save the code in a file named `solution.js`.  
2. Run with Node.js:  
   `node solution.js`

**TypeScript**  
1. Save the code in a file named `solution.ts`.  
2. Compile:  
   `tsc solution.ts`  
3. Run the generated JavaScript:  
   `node solution.js`

**Python3**  
1. Save the code in a file named `solution.py`.  
2. Run:  
   `python3 solution.py`

**Go**  
1. Save the code in a file named `main.go`.  
2. Run:  
   `go run main.go`

In each language you will need to add a small main function or driver code that creates a sample string and prints the result of the function.

## Notes & Optimizations

The two-pass counter method is optimal for both time and space.  

A common alternative is the stack-based solution that stores indices of unmatched brackets. It also runs in O(n) time but needs O(n) extra space in the worst case.  

Dynamic programming is another popular approach. It uses an array of size n and also runs in O(n) time and O(n) space.  

The counter method wins when memory is limited or when you want the simplest constant-space solution.  

Watch out for strings that consist of only opening or only closing brackets — the algorithm correctly returns 0 for those cases.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)