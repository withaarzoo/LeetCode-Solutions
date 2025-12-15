# 2110. Number of Smooth Descent Periods of a Stock

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
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given an integer array `prices` where `prices[i]` represents the stock price on the `i-th` day.

A **smooth descent period** is a group of **one or more continuous days** where:

* Each day’s price is **exactly 1 less** than the previous day’s price.
* A single day is always counted as a valid smooth descent period.

My task is to return the **total number of smooth descent periods**.

---

## Constraints

* `1 <= prices.length <= 10^5`
* `1 <= prices[i] <= 10^5`

The solution must be fast and memory-efficient.

---

## Intuition

When I first analyzed the problem, I realized something important.

Every single day by itself is already a smooth descent period.
So I don’t need to check all subarrays.

Then I noticed a pattern:

* If today’s price is exactly `1` less than yesterday’s price, I can **extend the previous descent**.
* Otherwise, the descent breaks and must restart.

So instead of checking all possible ranges, I just **count how long the current smooth descent is** and keep adding it to the answer.

This way, I solve the problem in **one pass**.

---

## Approach

1. I initialize:

   * `answer = 1` → first day is always valid
   * `length = 1` → current smooth descent length

2. I loop from day `1` to the end:

   * If `prices[i] == prices[i-1] - 1`, I extend the descent
   * Otherwise, I reset the descent length to `1`

3. For each day, I add `length` to the answer.

4. I return the final answer.

This approach is simple, fast, and very reliable.

---

## Data Structures Used

* No extra data structures
* Only integer variables for counting

---

## Operations & Behavior Summary

* Continuous decreasing by `1` → extend descent
* Any other change → reset descent
* Each day contributes all valid subarrays ending on that day

---

## Complexity

* **Time Complexity:** `O(n)`
  I traverse the prices array once, where `n` is the number of days.

* **Space Complexity:** `O(1)`
  I use only constant extra space.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long getDescentPeriods(vector<int>& prices) {
        long long answer = 1;
        long long length = 1;

        for (int i = 1; i < prices.size(); i++) {
            if (prices[i] == prices[i - 1] - 1) {
                length++;
            } else {
                length = 1;
            }
            answer += length;
        }
        return answer;
    }
};
```

---

### Java

```java
class Solution {
    public long getDescentPeriods(int[] prices) {
        long answer = 1;
        long length = 1;

        for (int i = 1; i < prices.length; i++) {
            if (prices[i] == prices[i - 1] - 1) {
                length++;
            } else {
                length = 1;
            }
            answer += length;
        }
        return answer;
    }
}
```

---

### JavaScript

```javascript
var getDescentPeriods = function(prices) {
    let answer = 1;
    let length = 1;

    for (let i = 1; i < prices.length; i++) {
        if (prices[i] === prices[i - 1] - 1) {
            length++;
        } else {
            length = 1;
        }
        answer += length;
    }
    return answer;
};
```

---

### Python3

```python
class Solution:
    def getDescentPeriods(self, prices):
        answer = 1
        length = 1

        for i in range(1, len(prices)):
            if prices[i] == prices[i - 1] - 1:
                length += 1
            else:
                length = 1
            answer += length

        return answer
```

---

### Go

```go
func getDescentPeriods(prices []int) int64 {
    var answer int64 = 1
    var length int64 = 1

    for i := 1; i < len(prices); i++ {
        if prices[i] == prices[i-1]-1 {
            length++
        } else {
            length = 1
        }
        answer += length
    }
    return answer
}
```

---

## Step-by-step Detailed Explanation (All Languages)

* I track how many continuous days are decreasing by exactly `1`
* Each time the chain continues, it forms new valid subarrays
* When the chain breaks, I reset the counter
* Adding `length` daily counts all subarrays ending at that position
* This avoids nested loops and heavy computation

---

## Examples

### Example 1

Input:

```
prices = [3, 2, 1, 4]
```

Output:

```
7
```

Explanation:

```
[3], [2], [1], [4], [3,2], [2,1], [3,2,1]
```

---

### Example 2

Input:

```
prices = [8, 6, 7, 7]
```

Output:

```
4
```

---

## How to Use / Run Locally

1. Clone the repository
2. Open the file for your preferred language
3. Compile or run using standard commands

Example (C++):

```
g++ solution.cpp
./a.out
```

Example (Python):

```
python solution.py
```

---

## Notes & Optimizations

* Brute force solutions fail due to time limits
* This greedy counting method is optimal
* Works efficiently for large inputs
* No extra memory required

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
