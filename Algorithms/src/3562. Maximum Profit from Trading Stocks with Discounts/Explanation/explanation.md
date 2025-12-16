# Maximum Profit from Trading Stocks with Discounts

**LeetCode Problem 3562**

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given a company with `n` employees arranged in a hierarchy tree.
Each employee can buy **one stock today** and sell it **tomorrow**.

* `present[i]` → price to buy stock today
* `future[i]` → price to sell stock tomorrow
* `hierarchy[u, v]` → `u` is the direct boss of `v`
* Total money available is `budget`

### Special Discount Rule

If an employee’s **direct boss buys their own stock**, then the employee can buy their stock at **half price**.

My task is to calculate the **maximum profit** I can make **without exceeding the budget**.

---

## Constraints

* `1 ≤ n ≤ 160`
* `1 ≤ budget ≤ 160`
* `1 ≤ present[i], future[i] ≤ 50`
* `hierarchy` forms a tree (no cycles)
* Employee `1` is the CEO (root)
* Each stock can be bought **at most once**
* Profit from future selling **cannot be reused** to buy more stocks

---

## Intuition

When I read the problem, I immediately noticed two things:

1. The hierarchy forms a **tree**
2. The price of an employee depends on whether their **parent bought stock or not**

So for every employee, I must consider **two situations**:

* Parent did **not** buy → no discount
* Parent **did** buy → discount applies

Because the **budget is small**, I realized this is a **Tree + Knapsack Dynamic Programming** problem.

---

## Approach

I solve the problem using **DFS + DP**.

For every employee `u`, I maintain:

```
dp[u][0][b] → max profit using budget b, parent did NOT buy
dp[u][1][b] → max profit using budget b, parent DID buy
```

At each node, I consider two choices:

### 1. Skip buying this stock

* Spend `0` budget
* Children do NOT get discount
* Merge children using knapsack

### 2. Buy this stock

* Spend `present[u]` or `present[u]/2`
* Gain `future[u] - cost`
* Children DO get discount
* Merge children with discounted state

I take the **maximum** of these two choices for each budget.

---

## Data Structures Used

* **Adjacency List** → represent company hierarchy
* **3D DP Array** → `dp[node][parentBought][budget]`
* **DFS (Depth First Search)** → bottom-up processing
* **Knapsack Merge Arrays** → combine child results

---

## Operations & Behavior Summary

* DFS ensures children are processed first
* Knapsack merge ensures budget is respected
* Discount is applied **only if parent bought**
* Each node is evaluated independently and correctly

---

## Complexity

**Time Complexity:**
`O(n × budget²)`

* `n` nodes
* each merge is knapsack-based

**Space Complexity:**
`O(n × budget)`

* DP arrays per node

---

## Multi-language Solutions

### C++

```cpp
// See solution implementation in repository
```

### Java

```java
// See solution implementation in repository
```

### JavaScript

```javascript
// See solution implementation in repository
```

### Python3

```python
# See solution implementation in repository
```

### Go

```go
// See solution implementation in repository
```

> All implementations follow **the same logic**, only syntax differs.

---

## Step-by-step Detailed Explanation (All Languages)

1. Build the hierarchy tree
2. Start DFS from CEO (node 0)
3. For each node:

   * Solve all children first
   * Compute two DP states:

     * parentBought = 0
     * parentBought = 1
4. For each state:

   * Try **skip**
   * Try **buy**
   * Merge children using knapsack
5. Store the best result
6. Final answer = max of `dp[0][0][b]` for all `b ≤ budget`

---

## Examples

### Example 1

```
Input:
n = 2
present = [1,2]
future = [4,3]
hierarchy = [[1,2]]
budget = 3

Output:
5
```

Explanation:

* Buy employee 1 → profit 3
* Employee 2 gets discount → profit 2
* Total profit = 5

---

## How to use / Run locally

1. Clone the repository
2. Choose your language folder
3. Compile and run

### Example (C++)

```bash
g++ solution.cpp -o solution
./solution
```

### Example (Python)

```bash
python3 solution.py
```

---

## Notes & Optimizations

* Budget is small → knapsack DP is safe
* Tree structure avoids cycles
* No greedy approach works here
* Order of operations (buy vs merge) is **very important**

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
