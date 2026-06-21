# 1833. Maximum Ice Cream Bars

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

LeetCode 1833: Maximum Ice Cream Bars is a classic greedy algorithm and counting sort problem.

You are given an array called `costs` where each element represents the price of an ice cream bar. You are also given a certain number of `coins`.

Your goal is to buy the maximum number of ice cream bars without spending more coins than you have.

The order of purchasing does not matter. To maximize the number of ice cream bars, you need to spend your coins wisely and always prioritize cheaper bars whenever possible.

### Input

* An integer array `costs`
* An integer `coins`

### Output

* The maximum number of ice cream bars that can be purchased

This problem is commonly asked in coding interviews because it combines greedy thinking with counting sort optimization.

---

## Constraints

| Constraint            | Value |
| --------------------- | ----- |
| costs.length == n     | True  |
| 1 ≤ n ≤ 100000        |       |
| 1 ≤ costs[i] ≤ 100000 |       |
| 1 ≤ coins ≤ 100000000 |       |

---

## Intuition

When I first looked at the problem, I noticed that the objective is not to maximize the amount spent. The objective is to maximize the number of ice cream bars purchased.

That immediately suggests buying the cheapest bars first.

If I spend my coins on expensive bars early, I may run out of money and miss the chance to buy several cheaper bars.

The problem also explicitly asks for a counting sort solution. Since every cost is limited to a maximum value of 100000, I can count how many bars exist at each price and process prices from smallest to largest.

This lets me avoid a traditional sorting algorithm and achieve better performance.

---

## Approach

1. Create a frequency array to count how many ice cream bars exist for every possible cost.
2. Traverse the `costs` array and populate the frequency array.
3. Start checking prices from the smallest possible cost.
4. For each price:

   * Determine how many bars of that price exist.
   * Calculate how many of them can be purchased with the remaining coins.
   * Buy as many as possible.
5. Reduce the remaining coins.
6. Keep track of the total number of bars purchased.
7. Return the final count.

This greedy strategy works because buying cheaper items first always leaves the most money available for future purchases.

---

## Data Structures Used

### Frequency Array

A frequency array stores how many ice cream bars exist for each possible price.

Why I used it:

* Fast counting of costs
* Avoids sorting the entire input
* Matches the counting sort requirement
* Provides efficient O(n + k) performance

Where:

* `n` = number of ice cream bars
* `k` = maximum possible cost value

---

## Operations & Behavior Summary

The algorithm performs the following major operations:

1. Read every ice cream cost.
2. Count occurrences of each cost.
3. Process prices from cheapest to most expensive.
4. Determine how many bars at each price can be afforded.
5. Purchase those bars.
6. Deduct spent coins.
7. Continue until all prices are processed.
8. Return the total number of purchased bars.

In simple terms:

* Count prices
* Buy cheap bars first
* Keep spending until coins run out
* Return the number of bars bought

---

## Complexity

| Metric           | Complexity | Explanation                                                       |
| ---------------- | ---------- | ----------------------------------------------------------------- |
| Time Complexity  | O(n + k)   | O(n) for counting costs and O(k) for scanning all possible prices |
| Space Complexity | O(k)       | Frequency array stores counts for all possible costs              |

Where:

* `n` = number of ice cream bars
* `k` = maximum possible cost value (100000)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxIceCream(vector<int>& costs, int coins) {
        // Maximum possible cost according to constraints
        const int MAX_COST = 100000;

        // Frequency array to count ice cream bars of each cost
        vector<int> freq(MAX_COST + 1, 0);

        // Count occurrences of every cost
        for (int cost : costs) {
            freq[cost]++;
        }

        // Stores total ice cream bars purchased
        int answer = 0;

        // Try buying from cheapest cost to most expensive
        for (int cost = 1; cost <= MAX_COST; cost++) {

            // Skip if no ice cream bar has this cost
            if (freq[cost] == 0) continue;

            // Maximum bars of this cost that can be afforded
            int canBuy = min(freq[cost], coins / cost);

            // Add purchased bars to answer
            answer += canBuy;

            // Deduct spent coins
            coins -= canBuy * cost;

            // No need to continue if coins are exhausted
            if (coins < cost) continue;
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    public int maxIceCream(int[] costs, int coins) {

        // Maximum possible cost according to constraints
        final int MAX_COST = 100000;

        // Frequency array to count ice cream bars of each cost
        int[] freq = new int[MAX_COST + 1];

        // Count occurrences of every cost
        for (int cost : costs) {
            freq[cost]++;
        }

        // Stores total ice cream bars purchased
        int answer = 0;

        // Process costs from smallest to largest
        for (int cost = 1; cost <= MAX_COST; cost++) {

            // Skip unavailable costs
            if (freq[cost] == 0) {
                continue;
            }

            // Maximum bars affordable at this cost
            int canBuy = Math.min(freq[cost], coins / cost);

            // Increase purchased count
            answer += canBuy;

            // Deduct spent coins
            coins -= canBuy * cost;
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} costs
 * @param {number} coins
 * @return {number}
 */
var maxIceCream = function(costs, coins) {

    // Maximum possible cost according to constraints
    const MAX_COST = 100000;

    // Frequency array to count occurrences of each cost
    const freq = new Array(MAX_COST + 1).fill(0);

    // Count every cost
    for (const cost of costs) {
        freq[cost]++;
    }

    // Total purchased ice cream bars
    let answer = 0;

    // Buy from cheapest to most expensive
    for (let cost = 1; cost <= MAX_COST; cost++) {

        // Skip if this cost does not exist
        if (freq[cost] === 0) continue;

        // Maximum bars affordable at current cost
        const canBuy = Math.min(freq[cost], Math.floor(coins / cost));

        // Add purchased bars
        answer += canBuy;

        // Remove spent coins
        coins -= canBuy * cost;
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def maxIceCream(self, costs: List[int], coins: int) -> int:
        
        # Maximum possible cost according to constraints
        MAX_COST = 100000

        # Frequency array to count occurrences of each cost
        freq = [0] * (MAX_COST + 1)

        # Count every ice cream cost
        for cost in costs:
            freq[cost] += 1

        # Stores total purchased bars
        answer = 0

        # Process costs from smallest to largest
        for cost in range(1, MAX_COST + 1):

            # Skip unavailable costs
            if freq[cost] == 0:
                continue

            # Maximum bars affordable at current cost
            can_buy = min(freq[cost], coins // cost)

            # Increase purchased count
            answer += can_buy

            # Deduct spent coins
            coins -= can_buy * cost

        return answer
```

### Go

```go
func maxIceCream(costs []int, coins int) int {

    // Maximum possible cost according to constraints
    const MAX_COST = 100000

    // Frequency array to count occurrences of each cost
    freq := make([]int, MAX_COST+1)

    // Count every ice cream cost
    for _, cost := range costs {
        freq[cost]++
    }

    // Stores total purchased bars
    answer := 0

    // Process costs from smallest to largest
    for cost := 1; cost <= MAX_COST; cost++ {

        // Skip unavailable costs
        if freq[cost] == 0 {
            continue
        }

        // Maximum bars affordable at current cost
        canBuy := freq[cost]
        affordable := coins / cost

        if canBuy > affordable {
            canBuy = affordable
        }

        // Add purchased bars
        answer += canBuy

        // Deduct spent coins
        coins -= canBuy * cost
    }

    return answer
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is identical across all five languages.

### Step 1: Create a Frequency Array

The first task is counting how many times each cost appears.

For example:

```text
costs = [1, 3, 2, 1]
```

The frequency array conceptually becomes:

```text
cost 1 → 2 bars
cost 2 → 1 bar
cost 3 → 1 bar
```

This gives instant access to the number of bars available at any price.

---

### Step 2: Process Costs in Increasing Order

The next step is to start from the smallest possible cost.

Why?

Because every coin matters.

If a bar costs 1 coin and another costs 10 coins, purchasing the cheaper one helps maximize the total number of bars.

This is the key greedy observation.

---

### Step 3: Determine Affordable Quantity

For each price:

```text
remaining coins ÷ current cost
```

tells us how many bars of that price can be purchased.

However, there may not be that many bars available.

So we purchase:

```text
minimum(
    available bars,
    affordable bars
)
```

This guarantees that we never overspend.

---

### Step 4: Update Answer

Whenever bars are purchased, the total count increases.

The answer simply tracks the total number of successful purchases throughout the process.

---

### Step 5: Deduct Coins

After buying bars:

```text
spent coins = purchased bars × current cost
```

The remaining coin balance is updated accordingly.

This ensures future purchases use the correct remaining budget.

---

### Step 6: Continue Until All Costs Are Processed

The algorithm continues moving through all possible prices from smallest to largest.

Since cheaper prices are handled first, the final count is guaranteed to be optimal.

---

### Why This Works

The greedy choice is always safe.

If I ever skip a cheaper bar and buy a more expensive one instead, I spend more money while purchasing the same number of bars.

That can never help maximize the total count.

Therefore, buying the cheapest available bars first always leads to the best answer.

---

## Examples

### Example 1

#### Input

```text
costs = [1,3,2,4,1]
coins = 7
```

#### Output

```text
4
```

#### Trace

Sorted purchasing order:

```text
1, 1, 2, 3, 4
```

Buy:

```text
1 + 1 + 2 + 3 = 7
```

Purchased bars:

```text
4
```

---

### Example 2

#### Input

```text
costs = [10,6,8,7,7,8]
coins = 5
```

#### Output

```text
0
```

#### Trace

The cheapest bar costs:

```text
6
```

Available coins:

```text
5
```

No purchase is possible.

---

### Example 3

#### Input

```text
costs = [1,6,3,1,2,5]
coins = 20
```

#### Output

```text
6
```

#### Trace

Total cost:

```text
1 + 6 + 3 + 1 + 2 + 5 = 18
```

Since:

```text
18 ≤ 20
```

all ice cream bars can be purchased.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

---

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

Build:

```bash
go build solution.go
```

---

## Notes & Optimizations

* This problem specifically requires a counting sort based solution.
* A traditional sorting solution would use a greedy algorithm after sorting.
* Sorting produces O(n log n) complexity.
* The counting sort approach improves this to O(n + k).
* Since the maximum cost value is fixed at 100000, the frequency array approach is very efficient.
* Always buying the cheapest available bar first is the critical greedy observation.
* The solution handles large inputs efficiently and works within all given constraints.
* If the cost range were much larger, sorting might become a more practical choice than maintaining a large frequency array.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
