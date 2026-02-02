# 3013. Divide an Array Into Subarrays With Minimum Cost II

## Table of Contents

* [Problem Summary]()
* [Constraints]()
* [Intuition]()
* [Approach]()
* [Data Structures Used]()
* [Operations & Behavior Summary]()
* [Complexity]()
* [Multi-language Solutions]()
* [C++]()
* [Java]()
* [JavaScript]()
* [Python3]()
* [Go]()

* [Step-by-step Detailed Explanation]()
* [Examples]()
* [How to use / Run locally]()
* [Notes & Optimizations]()
* [Author]()

## Problem Summary

You are given an array `nums` of length `n`, and two integers `k` and `dist`. The goal is to divide `nums` into `k` disjoint contiguous subarrays such that the difference between the starting index of the **second** subarray and the starting index of the **k-th** subarray is at most `dist`.

The **cost** of an array is the value of its first element. You need to return the **minimum possible sum** of the costs of these `k` subarrays.

Effectively, the first element `nums[0]` is always selected. You need to select `k-1` additional "split points" (indices) such that if the second split starts at index `i`, the -th split must start within `[i + 1, i + dist]`.

## Constraints

* `3 <= n <= 10^5`
* `1 <= nums[i] <= 10^9`
* `3 <= k <= n`
* `k - 2 <= dist <= n - 2`

## Intuition

The problem asks us to pick `nums[0]` (fixed cost) and then choose `k-1` more starting positions. Let the second subarray start at index `i`.
The constraint implies that we must pick the remaining `k-2` starting positions from the range `[i + 1, i + dist]`.

To minimize the total cost, for every valid second starting position `i`, we should pick the **smallest** `k-2` numbers available in the window `[i + 1, i + dist]`.

This transforms the problem into a **Sliding Window** challenge:

1. Fix `nums[0]`.
2. Iterate `i` (the start of the 2nd subarray).
3. Efficiently maintain the sum of the smallest `k-2` numbers in the sliding window to the right of `i`.

## Approach

1. **Window Definition:** We maintain a window of potential candidates for the "other" `k-2` split points.
2. **Two-Set / Two-Heap Strategy:** To efficiently calculate the sum of the smallest  elements () in a dynamic window, we use two data structures:

* **L (Left Set):** Stores the  smallest elements currently in the window. We track their sum.
* **R (Right Set):** Stores the remaining elements in the window.

1. **Sliding Logic:**

* As we increment `i` (move the second subarray's start):
* The element `nums[i+1]` leaves the candidate pool (it becomes the new `i`).
* The element `nums[i + 1 + dist]` enters the candidate pool.

* We update our sets `L` and `R` to handle this removal and addition.

1. **Balancing:** After every update, we ensure `L` contains exactly  elements and that every element in `L` is smaller than or equal to the smallest element in `R`.

## Data Structures Used

* **C++:** `std::multiset` (Self-balancing BST). Handles duplicates and ordering automatically.
* **Java:** `TreeMap` (Red-Black Tree). Used to simulate a multiset by mapping values to counts.
* **Python / Go / JavaScript:** **Two Heaps** (Max Heap for `L`, Min Heap for `R`) with **Lazy Removal**. Since standard libraries often lack a "remove arbitrary element" function for heaps, we keep a record of "deleted" elements and only remove them physically when they float to the top of the heap.

## Operations & Behavior Summary

| Operation | Description | Complexity |
| --- | --- | --- |
| **Add(val)** | Insert a new number into the appropriate set (`L` or `R`). |  |
| **Remove(val)** | Remove a number leaving the window from `L` or `R` (or mark as lazy deleted). |  |
| **Balance()** | Move elements between `L` and `R` to ensure `L.size() == k-2` and `max(L) <= min(R)`. |  |

## Complexity

* **Time Complexity:**  or .
* We iterate through the array once.
* In each iteration, we perform heap/tree operations (insert/delete) which take logarithmic time.

* **Space Complexity:** .
* We store the window elements in our data structures. In the worst case (lazy removal), the heap size might grow proportional to .

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long minimumCost(vector<int>& nums, int k, int dist) {
        int target = k - 2;
        long long current_window_sum = 0;
        multiset<int> L, R;

        // Initialize window [2, dist + 1]
        for (int j = 2; j <= min((int)nums.size() - 1, dist + 1); ++j) {
            L.insert(nums[j]);
            current_window_sum += nums[j];
        }

        auto balance = [&]() {
            while (L.size() > target) {
                int x = *L.rbegin();
                L.erase(prev(L.end()));
                current_window_sum -= x;
                R.insert(x);
            }
            while (L.size() < target && !R.empty()) {
                int x = *R.begin();
                R.erase(R.begin());
                L.insert(x);
                current_window_sum += x;
            }
            while (!L.empty() && !R.empty() && *L.rbegin() > *R.begin()) {
                int l_val = *L.rbegin();
                int r_val = *R.begin();
                L.erase(prev(L.end()));
                R.erase(R.begin());
                L.insert(r_val);
                R.insert(l_val);
                current_window_sum = current_window_sum - l_val + r_val;
            }
        };

        balance();
        long long min_cost = (long long)nums[0] + nums[1] + current_window_sum;

        // Slide window
        for (int i = 2; i <= nums.size() - (k - 1); ++i) {
            int out_val = nums[i];
            auto it = L.find(out_val);
            if (it != L.end()) {
                L.erase(it);
                current_window_sum -= out_val;
            } else {
                R.erase(R.find(out_val));
            }

            if (i + dist < nums.size()) {
                int in_val = nums[i + dist];
                L.insert(in_val);
                current_window_sum += in_val;
            }

            balance();
            long long current_cost = (long long)nums[0] + nums[i] + current_window_sum;
            min_cost = min(min_cost, current_cost);
        }
        return min_cost;
    }
};

```

### Java

```java
import java.util.TreeMap;

class Solution {
    private void add(TreeMap<Integer, Integer> map, int val) {
        map.put(val, map.getOrDefault(val, 0) + 1);
    }
    
    private void remove(TreeMap<Integer, Integer> map, int val) {
        if (map.get(val) == 1) map.remove(val);
        else map.put(val, map.get(val) - 1);
    }

    TreeMap<Integer, Integer> L = new TreeMap<>();
    TreeMap<Integer, Integer> R = new TreeMap<>();
    long currentLSum = 0;
    int sizeL = 0;
    
    public long minimumCost(int[] nums, int k, int dist) {
        int target = k - 2;
        int n = nums.length;
        
        for (int j = 2; j <= Math.min(n - 1, dist + 1); j++) {
            add(L, nums[j]);
            currentLSum += nums[j];
            sizeL++;
        }
        
        balance(target);
        
        long minCost = (long)nums[0] + nums[1] + currentLSum;
        
        for (int i = 2; i <= n - (k - 1); i++) {
            int outVal = nums[i];
            if (L.containsKey(outVal)) {
                remove(L, outVal);
                currentLSum -= outVal;
                sizeL--;
            } else {
                remove(R, outVal);
            }
            
            if (i + dist < n) {
                int inVal = nums[i + dist];
                add(L, inVal);
                currentLSum += inVal;
                sizeL++;
            }
            
            balance(target);
            minCost = Math.min(minCost, (long)nums[0] + nums[i] + currentLSum);
        }
        return minCost;
    }
    
    private void balance(int target) {
        while (sizeL > target) {
            int maxL = L.lastKey();
            remove(L, maxL);
            currentLSum -= maxL;
            sizeL--;
            add(R, maxL);
        }
        while (sizeL < target && !R.isEmpty()) {
            int minR = R.firstKey();
            remove(R, minR);
            add(L, minR);
            currentLSum += minR;
            sizeL++;
        }
        while (!L.isEmpty() && !R.isEmpty() && L.lastKey() > R.firstKey()) {
            int maxL = L.lastKey();
            int minR = R.firstKey();
            remove(L, maxL); currentLSum -= maxL; sizeL--; add(R, maxL);
            remove(R, minR); add(L, minR); currentLSum += minR; sizeL++;
        }
    }
}

```

### JavaScript

```javascript

```

### Python3

```python
import heapq

class Solution:
    def minimumCost(self, nums: List[int], k: int, dist: int) -> int:
        target = k - 2
        
        # L: Max-heap (-val), R: Min-heap (val)
        L, R = [], []
        L_sum = 0
        L_size = 0 
        
        # Lazy removal tracking
        to_remove = {}
        
        def push(val):
            nonlocal L_sum, L_size
            # If L is empty or val < max(L)
            if not L or val < -L[0]:
                heapq.heappush(L, -val)
                L_sum += val
                L_size += 1
            else:
                heapq.heappush(R, val)
                
        def clean(heap, is_max_heap):
            while heap:
                val = -heap[0] if is_max_heap else heap[0]
                if to_remove.get(val, 0) > 0:
                    to_remove[val] -= 1
                    if to_remove[val] == 0: del to_remove[val]
                    heapq.heappop(heap)
                else:
                    break

        def rebalance():
            nonlocal L_sum, L_size
            clean(L, True)
            clean(R, False)
            
            while L_size > target:
                clean(L, True)
                val = -heapq.heappop(L)
                L_sum -= val
                L_size -= 1
                heapq.heappush(R, val)
                clean(L, True)
            
            while L_size < target and R:
                clean(R, False)
                val = heapq.heappop(R)
                heapq.heappush(L, -val)
                L_sum += val
                L_size += 1
                clean(R, False)
                
        # Initial Window
        for j in range(2, min(len(nums), dist + 2)):
            push(nums[j])
        rebalance()
        
        min_cost = nums[0] + nums[1] + L_sum
        
        for i in range(2, len(nums) - (k - 1) + 1):
            out_val = nums[i]
            
            clean(L, True)
            in_L = L and out_val <= -L[0]
            to_remove[out_val] = to_remove.get(out_val, 0) + 1
            
            if in_L:
                L_sum -= out_val
                L_size -= 1
            
            if i + dist < len(nums):
                push(nums[i + dist])
                
            rebalance()
            min_cost = min(min_cost, nums[0] + nums[i] + L_sum)
            
        return min_cost

```

### Go

```go
```

## Step-by-step Detailed Explanation

1. **Initialization**:

* We define `target` as `k-2`. This is the number of extra elements we need to pick from the window.
* We initialize two heaps (or sets): `L` (for the chosen small numbers) and `R` (for the rejected large numbers).
* We look at the first possible window of candidates, which starts at `nums[2]` (since `nums[1]` is the first possible pivot). We insert these candidates into our heaps.

1. **Balancing**:

* The `balance()` or `rebalance()` function is the core logic. It moves numbers between `L` and `R` to satisfy two conditions:

1. `L` must have exactly `target` valid elements.
2. Every element in `L` must be smaller than the smallest element in `R`.

3. **The Loop (Sliding Window)**:

* We iterate `i` from `2` up to the last valid pivot index. `i` represents the starting index of the *second* subarray.
* **Departure**: `nums[i]` was previously in our candidate pool (window). Now it becomes the pivot itself, so we must remove it from `L` or `R`.
* **Arrival**: As `i` moves right, the window end extends. `nums[i + dist]` enters the window. We add it to `L` or `R`.

1. **Lazy Removal (Python, Go, JS)**:

* Since standard heaps don't support removing a specific value efficiently, we use a "Lazy Removal" technique.
* We keep a `delayed` map (frequency counter). When we want to "remove" `x`, we just increment its count in `delayed`.
* Whenever we access the `top` of a heap, we check if that value is in `delayed`. If it is, we discard it and look at the next top.

1. **Cost Calculation**:

* At each step `i`, the cost is: `nums[0] (fixed) + nums[i] (current pivot) + sum(L) (smallest neighbors)`.
* We update our global `minCost` if the current cost is lower.

## Examples

**Example 1:**

```
Input: nums = [1,3,2,6,4,2], k = 3, dist = 3
Output: 5
Explanation:
1. nums[0]=1 is fixed.
2. We need 2 more subarrays.
3. If split at nums[2]=2, window is [6,4,2]. Smallest 1 element is 2. Total: 1 + 2 + 2 = 5.

```

## How to use / Run locally

1. Clone this repository.
2. Choose your preferred language file (e.g., `Solution.cpp`, `Solution.py`).
3. Instantiate the class/function in your main method.
4. Pass the input array `nums`, integer `k`, and integer `dist`.
5. Print the result.

**Example (Python):**

```python
sol = Solution()
print(sol.minimumCost([1,3,2,6,4,2], 3, 3)) # Output: 5

```

## Notes & Optimizations

* **Lazy Removal vs. Multisets:** In C++, `std::multiset` makes the code cleaner because it supports direct removal. In other languages, the "Two Heaps with Lazy Removal" pattern is a standard optimization for median/sliding-window problems.
* **Edge Cases:** The solution naturally handles cases where `dist` is large (covering the whole array) or minimal, thanks to the loop bounds.

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
