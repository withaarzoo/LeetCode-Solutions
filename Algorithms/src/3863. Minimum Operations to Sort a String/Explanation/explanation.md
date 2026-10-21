# Minimum Operations to Sort a String - LeetCode 3863 Solution

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

You are given a string made of only lowercase English letters. In one operation you are allowed to pick any substring that is not the entire string and sort that substring into non-descending alphabetical order.  

The goal is to find the minimum number of such operations needed to make the whole string sorted in non-descending order. If it is impossible, return -1.  

This is a classic competitive programming problem that tests careful case analysis on strings rather than heavy data structures. The key challenge is that you can never sort the full string in a single move, so the positions of the smallest and largest characters become the deciding factor.

## Constraints

- 1 <= s.length <= 10^5
- s consists of only lowercase English letters

Because the length can reach 100000, any solution that sorts the string multiple times or uses extra memory proportional to n will be too slow. A linear scan with constant extra space is required.

## Intuition

The first thing I noticed is that for a string of length 2 the only possible operations do nothing. So an unsorted pair is simply impossible.  

For longer strings I realized I never need more than three operations. Sorting a large proper substring can rearrange almost everything at once. The only real obstacles are the global minimum character and the global maximum character. If either of them is already sitting at the correct end, one operation finishes the job. If one of them sits somewhere in the middle, two operations are enough. Only when both extremes are stuck at the wrong ends do I need three operations.  

That observation turns the whole problem into a short decision tree based on the positions of the minimum and maximum characters.

## Approach

1. Check whether the string is already non-decreasing. If it is, return 0.  
2. If the length is exactly 2, return -1.  
3. Find the overall minimum character and overall maximum character.  
4. If the first character is already the minimum or the last character is already the maximum, return 1.  
5. Scan the middle part of the string (everything except the first and last characters). If either the minimum or the maximum appears there, return 2.  
6. Otherwise return 3.  

The logic never actually performs any sorting. It only inspects the string a few times and decides the answer from the positions of the extreme characters.

## Data Structures Used

No extra data structures are needed. The algorithm works directly on the input string and a handful of character variables that store the minimum and maximum values. This keeps both time and space optimal for the given constraints.

## Operations & Behavior Summary

- Walk through consecutive characters to test if the string is already sorted.  
- Early exit for the impossible length-2 case.  
- Single pass to discover the global min and max characters.  
- Constant-time checks on the first and last positions.  
- One more linear scan over the interior characters to see whether an extreme value is present.  
- Return the corresponding answer (0, 1, 2, 3 or -1) according to the decision tree above.

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(n)  | A constant number of linear scans over the string of length n. |
| Space Complexity | O(1)  | Only a few character variables and loop indices are used. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minOperations(string s) {
        int n = s.size();
        bool sorted = true;
        for (int i = 1; i < n; ++i) {
            if (s[i] < s[i - 1]) {
                sorted = false;
                break;
            }
        }
        if (sorted) return 0;
        if (n == 2) return -1;
        char mn = *min_element(s.begin(), s.end());
        char mx = *max_element(s.begin(), s.end());
        if (s[0] == mn || s[n - 1] == mx) return 1;
        for (int i = 1; i < n - 1; ++i) {
            if (s[i] == mn || s[i] == mx) return 2;
        }
        return 3;
    }
};
```

### Java
```java
class Solution {
    public int minOperations(String s) {
        int n = s.length();
        boolean sorted = true;
        for (int i = 1; i < n; ++i) {
            if (s.charAt(i) < s.charAt(i - 1)) {
                sorted = false;
                break;
            }
        }
        if (sorted) return 0;
        if (n == 2) return -1;
        char mn = s.charAt(0), mx = s.charAt(0);
        for (int i = 1; i < n; ++i) {
            char c = s.charAt(i);
            if (c < mn) mn = c;
            if (c > mx) mx = c;
        }
        if (s.charAt(0) == mn || s.charAt(n - 1) == mx) return 1;
        for (int i = 1; i < n - 1; ++i) {
            char c = s.charAt(i);
            if (c == mn || c == mx) return 2;
        }
        return 3;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {number}
 */
var minOperations = function(s) {
    const n = s.length;
    let sorted = true;
    for (let i = 1; i < n; ++i) {
        if (s[i] < s[i - 1]) {
            sorted = false;
            break;
        }
    }
    if (sorted) return 0;
    if (n === 2) return -1;
    let mn = s[0], mx = s[0];
    for (let i = 1; i < n; ++i) {
        if (s[i] < mn) mn = s[i];
        if (s[i] > mx) mx = s[i];
    }
    if (s[0] === mn || s[n - 1] === mx) return 1;
    for (let i = 1; i < n - 1; ++i) {
        if (s[i] === mn || s[i] === mx) return 2;
    }
    return 3;
};
```

### TypeScript
```typescript
function minOperations(s: string): number {
    const n = s.length;
    let sorted = true;
    for (let i = 1; i < n; ++i) {
        if (s[i] < s[i - 1]) {
            sorted = false;
            break;
        }
    }
    if (sorted) return 0;
    if (n === 2) return -1;
    let mn = s[0], mx = s[0];
    for (let i = 1; i < n; ++i) {
        if (s[i] < mn) mn = s[i];
        if (s[i] > mx) mx = s[i];
    }
    if (s[0] === mn || s[n - 1] === mx) return 1;
    for (let i = 1; i < n - 1; ++i) {
        if (s[i] === mn || s[i] === mx) return 2;
    }
    return 3;
};
```

### Python3
```python
class Solution:
    def minOperations(self, s: str) -> int:
        n = len(s)
        if all(s[i] <= s[i + 1] for i in range(n - 1)):
            return 0
        if n == 2:
            return -1
        mn = min(s)
        mx = max(s)
        if s[0] == mn or s[-1] == mx:
            return 1
        if any(c == mn or c == mx for c in s[1:-1]):
            return 2
        return 3
```

### Go
```go
func minOperations(s string) int {
    n := len(s)
    sorted := true
    for i := 1; i < n; i++ {
        if s[i] < s[i-1] {
            sorted = false
            break
        }
    }
    if sorted {
        return 0
    }
    if n == 2 {
        return -1
    }
    mn, mx := s[0], s[0]
    for i := 1; i < n; i++ {
        if s[i] < mn {
            mn = s[i]
        }
        if s[i] > mx {
            mx = s[i]
        }
    }
    if s[0] == mn || s[n-1] == mx {
        return 1
    }
    for i := 1; i < n-1; i++ {
        if s[i] == mn || s[i] == mx {
            return 2
        }
    }
    return 3
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the same decision tree, so the reasoning is identical across languages.  

First the code measures the length of the string. It then walks from the second character to the end, comparing each character with the one before it. As soon as a descent is found the sorted flag is turned off. If the flag stays true the answer is immediately 0.  

Next comes the length-2 guard. Because no proper substring can change a two-character string, any unsorted pair returns -1.  

A linear pass then discovers the smallest and largest characters present in the whole string. These two values are stored in simple variables.  

The ends are examined next. If the first character already equals the global minimum, sorting the remaining suffix is enough, so the answer is 1. The same logic applies when the last character equals the global maximum.  

If neither end is helpful, the interior characters are examined. Finding either extreme value anywhere in the middle means that value can be moved to an end with one operation and the rest of the string can be finished with a second operation. The answer becomes 2.  

Only when both extremes sit exclusively at the two ends (and in the wrong order) does the algorithm fall through to the final case and return 3.  

Because every language performs the same constant number of linear passes and uses only a few scalar variables, the asymptotic complexity stays O(n) time and O(1) space in all six solutions.

## Examples

**Example 1**  
Input: s = "dog"  
Output: 1  
The string is not sorted. The minimum character 'd' is already at the front, so sorting the suffix "og" produces "dgo", which is sorted. One operation is enough.

**Example 2**  
Input: s = "card"  
Output: 2  
Neither end holds the global minimum or maximum in the right place, but both extremes appear in the middle. Two operations suffice: first bring an extreme character to an end, then sort the remaining proper substring.

**Example 3**  
Input: s = "gf"  
Output: -1  
Length is 2 and the characters are out of order. No useful operation exists, so the answer is -1.

## How to Use / Run Locally

Clone the repository and open the file that matches the language you want to test.  

- **C++**: compile with `g++ -std=c++17 solution.cpp -o solution` and run `./solution`.  
- **Java**: compile with `javac Solution.java` and run `java Solution`.  
- **JavaScript**: run with `node solution.js`.  
- **TypeScript**: compile with `tsc solution.ts` then run the generated JavaScript, or use `ts-node solution.ts`.  
- **Python3**: run with `python3 solution.py`.  
- **Go**: run with `go run solution.go`.  

Each file contains a simple main or test harness that reads a string from standard input and prints the result. You can also copy the core function into any online judge that accepts the corresponding language.

## Notes & Optimizations

The solution is already optimal for the given constraints. No further asymptotic improvement is possible because the whole string must be examined at least once.  

Edge cases worth remembering:  
- A single character is always sorted and returns 0.  
- Strings where the minimum appears only at the very end and the maximum appears only at the very beginning force the three-operation path.  
- Duplicate minimum or maximum characters do not change the logic; the presence of the value anywhere in the middle is enough to return 2.  

An alternative approach that actually sorts the prefix or suffix and then checks the result works but is slower (O(n log n)) and unnecessary. The pure case-analysis version is both faster and clearer.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)