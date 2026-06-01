# 2144. Minimum Cost of Buying Candies With Discount

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

The **Minimum Cost of Buying Candies With Discount** problem asks us to find the minimum amount of money required to buy all candies from a shop offering a special discount.

The discount rule is simple:

* For every two candies purchased, one additional candy can be taken for free.
* The free candy must have a price less than or equal to the cheaper of the two purchased candies.

We are given an integer array `cost`, where each element represents the price of a candy.

Our goal is to calculate the minimum total cost needed to obtain every candy while using the discount as efficiently as possible.

This is a classic **greedy algorithm** problem where making the best local choice leads to the optimal global result.

---

## Constraints

| Constraint                | Value              |
| ------------------------- | ------------------ |
| `1 <= cost.length <= 100` | Number of candies  |
| `1 <= cost[i] <= 100`     | Cost of each candy |

---

## Intuition

My first observation was that the discount becomes more valuable when the free candy is expensive.

Since I want to minimize the money I spend, I should try to make the most expensive candies possible become free.

After thinking about the discount rule, I realized that sorting the candy prices in descending order creates the best situation.

If the candies are arranged from highest price to lowest price, then every group of three candies naturally becomes:

* Pay for the first candy
* Pay for the second candy
* Get the third candy for free

This guarantees that the free candy is as expensive as possible while still satisfying the discount condition.

---

## Approach

1. Sort the candy prices in descending order.
2. Traverse the sorted array.
3. For every group of three candies:

   * Pay for the first candy.
   * Pay for the second candy.
   * Skip the third candy because it becomes free.
4. Continue until all candies are processed.
5. Return the total amount paid.

The key observation is that after sorting, every third candy contributes nothing to the final cost.

---

## Data Structures Used

### Array

The input array itself is enough for this solution.

Why I used it:

* Stores candy prices.
* Can be sorted efficiently.
* No additional complex data structure is needed.

### Integer Variable

Used to keep track of the running total cost.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Read all candy prices.
2. Sort prices from highest to lowest.
3. Start scanning from left to right.
4. Add the first two candies of every group of three to the answer.
5. Ignore the third candy because it is free.
6. Continue until the array ends.
7. Return the final cost.

Conceptually:

```text
Sorted prices:
[9, 7, 6, 5, 2, 2]

Group 1:
Pay -> 9
Pay -> 7
Free -> 6

Group 2:
Pay -> 5
Pay -> 2
Free -> 2

Answer = 9 + 7 + 5 + 2
```

---

## Complexity

| Metric           | Complexity | Explanation                                                    |
| ---------------- | ---------- | -------------------------------------------------------------- |
| Time Complexity  | O(n log n) | Sorting the array dominates the runtime                        |
| Space Complexity | O(1)       | No extra data structure is required apart from a few variables |

Where:

* `n` = total number of candies

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumCost(vector<int>& cost) {
        // Sort candies from highest cost to lowest cost
        sort(cost.begin(), cost.end(), greater<int>());

        int total = 0;

        // Add all candies except every third candy
        for (int i = 0; i < cost.size(); i++) {
            // Index 2,5,8,... are free candies
            if (i % 3 == 2) continue;

            total += cost[i];
        }

        return total;
    }
};
```

### Java

```java
class Solution {
    public int minimumCost(int[] cost) {
        // Sort in ascending order first
        Arrays.sort(cost);

        int total = 0;

        // Traverse from largest to smallest
        int position = 0;

        for (int i = cost.length - 1; i >= 0; i--) {
            // Every third candy is free
            if (position % 3 != 2) {
                total += cost[i];
            }

            position++;
        }

        return total;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} cost
 * @return {number}
 */
var minimumCost = function(cost) {
    // Sort from highest cost to lowest cost
    cost.sort((a, b) => b - a);

    let total = 0;

    // Skip every third candy
    for (let i = 0; i < cost.length; i++) {
        if (i % 3 === 2) continue;

        total += cost[i];
    }

    return total;
};
```

### Python3

```python
class Solution:
    def minimumCost(self, cost: List[int]) -> int:
        # Sort candies from highest cost to lowest cost
        cost.sort(reverse=True)

        total = 0

        # Every third candy becomes free
        for i in range(len(cost)):
            if i % 3 == 2:
                continue

            total += cost[i]

        return total
```

### Go

```go
func minimumCost(cost []int) int {
    // Sort candies from highest cost to lowest cost
    sort.Slice(cost, func(i, j int) bool {
        return cost[i] > cost[j]
    })

    total := 0

    // Skip every third candy
    for i := 0; i < len(cost); i++ {
        if i%3 == 2 {
            continue
        }

        total += cost[i]
    }

    return total
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages.

### Step 1: Sort the Candy Prices

I first sort the prices in descending order.

Example:

```text
Original:
[6,5,7,9,2,2]

Sorted:
[9,7,6,5,2,2]
```

Sorting is important because it allows me to organize candies into the most profitable groups.

---

### Step 2: Process Candies in Groups of Three

After sorting, every three consecutive candies form a valid discount group.

```text
[9,7,6]
[5,2,2]
```

For each group:

```text
Pay for first
Pay for second
Get third free
```

---

### Step 3: Skip Every Third Candy

While traversing the sorted array:

```text
Index: 0 1 2 3 4 5 6 7 8
Role : P P F P P F P P F
```

Where:

* P = Paid candy
* F = Free candy

Every third candy can be ignored when calculating the answer.

---

### Step 4: Build the Final Cost

Whenever a candy is not free:

```text
total += candyCost
```

Whenever a candy is free:

```text
skip it
```

By the end of the traversal, the total variable contains the minimum possible cost.

---

### Why This Greedy Strategy Works

The discount allows only one free candy for every two purchased candies.

To maximize savings, I want the free candies to be as expensive as possible.

Sorting in descending order guarantees exactly that.

Any other arrangement would force cheaper candies to become free and would increase the final amount paid.

---

## Examples

### Example 1

Input

```text
cost = [1,2,3]
```

Output

```text
5
```

Explanation

```text
Sorted = [3,2,1]

Pay 3
Pay 2
Get 1 free

Total = 5
```

---

### Example 2

Input

```text
cost = [6,5,7,9,2,2]
```

Output

```text
23
```

Explanation

```text
Sorted = [9,7,6,5,2,2]

Pay 9
Pay 7
Free 6

Pay 5
Pay 2
Free 2

Total = 23
```

---

### Example 3

Input

```text
cost = [5,5]
```

Output

```text
10
```

Explanation

```text
Only two candies exist.

No free candy can be taken.

Total = 5 + 5 = 10
```

---

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
```

---

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

---

### JavaScript

Run

```bash
node solution.js
```

---

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run

```bash
go run solution.go
```

Build

```bash
go build solution.go
```

---

## Notes & Optimizations

* Sorting is the most expensive operation in this solution.
* Since `n <= 100`, performance is more than sufficient.
* The greedy approach guarantees the optimal answer.
* No dynamic programming is required.
* No additional arrays or hash maps are needed.
* Edge cases such as one candy or two candies are automatically handled.
* If the array length is not a multiple of three, the remaining candies are simply paid for normally.

Alternative approaches exist, but they do not improve the asymptotic complexity compared to this clean greedy solution.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
