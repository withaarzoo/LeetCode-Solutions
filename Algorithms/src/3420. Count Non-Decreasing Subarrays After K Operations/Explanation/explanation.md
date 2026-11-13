# Count Non-Decreasing Subarrays After K Operations - LeetCode 3420 Solution

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

You are given an array of integers called nums and an integer k.  

For every possible subarray of nums you are allowed to perform at most k increment operations. In one operation you may increase any single element of that subarray by 1. Each subarray is independent, so changes made to one subarray do not affect any other subarray.  

The task is to count how many subarrays can be turned into a non-decreasing sequence using at most k such operations.  

A sequence is non-decreasing when every element is greater than or equal to the element that comes before it.  

This is LeetCode problem 3420, a hard-level array problem that appears in weekly contests and is often searched under keywords such as “count non-decreasing subarrays after k operations”, “minimum increments to make subarray non-decreasing”, and “sliding window + monotonic deque solution”.

## Constraints

- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^9
- 1 <= k <= 10^9

Because the array can be as long as 100 000 elements, any O(n²) solution will time out. An efficient linear or near-linear approach is required.

## Intuition

The first observation is that the cost of fixing a subarray is completely decided once the left endpoint is fixed.  

To make a subarray non-decreasing you only ever need to raise each element up to the running maximum that starts from the left end of the subarray. The total number of increments is exactly the sum of (running-max – original value) for every position.  

That cost never decreases when the subarray is extended to the right. Therefore, for every possible left index there exists a farthest right index that can still be afforded with the given budget k.  

If we process the array from right to left we can maintain a sliding window [i \ldots j] and only shrink the right side when the cost exceeds k. The remaining challenge is to update the cost quickly when a new left element is added. A monotonic deque that stores groups of equal final values solves this update in amortized constant time.

## Approach

We walk the array from the rightmost index toward the leftmost index.  

At every step the variable i represents the current left endpoint of the window and j represents the farthest right endpoint that still costs at most k.  

We keep a deque that stores the final adjusted values of the current window as consecutive groups of identical numbers. The back of the deque always holds the leftmost group (closest to the newly inserted element) and the front holds the rightmost group.  

When a new value nums[i] is inserted we repeatedly pop groups from the back of the deque while those groups are strictly smaller than nums[i]. Each popped group must be raised to the new higher value, so we add (newValue – oldValue) × groupSize to the total cost and merge the sizes. After the pops we push the new group onto the deque.  

If the cost is still larger than k we remove positions from the right end of the window one by one, subtracting the contribution of each removed position and moving j leftward, until the cost becomes acceptable again.  

Once the window is valid we know that every subarray starting at i and ending anywhere between i and j inclusive can be made non-decreasing with at most k operations. We simply add (j – i + 1) to the answer.  

Because every index is pushed onto and popped from the deque at most once, the whole algorithm runs in linear time.

## Data Structures Used

- Deque (double-ended queue) of pairs (value, count)  
  Stores consecutive groups of identical final values. Allows O(1) insertion and deletion from both ends, which is essential for the amortized linear runtime.

- Two integer pointers i and j  
  Represent the current left and right ends of the sliding window.

- Long-integer variables for the running cost and the answer  
  Necessary because the cost and the number of subarrays can exceed the range of a 32-bit integer.

## Operations & Behavior Summary

1. Initialise answer = 0, cost = 0, and an empty deque.  
2. Set both i and j to the last index of the array.  
3. For every index i moving from right to left:  
   a. Insert nums[i] into the deque, raising and merging any smaller groups that sit at the back.  
   b. While the current cost exceeds k, remove one position from the right end of the window and subtract its contribution.  
   c. Add the number of valid ending positions (j – i + 1) to the answer.  
4. Return the final answer.

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(n)  | Each element is pushed and popped from the deque a constant number of times. The outer loop runs exactly n times. |
| Space Complexity | O(n)  | In the worst case the deque can hold a linear number of groups. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    long long countNonDecreasingSubarrays(vector<int>& nums, int k) {
        long long ans = 0;
        long long cost = 0;
        deque<pair<int, int>> dq;
        int n = nums.size();
        for (int i = n - 1, j = n - 1; i >= 0; --i) {
            int num = nums[i];
            int count = 1;
            while (!dq.empty() && dq.back().first < num) {
                auto [nextNum, nextCount] = dq.back();
                dq.pop_back();
                count += nextCount;
                cost += 1LL * (num - nextNum) * nextCount;
            }
            dq.emplace_back(num, count);
            while (cost > k) {
                auto [rightmostNum, rightmostCount] = dq.front();
                dq.pop_front();
                cost -= 1LL * (rightmostNum - nums[j]);
                --j;
                if (rightmostCount > 1) {
                    dq.emplace_front(rightmostNum, rightmostCount - 1);
                }
            }
            ans += j - i + 1;
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public long countNonDecreasingSubarrays(int[] nums, int k) {
        long ans = 0;
        long cost = 0;
        Deque<int[]> dq = new ArrayDeque<>();
        int n = nums.length;
        for (int i = n - 1, j = n - 1; i >= 0; --i) {
            int num = nums[i];
            int count = 1;
            while (!dq.isEmpty() && dq.peekLast()[0] < num) {
                int[] back = dq.pollLast();
                count += back[1];
                cost += (long) (num - back[0]) * back[1];
            }
            dq.offerLast(new int[]{num, count});
            while (cost > k) {
                int[] front = dq.pollFirst();
                cost -= (long) (front[0] - nums[j]);
                --j;
                if (front[1] > 1) {
                    dq.offerFirst(new int[]{front[0], front[1] - 1});
                }
            }
            ans += j - i + 1;
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var countNonDecreasingSubarrays = function(nums, k) {
    let ans = 0;
    let cost = 0;
    const dq = [];
    const n = nums.length;
    for (let i = n - 1, j = n - 1; i >= 0; --i) {
        let num = nums[i];
        let count = 1;
        while (dq.length && dq[dq.length - 1][0] < num) {
            const [nextNum, nextCount] = dq.pop();
            count += nextCount;
            cost += (num - nextNum) * nextCount;
        }
        dq.push([num, count]);
        while (cost > k) {
            const [rightmostNum, rightmostCount] = dq.shift();
            cost -= rightmostNum - nums[j];
            --j;
            if (rightmostCount > 1) {
                dq.unshift([rightmostNum, rightmostCount - 1]);
            }
        }
        ans += j - i + 1;
    }
    return ans;
};
```

### TypeScript
```typescript
function countNonDecreasingSubarrays(nums: number[], k: number): number {
    let ans = 0;
    let cost = 0;
    const dq: [number, number][] = [];
    const n = nums.length;
    for (let i = n - 1, j = n - 1; i >= 0; --i) {
        let num = nums[i];
        let count = 1;
        while (dq.length && dq[dq.length - 1][0] < num) {
            const [nextNum, nextCount] = dq.pop()!;
            count += nextCount;
            cost += (num - nextNum) * nextCount;
        }
        dq.push([num, count]);
        while (cost > k) {
            const [rightmostNum, rightmostCount] = dq.shift()!;
            cost -= rightmostNum - nums[j];
            --j;
            if (rightmostCount > 1) {
                dq.unshift([rightmostNum, rightmostCount - 1]);
            }
        }
        ans += j - i + 1;
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def countNonDecreasingSubarrays(self, nums: List[int], k: int) -> int:
        ans = 0
        cost = 0
        dq = collections.deque()
        n = len(nums)
        j = n - 1
        for i in range(n - 1, -1, -1):
            num = nums[i]
            count = 1
            while dq and dq[-1][0] < num:
                nextNum, nextCount = dq.pop()
                count += nextCount
                cost += (num - nextNum) * nextCount
            dq.append((num, count))
            while cost > k:
                rightmostNum, rightmostCount = dq.popleft()
                cost -= rightmostNum - nums[j]
                j -= 1
                if rightmostCount > 1:
                    dq.appendleft((rightmostNum, rightmostCount - 1))
            ans += j - i + 1
        return ans
```

### Go
```go
func countNonDecreasingSubarrays(nums []int, k int) int64 {
    ans := int64(0)
    cost := int64(0)
    type pair struct{ val, cnt int }
    dq := []pair{}
    n := len(nums)
    j := n - 1
    for i := n - 1; i >= 0; i-- {
        num := nums[i]
        count := 1
        for len(dq) > 0 && dq[len(dq)-1].val < num {
            back := dq[len(dq)-1]
            dq = dq[:len(dq)-1]
            count += back.cnt
            cost += int64(num-back.val) * int64(back.cnt)
        }
        dq = append(dq, pair{num, count})
        for cost > int64(k) {
            front := dq[0]
            dq = dq[1:]
            cost -= int64(front.val - nums[j])
            j--
            if front.cnt > 1 {
                dq = append([]pair{{front.val, front.cnt - 1}}, dq...)
            }
        }
        ans += int64(j - i + 1)
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the same logic; only the syntax for the deque and integer types differs.

- We start from the right end of the array so that the right pointer j only ever moves left. This guarantees that the window shrinks monotonically.  
- When a new left value is larger than some groups already present, those groups must be raised. Raising a whole group of size s costs (newValue – oldValue) × s. Merging the groups keeps the deque size small.  
- The front of the deque always corresponds to the rightmost position still inside the window. Removing that position subtracts exactly the amount that was previously added for it, so the cost stays accurate.  
- Because the cost is non-decreasing when the window grows to the right, once the cost exceeds k we never need to reconsider any larger window that starts at the same left endpoint.  
- Edge cases such as a single-element array or an already non-decreasing array are handled automatically: the cost stays zero and every possible subarray is counted.  
- Large values of k (up to 10^9) never cause overflow because all arithmetic is performed with 64-bit integers.

## Examples

**Example 1**  
Input: nums = [6, 3, 1, 2, 4, 4], k = 7  
Output: 17  

There are 21 possible subarrays in total. Only four of them require more than 7 increments:  
[6,3,1], [6,3,1,2], [6,3,1,2,4] and [6,3,1,2,4,4].  
All remaining 17 subarrays can be fixed with at most 7 operations.

**Example 2**  
Input: nums = [6, 3, 1, 3, 6], k = 4  
Output: 12  

All single-element and two-element subarrays are valid. Among the three-element subarrays only [6,3,1] is too expensive. The longer subarray [3,1,3,6] can still be fixed within the budget, giving a total of 12 valid subarrays.

**Example 3**  
Input: nums = [5], k = 0  
Output: 1  

A single element is always non-decreasing and needs zero operations.

## How to Use / Run Locally

**C++**  
```bash
g++ -std=c++17 -O2 solution.cpp -o solution
./solution
```

**Java**  
```bash
javac Solution.java
java Solution
```

**JavaScript**  
```bash
node solution.js
```

**TypeScript**  
```bash
tsc solution.ts
node solution.js
```

**Python3**  
```bash
python3 solution.py
```

**Go**  
```bash
go run solution.go
```

In each case replace the file name with the actual source file that contains the class or function shown above. You can hard-code the sample inputs inside the main function or read them from standard input.

## Notes & Optimizations

- The algorithm is already optimal in both time and space for the given constraints.  
- An alternative sparse-table / binary-lifting approach also works in O(n log n) time and can be useful if you need to answer many independent queries, but it is slower and more complex for the single-query version of this problem.  
- Be careful with 32-bit integers; the cost can easily exceed 2^31 – 1 when nums[i] and the array length are both large.  
- When two consecutive groups have the same final value they can be merged, but the current implementation leaves them separate. The asymptotic complexity stays the same either way.  
- The same sliding-window + monotonic-deque idea appears in several other “minimum operations to make array non-decreasing / non-increasing” problems, so mastering this pattern is useful beyond LeetCode 3420.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
