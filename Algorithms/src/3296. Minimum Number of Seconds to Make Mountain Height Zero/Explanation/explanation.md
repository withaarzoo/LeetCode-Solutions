# 3296. Minimum Number of Seconds to Make Mountain Height Zero

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

You are given an integer `mountainHeight` representing the height of a mountain.

You are also given an integer array `workerTimes` where `workerTimes[i]` represents the time taken by the `i-th` worker to remove **1 unit** of mountain height initially.

Workers remove the mountain **simultaneously**.

For worker `i`, removing `x` height takes:

```
workerTimes[i] * (1 + 2 + 3 + ... + x)
```

Which is the sum of the first `x` natural numbers.

Your task is to determine the **minimum number of seconds required** for all workers together to reduce the mountain height to **0**.

---

## Constraints

```
1 <= mountainHeight <= 10^5
1 <= workerTimes.length <= 10^4
1 <= workerTimes[i] <= 10^6
```

---

## Intuition

At first glance the problem looks like a simulation problem, but the time grows very quickly because removing each additional unit takes more time than the previous one.

For a worker with time `t`, removing `x` units takes:

```
t * (1 + 2 + ... + x)
```

Using the arithmetic series formula:

```
1 + 2 + ... + x = x(x+1)/2
```

So the time required becomes:

```
t * x(x+1)/2
```

Instead of directly assigning heights to workers, I realized it is easier to **binary search on time**.

Idea:

1. Guess a time `T`.
2. Check how much height all workers can remove in `T` seconds.
3. If they can remove at least `mountainHeight`, then `T` is possible.
4. Try to minimize `T`.

---

## Approach

1. Use **Binary Search on time**.
2. Search range:

```
left = 1
right = 1e18
```

1. For each candidate time `mid`:

* For every worker
* Find maximum height they can remove within `mid` seconds
* Use binary search for that worker

Condition:

```
t * x(x+1)/2 <= mid
```

1. Add the heights removed by all workers.

If total height >= mountainHeight:

```
mid is feasible
```

1. Continue binary search until the minimum valid time is found.

---

## Data Structures Used

The solution mainly uses primitive variables.

* Integers
* Long integers
* Arrays (workerTimes)

No additional complex data structures are required.

---

## Operations & Behavior Summary

Main operations performed in the algorithm:

1. Binary search on the answer (time)
2. Binary search per worker to compute removable height
3. Summation of heights from all workers
4. Early stopping if mountain height requirement is reached

---

## Complexity

### Time Complexity

```
O(n * log(H) * log(T))
```

Where:

* `n` = number of workers
* `H` = mountain height
* `T` = search range of time

Explanation:

Binary search on time → `log(T)`
Binary search per worker → `log(H)`
Workers iteration → `n`

### Space Complexity

```
O(1)
```

Only constant extra space is used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:

    bool can(long long time, int mountainHeight, vector<int>& workerTimes) {
        long long totalHeight = 0;

        for (int t : workerTimes) {
            long long left = 0, right = mountainHeight;

            while (left <= right) {
                long long mid = (left + right) / 2;
                long long required = (long long)t * (mid * (mid + 1) / 2);

                if (required <= time)
                    left = mid + 1;
                else
                    right = mid - 1;
            }

            totalHeight += right;

            if (totalHeight >= mountainHeight)
                return true;
        }

        return false;
    }

    long long minNumberOfSeconds(int mountainHeight, vector<int>& workerTimes) {

        long long left = 1, right = 1e18;
        long long ans = right;

        while (left <= right) {
            long long mid = (left + right) / 2;

            if (can(mid, mountainHeight, workerTimes)) {
                ans = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {

    private boolean can(long time, int mountainHeight, int[] workerTimes) {

        long totalHeight = 0;

        for (int t : workerTimes) {

            long left = 0, right = mountainHeight;

            while (left <= right) {

                long mid = (left + right) / 2;
                long required = (long) t * (mid * (mid + 1) / 2);

                if (required <= time)
                    left = mid + 1;
                else
                    right = mid - 1;
            }

            totalHeight += right;

            if (totalHeight >= mountainHeight)
                return true;
        }

        return false;
    }

    public long minNumberOfSeconds(int mountainHeight, int[] workerTimes) {

        long left = 1, right = (long)1e18;
        long ans = right;

        while (left <= right) {

            long mid = (left + right) / 2;

            if (can(mid, mountainHeight, workerTimes)) {
                ans = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
var minNumberOfSeconds = function(mountainHeight, workerTimes) {

    const can = (time) => {

        let totalHeight = 0n;

        for (let t of workerTimes) {

            let left = 0n;
            let right = BigInt(mountainHeight);

            while (left <= right) {

                let mid = (left + right) / 2n;
                let required = BigInt(t) * (mid * (mid + 1n) / 2n);

                if (required <= time)
                    left = mid + 1n;
                else
                    right = mid - 1n;
            }

            totalHeight += right;

            if (totalHeight >= BigInt(mountainHeight))
                return true;
        }

        return false;
    };

    let left = 1n;
    let right = 10n ** 18n;
    let ans = right;

    while (left <= right) {

        let mid = (left + right) / 2n;

        if (can(mid)) {
            ans = mid;
            right = mid - 1n;
        } else {
            left = mid + 1n;
        }
    }

    return Number(ans);
};
```

### Python3

```python
class Solution:
    def minNumberOfSeconds(self, mountainHeight: int, workerTimes: List[int]) -> int:

        def can(time):

            total = 0

            for t in workerTimes:

                left, right = 0, mountainHeight

                while left <= right:

                    mid = (left + right) // 2
                    required = t * (mid * (mid + 1) // 2)

                    if required <= time:
                        left = mid + 1
                    else:
                        right = mid - 1

                total += right

                if total >= mountainHeight:
                    return True

            return False

        left, right = 1, 10**18
        ans = right

        while left <= right:

            mid = (left + right) // 2

            if can(mid):
                ans = mid
                right = mid - 1
            else:
                left = mid + 1

        return ans
```

### Go

```go
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {

    can := func(time int64) bool {

        totalHeight := int64(0)

        for _, t := range workerTimes {

            left := int64(0)
            right := int64(mountainHeight)

            for left <= right {

                mid := (left + right) / 2
                required := int64(t) * (mid * (mid + 1) / 2)

                if required <= time {
                    left = mid + 1
                } else {
                    right = mid - 1
                }
            }

            totalHeight += right

            if totalHeight >= int64(mountainHeight) {
                return true
            }
        }

        return false
    }

    left := int64(1)
    right := int64(1e18)
    ans := right

    for left <= right {

        mid := (left + right) / 2

        if can(mid) {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Start binary search on time.

```
left = 1
right = 1e18
```

1. Pick middle value.

```
mid = (left + right) / 2
```

1. Check if workers can remove the mountain within `mid` seconds.

2. For each worker, compute maximum height removable using binary search.

Condition:

```
t * x(x+1) / 2 <= mid
```

1. Add heights from all workers.

If:

```
totalHeight >= mountainHeight
```

Then the time works.

1. Update binary search range accordingly.

If possible:

```
right = mid - 1
```

Else:

```
left = mid + 1
```

1. The smallest feasible time is the answer.

---

## Examples

### Example 1

Input

```
mountainHeight = 4
workerTimes = [2,1,1]
```

Output

```
3
```

Explanation

Workers reduce heights simultaneously and the maximum time among them becomes the total time.

---

### Example 2

Input

```
mountainHeight = 10
workerTimes = [3,2,2,4]
```

Output

```
12
```

---

## How to Use / Run Locally

1. Clone the repository

```
git clone https://github.com/yourusername/repository.git
```

1. Navigate to the project folder

```
cd repository
```

1. Compile and run

Example for C++:

```
g++ solution.cpp
./a.out
```

Example for Python:

```
python solution.py
```

---

## Notes & Optimizations

Possible improvements:

1. Use a mathematical formula to compute height directly instead of binary searching per worker.
2. Early stopping when accumulated height already reaches `mountainHeight`.
3. Using `long long` or `BigInt` to avoid overflow.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
