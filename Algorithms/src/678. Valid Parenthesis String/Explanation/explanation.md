# LeetCode 678. Valid Parenthesis String - Greedy Balance Solution

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

This is LeetCode problem 678 called Valid Parenthesis String. You are given a string that only contains three kinds of characters: opening parenthesis '(', closing parenthesis ')', and star '*'.  

Your job is to check whether the string can be made into a valid parentheses string by treating each star as either an opening parenthesis, a closing parenthesis, or an empty string.  

A string is valid when every opening parenthesis has a matching closing one, every closing parenthesis has a matching opening one, and the openings always appear before their matching closings.  

The function should return true if it is possible to make the string valid, otherwise false.

## Constraints

- 1 <= s.length <= 100
- s[i] is either '(', ')' or '*'

## Intuition

When I first saw Valid Parenthesis String I thought about the classic parentheses matching problem that uses a stack. The stars made that idea messy because each star has three possible meanings. Trying every combination would work for a length of 100 but felt heavy.  

Then I noticed that I only really care about the balance of unmatched opening parentheses at every position. Instead of picking one meaning for each star, I can keep track of the lowest possible balance and the highest possible balance I could have reached so far. If the highest balance ever goes negative, no choice of stars can save the string. If at the end the lowest balance is zero, there is at least one way to assign the stars that makes everything match.

## Approach

I walk through the string once while maintaining two counters called low and high.  

Low stores the smallest number of unmatched openings I could still have. High stores the largest number of unmatched openings I could still have.  

- When I see '(', both low and high increase by one.  
- When I see ')', both low and high decrease by one.  
- When I see '*', low decreases by one (treating the star as a close) and high increases by one (treating the star as an open).  

After every character I force low to stay at least zero, because I can always choose to treat extra stars as empty instead of closing parentheses. If high ever becomes negative I know the string is impossible and I return false right away.  

At the end I simply check whether low is exactly zero. That tells me there exists at least one valid way to interpret the stars.

## Data Structures Used

Only two integer variables (low and high) are used. No arrays, stacks, or other data structures are required. These two counters are enough to represent the entire range of possible balances, which keeps the solution both simple and memory-efficient.

## Operations & Behavior Summary

1. Initialize low = 0 and high = 0.  
2. For each character in the string:  
   - If it is '(', increase both low and high.  
   - If it is ')', decrease both low and high.  
   - If it is '*', decrease low and increase high.  
3. After the update, if high is negative, return false.  
4. If low has become negative, reset it to zero.  
5. After processing the whole string, return true only if low equals zero.

## Complexity

| Complexity       | Value | Explanation                                                                 |
|------------------|-------|-----------------------------------------------------------------------------|
| Time Complexity  | O(n)  | We examine each of the n characters in the string exactly once.             |
| Space Complexity | O(1)  | Only two integer variables are used regardless of the input size.           |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    bool checkValidString(string s) {
        int low = 0, high = 0;
        for (char c : s) {
            if (c == '(') {
                low++;
                high++;
            } else if (c == ')') {
                low--;
                high--;
            } else {
                low--;
                high++;
            }
            if (high < 0) return false;
            if (low < 0) low = 0;
        }
        return low == 0;
    }
};
```

### Java
```java
class Solution {
    public boolean checkValidString(String s) {
        int low = 0, high = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                low++;
                high++;
            } else if (c == ')') {
                low--;
                high--;
            } else {
                low--;
                high++;
            }
            if (high < 0) return false;
            if (low < 0) low = 0;
        }
        return low == 0;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {boolean}
 */
var checkValidString = function(s) {
    let low = 0, high = 0;
    for (let c of s) {
        if (c === '(') {
            low++;
            high++;
        } else if (c === ')') {
            low--;
            high--;
        } else {
            low--;
            high++;
        }
        if (high < 0) return false;
        if (low < 0) low = 0;
    }
    return low === 0;
};
```

### TypeScript
```typescript
function checkValidString(s: string): boolean {
    let low = 0, high = 0;
    for (let c of s) {
        if (c === '(') {
            low++;
            high++;
        } else if (c === ')') {
            low--;
            high--;
        } else {
            low--;
            high++;
        }
        if (high < 0) return false;
        if (low < 0) low = 0;
    }
    return low === 0;
};
```

### Python3
```python
class Solution:
    def checkValidString(self, s: str) -> bool:
        low = high = 0
        for c in s:
            if c == '(':
                low += 1
                high += 1
            elif c == ')':
                low -= 1
                high -= 1
            else:
                low -= 1
                high += 1
            if high < 0:
                return False
            if low < 0:
                low = 0
        return low == 0
```

### Go
```go
func checkValidString(s string) bool {
    low, high := 0, 0
    for _, c := range s {
        if c == '(' {
            low++
            high++
        } else if c == ')' {
            low--
            high--
        } else {
            low--
            high++
        }
        if high < 0 {
            return false
        }
        if low < 0 {
            low = 0
        }
    }
    return low == 0
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages, so the reasoning below applies to every version.

I begin by declaring two integer variables, low and high, and set both to zero. These two numbers will always describe the possible range of unmatched opening parentheses after seeing the characters so far.

I then loop over every character of the input string.  

When the current character is an opening parenthesis I add one to both counters because every possible interpretation now has one extra unmatched open.  

When the current character is a closing parenthesis I subtract one from both counters. This shrinks the whole range by one.  

When the current character is a star I treat it in the most flexible way: I subtract one from low (as if the star closed a parenthesis) and add one to high (as if the star opened a parenthesis).  

Immediately after updating the counters I check whether high has dropped below zero. If it has, even the most optimistic choice of stars still produced more closing parentheses than opening ones, so the string can never be valid and I return false.  

I also guard low so it never stays negative. A negative low would mean I closed more parentheses than necessary; I could always have treated those extra stars as empty, so the true minimum possible balance is zero. Resetting low to zero keeps the range realistic.  

After the loop finishes I look at the final value of low. If low is exactly zero then there is at least one sequence of choices for the stars that leaves no unmatched openings, and the function returns true. If low is still greater than zero then even the greediest closing strategy left some openings unmatched, so the function returns false.

Edge cases such as a string that starts with a closing parenthesis, a string of only stars, or a string that ends with too many openings are all handled naturally by the same two-counter logic.

## Examples

**Example 1**  
Input: s = "()"  
- Start: low = 0, high = 0  
- See '(': low = 1, high = 1  
- See ')': low = 0, high = 0  
- Final low is 0 → return true  

**Example 2**  
Input: s = "(*)"  
- Start: low = 0, high = 0  
- See '(': low = 1, high = 1  
- See '*': low = 0, high = 2  
- See ')': low = -1 → reset to 0, high = 1  
- Final low is 0 → return true  

**Example 3**  
Input: s = "(*))"  
- Start: low = 0, high = 0  
- See '(': low = 1, high = 1  
- See '*': low = 0, high = 2  
- See ')': low = -1 → reset to 0, high = 1  
- See ')': low = -1 → reset to 0, high = 0  
- Final low is 0 → return true  

**Example 4**  
Input: s = "("  
- Start: low = 0, high = 0  
- See '(': low = 1, high = 1  
- Final low is 1 → return false  

## How to Use / Run Locally

**C++**  
1. Copy the C++ code into a file named `main.cpp`.  
2. Open a terminal and run: `g++ -std=c++17 main.cpp -o main`  
3. Run the program: `./main`  

**Java**  
1. Copy the Java code into a file named `Solution.java`.  
2. Compile with: `javac Solution.java`  
3. Run with: `java Solution`  

**JavaScript**  
1. Copy the JavaScript code into a file named `solution.js`.  
2. Run with Node.js: `node solution.js`  

**TypeScript**  
1. Copy the TypeScript code into a file named `solution.ts`.  
2. Compile with: `tsc solution.ts`  
3. Run the generated JavaScript: `node solution.js`  

**Python3**  
1. Copy the Python code into a file named `solution.py`.  
2. Run with: `python3 solution.py`  

**Go**  
1. Copy the Go code into a file named `main.go`.  
2. Run with: `go run main.go`  

In each case you will need to add a small main function or test harness that creates a string, calls the solution method, and prints the boolean result.

## Notes & Optimizations

This greedy two-counter approach is optimal for the given constraints. Because the length never exceeds 100, even a backtracking solution would pass, but the O(n) time and O(1) space version is cleaner and teaches a useful balance-tracking technique that appears in many other parentheses problems.  

One alternative is dynamic programming where dp[i][j] means whether the substring from i to j can be valid. That solution uses O(n²) time and space and is unnecessary here.  

Another common idea is to use two stacks—one for open parentheses and one for stars—and match them from left to right and then from right to left. That also works in O(n) time but needs extra space for the stacks. The low/high method avoids that overhead.  

Watch out for strings that consist only of stars; they are always valid because every star can be treated as empty. Also watch strings that start with a closing parenthesis; the high counter will immediately go negative and correctly reject them.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)