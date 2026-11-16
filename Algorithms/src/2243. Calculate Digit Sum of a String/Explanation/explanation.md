# Calculate Digit Sum of a String - LeetCode 2243 Solution

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

You are given a string s that contains only digits and an integer k. The task is to keep reducing the string by repeatedly grouping its characters into consecutive groups of size k (the last group can be smaller) and replacing every group with the sum of its digits written as a string. You merge these sums back into a new string and repeat the process as long as the length of the current string is greater than k. Once the string length becomes less than or equal to k, you return that final string.

This is a classic simulation problem on strings that appears in coding interviews and competitive programming contests. The goal is to correctly implement the digit sum grouping process without missing edge cases such as leading zeros or uneven final groups.

## Constraints

- 1 <= s.length <= 100
- 1 <= k <= 100
- s consists of digits only (0-9)

## Intuition

When I first read the problem I noticed that the string is short (maximum length 100). That immediately told me a straightforward simulation would be fast enough. There is no need for clever math or prefix sums. I just need to keep chopping the current string into groups of size k, calculate the digit sum of each group, turn those sums back into strings, and stick them together until the length drops to k or less. The process is described almost like a loop in the problem statement itself, so following it directly feels natural.

## Approach

I start with the input string and keep looping while its length is greater than k. Inside each round I create a new empty result string. I walk through the current string in steps of k. For every group I add up the numeric values of its characters. After finishing a group I convert the sum into a string and append it to the result. When all groups are processed I replace the old string with this new shorter string and check the length again. The moment the length is no longer greater than k I stop and return whatever is left.

This approach stays close to the problem description and makes it easy to verify against the sample cases.

## Data Structures Used

- String (or StringBuilder in Java): The main data structure. I use a temporary string in every round to collect the digit sums of the groups. No arrays, maps, or trees are required because the operations are purely sequential.

## Operations & Behavior Summary

1. Check whether the current string length is still greater than k.
2. If yes, prepare an empty collector string for the next round.
3. Move through the string in fixed steps of size k.
4. For each group calculate the sum of its digits by converting characters to numbers.
5. Append the string representation of that sum to the collector.
6. After processing every group, replace the original string with the collector.
7. Repeat until the length condition fails.
8. Return the final string.

## Complexity

| Type              | Value     | Explanation |
|-------------------|-----------|-------------|
| Time Complexity   | O(n)      | Each character is processed a constant number of times across all rounds because every round produces a strictly shorter string. With n ≤ 100 the total work stays linear. |
| Space Complexity  | O(n)      | Only a temporary string of size at most the current length is created in each round. No extra data structures grow with input size. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    string digitSum(string s, int k) {
        while (s.length() > k) {
            string next = "";
            for (int i = 0; i < s.length(); i += k) {
                int sum = 0;
                for (int j = i; j < s.length() && j < i + k; j++) {
                    sum += s[j] - '0';
                }
                next += to_string(sum);
            }
            s = next;
        }
        return s;
    }
};
```

### Java
```java
class Solution {
    public String digitSum(String s, int k) {
        while (s.length() > k) {
            StringBuilder next = new StringBuilder();
            for (int i = 0; i < s.length(); i += k) {
                int sum = 0;
                for (int j = i; j < s.length() && j < i + k; j++) {
                    sum += s.charAt(j) - '0';
                }
                next.append(sum);
            }
            s = next.toString();
        }
        return s;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
var digitSum = function(s, k) {
    while (s.length > k) {
        let next = "";
        for (let i = 0; i < s.length; i += k) {
            let sum = 0;
            for (let j = i; j < s.length && j < i + k; j++) {
                sum += Number(s[j]);
            }
            next += sum.toString();
        }
        s = next;
    }
    return s;
};
```

### TypeScript
```typescript
function digitSum(s: string, k: number): string {
    while (s.length > k) {
        let next = "";
        for (let i = 0; i < s.length; i += k) {
            let sum = 0;
            for (let j = i; j < s.length && j < i + k; j++) {
                sum += Number(s[j]);
            }
            next += sum.toString();
        }
        s = next;
    }
    return s;
};
```

### Python3
```python
class Solution:
    def digitSum(self, s: str, k: int) -> str:
        while len(s) > k:
            next_s = ""
            for i in range(0, len(s), k):
                total = 0
                for j in range(i, min(i + k, len(s))):
                    total += int(s[j])
                next_s += str(total)
            s = next_s
        return s
```

### Go
```go
func digitSum(s string, k int) string {
    for len(s) > k {
        next := ""
        for i := 0; i < len(s); i += k {
            sum := 0
            for j := i; j < len(s) && j < i+k; j++ {
                sum += int(s[j] - '0')
            }
            next += strconv.Itoa(sum)
        }
        s = next
    }
    return s
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

In every language the core idea is the same. I keep a while loop that continues as long as the length of the current string is greater than k. Inside the loop I allocate a fresh empty string (or StringBuilder in Java for efficiency). Then I run an outer loop that starts at index 0 and advances by k each time. This outer loop isolates one group per iteration.

For the current group I run an inner loop that adds the digit values. I convert each character to an integer by subtracting the character '0' (or using a language-specific conversion such as Number() or int()). After the sum is ready I convert it back to a string and append it to the collector. Once the outer loop finishes, the collector becomes the new string for the next round.

The only small language differences are:
- C++ and Go use character arithmetic (s[j] - '0').
- Java uses charAt and StringBuilder.append.
- JavaScript and TypeScript use Number() or unary plus.
- Python uses int() and string concatenation.
- Go needs strconv.Itoa for the conversion.

Edge cases such as a final group shorter than k or a string that consists only of zeros are handled automatically because the inner loop stops at the string end and zero sums simply produce the character '0'.

## Examples

**Example 1**

Input: s = "11111222223", k = 3  
Output: "135"

Trace:
- First round groups: "111", "112", "222", "23"  
  Digit sums: 3, 4, 6, 5 → new string "3465"
- Second round groups: "346", "5"  
  Digit sums: 13, 5 → new string "135"
- Length of "135" is 3 which is not greater than k, so return "135"

**Example 2**

Input: s = "00000000", k = 3  
Output: "000"

Trace:
- Groups: "000", "000", "00"  
  Digit sums: 0, 0, 0 → new string "000"
- Length equals k, stop and return "000"

**Example 3**

Input: s = "1234", k = 2  
Output: "37"

Trace:
- Groups: "12", "34"  
  Digit sums: 3, 7 → new string "37"
- Length equals k, return "37"

## How to Use / Run Locally

1. Copy the code for the language you want into a file (for example main.cpp, Solution.java, digitSum.js, etc.).
2. C++:  
   g++ -std=c++17 main.cpp -o main && ./main
3. Java:  
   javac Solution.java && java Solution
4. JavaScript:  
   node digitSum.js
5. TypeScript:  
   tsc digitSum.ts && node digitSum.js
6. Python3:  
   python3 solution.py
7. Go:  
   go run main.go

Add a small main function or test harness that creates a Solution object (or calls the function) and prints the result for the sample inputs.

## Notes & Optimizations

Because the maximum length is only 100, the simple simulation is already optimal in practice. There is no need for further optimization. One possible alternative is to keep converting the string into a list of integers, but that adds unnecessary complexity without improving the asymptotic bounds. Always watch for the last incomplete group and for strings that contain many zeros; both are correctly handled by the loops shown above.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
