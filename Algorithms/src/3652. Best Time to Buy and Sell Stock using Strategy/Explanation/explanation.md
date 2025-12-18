# 📈 Best Time to Buy and Sell Stock using Strategy (LeetCode 3652)

---

## 📑 Table of Contents

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## 🧩 Problem Summary

I am given two arrays:

* `prices[i]` → stock price on day `i`
* `strategy[i]` → action on day `i`

  * `-1` → buy
  * `0` → hold
  * `1` → sell

Profit is calculated as:

```
profit = Σ(strategy[i] * prices[i])
```

I am allowed to modify the strategy **at most once**:

* Select exactly `k` consecutive days
* First `k/2` days → set to `0` (hold)
* Last `k/2` days → set to `1` (sell)

My task is to return the **maximum possible profit**.

Important note:
There is **no restriction on buying or selling**. Every action is allowed independently.

---

## 📏 Constraints

* `2 ≤ prices.length = strategy.length ≤ 10⁵`
* `1 ≤ prices[i] ≤ 10⁵`
* `-1 ≤ strategy[i] ≤ 1`
* `2 ≤ k ≤ prices.length`
* `k` is always even

---

## 💡 Intuition

When I read the problem, I realized something simple:

The original profit is already known if I just multiply `strategy[i]` with `prices[i]`.

The real challenge is:

> How much extra profit can I gain by modifying **one window of length `k`**?

Instead of rebuilding the strategy array again and again, I decided to:

* Keep the base profit
* Calculate only the **difference (delta)** caused by each possible window

This made the problem perfect for **prefix sums + sliding window**.

---

## 🛠 Approach

I solved the problem in these steps:

1. Calculate the base profit using the original strategy.
2. Build prefix sums for:

   * prices
   * strategy × prices
3. Slide a window of size `k` across the array.
4. For each window:

   * Remove the old contribution.
   * Add the new forced contribution:

     * First half → 0
     * Second half → sell (`prices[i]`)
5. Track the maximum profit improvement.
6. Add it to the base profit.

---

## 🧮 Data Structures Used

* Prefix Sum Arrays

  * To calculate range sums in O(1)
* Simple variables for tracking maximum delta

---

## 🔁 Operations & Behavior Summary

* No backtracking
* No dynamic programming
* Single pass + sliding window
* Efficient even for large inputs

---

## ⏱ Complexity

**Time Complexity:**
`O(n)`

* One pass to calculate base profit
* One pass to slide the window

**Space Complexity:**
`O(n)`

* Prefix sum arrays

---

## 🌐 Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long maxProfit(vector<int>& prices, vector<int>& strategy, int k) {
        int n = prices.size();
        long long base = 0;

        for (int i = 0; i < n; i++)
            base += 1LL * strategy[i] * prices[i];

        vector<long long> prefPrice(n + 1, 0), prefProfit(n + 1, 0);
        for (int i = 0; i < n; i++) {
            prefPrice[i + 1] = prefPrice[i] + prices[i];
            prefProfit[i + 1] = prefProfit[i] + 1LL * strategy[i] * prices[i];
        }

        long long bestDelta = 0;
        int half = k / 2;

        for (int l = 0; l + k <= n; l++) {
            int m = l + half, r = l + k;
            long long oldProfit = prefProfit[r] - prefProfit[l];
            long long newProfit = prefPrice[r] - prefPrice[m];
            bestDelta = max(bestDelta, newProfit - oldProfit);
        }

        return base + bestDelta;
    }
};
```

---

### Java

```java
class Solution {
    public long maxProfit(int[] prices, int[] strategy, int k) {
        int n = prices.length;
        long base = 0;

        for (int i = 0; i < n; i++)
            base += (long) strategy[i] * prices[i];

        long[] prefPrice = new long[n + 1];
        long[] prefProfit = new long[n + 1];

        for (int i = 0; i < n; i++) {
            prefPrice[i + 1] = prefPrice[i] + prices[i];
            prefProfit[i + 1] = prefProfit[i] + (long) strategy[i] * prices[i];
        }

        long bestDelta = 0;
        int half = k / 2;

        for (int l = 0; l + k <= n; l++) {
            int m = l + half, r = l + k;
            long oldProfit = prefProfit[r] - prefProfit[l];
            long newProfit = prefPrice[r] - prefPrice[m];
            bestDelta = Math.max(bestDelta, newProfit - oldProfit);
        }

        return base + bestDelta;
    }
}
```

---

### JavaScript

```javascript
var maxProfit = function(prices, strategy, k) {
    const n = prices.length;
    let base = 0;

    for (let i = 0; i < n; i++)
        base += strategy[i] * prices[i];

    const prefPrice = Array(n + 1).fill(0);
    const prefProfit = Array(n + 1).fill(0);

    for (let i = 0; i < n; i++) {
        prefPrice[i + 1] = prefPrice[i] + prices[i];
        prefProfit[i + 1] = prefProfit[i] + strategy[i] * prices[i];
    }

    let bestDelta = 0;
    const half = k / 2;

    for (let l = 0; l + k <= n; l++) {
        const m = l + half, r = l + k;
        const oldProfit = prefProfit[r] - prefProfit[l];
        const newProfit = prefPrice[r] - prefPrice[m];
        bestDelta = Math.max(bestDelta, newProfit - oldProfit);
    }

    return base + bestDelta;
};
```

---

### Python3

```python
class Solution:
    def maxProfit(self, prices, strategy, k):
        n = len(prices)
        base = sum(strategy[i] * prices[i] for i in range(n))

        pref_price = [0] * (n + 1)
        pref_profit = [0] * (n + 1)

        for i in range(n):
            pref_price[i + 1] = pref_price[i] + prices[i]
            pref_profit[i + 1] = pref_profit[i] + strategy[i] * prices[i]

        best_delta = 0
        half = k // 2

        for l in range(n - k + 1):
            m, r = l + half, l + k
            old = pref_profit[r] - pref_profit[l]
            new = pref_price[r] - pref_price[m]
            best_delta = max(best_delta, new - old)

        return base + best_delta
```

---

### Go

```go
func maxProfit(prices []int, strategy []int, k int) int64 {
 n := len(prices)
 var base int64 = 0

 for i := 0; i < n; i++ {
  base += int64(strategy[i] * prices[i])
 }

 prefPrice := make([]int64, n+1)
 prefProfit := make([]int64, n+1)

 for i := 0; i < n; i++ {
  prefPrice[i+1] = prefPrice[i] + int64(prices[i])
  prefProfit[i+1] = prefProfit[i] + int64(strategy[i]*prices[i])
 }

 var bestDelta int64 = 0
 half := k / 2

 for l := 0; l+k <= n; l++ {
  m, r := l+half, l+k
  oldProfit := prefProfit[r] - prefProfit[l]
  newProfit := prefPrice[r] - prefPrice[m]
  if newProfit-oldProfit > bestDelta {
   bestDelta = newProfit - oldProfit
  }
 }

 return base + bestDelta
}
```

---

## 🧠 Step-by-step Detailed Explanation

1. Calculate original profit.
2. Build prefix sums for fast range queries.
3. Slide a window of size `k`.
4. Replace its behavior logically.
5. Track best improvement.
6. Add improvement to base profit.

---

## 📌 Examples

**Input**

```
prices = [4,2,8]
strategy = [-1,0,1]
k = 2
```

**Output**

```
10
```

---

## ▶ How to use / Run locally

1. Clone the repository
2. Choose your language file
3. Run using standard compiler/interpreter
4. Test with custom inputs

---

## 📝 Notes & Optimizations

* Sliding window avoids O(n²) brute force
* Prefix sums make range calculations instant
* Works efficiently for large inputs

---

## 👤 Author

* **Md Aarzoo Islam**
  👉 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
