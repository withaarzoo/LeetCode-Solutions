# Minimum Cost to Acquire Required Items - LeetCode 3789 Solution

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

You are given five integers: cost1, cost2, costBoth, need1, and need2.

There are three kinds of items you can buy:

- Type 1 costs cost1 and gives you 1 unit toward the type-1 requirement only.
- Type 2 costs cost2 and gives you 1 unit toward the type-2 requirement only.
- Type 3 costs costBoth and gives you 1 unit toward both requirements at the same time.

Your goal is to collect enough items so that the total units for type 1 are at least need1 and the total units for type 2 are at least need2. Return the smallest possible total cost that achieves this.

This is a classic minimum cost problem that appears in coding interviews and competitive programming contests. The challenge is to decide the best mix of the three item types without exceeding time limits when the needs become very large.

## Constraints

- 1 <= cost1, cost2, costBoth <= 10^6
- 0 <= need1, need2 <= 10^9

Because the needs can reach one billion, any solution that loops over the actual counts will be too slow. The algorithm must run in constant time.

## Intuition

The first useful observation is that type-3 items are only worth buying when they are cheaper than buying one type-1 and one type-2 separately.  

Once that decision is clear, the problem splits into two parts: the overlapping portion of the two needs, and whatever is left over on the larger side.  

You never need more type-3 items than the maximum of the two needs, and buying extra beyond that only adds cost. So the optimal count of type-3 items will always sit at one of a few critical points: zero, the minimum of the two needs, or the maximum of the two needs.  

This insight removes the need to try every possible count and lets us compute the answer with a handful of min and multiplication operations.

## Approach

1. Find the smaller of the two needs. Call this value pairs. These pairs can be covered either by type-3 items or by separate type-1 and type-2 items.
2. For each of those pairs, pay the cheaper of (costBoth) or (cost1 + cost2).
3. After the pairs are covered, one need is fully satisfied. The other need may still have a remainder equal to the difference between the two needs.
4. For every remaining unit, again choose the cheaper option: the pure single-item cost or another type-3 item.
5. Add the cost of the pairs and the cost of the remainders. That sum is the minimum total cost.

When both needs are zero the answer is simply zero. The same formulas handle this case without any special checks.

## Data Structures Used

No extra data structures are required.  

All work is done with a few ordinary integer variables. This keeps both time and memory usage constant and makes the solution easy to implement in any language.

## Operations & Behavior Summary

- Convert the five input values to 64-bit integers so that the later multiplications cannot overflow.
- Compute pairs = min(need1, need2).
- Calculate the cost of covering those pairs using the cheaper of costBoth and cost1 + cost2.
- Compute the two remainders: rem1 = need1 - pairs and rem2 = need2 - pairs (one of them will be zero).
- For each remainder, multiply by the cheaper of the corresponding single cost and costBoth.
- Return the sum of the pair cost and the remainder cost.

The algorithm never iterates over the actual need values, so it stays fast even when need1 or need2 equals one billion.

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(1)  | Only a fixed number of arithmetic operations and comparisons are performed. The size of need1 and need2 never appears in a loop. |
| Space Complexity | O(1)  | A constant number of variables are used. No arrays, maps, or other growing structures are allocated. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    long long minimumCost(int cost1, int cost2, int costBoth, int need1, int need2) {
        long long c1 = cost1, c2 = cost2, cb = costBoth;
        long long n1 = need1, n2 = need2;
        long long pairs = min(n1, n2);
        long long costPairs = pairs * min(cb, c1 + c2);
        long long rem1 = n1 - pairs;
        long long rem2 = n2 - pairs;
        long long costRem = rem1 * min(c1, cb) + rem2 * min(c2, cb);
        return costPairs + costRem;
    }
};
```

### Java
```java
class Solution {
    public long minimumCost(int cost1, int cost2, int costBoth, int need1, int need2) {
        long c1 = cost1, c2 = cost2, cb = costBoth;
        long n1 = need1, n2 = need2;
        long pairs = Math.min(n1, n2);
        long costPairs = pairs * Math.min(cb, c1 + c2);
        long rem1 = n1 - pairs;
        long rem2 = n2 - pairs;
        long costRem = rem1 * Math.min(c1, cb) + rem2 * Math.min(c2, cb);
        return costPairs + costRem;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number} cost1
 * @param {number} cost2
 * @param {number} costBoth
 * @param {number} need1
 * @param {number} need2
 * @return {number}
 */
var minimumCost = function(cost1, cost2, costBoth, need1, need2) {
    let c1 = cost1, c2 = cost2, cb = costBoth;
    let n1 = need1, n2 = need2;
    let pairs = Math.min(n1, n2);
    let costPairs = pairs * Math.min(cb, c1 + c2);
    let rem1 = n1 - pairs;
    let rem2 = n2 - pairs;
    let costRem = rem1 * Math.min(c1, cb) + rem2 * Math.min(c2, cb);
    return costPairs + costRem;
};
```

### TypeScript
```typescript
function minimumCost(cost1: number, cost2: number, costBoth: number, need1: number, need2: number): number {
    let c1 = cost1, c2 = cost2, cb = costBoth;
    let n1 = need1, n2 = need2;
    let pairs = Math.min(n1, n2);
    let costPairs = pairs * Math.min(cb, c1 + c2);
    let rem1 = n1 - pairs;
    let rem2 = n2 - pairs;
    let costRem = rem1 * Math.min(c1, cb) + rem2 * Math.min(c2, cb);
    return costPairs + costRem;
};
```

### Python3
```python
class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        c1, c2, cb = cost1, cost2, costBoth
        n1, n2 = need1, need2
        pairs = min(n1, n2)
        cost_pairs = pairs * min(cb, c1 + c2)
        rem1 = n1 - pairs
        rem2 = n2 - pairs
        cost_rem = rem1 * min(c1, cb) + rem2 * min(c2, cb)
        return cost_pairs + cost_rem
```

### Go
```go
func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
    c1, c2, cb := int64(cost1), int64(cost2), int64(costBoth)
    n1, n2 := int64(need1), int64(need2)
    pairs := n1
    if n2 < pairs {
        pairs = n2
    }
    costPairs := pairs * min(cb, c1+c2)
    rem1 := n1 - pairs
    rem2 := n2 - pairs
    costRem := rem1*min(c1, cb) + rem2*min(c2, cb)
    return costPairs + costRem
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages. The only differences are syntax and the way large integers are handled.

First the five parameters are promoted to 64-bit integers (long long in C++, long in Java, Number in JavaScript/TypeScript, unlimited int in Python, int64 in Go). This prevents overflow when a need of 10^9 is multiplied by a cost of 10^6.

Next the smaller of the two needs is stored in a variable called pairs. These pairs represent the units that can be covered by a single type-3 item or by one type-1 plus one type-2 item. The cost of covering them is pairs multiplied by the minimum of costBoth and cost1 + cost2.

After the pairs are paid for, the smaller need is completely finished. The larger need still has a leftover equal to the absolute difference of the two original needs. That leftover is stored in two variables (rem1 and rem2); one of them will be zero.

For each leftover unit the algorithm again picks the cheaper price: the pure single-item cost or another type-3 item. The two remainder costs are added together.

Finally the pair cost and the remainder cost are summed and returned.  

Edge cases are handled automatically. When both needs are zero, pairs is zero and both remainders are zero, so the result is zero. When one need is zero the algorithm simply buys the required number of the other single item (or type-3 items if they happen to be cheaper).

Because the same sequence of min and multiplication steps is used in every language, the numerical result is identical no matter which implementation you run.

## Examples

**Example 1**

Input: cost1 = 3, cost2 = 2, costBoth = 1, need1 = 3, need2 = 2  
Output: 3  

Trace:  
pairs = min(3, 2) = 2  
cost of pairs = 2 * min(1, 3 + 2) = 2 * 1 = 2  
remainders: rem1 = 1, rem2 = 0  
cost of remainders = 1 * min(3, 1) + 0 = 1  
total = 2 + 1 = 3  

Buying three type-3 items also works and costs the same, confirming the answer is optimal.

**Example 2**

Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 2, need2 = 3  
Output: 22  

Trace:  
pairs = min(2, 3) = 2  
cost of pairs = 2 * min(15, 5 + 4) = 2 * 9 = 18  
remainders: rem1 = 0, rem2 = 1  
cost of remainders = 0 + 1 * min(4, 15) = 4  
total = 18 + 4 = 22  

Here type-3 items are more expensive than buying the singles separately, so the algorithm correctly avoids them for the pairs and also for the leftover.

**Example 3**

Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 0, need2 = 0  
Output: 0  

Trace:  
pairs = 0  
everything multiplies by zero, so the total is 0.

## How to Use / Run Locally

**C++**  
Save the code in a file named solution.cpp.  
Compile with: g++ -std=c++17 solution.cpp -o solution  
Run with: ./solution  

**Java**  
Save the code in a file named Solution.java.  
Compile with: javac Solution.java  
Run with: java Solution  

**JavaScript**  
Save the code in a file named solution.js.  
Run with: node solution.js  

**TypeScript**  
Save the code in a file named solution.ts.  
Compile with: tsc solution.ts  
Run the generated JavaScript with: node solution.js  

**Python3**  
Save the code in a file named solution.py.  
Run with: python3 solution.py  

**Go**  
Save the code in a file named solution.go.  
Run with: go run solution.go  

In every language you will need to add a small main or driver function that reads the five integers and prints the result if you want to test from the command line. The core function itself is ready to drop into an online judge.

## Notes & Optimizations

The solution is already optimal in both time and space. No further asymptotic improvement is possible because the answer depends on every input value.

One alternative approach is to iterate over the possible number of type-3 items from 0 to max(need1, need2). That works for small needs but fails the time limit when the needs reach 10^9. The constant-time method shown here avoids that problem completely.

Another possible formulation is to compute the cost for three candidate values of type-3 items (0, min(need1, need2), max(need1, need2)) and take the minimum among them. The formulas used above produce the same result with fewer operations.

Watch out for integer overflow. Always promote the inputs to 64-bit integers before any multiplication.

When costBoth is extremely cheap the algorithm correctly buys type-3 items for the leftovers as well. When costBoth is expensive it falls back to the pure single items. Both behaviors are covered by the same min expressions.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
