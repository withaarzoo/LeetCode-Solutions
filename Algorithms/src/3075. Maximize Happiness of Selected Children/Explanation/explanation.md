# 3075. Maximize Happiness of Selected Children

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

I am given an array `happiness[]` where each value represents a child’s happiness level.
There are `n` children standing in a queue.

I need to select exactly `k` children, one by one.

Each time I select a child:

* All **unselected** children lose `1` happiness.
* Happiness never goes below `0`.

My task is to **maximize the total happiness** of the `k` selected children.

---

## Constraints

* `1 ≤ n ≤ 2 × 10⁵`
* `1 ≤ happiness[i] ≤ 10⁸`
* `1 ≤ k ≤ n`

---

## Intuition

When I read the problem, I focused on one important thing:

Every time I pick a child, **all remaining children lose happiness**.

So if I delay picking a child who has high happiness, I lose value.

That’s why I thought:

> “I should always pick the happiest child first.”

This immediately pointed me toward a **greedy approach**.

---

## Approach

1. I sort the `happiness` array in **descending order**.
2. I pick the first `k` children.
3. When I pick the `i-th` child:

   * Its happiness is reduced by `i`
   * Because `i` children were already picked before it.
4. If the value becomes negative, I take `0`.
5. I keep adding these values to get the final answer.

---

## Data Structures Used

* Array / List
* Sorting (built-in)

No extra data structures like heap or map are required.

---

## Operations & Behavior Summary

| Operation   | Description                                     |
| ----------- | ----------------------------------------------- |
| Sort        | Arrange happiness values from highest to lowest |
| Greedy Pick | Always pick the next highest value              |
| Decrement   | Reduce happiness by number of previous picks    |
| Clamp       | Ensure happiness never goes below zero          |

---

## Complexity

* **Time Complexity:** `O(n log n)`

  * Sorting the array takes `O(n log n)`
* **Space Complexity:** `O(1)`

  * Only constant extra variables are used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long maximumHappinessSum(vector<int>& happiness, int k) {
        sort(happiness.begin(), happiness.end(), greater<int>());
        long long ans = 0;

        for (int i = 0; i < k; i++) {
            long long curr = happiness[i] - i;
            if (curr > 0) ans += curr;
        }
        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public long maximumHappinessSum(int[] happiness, int k) {
        Arrays.sort(happiness);
        long ans = 0;
        int n = happiness.length;

        for (int i = 0; i < k; i++) {
            long curr = happiness[n - 1 - i] - i;
            if (curr > 0) ans += curr;
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
var maximumHappinessSum = function(happiness, k) {
    happiness.sort((a, b) => b - a);
    let ans = 0;

    for (let i = 0; i < k; i++) {
        let curr = happiness[i] - i;
        if (curr > 0) ans += curr;
    }
    return ans;
};
```

---

### Python3

```python
class Solution:
    def maximumHappinessSum(self, happiness, k):
        happiness.sort(reverse=True)
        ans = 0

        for i in range(k):
            curr = happiness[i] - i
            if curr > 0:
                ans += curr
        return ans
```

---

### Go

```go
func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Slice(happiness, func(i, j int) bool {
        return happiness[i] > happiness[j]
    })

    var ans int64 = 0
    for i := 0; i < k; i++ {
        curr := happiness[i] - i
        if curr > 0 {
            ans += int64(curr)
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (All Languages)

1. I sort the happiness array in descending order.
2. I start selecting children from index `0`.
3. Before selecting the `i-th` child:

   * Happiness already reduced `i` times.
4. I calculate:

   ```
   effective_happiness = happiness[i] - i
   ```

5. If this value is negative, I ignore it.
6. I add valid happiness values to my answer.
7. After `k` selections, I return the result.

---

## Examples

### Example 1

```
Input: happiness = [1,2,3], k = 2
Output: 4
Explanation:
Pick 3 → remaining [0,1]
Pick 1 → total = 4
```

### Example 2

```
Input: happiness = [1,1,1,1], k = 2
Output: 1
```

### Example 3

```
Input: happiness = [2,3,4,5], k = 1
Output: 5
```

---

## How to use / Run locally

1. Clone the repository
2. Open the file for your preferred language
3. Run using:

   * `g++` for C++
   * `javac` for Java
   * `node` for JavaScript
   * `python3` for Python
   * `go run` for Go

---

## Notes & Optimizations

* No need to simulate the entire process.
* Greedy + sorting is enough.
* This solution easily handles large inputs.
* Works perfectly for competitive programming and interviews.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
