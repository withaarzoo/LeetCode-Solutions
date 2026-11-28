# Maximum Number of Items From Sale II - LeetCode 3947 Solution

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

You are given a list of items. Each item has a factor and a price. You also get a budget. You can buy as many copies of any item as you want, as long as the total money spent stays within the budget.

After you buy items, you can receive free copies under special rules. Every purchased copy of an item can unlock at most one free copy of another item, but only if the factor of the purchased item divides the factor of the free item. The same ordered pair of items can give you a free copy only once, no matter how many times you buy the first item. Different items can still give free copies of the same target item.

The goal is to return the maximum total number of item copies you can end up with (purchased plus free) without spending more than the given budget.

This is a classic competitive programming problem that mixes greedy selection with number-theoretic counting of multiples. It appears in weekly contests and tests how well you can turn a free-item rule into an efficient marginal-value calculation.

## Constraints

- 1 <= items.length <= 10^5
- items[i] = [factor_i, price_i]
- 1 <= factor_i <= items.length
- 1 <= price_i <= 10^9
- 1 <= budget <= 10^9

## Intuition

The free-copy rule looks complicated at first, but it quickly simplifies. For any item the maximum number of free copies it can ever produce is fixed. That number equals how many other items have a factor that is a multiple of its own factor. Once you know this “gain” value for every item, the problem becomes a pure budget-allocation task.

Each of the first `gain` purchases of an item gives you two copies in total (the one you paid for plus one free). Every extra purchase beyond that gain gives only one copy. This marginal-value view immediately suggests a greedy strategy: grab the cheapest double-value purchases first, then fill whatever money is left with the absolute cheapest single-value purchases.

Because the budget and the number of items are both large, any DP or expansion of every possible free copy would be too slow. Working with whole batches of the same price keeps the solution fast enough for the constraints.

## Approach

1. Count how many items share each possible factor value.  
2. For every factor, count how many items have a multiple of that factor. Subtract one to remove the item itself; the result is the gain of that item.  
3. While scanning the items, also record the global minimum price. This will be the cost of every single-value copy later.  
4. Collect every item that has a positive gain as a (price, gain) pair.  
5. Sort those pairs by increasing price.  
6. Walk through the sorted list and buy as many double-value copies as the remaining budget allows. Stop as soon as the price exceeds twice the minimum price (two minimum-price singles would be cheaper or equal).  
7. Spend every leftover unit of budget on minimum-price singles.  
8. Return the total number of copies obtained.

The same logic works in every language; only the syntax for arrays, sorting and 64-bit arithmetic changes.

## Data Structures Used

- Frequency array of size n+1 – stores how many times each factor appears.  
- Multiple-count array of size n+1 – stores, for every factor, the total number of items whose factor is a multiple of it.  
- List / vector of (price, gain) pairs – holds only the candidates that can produce free copies.  
- Integer variables for the global minimum price, spent money and running total – ordinary scalars, kept in 64-bit form to avoid overflow.

These structures are chosen because they give O(1) lookups after an O(n log n) pre-computation and need only linear extra memory.

## Operations & Behavior Summary

- Build factor frequencies.  
- Build multiple counts by iterating over every factor and stepping through its multiples.  
- Extract the minimum price and the list of positive-gain items.  
- Sort the candidate list by price.  
- Greedily consume double-value batches while the price stays attractive.  
- Convert remaining budget into single-value copies using the cheapest item.  
- Cast the final 64-bit total back to a 32-bit integer (safe under the given constraints).

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time | O(n log n) | Building the multiple counts costs O(n log n) because of the harmonic series of multiples. Sorting the candidate pairs costs another O(n log n). Everything else is linear. |
| Space | O(n) | Two arrays of size n+1 plus a list that holds at most n pairs. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int maximumSaleItems(vector<vector<int>>& items, int budget) {
        int n = items.size();
        vector<int> cnt(n + 1, 0);
        for (auto& it : items) cnt[it[0]]++;
        vector<int> mult(n + 1, 0);
        for (int f = 1; f <= n; f++) {
            for (int m = f; m <= n; m += f) mult[f] += cnt[m];
        }
        int minP = INT_MAX;
        vector<pair<int, int>> cand;
        for (auto& it : items) {
            int f = it[0], p = it[1];
            minP = min(minP, p);
            int g = mult[f] - 1;
            if (g > 0) cand.emplace_back(p, g);
        }
        sort(cand.begin(), cand.end());
        long long spent = 0, total = 0, B = budget;
        for (auto& [p, g] : cand) {
            if (p > 2LL * minP) break;
            long long can = min((long long)g, (B - spent) / p);
            total += can * 2;
            spent += can * p;
        }
        total += (B - spent) / minP;
        return (int)total;
    }
};
```

### Java
```java
class Solution {
    public int maximumSaleItems(int[][] items, int budget) {
        int n = items.length;
        int[] cnt = new int[n + 1];
        for (int[] it : items) cnt[it[0]]++;
        int[] mult = new int[n + 1];
        for (int f = 1; f <= n; f++) {
            for (int m = f; m <= n; m += f) mult[f] += cnt[m];
        }
        int minP = Integer.MAX_VALUE;
        List<int[]> cand = new ArrayList<>();
        for (int[] it : items) {
            int f = it[0], p = it[1];
            minP = Math.min(minP, p);
            int g = mult[f] - 1;
            if (g > 0) cand.add(new int[]{p, g});
        }
        cand.sort((a, b) -> Integer.compare(a[0], b[0]));
        long spent = 0, total = 0, B = budget;
        for (int[] c : cand) {
            int p = c[0], g = c[1];
            if (p > 2L * minP) break;
            long can = Math.min((long)g, (B - spent) / p);
            total += can * 2;
            spent += can * p;
        }
        total += (B - spent) / minP;
        return (int)total;
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
    const n = items.length;
    const cnt = new Array(n + 1).fill(0);
    for (const it of items) cnt[it[0]]++;
    const mult = new Array(n + 1).fill(0);
    for (let f = 1; f <= n; f++) {
        for (let m = f; m <= n; m += f) mult[f] += cnt[m];
    }
    let minP = Infinity;
    const cand = [];
    for (const it of items) {
        const f = it[0], p = it[1];
        minP = Math.min(minP, p);
        const g = mult[f] - 1;
        if (g > 0) cand.push([p, g]);
    }
    cand.sort((a, b) => a[0] - b[0]);
    let spent = 0, total = 0, B = budget;
    for (const [p, g] of cand) {
        if (p > 2 * minP) break;
        const can = Math.min(g, Math.floor((B - spent) / p));
        total += can * 2;
        spent += can * p;
    }
    total += Math.floor((B - spent) / minP);
    return total;
};
```

### TypeScript
```typescript
function maximumSaleItems(items: number[][], budget: number): number {
    const n = items.length;
    const cnt = new Array(n + 1).fill(0);
    for (const it of items) cnt[it[0]]++;
    const mult = new Array(n + 1).fill(0);
    for (let f = 1; f <= n; f++) {
        for (let m = f; m <= n; m += f) mult[f] += cnt[m];
    }
    let minP = Infinity;
    const cand: number[][] = [];
    for (const it of items) {
        const f = it[0], p = it[1];
        minP = Math.min(minP, p);
        const g = mult[f] - 1;
        if (g > 0) cand.push([p, g]);
    }
    cand.sort((a, b) => a[0] - b[0]);
    let spent = 0, total = 0, B = budget;
    for (const [p, g] of cand) {
        if (p > 2 * minP) break;
        const can = Math.min(g, Math.floor((B - spent) / p));
        total += can * 2;
        spent += can * p;
    }
    total += Math.floor((B - spent) / minP);
    return total;
};
```

### Python3
```python
class Solution:
    def maximumSaleItems(self, items: List[List[int]], budget: int) -> int:
        n = len(items)
        cnt = [0] * (n + 1)
        for f, p in items:
            cnt[f] += 1
        mult = [0] * (n + 1)
        for f in range(1, n + 1):
            for m in range(f, n + 1, f):
                mult[f] += cnt[m]
        min_p = min(p for f, p in items)
        cand = []
        for f, p in items:
            g = mult[f] - 1
            if g > 0:
                cand.append((p, g))
        cand.sort()
        spent = 0
        total = 0
        B = budget
        for p, g in cand:
            if p > 2 * min_p:
                break
            can = min(g, (B - spent) // p)
            total += can * 2
            spent += can * p
        total += (B - spent) // min_p
        return total
```

### Go
```go
func maximumSaleItems(items [][]int, budget int) int {
    n := len(items)
    cnt := make([]int, n+1)
    for _, it := range items {
        cnt[it[0]]++
    }
    mult := make([]int, n+1)
    for f := 1; f <= n; f++ {
        for m := f; m <= n; m += f {
            mult[f] += cnt[m]
        }
    }
    minP := int(^uint(0) >> 1)
    type pair struct{ p, g int }
    cand := make([]pair, 0)
    for _, it := range items {
        f, p := it[0], it[1]
        if p < minP {
            minP = p
        }
        g := mult[f] - 1
        if g > 0 {
            cand = append(cand, pair{p, g})
        }
    }
    sort.Slice(cand, func(i, j int) bool { return cand[i].p < cand[j].p })
    spent, total, B := int64(0), int64(0), int64(budget)
    for _, c := range cand {
        if int64(c.p) > 2*int64(minP) {
            break
        }
        can := int64(c.g)
        if rem := (B - spent) / int64(c.p); rem < can {
            can = rem
        }
        total += can * 2
        spent += can * int64(c.p)
    }
    total += (B - spent) / int64(minP)
    return int(total)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core reasoning is identical in every language. Only the way we declare arrays, sort them and handle large integers differs.

First we allocate a frequency array of length n+1 and fill it by walking through the input. Because every factor is guaranteed to lie between 1 and n, no bounds checks are required.

Next we allocate a second array of the same size. For every possible factor f we start at f and keep adding f until we pass n, accumulating the frequency of each multiple. After this nested loop finishes, the entry for f already contains the total number of items whose factor is a multiple of f. Subtracting one later removes the item itself and yields the exact gain.

While we scan the items a second time we keep the smallest price seen so far and, for every item whose gain is positive, we push a (price, gain) pair into a candidate list. Sorting that list by price puts the cheapest double-value opportunities first.

We then walk the sorted list. For each pair we compute how many of its gain copies still fit into the remaining budget. We add twice that many to the answer and increase the spent amount. The moment the current price exceeds twice the global minimum price we stop; any further double would cost more than two ordinary singles.

Finally the leftover money is divided by the minimum price. Each of those purchases adds exactly one copy, so we simply add the quotient to the total and return it.

In languages that do not promote integers automatically (C++, Java, Go) we keep the spent and total variables in 64-bit form so that budget * 2 never overflows. In JavaScript and TypeScript the native Number type already has enough precision for the given constraints. Python integers grow automatically, so no extra care is needed.

Edge cases such as a zero-gain item, a budget smaller than the cheapest price, or many items sharing the same factor are all handled naturally by the same loops; no special branches are required.

## Examples

**Example 1**  
Input: items = [[1,6],[2,4],[3,5]], budget = 19  
Output: 5  

Trace:  
- Gains are 2, 0, 0.  
- Minimum price is 4.  
- We can buy two copies of the first item (cost 12) and one copy of the second item (cost 4). Total spent 16.  
- The two purchases of the first item unlock two free copies.  
- Total items = 3 purchased + 2 free = 5.

**Example 2**  
Input: items = [[2,8],[1,10],[6,6],[4,12],[5,20],[5,17]], budget = 35  
Output: 7  

Trace:  
- Gains become 2, 5, 0, 0, 1, 1.  
- Minimum price is 6.  
- Cheapest useful doubles are taken first. After spending 32 we still have 3 left, which is not enough for another item.  
- Four purchased copies plus three free copies give the answer 7.

**Example 3 (edge)**  
Input: items = [[1,100]], budget = 50  
Output: 0  

Trace:  
- The single item costs more than the budget, so nothing can be bought and the answer stays zero.

## How to Use / Run Locally

**C++**  
Save the code in a file named `main.cpp`. Compile with  
`g++ -std=c++17 main.cpp -o main`  
Run with  
`./main`  
(You will need to add a small driver that reads the input and calls the function.)

**Java**  
Save the code in `Solution.java`. Compile with  
`javac Solution.java`  
Run with  
`java Solution`  
(Again add a main method that feeds sample input.)

**JavaScript**  
Save the code in `solution.js`. Run with  
`node solution.js`  
(Wrap the function in a short test harness.)

**TypeScript**  
Save the code in `solution.ts`. Compile with  
`tsc solution.ts`  
then run the generated JavaScript, or use `ts-node solution.ts` if you have it installed.

**Python3**  
Save the code in `solution.py`. Run with  
`python3 solution.py`  
(Add a few lines that create the input lists and print the result.)

**Go**  
Save the code in `main.go`. Run with  
`go run main.go`  
(You need a main function that prepares the slice and prints the answer.)

In every language the function signature already matches the LeetCode judge, so you can also paste the code directly into the online editor.

## Notes & Optimizations

- The harmonic-loop that builds the multiple counts is the classic O(n log n) sieve-style technique; it is already optimal for the given constraints.  
- Filtering candidates whose price is greater than twice the minimum price is a small but useful pruning step that avoids unnecessary work.  
- If the minimum price itself has a positive gain, those gain copies are automatically taken as doubles because the candidate list is sorted. Any further copies of the same item become ordinary singles, which is exactly what we want.  
- An alternative approach that expands every double into a separate unit would be correct but would time out when gains are large. Batch processing is therefore essential.  
- The solution never materializes the free items themselves; it only counts them. This keeps both time and memory linear.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
