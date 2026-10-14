# 3723. Maximize Sum of Squares of Digits – LeetCode Solution

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

We are given two positive integers called `num` and `sum`. We need to find a positive integer that has exactly `num` digits and whose digits add up to exactly `sum`. Among all such numbers, we want the one that gives the highest score. The score of a number is the sum of the squares of its digits.

If several numbers achieve the same maximum score, we must return the largest one among them (as a string). If no number satisfies the conditions, we return an empty string.

This is a classic digit-construction problem that appears in many competitive programming contests and LeetCode-style interviews. The goal is to maximize the sum of squares of digits while respecting a fixed digit length and a fixed digit sum.

## Constraints

- 1 <= num <= 2 * 10^5
- 1 <= sum <= 2 * 10^5

These limits rule out any solution that tries to generate all possible numbers or use heavy dynamic programming. We need a linear-time approach that builds the answer directly.

## Intuition

The first thing I noticed is that the score is the sum of squares of the digits. For a fixed total sum of digits, the sum of squares becomes larger when the digits are as unbalanced as possible. Putting as many 9s as we can, then one leftover digit, and the rest zeros gives the highest possible score.

Once we have that optimal set of digits, forming the largest number is straightforward: just place the digits from largest to smallest. This automatically avoids a leading zero and satisfies the “return the maximum one” requirement.

## Approach

1. Check whether it is even possible to form such a number. The maximum digit sum we can achieve with `num` digits is 9 * num. If the given `sum` is larger than that, return an empty string.
2. Calculate how many full 9s we need: let `q = sum / 9` and `r = sum % 9`.
3. Build the answer string by writing `q` nines first, then the single remainder digit `r` (if it is greater than zero), and finally fill the remaining positions with zeros so the total length becomes exactly `num`.
4. The resulting string is already the largest number that achieves the maximum score.

This greedy construction works because any more balanced distribution of digits would produce a strictly smaller sum of squares.

## Data Structures Used

- A simple string (or string builder / byte slice) to hold the final answer.  
  No extra arrays, maps, or trees are required. We only need a place to accumulate the digits while we construct the number from left to right.

## Operations & Behavior Summary

- First perform the impossibility check using a single multiplication and comparison.
- Compute the quotient and remainder when the target sum is divided by 9.
- Append the correct number of ‘9’ characters.
- Append the remainder digit if it is positive.
- Append the remaining zeros until the length equals `num`.
- Return the finished string (or an empty string if the check failed).

The whole process is a single left-to-right pass that never looks back.

## Complexity

| Complexity          | Value   | Explanation |
|---------------------|---------|-------------|
| Time Complexity     | O(num)  | We visit each of the `num` positions exactly once while building the result string. All other operations are constant time. |
| Space Complexity    | O(num)  | The only extra memory is the string that stores the answer. No auxiliary data structures grow with the input size. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    string maxSumOfSquares(int num, int sum) {
        if (sum > 9LL * num) return "";
        int q = sum / 9;
        int r = sum % 9;
        string ans;
        ans.reserve(num);
        ans.append(q, '9');
        if (r) ans.push_back('0' + r);
        ans.append(num - q - (r > 0), '0');
        return ans;
    }
};
```

### Java
```java
class Solution {
    public String maxSumOfSquares(int num, int sum) {
        if (sum > 9L * num) return "";
        int q = sum / 9;
        int r = sum % 9;
        StringBuilder sb = new StringBuilder(num);
        for (int i = 0; i < q; i++) sb.append('9');
        if (r > 0) sb.append((char)('0' + r));
        for (int i = q + (r > 0 ? 1 : 0); i < num; i++) sb.append('0');
        return sb.toString();
    }
}
```

### JavaScript
```javascript
/**
 * @param {number} num
 * @param {number} sum
 * @return {string}
 */
var maxSumOfSquares = function(num, sum) {
    if (sum > 9 * num) return "";
    let q = Math.floor(sum / 9);
    let r = sum % 9;
    let ans = '9'.repeat(q);
    if (r) ans += String(r);
    ans += '0'.repeat(num - q - (r ? 1 : 0));
    return ans;
};
```

### TypeScript
```typescript
function maxSumOfSquares(num: number, sum: number): string {
    if (sum > 9 * num) return "";
    let q = Math.floor(sum / 9);
    let r = sum % 9;
    let ans = '9'.repeat(q);
    if (r) ans += String(r);
    ans += '0'.repeat(num - q - (r ? 1 : 0));
    return ans;
};
```

### Python3
```python
class Solution:
    def maxSumOfSquares(self, num: int, sum: int) -> str:
        if sum > 9 * num:
            return ""
        q, r = divmod(sum, 9)
        return '9' * q + (str(r) if r else '') + '0' * (num - q - (1 if r else 0))
```

### Go
```go
func maxSumOfSquares(num int, sum int) string {
    if sum > 9*num {
        return ""
    }
    q := sum / 9
    r := sum % 9
    b := make([]byte, 0, num)
    for i := 0; i < q; i++ {
        b = append(b, '9')
    }
    if r > 0 {
        b = append(b, byte('0'+r))
    }
    for i := q + map[bool]int{true: 1, false: 0}[r > 0]; i < num; i++ {
        b = append(b, '0')
    }
    return string(b)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages; only the syntax for string building changes.

We begin by testing whether `sum` can ever be reached with `num` digits. If `sum > 9 * num` we immediately return the empty string. In languages where integer overflow is possible we cast to a wider type for the multiplication, but the idea stays the same.

Next we obtain the number of 9s (`q`) and the leftover digit (`r`) with a single division and modulo. Because of the earlier check we already know that `q` never exceeds `num`, and when `q` equals `num` the remainder must be zero.

We then construct the answer from the most significant digit to the least significant digit:

- Write `q` characters ‘9’.
- If `r` is positive, write the character that represents `r`.
- Fill every remaining position with ‘0’ until the length becomes exactly `num`.

This order guarantees two things at once: the sum of squares is maximal (because the digits are as unbalanced as possible) and the numeric value is the largest possible among all such maximal-score numbers (because larger digits appear earlier).

Edge cases are handled automatically. When the remainder is zero we simply write pure 9s followed by zeros. When we need only a single non-zero digit we write that digit followed by zeros. Leading zeros never appear because any non-zero digit is placed at the front.

The only differences between languages are the concrete tools used to build the string (std::string::append, StringBuilder, string concatenation, byte slices, etc.). The algorithmic decisions remain exactly the same.

## Examples

**Example 1**  
Input: num = 2, sum = 3  
Output: "30"  

Trace: Maximum possible sum is 18, so the request is feasible. We need zero full 9s and a remainder of 3. The string becomes “3” followed by one zero → “30”. The score is 9, which is higher than the scores of 12 and 21.

**Example 2**  
Input: num = 2, sum = 17  
Output: "98"  

Trace: One full 9 and remainder 8. Placing them in descending order gives “98”. Both 89 and 98 have the same score 145, but 98 is larger, so it is chosen.

**Example 3**  
Input: num = 1, sum = 10  
Output: ""  

Trace: A one-digit number cannot have digit sum 10, because the largest possible digit is 9. The impossibility check returns the empty string.

## How to Use / Run Locally

**C++**  
Save the code in a file named `main.cpp`. Compile with  
`g++ -std=c++17 main.cpp -o main`  
then run `./main`. You can add a simple driver that reads `num` and `sum` from standard input and prints the result.

**Java**  
Save the code inside a class named `Solution` in a file `Solution.java`. Compile with  
`javac Solution.java`  
and run with a small main method that creates an instance and calls the function.

**JavaScript**  
Copy the function into a file `solution.js`. Run it with Node.js:  
`node solution.js`  
(after adding a few console.log test calls).

**TypeScript**  
Save the code as `solution.ts`. Compile with  
`tsc solution.ts`  
then run the generated JavaScript with Node, or use `ts-node` directly.

**Python3**  
Save the code in `solution.py`. Run with  
`python3 solution.py`  
after adding a short test harness that prints the results for the sample inputs.

**Go**  
Save the function in a file `main.go` inside a package main. Build and run with  
`go run main.go`.

In every language the function signature already matches the LeetCode template, so you can also paste the code directly into the online judge.

## Notes & Optimizations

The solution is already optimal for the given constraints: both time and space are linear in `num`, which is the best we can hope for when we must output a string of that length.

One subtle point is that the greedy choice of “as many 9s as possible” is not only sufficient for the maximum score; it is the unique multiset of digits that achieves it. Consequently we never need a second pass to compare different candidates.

An alternative approach would be to place the remainder first and then the 9s, but that would produce a smaller number and violate the problem statement. Keeping the digits in strictly descending order is therefore both necessary and sufficient.

Because `num` can be as large as 200 000, any solution that allocates extra arrays of size proportional to `num` would still be acceptable, yet the pure string-building method uses the absolute minimum extra memory.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)