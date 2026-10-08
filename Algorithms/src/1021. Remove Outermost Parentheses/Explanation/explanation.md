# Remove Outermost Parentheses | LeetCode 1021 Solution

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

You are given a valid parentheses string `s`. A valid parentheses string is either empty, or formed by wrapping a valid string inside a pair of parentheses, or by joining two valid strings together.

The string can be broken into one or more "primitive" parts. A primitive part is a valid parentheses string that cannot be split further into two non-empty valid parentheses strings.

The task is to remove the outermost pair of parentheses from every primitive part and return the resulting string.

In short, take a valid parentheses string, find its top-level groups, strip the outer parentheses from each group, and join what remains.

## Constraints

- 1 <= s.length <= 10^5
- s[i] is either '(' or ')'
- s is a valid parentheses string

## Intuition

When I first looked at the problem, I noticed that every primitive group starts when the balance of open parentheses goes from zero to one, and it ends when the balance returns to zero. Those starting and ending parentheses are the ones that should be removed. Everything inside them should stay.

I realized I do not need a stack. A simple integer that tracks the current depth is enough, because the input is already guaranteed to be valid. This observation led me to a clean single-pass solution that uses constant extra memory besides the result string.

## Approach

I walk through the string once while keeping a balance counter.

- When I meet an opening parenthesis, I check the current balance. If the balance is already greater than zero, the parenthesis is not outermost, so I keep it. Then I increase the balance.
- When I meet a closing parenthesis, I first decrease the balance. After the decrease, if the balance is still greater than zero, the parenthesis is not outermost, so I keep it.

By the end of the scan I have collected exactly the characters that belong inside the primitive groups. The outermost parentheses of every group have been skipped.

## Data Structures Used

- A single integer (balance / depth counter) to track how many unmatched opening parentheses are currently open.
- A result string (or StringBuilder / dynamic array) to collect the characters that should be kept.

No stack is required because the string is already valid. The integer alone is sufficient to know the current nesting level.

## Operations & Behavior Summary

1. Initialize an empty result and set balance to 0.
2. For each character in the input:
   - If the character is '(':
     - If balance > 0, append it to the result.
     - Increase balance by 1.
   - If the character is ')':
     - Decrease balance by 1.
     - If balance > 0, append it to the result.
3. Return the result string.

The algorithm never looks ahead or backtracks. It decides for each character in constant time whether it should be kept.

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time       | O(n)  | We examine each of the n characters exactly once. |
| Space      | O(n)  | The result string can grow up to length n in the worst case. Extra memory used by the algorithm itself is O(1). |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    string removeOuterParentheses(string s) {
        string res;
        int bal = 0;
        for (char c : s) {
            if (c == '(') {
                if (bal > 0) res += c;
                bal++;
            } else {
                bal--;
                if (bal > 0) res += c;
            }
        }
        return res;
    }
};
```

### Java
```java
class Solution {
    public String removeOuterParentheses(String s) {
        StringBuilder res = new StringBuilder();
        int bal = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                if (bal > 0) res.append(c);
                bal++;
            } else {
                bal--;
                if (bal > 0) res.append(c);
            }
        }
        return res.toString();
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {string}
 */
var removeOuterParentheses = function(s) {
    let res = "";
    let bal = 0;
    for (let c of s) {
        if (c === '(') {
            if (bal > 0) res += c;
            bal++;
        } else {
            bal--;
            if (bal > 0) res += c;
        }
    }
    return res;
};
```

### TypeScript
```typescript
function removeOuterParentheses(s: string): string {
    let res = "";
    let bal = 0;
    for (let c of s) {
        if (c === '(') {
            if (bal > 0) res += c;
            bal++;
        } else {
            bal--;
            if (bal > 0) res += c;
        }
    }
    return res;
};
```

### Python3
```python
class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        res = []
        bal = 0
        for c in s:
            if c == '(':
                if bal > 0:
                    res.append(c)
                bal += 1
            else:
                bal -= 1
                if bal > 0:
                    res.append(c)
        return ''.join(res)
```

### Go
```go
func removeOuterParentheses(s string) string {
    res := make([]byte, 0, len(s))
    bal := 0
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' {
            if bal > 0 {
                res = append(res, c)
            }
            bal++
        } else {
            bal--
            if bal > 0 {
                res = append(res, c)
            }
        }
    }
    return string(res)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages. Only the way we build the result string changes.

I start with a balance variable set to zero and an empty result container.

For every character I look at:

- Opening parenthesis case  
  I first ask whether the current balance is already positive. A positive balance means I am already inside a primitive group, so this opening parenthesis belongs in the answer. After making that decision I always increase the balance, because a new level has been opened.

- Closing parenthesis case  
  I first decrease the balance. Now the balance reflects the depth after this closing parenthesis has been matched. If the balance is still positive, the closing parenthesis is still inside a group and must be kept. If the balance has become zero, I have just finished a whole primitive group and this closing parenthesis was its outermost pair, so I discard it.

Because the input is guaranteed to be valid, the balance never goes negative and always ends at zero. That guarantee lets me avoid any extra checks or a stack.

Edge cases are handled naturally:

- A single pair "()" produces an empty result, which is correct.
- Nested groups such as "(()())" keep the inner content "()()".
- Consecutive primitive groups are processed one after another without any special handling.

The only language-specific detail is how the result is accumulated (string concatenation, StringBuilder, list + join, or byte slice). All of them produce the same final string.

## Examples

**Example 1**  
Input: `s = "(()())(())"`  
Output: `"()()()"`  

Trace:  
- First group `(()())` → keep `()()`  
- Second group `(())` → keep `()`  
- Final result: `()()()`

**Example 2**  
Input: `s = "(()())(())(()(()))"`  
Output: `"()()()()(())"`  

Trace:  
- Group 1 `(()())` → `()()`  
- Group 2 `(())` → `()`  
- Group 3 `(()(()))` → `()(())`  
- Final result: `()()()()(())`

**Example 3**  
Input: `s = "()()"`  
Output: `""`  

Trace:  
- First `()` → empty  
- Second `()` → empty  
- Final result is the empty string.

## How to Use / Run Locally

**C++**  
1. Copy the solution into a file named `main.cpp`.  
2. Compile: `g++ -std=c++17 main.cpp -o main`  
3. Run: `./main` (add a small driver that reads a string and prints the result if you want to test).

**Java**  
1. Place the class in `Solution.java`.  
2. Compile: `javac Solution.java`  
3. Run with a test harness or from an online judge.

**JavaScript**  
1. Save the function in `solution.js`.  
2. Run with Node: `node solution.js` (add a few console.log test cases).

**TypeScript**  
1. Save the function in `solution.ts`.  
2. Compile: `tsc solution.ts`  
3. Run the generated JavaScript with Node.

**Python3**  
1. Save the class in `solution.py`.  
2. Run: `python3 solution.py` (add a few print statements for testing).

**Go**  
1. Save the function in `main.go`.  
2. Run: `go run main.go` (add a main function that calls the solution with sample inputs).

## Notes & Optimizations

The balance-counter method is already optimal: linear time and constant extra space (apart from the output).  

A stack-based approach also works, but it uses more memory and is unnecessary because the string is valid.  

If the problem ever required reconstructing the original groups, a stack would become useful. For the current requirement of simply stripping outer parentheses, the integer counter is the cleanest solution.

The algorithm correctly handles the full range of constraints (up to 10^5 characters) without any performance issues.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)