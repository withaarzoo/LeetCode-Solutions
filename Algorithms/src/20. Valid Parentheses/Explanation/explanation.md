# Valid Parentheses - LeetCode 20 Solution | Stack Based Balanced Brackets Check

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

This is the classic Valid Parentheses problem from LeetCode (problem number 20). You are given a string that contains only the characters '(', ')', '{', '}', '[' and ']'. Your job is to check whether the string is valid.

A string is considered valid only when three simple rules hold true:

1. Every opening bracket is closed by the same type of bracket.
2. Opening brackets are closed in the correct order (the most recent open must be closed first).
3. Every closing bracket has a matching opening bracket of the same type.

The function should return true if the input string is valid and false otherwise. This is one of the most common stack based problems asked in coding interviews and forms the foundation for many balanced parentheses and expression validation tasks in DSA.

## Constraints

- 1 <= s.length <= 10^4
- s consists only of the characters '(', ')', '{', '}', '[' and ']'

## Intuition

When I first looked at the Valid Parentheses problem I noticed that the order of brackets is everything. The last opening bracket I see must be the first one that gets closed. That "last in, first out" behavior is exactly what a stack gives me.

I realized I could push every opening bracket onto a stack and, whenever I meet a closing bracket, check whether the top of the stack matches it. If it does, I pop. If it does not match or the stack is empty, the string is already invalid. At the end the stack must be empty for the whole string to be balanced.

This stack based approach feels natural because it mirrors how we mentally track nested brackets when we read code or mathematical expressions.

## Approach

I create an empty stack that will hold opening brackets.

I walk through every character of the input string from left to right.

- If the character is an opening bracket ('(', '{', or '['), I push it onto the stack.
- If the character is a closing bracket, I first check whether the stack is empty. An empty stack means there is no matching open, so I immediately return false.
- I then look at the top of the stack and pop it. I verify that the popped opening bracket is the correct pair for the current closing bracket. If the pair does not match I return false.

After processing every character I check the stack one last time. If it is empty, every opening found its matching closer in the right order and the string is valid. Anything left on the stack means there are unmatched openings, so the answer is false.

This linear scan with a stack gives a clean and efficient solution for the balanced brackets problem.

## Data Structures Used

- Stack: I use a stack (or an array used as a stack) to keep track of unmatched opening brackets. The stack is ideal because the most recently seen opening bracket must be closed first. No other data structure gives this LIFO behavior as simply.

## Operations & Behavior Summary

1. Initialize an empty stack.
2. For each character in the string:
   - Opening bracket → push onto the stack.
   - Closing bracket → if stack is empty return false; otherwise pop the top and check whether it forms a valid pair with the current character. If not, return false.
3. After the loop, return true only if the stack is empty.

The algorithm never revisits a character and never needs extra bookkeeping beyond the stack itself.

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(n)  | n is the length of the input string. Each character is examined exactly once. |
| Space Complexity | O(n)  | In the worst case the stack stores up to n/2 opening brackets (for example a long sequence of opening brackets). |

## Multi-language Solutions

### C++
```cpp
class Solution {

public:

    bool isValid(string s) {

        stack<char> st;

        for (char c : s) {

            if (c == '(' || c == '{' || c == '[') {

                st.push(c);

            } else {

                if (st.empty()) return false;

                char top = st.top();

                st.pop();

                if ((c == ')' && top != '(') || (c == '}' && top != '{') || (c == ']' && top != '[')) {

                    return false;

                }

            }

        }

        return st.empty();

    }

};
```

### Java
```java
class Solution {

    public boolean isValid(String s) {

        Stack<Character> st = new Stack<>();

        for (char c : s.toCharArray()) {

            if (c == '(' || c == '{' || c == '[') {

                st.push(c);

            } else {

                if (st.isEmpty()) return false;

                char top = st.pop();

                if ((c == ')' && top != '(') || (c == '}' && top != '{') || (c == ']' && top != '[')) {

                    return false;

                }

            }

        }

        return st.isEmpty();

    }

}
```

### JavaScript
```javascript
/**

 * @param {string} s

 * @return {boolean}

 */

var isValid = function(s) {

    const st = [];

    for (const c of s) {

        if (c === '(' || c === '{' || c === '[') {

            st.push(c);

        } else {

            if (st.length === 0) return false;

            const top = st.pop();

            if ((c === ')' && top !== '(') || (c === '}' && top !== '{') || (c === ']' && top !== '[')) {

                return false;

            }

        }

    }

    return st.length === 0;

};
```

### TypeScript
```typescript
function isValid(s: string): boolean {

    const st: string[] = [];

    for (const c of s) {

        if (c === '(' || c === '{' || c === '[') {

            st.push(c);

        } else {

            if (st.length === 0) return false;

            const top = st.pop()!;

            if ((c === ')' && top !== '(') || (c === '}' && top !== '{') || (c === ']' && top !== '[')) {

                return false;

            }

        }

    }

    return st.length === 0;

};
```

### Python3
```python
class Solution:

    def isValid(self, s: str) -> bool:

        st = []

        for c in s:

            if c in '({[':

                st.append(c)

            else:

                if not st:

                    return False

                top = st.pop()

                if (c == ')' and top != '(') or (c == '}' and top != '{') or (c == ']' and top != '['):

                    return False

        return not st
```

### Go
```go
func isValid(s string) bool {

    st := []rune{}

    for _, c := range s {

        if c == '(' || c == '{' || c == '[' {

            st = append(st, c)

        } else {

            if len(st) == 0 {

                return false

            }

            top := st[len(st)-1]

            st = st[:len(st)-1]

            if (c == ')' && top != '(') || (c == '}' && top != '{') || (c == ']' && top != '[') {

                return false

            }

        }

    }

    return len(st) == 0

}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages. Here is the reasoning that sits behind every line.

I begin by creating an empty stack. In C++ and Java I use the language’s built-in Stack or stack container. In JavaScript, TypeScript, Python and Go I simply use a dynamic array and treat the end of the array as the top of the stack.

Next I iterate over every character of the input string. For each character I decide whether it is an opening or a closing bracket.

When I see an opening bracket I push it. This records that I still need a matching closer later. The push operation is constant time in every language.

When I see a closing bracket I first test whether the stack is empty. If it is empty there is no opening left to match this closer, so the string is invalid and I return false immediately. This early exit handles cases such as a string that starts with a closing bracket or has more closers than openers.

If the stack is not empty I remove the top element (the most recent unmatched opening). I then compare that opening with the current closing character:

- ')' is only allowed to match '('
- '}' is only allowed to match '{'
- ']' is only allowed to match '['

If any of those three checks fail I return false. This single comparison catches both wrong-type mismatches and out-of-order closings.

After the loop finishes I look at the stack once more. An empty stack means every opening found its matching closer in the correct nested order, so I return true. Any remaining items mean there are leftover openings that were never closed, so I return false.

Edge cases that this logic covers automatically:

- Empty string is never given because the length constraint starts at 1, but the same code would correctly return true for an empty string.
- Single closing bracket returns false because the stack is empty.
- Single opening bracket returns false because the stack is not empty at the end.
- Properly nested and sequential mixtures such as "()[]{}" and "([])" all return true.
- Crossed brackets such as "([)]" fail the pair check and return false.

Because every language uses the same control flow, the explanation above applies equally to the C++, Java, JavaScript, TypeScript, Python3 and Go implementations. The only surface differences are the exact syntax for pushing, popping and checking emptiness, but the decision points remain identical.

## Examples

**Example 1**  
Input: s = "()"  
Output: true  

Trace:  
- See '(' → push. Stack = ['(']  
- See ')' → pop '(' and it matches. Stack becomes empty.  
- End of string, stack empty → true.

**Example 2**  
Input: s = "()[]{}"  
Output: true  

Trace:  
- '(' → push  
- ')' → matches, pop  
- '[' → push  
- ']' → matches, pop  
- '{' → push  
- '}' → matches, pop  
- Stack empty → true.

**Example 3**  
Input: s = "(]"  
Output: false  

Trace:  
- '(' → push. Stack = ['(']  
- ']' → top is '(', which does not match ']' → return false.

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `valid_parentheses.cpp`.  
2. Compile with `g++ -std=c++17 valid_parentheses.cpp -o valid_parentheses`.  
3. Run with `./valid_parentheses`.  
4. You will need a small main function that reads a string and prints the result of `isValid`.

**Java**  
1. Save the code in `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run with `java Solution`.  
4. Add a main method that creates a Solution object and calls `isValid` with a test string.

**JavaScript**  
1. Save the code in `validParentheses.js`.  
2. Run with `node validParentheses.js`.  
3. Call the function with a test string and log the boolean result.

**TypeScript**  
1. Save the code in `validParentheses.ts`.  
2. Compile with `tsc validParentheses.ts`.  
3. Run the generated JavaScript with `node validParentheses.js`.  
4. Or use `ts-node` if you have it installed.

**Python3**  
1. Save the code in `valid_parentheses.py`.  
2. Run with `python3 valid_parentheses.py`.  
3. Inside an `if __name__ == "__main__":` block create a Solution instance and print the result of `isValid` for a few test cases.

**Go**  
1. Save the code in `valid_parentheses.go`.  
2. Run with `go run valid_parentheses.go`.  
3. Add a main function that calls `isValid` and prints the boolean answer.

In every language you can replace the hard-coded test strings with input read from the console if you want interactive testing.

## Notes & Optimizations

The stack solution is already optimal for both time and space under the given constraints. No faster asymptotic approach exists because every character must be examined at least once.

A common alternative is to use a hash map that stores the matching pairs (closing → opening). This makes the pair-check cleaner and avoids a long chain of if conditions, but it does not change the complexity.

Watch out for the edge case where the string begins with a closing bracket or contains more closing brackets than opening ones; both are caught by the empty-stack check.

If the problem ever allowed other characters you would simply ignore them, but the current constraints guarantee that every character is a bracket, so no extra filtering is required.

This Valid Parentheses solution is a perfect building block for more advanced problems such as longest valid parentheses substring, removing invalid parentheses, or evaluating arithmetic expressions with brackets.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)