# Maximum Number of Items From Sale I - LeetCode 3946 Solution

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

You are given a list of items. Each item has two values: a factor and a price. You also get a budget.

You can buy as many copies of any item as you want, as long as the total money spent stays within the budget. After you buy at least one copy of an item, you receive free copies of other items whose factors are divisible by the factor of the item you just bought.

The goal is to find the maximum total number of item copies you can collect. This total includes both the items you actually paid for and the free copies you received.

This problem combines a classic knapsack-style decision with a free-item unlocking rule based on factors, making it a good example of dynamic programming on competitive programming platforms like LeetCode.

## Constraints

- Number of items: 1 <= items.length <= 1000
- Each item is given as [factor, price]
- 1 <= factor, price <= 1500
- 1 <= budget <= 1500

These limits allow an efficient O(n * budget) solution that runs comfortably within typical time limits.

## Intuition

The key observation is that buying the first copy of any item is special. That single purchase unlocks free copies of every item whose factor is divisible by the current item’s factor. Every extra copy of the same item after that only gives you one additional item and no extra freebies.

Because the budget is small (only 1500), a dynamic programming approach over the budget becomes practical. I can precompute, for every possible factor value, how many items would be unlocked by buying an item with that factor. Once I have that information, each item can be treated as a special first purchase followed by unlimited extra purchases.

## Approach

I start by building a simple array that tells me, for every possible factor f, how many items have a factor divisible by f. Call this the mult array.

I keep a DP array where dp[b] stores the maximum number of items I can obtain while spending exactly b money.

I process the items one by one. For each item I create a temporary copy of the current DP. Using the old values I try adding the first purchase of the current item: I pay its price and add the precomputed mult value (the bought copy plus all free copies). After that I treat the same item as unbounded and keep adding extra copies, each giving +1 item.

Once both updates are done I replace the main DP with the temporary version and move to the next item. At the end the answer is simply the largest value present in the DP array for any spending amount up to the given budget.

This way every possible combination of activations and extra purchases is considered without exceeding the time limits.

## Data Structures Used

- Integer array of size 1501 for the mult counts. It stores how many items are unlocked by each possible factor.
- Integer array (or vector) of size budget + 1 for the DP values. It tracks the best result for every possible spending amount.
- A temporary copy of the DP array is created for each item so that the first-purchase decision does not interfere with later unbounded updates of the same item.

These structures are chosen because the factor and budget ranges are tiny, so fixed-size arrays are both fast and memory-efficient.

## Operations & Behavior Summary

1. Precompute the mult array by checking, for every item and every possible factor, whether the item’s factor is divisible by that value.
2. Initialize a DP array of zeros. Zero items are always achievable with any budget.
3. For each item:
   - Copy the current DP into a temporary array.
   - Run a backward loop to consider the first purchase and add the corresponding mult value.
   - Run a forward loop to add as many extra copies as the remaining budget allows.
   - Replace the main DP with the temporary array.
4. After processing all items, scan the DP array and return the maximum value.

The separation of the first-purchase loop and the extra-copy loop guarantees that each item is activated at most once while still allowing unlimited extra copies after activation.

## Complexity

| Type              | Value              | Explanation                                                                 |
|-------------------|--------------------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n * B + n * F)   | n = number of items (≤ 1000), B = budget (≤ 1500), F = max factor (≤ 1500). Precomputation is O(n * F) and the DP runs in O(n * B). |
| Space Complexity  | O(B)               | Only the DP arrays of size budget + 1 are needed. The mult array is constant size under the given constraints. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int maximumSaleItems(vector<vector<int>>& items, int budget) {
        int mult[1501] = {};
        for (auto& it : items) {
            int fac = it[0];
            for (int f = 1; f <= 1500; ++f) {
                if (fac % f == 0) ++mult[f];
            }
        }
        vector<int> dp(budget + 1, 0);
        for (auto& it : items) {
            int f = it[0], p = it[1];
            vector<int> ndp = dp;
            for (int b = budget; b >= p; --b) {
                ndp[b] = max(ndp[b], dp[b - p] + mult[f]);
            }
            for (int b = p; b <= budget; ++b) {
                ndp[b] = max(ndp[b], ndp[b - p] + 1);
            }
            dp = move(ndp);
        }
        return *max_element(dp.begin(), dp.end());
    }
};
```

### Java
```java
class Solution {
    public int maximumSaleItems(int[][] items, int budget) {
        int[] mult = new int[1501];
        for (int[] it : items) {
            int fac = it[0];
            for (int f = 1; f <= 1500; ++f) {
                if (fac % f == 0) ++mult[f];
            }
        }
        int[] dp = new int[budget + 1];
        for (int[] it : items) {
            int f = it[0], p = it[1];
            int[] ndp = dp.clone();
            for (int b = budget; b >= p; --b) {
                ndp[b] = Math.max(ndp[b], dp[b - p] + mult[f]);
            }
            for (int b = p; b <= budget; ++b) {
                ndp[b] = Math.max(ndp[b], ndp[b - p] + 1);
            }
            dp = ndp;
        }
        int ans = 0;
        for (int v : dp) ans = Math.max(ans, v);
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[][]} items
 * @param {number} budget
 * @return {number}
 */
var maximumSaleItems = function(items, budget) {
    const mult = new Array(1501).fill(0);
    for (const it of items) {
        const fac = it[0];
        for (let f = 1; f <= 1500; ++f) {
            if (fac % f === 0) ++mult[f];
        }
    }
    let dp = new Array(budget + 1).fill(0);
    for (const it of items) {
        const f = it[0], p = it[1];
        const ndp = dp.slice();
        for (let b = budget; b >= p; --b) {
            ndp[b] = Math.max(ndp[b], dp[b - p] + mult[f]);
        }
        for (let b = p; b <= budget; ++b) {
            ndp[b] = Math.max(ndp[b], ndp[b - p] + 1);
        }
        dp = ndp;
    }
    return Math.max(...dp);
};
```

### TypeScript
```typescript
function maximumSaleItems(items: number[][], budget: number): number {
    const mult = new Array(1501).fill(0);
    for (const it of items) {
        const fac = it[0];
        for (let f = 1; f <= 1500; ++f) {
            if (fac % f === 0) ++mult[f];
        }
    }
    let dp = new Array(budget + 1).fill(0);
    for (const it of items) {
        const f = it[0], p = it[1];
        const ndp = dp.slice();
        for (let b = budget; b >= p; --b) {
            ndp[b] = Math.max(ndp[b], dp[b - p] + mult[f]);
        }
        for (let b = p; b <= budget; ++b) {
            ndp[b] = Math.max(ndp[b], ndp[b - p] + 1);
        }
        dp = ndp;
    }
    return Math.max(...dp);
};
```

### Python3
```python
class Solution:
    def maximumSaleItems(self, items: List[List[int]], budget: int) -> int:
        mult = [0] * 1501
        for it in items:
            fac = it[0]
            for f in range(1, 1501):
                if fac % f == 0:
                    mult[f] += 1
        dp = [0] * (budget + 1)
        for it in items:
            f, p = it[0], it[1]
            ndp = dp[:]
            for b in range(budget, p - 1, -1):
                ndp[b] = max(ndp[b], dp[b - p] + mult[f])
            for b in range(p, budget + 1):
                ndp[b] = max(ndp[b], ndp[b - p] + 1)
            dp = ndp
        return max(dp)
```

### Go
```go
func maximumSaleItems(items [][]int, budget int) int {
    mult := make([]int, 1501)
    for _, it := range items {
        fac := it[0]
        for f := 1; f <= 1500; f++ {
            if fac%f == 0 {
                mult[f]++
            }
        }
    }
    dp := make([]int, budget+1)
    for _, it := range items {
        f, p := it[0], it[1]
        ndp := make([]int, budget+1)
        copy(ndp, dp)
        for b := budget; b >= p; b-- {
            if dp[b-p]+mult[f] > ndp[b] {
                ndp[b] = dp[b-p] + mult[f]
            }
        }
        for b := p; b <= budget; b++ {
            if ndp[b-p]+1 > ndp[b] {
                ndp[b] = ndp[b-p] + 1
            }
        }
        dp = ndp
    }
    ans := 0
    for _, v := range dp {
        if v > ans {
            ans = v
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages. Only the syntax for arrays, loops, and maximum operations changes.

First I allocate a mult array of size 1501 and fill it with zeros. Then I walk through every item. For its factor I test every candidate f from 1 to 1500. Whenever the factor is divisible by f I increment mult[f]. After this pass mult[f] correctly holds the number of free copies plus the purchased copy that will be obtained when any item with that factor is activated for the first time.

I create a DP array of size budget + 1 and set every entry to zero. This represents the fact that spending any amount of money can still result in zero items if I choose to buy nothing.

For every item I make a fresh copy of the current DP. The copy is essential: the decision to buy the first copy of the current item must be based on states that have never seen this item before. I run a backward loop from the full budget down to the item’s price. For each possible remaining budget I look at the old DP value for (budget - price) and add mult[factor]. I keep the maximum of that candidate and whatever is already stored in the temporary array.

After the first-purchase pass I switch to a forward loop. Starting from the item’s price and going up to the full budget I repeatedly try adding one extra copy. Because the loop runs forward, each newly written value can be reused, allowing any number of extra copies. Each extra copy contributes exactly +1 to the item count and costs the same price again.

When both loops finish I replace the main DP with the temporary array and continue to the next item. After the last item has been processed I simply scan the entire DP array and return the highest value found. That value is the maximum number of items obtainable under the given budget.

Edge cases such as budget equal to zero, a single item, or items whose factor unlocks no extra free copies are handled automatically by the same loops; no special branches are required.

## Examples

**Example 1**

Input: items = [[6,2],[2,6],[3,4]], budget = 9  
Output: 4

I can buy two copies of the first item and one copy of the third item for a total cost of 8. The third item unlocks one free copy of the first item, giving a total of four items.

**Example 2**

Input: items = [[2,4],[3,2],[4,1],[6,4],[12,4]], budget = 8  
Output: 10

Buying one copy of the first item, one copy of the second item and two copies of the third item costs exactly 8. These three activations together unlock six free copies, producing a total of ten items.

**Example 3 (edge case)**

Input: items = [[5,10]], budget = 9  
Output: 0

The only item costs more than the budget, so the best answer is zero.

## How to Use / Run Locally

**C++**  
Save the code in a file named main.cpp. Compile with g++ -std=c++17 main.cpp -o main and run ./main. You can hard-code the input arrays inside the main function or read them from standard input.

**Java**  
Save the code in a file named Solution.java. Compile with javac Solution.java and run java Solution. Add a main method that creates an instance and calls the method with sample inputs.

**JavaScript**  
Save the code in a file named solution.js. Run it with node solution.js. You can test the function by calling it directly with sample arrays at the bottom of the file.

**TypeScript**  
Save the code in a file named solution.ts. Compile with tsc solution.ts and run the resulting JavaScript with node, or use ts-node solution.ts if you have it installed.

**Python3**  
Save the code in a file named solution.py. Run it with python3 solution.py. Add a few test calls at the bottom of the file to verify the output.

**Go**  
Save the code in a file named main.go. Run it with go run main.go. Add a main function that prepares the input slices and prints the result of the function call.

In every language you can replace the sample inputs with your own test cases to explore different scenarios.

## Notes & Optimizations

The solution already runs in optimal time for the given constraints. Because both the factor range and the budget are capped at 1500, the simple nested loops are fast enough and easy to understand.

One possible alternative is to group items that share the same factor and process each unique factor only once. That can reduce the constant factor slightly when many items have identical factors, but the asymptotic complexity stays the same.

Be careful with the order of the two DP updates for each item. Performing the unbounded update before the first-purchase update (or using the same array for both) can cause an item to be activated more than once, producing incorrect results.

The approach works for the empty-item or zero-budget cases without any extra checks, which keeps the code clean.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
