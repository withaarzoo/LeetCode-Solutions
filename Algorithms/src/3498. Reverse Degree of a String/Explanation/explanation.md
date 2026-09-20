# LeetCode 3498: Reverse Degree of a String - Complete Solution Guide

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

In this LeetCode problem called Reverse Degree of a String, you are given a string made of only lowercase English letters. Your task is to calculate a special value called the reverse degree.

For every character in the string you do two things. First, find its position in the reversed alphabet. That means 'a' is worth 26, 'b' is worth 25, and so on until 'z' which is worth 1. Second, multiply that value by the character's position in the string (counting from 1). Add up all these products and return the final sum.

This is a classic string manipulation problem that appears in competitive programming and DSA practice. It tests how carefully you handle character indexing and simple arithmetic while keeping the code clean and efficient.

## Constraints

- 1 <= s.length <= 1000
- s contains only lowercase English letters

## Intuition

When I first read the problem, I noticed that every letter has a fixed reverse-alphabet score that never changes. I also saw that the string index is simply 1-based and easy to calculate while looping. There is no need for sorting, hashing, or any complex structure. A single pass through the string is enough to compute everything. That observation immediately pointed me toward a linear scan solution for this Reverse Degree of a String problem.

## Approach

I start with a variable that will hold the final answer and set it to zero.  
I walk through the string from left to right using a normal index.  
For each character I calculate its reverse-alphabet value by subtracting its distance from 'a' from 26.  
I multiply that value by the current 1-based position.  
I add the product to the running total.  
When the loop finishes, the total is the reverse degree, so I return it.

This approach stays simple, avoids extra memory, and works for every valid input under the given constraints.

## Data Structures Used

No special data structures are required.  
I only use a few integer variables to store the running sum and the loop index.  
The input string itself is read-only, so the solution stays constant in extra space.

## Operations & Behavior Summary

1. Initialize a sum variable to 0.  
2. For each character at position i (0-based):  
   - Compute reverse value = 26 - (character - 'a')  
   - Multiply by (i + 1)  
   - Add the result to the sum  
3. After processing every character, return the sum.

This plain-English walkthrough matches the actual code flow in every language.

## Complexity

| Type              | Value | Explanation                                      |
|-------------------|-------|--------------------------------------------------|
| Time Complexity   | O(n)  | We examine each of the n characters exactly once |
| Space Complexity  | O(1)  | Only a constant number of variables are used     |

## Multi-language Solutions

### C++
```cpp
class Solution { 
public: 
    int reverseDegree(string s) { 
        int ans = 0;
        for (int i = 0; i < s.size(); ++i) {
            ans += (26 - (s[i] - 'a')) * (i + 1);
        }
        return ans;
    } 
};
```

### Java
```java
class Solution { 
    public int reverseDegree(String s) { 
        int ans = 0;
        for (int i = 0; i < s.length(); ++i) {
            ans += (26 - (s.charAt(i) - 'a')) * (i + 1);
        }
        return ans;
    } 
}
```

### JavaScript
```javascript
/** 
 * @param {string} s 
 * @return {number} 
 */ 
var reverseDegree = function(s) { 
    let ans = 0;
    for (let i = 0; i < s.length; ++i) {
        ans += (26 - (s.charCodeAt(i) - 97)) * (i + 1);
    }
    return ans;
};
```

### TypeScript
```typescript
function reverseDegree(s: string): number {
    let ans = 0;
    for (let i = 0; i < s.length; ++i) {
        ans += (26 - (s.charCodeAt(i) - 97)) * (i + 1);
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def reverseDegree(self, s: str) -> int:
        ans = 0
        for i, c in enumerate(s):
            ans += (26 - (ord(c) - ord('a'))) * (i + 1)
        return ans
```

### Go
```go
func reverseDegree(s string) int {
    ans := 0
    for i, c := range s {
        ans += (26 - int(c - 'a')) * (i + 1)
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea is identical across all six languages, so the reasoning stays the same.

I keep one integer that accumulates the answer.  
Inside the loop I convert the current character into a number between 0 and 25 by subtracting the code of 'a'. Subtracting that number from 26 gives the reverse-alphabet score.  

Multiplying by the 1-based index is done next. Using i+1 instead of i is important; forgetting the +1 would produce wrong answers on every test case.  

In C++ and Java I use the built-in length or size methods and simple char arithmetic.  
In JavaScript and TypeScript I use charCodeAt and subtract 97 (the code for 'a').  
In Python I use enumerate so I get both the index and the character in one step, then ord() for the character value.  
In Go I range over the string, which gives the index and the rune; converting the rune to int works the same way.

Edge cases are handled automatically. A string of length 1 simply multiplies the reverse value by 1. A string of all 'a's produces the largest possible sum for its length. Because the length never exceeds 1000, integer overflow is not a concern in any of the languages.

## Examples

**Example 1**  
Input: s = "abc"  
Output: 148  

Trace:  
- 'a' → 26 × 1 = 26  
- 'b' → 25 × 2 = 50  
- 'c' → 24 × 3 = 72  
- Total = 26 + 50 + 72 = 148  

**Example 2**  
Input: s = "zaza"  
Output: 160  

Trace:  
- 'z' → 1 × 1 = 1  
- 'a' → 26 × 2 = 52  
- 'z' → 1 × 3 = 3  
- 'a' → 26 × 4 = 104  
- Total = 1 + 52 + 3 + 104 = 160  

**Example 3**  
Input: s = "a"  
Output: 26  

Trace:  
- 'a' → 26 × 1 = 26  

## How to Use / Run Locally

**C++**  
Save the code in a file named main.cpp.  
Compile with: g++ -std=c++17 main.cpp -o main  
Run with: ./main  

**Java**  
Save the code in a file named Solution.java.  
Compile with: javac Solution.java  
Run with: java Solution  

**JavaScript**  
Save the code in a file named reverseDegree.js.  
Run with: node reverseDegree.js  

**TypeScript**  
Save the code in a file named reverseDegree.ts.  
Compile with: tsc reverseDegree.ts  
Run the generated JavaScript file with node.  

**Python3**  
Save the code in a file named reverse_degree.py.  
Run with: python3 reverse_degree.py  

**Go**  
Save the code in a file named reverse_degree.go.  
Run with: go run reverse_degree.go  

In each case you will need to add a small main function or driver code that creates a Solution object (or calls the function) and prints the result for a test string.

## Notes & Optimizations

The solution is already optimal for both time and space. No further asymptotic improvement is possible because every character must be examined at least once.

One common mistake is forgetting to use 1-based indexing. Another is mixing up normal alphabet order with reverse order. Both mistakes produce wrong answers on the sample cases.

Because the string length is at most 1000, even a less efficient approach would still pass, but the linear scan is the cleanest and most natural way to solve the Reverse Degree of a String problem.

If the constraints ever grew much larger, the same O(n) approach would still be the correct choice.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)