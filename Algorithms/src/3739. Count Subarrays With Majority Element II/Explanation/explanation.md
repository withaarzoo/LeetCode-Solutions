# 3739. Count Subarrays With Majority Element II

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

## Problem Summary

In this LeetCode problem, I am given an array `nums` and a number `target`.

My job is to count how many subarrays have `target` as the majority element.

A majority element means the value appears strictly more than half of the length of that subarray.

So for every subarray, I need to check whether `target` appears more than all the other values combined.

The output is the total number of such subarrays.

## Constraints

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^9`
* `1 <= target <= 10^9`

## Intuition

The key idea is to stop thinking about the subarray as a normal array and instead turn it into a balance problem.

I noticed that every time I see `target`, it helps the answer. Every other number works against it.

So I can convert:

* `target` into `+1`
* everything else into `-1`

After that, a subarray has `target` as the majority element only when the total sum is positive.

That turns the problem into a prefix sum counting problem, which is much easier to manage efficiently.

## Approach

I first convert the array into `+1` and `-1`.

Then I build prefix sums over that transformed array.

For any subarray, the sum is the difference between two prefix sums.

A subarray is valid only when that difference is greater than zero.

That means for every ending position, I need to count how many earlier prefix sums are smaller than the current one.

Since prefix sums can be negative and large, I compress them into a smaller range.

Then I use a Fenwick Tree, also called a Binary Indexed Tree, to count how many prefix sums I have seen so far.

For each prefix sum:

1. I find its compressed index.
2. I ask how many earlier prefix sums are smaller.
3. I add that count to the answer.
4. I insert the current prefix sum into the Fenwick Tree.

This gives an `O(n log n)` solution, which is fast enough for `10^5` elements.

## Data Structures Used

* Prefix Sum Array: I use it to convert subarray sum checks into difference checks between two positions.
* Coordinate Compression: I use it because prefix sums can be negative and large, so I need a compact index range.
* Fenwick Tree / Binary Indexed Tree: I use it to count how many previous prefix sums are smaller in logarithmic time.

## Operations & Behavior Summary

* Convert each number to `+1` if it equals `target`, otherwise `-1`.
* Build prefix sums from the transformed values.
* Count how many pairs of prefix sums satisfy `prefix[left] < prefix[right]`.
* Use coordinate compression so all prefix sums fit into a small index range.
* Use Fenwick Tree queries to count smaller prefix sums already seen.
* Insert the current prefix sum after querying it.
* Sum all valid counts to get the final answer.

## Complexity

| Metric           |   Complexity | Explanation                                                                                                          |
| ---------------- | -----------: | -------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n log n)` | I build prefix sums in `O(n)`, compress values in `O(n log n)`, and each Fenwick Tree query/update takes `O(log n)`. |
| Space Complexity |       `O(n)` | I store the prefix sums, the compressed values, and the Fenwick Tree array.                                          |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    // Fenwick Tree for prefix frequency counting
    class Fenwick {
    public:
        vector<int> bit;

        Fenwick(int n) {
            bit.assign(n + 1, 0);
        }

        // Add one occurrence at index
        void update(int idx) {
            while (idx < bit.size()) {
                bit[idx]++;
                idx += idx & -idx;
            }
        }

        // Count frequencies from 1...idx
        long long query(int idx) {
            long long sum = 0;
            while (idx > 0) {
                sum += bit[idx];
                idx -= idx & -idx;
            }
            return sum;
        }
    };

    long long countMajoritySubarrays(vector<int>& nums, int target) {

        int n = nums.size();

        // Build prefix sums after converting target -> +1, others -> -1
        vector<int> pref(n + 1, 0);
        for (int i = 0; i < n; i++) {
            pref[i + 1] = pref[i] + (nums[i] == target ? 1 : -1);
        }

        // Coordinate compression because prefix sums may be negative
        vector<int> values = pref;
        sort(values.begin(), values.end());
        values.erase(unique(values.begin(), values.end()), values.end());

        Fenwick ft(values.size());

        long long ans = 0;

        // Process every prefix sum
        for (int x : pref) {

            // Compressed index (1-based)
            int idx = lower_bound(values.begin(), values.end(), x) - values.begin() + 1;

            // Count previous prefix sums strictly smaller than current
            ans += ft.query(idx - 1);

            // Insert current prefix sum
            ft.update(idx);
        }

        return ans;
    }
};
```

### Java

```java
class Solution {

    // Fenwick Tree for frequency counting
    static class Fenwick {
        int[] bit;

        Fenwick(int n) {
            bit = new int[n + 1];
        }

        // Add one occurrence
        void update(int idx) {
            while (idx < bit.length) {
                bit[idx]++;
                idx += idx & -idx;
            }
        }

        // Prefix frequency
        long query(int idx) {
            long sum = 0;
            while (idx > 0) {
                sum += bit[idx];
                idx -= idx & -idx;
            }
            return sum;
        }
    }

    public long countMajoritySubarrays(int[] nums, int target) {

        int n = nums.length;

        // Prefix sums after transformation
        int[] pref = new int[n + 1];
        for (int i = 0; i < n; i++) {
            pref[i + 1] = pref[i] + (nums[i] == target ? 1 : -1);
        }

        // Coordinate compression
        int[] values = pref.clone();
        java.util.Arrays.sort(values);

        int m = 0;
        for (int x : values) {
            if (m == 0 || values[m - 1] != x) {
                values[m++] = x;
            }
        }

        Fenwick ft = new Fenwick(m);

        long ans = 0;

        for (int x : pref) {

            int idx = java.util.Arrays.binarySearch(values, 0, m, x) + 1;

            ans += ft.query(idx - 1);

            ft.update(idx);
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var countMajoritySubarrays = function(nums, target) {

    const n = nums.length;

    // Prefix sums after converting target -> +1 and others -> -1
    const pref = new Array(n + 1).fill(0);

    for (let i = 0; i < n; i++) {
        pref[i + 1] = pref[i] + (nums[i] === target ? 1 : -1);
    }

    // Coordinate compression
    const values = [...new Set([...pref].sort((a, b) => a - b))];

    // Fenwick Tree
    const bit = new Array(values.length + 2).fill(0);

    function update(idx) {
        while (idx < bit.length) {
            bit[idx]++;
            idx += idx & -idx;
        }
    }

    function query(idx) {
        let sum = 0;
        while (idx > 0) {
            sum += bit[idx];
            idx -= idx & -idx;
        }
        return sum;
    }

    function lowerBound(arr, x) {
        let l = 0, r = arr.length;
        while (l < r) {
            const mid = (l + r) >> 1;
            if (arr[mid] < x) l = mid + 1;
            else r = mid;
        }
        return l;
    }

    let ans = 0;

    for (const x of pref) {

        const idx = lowerBound(values, x) + 1;

        ans += query(idx - 1);

        update(idx);
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:

        n = len(nums)

        # Build prefix sums after transforming the array
        pref = [0]
        for x in nums:
            pref.append(pref[-1] + (1 if x == target else -1))

        # Coordinate compression
        values = sorted(set(pref))

        # Fenwick Tree
        bit = [0] * (len(values) + 2)

        # Insert one prefix sum
        def update(idx):
            while idx < len(bit):
                bit[idx] += 1
                idx += idx & -idx

        # Count prefix sums up to idx
        def query(idx):
            s = 0
            while idx > 0:
                s += bit[idx]
                idx -= idx & -idx
            return s

        ans = 0

        for x in pref:

            # Compressed index
            idx = bisect_left(values, x) + 1

            # Count smaller prefix sums
            ans += query(idx - 1)

            # Store current prefix sum
            update(idx)

        return ans
```

### Go

```go
func countMajoritySubarrays(nums []int, target int) int64 {

 n := len(nums)

 // Prefix sums after transformation
 pref := make([]int, n+1)
 for i := 0; i < n; i++ {
  if nums[i] == target {
   pref[i+1] = pref[i] + 1
  } else {
   pref[i+1] = pref[i] - 1
  }
 }

 // Coordinate compression
 values := append([]int{}, pref...)
 sort.Ints(values)

 k := 0
 for _, x := range values {
  if k == 0 || values[k-1] != x {
   values[k] = x
   k++
  }
 }
 values = values[:k]

 // Fenwick Tree
 bit := make([]int, k+2)

 update := func(idx int) {
  for idx < len(bit) {
   bit[idx]++
   idx += idx & -idx
  }
 }

 query := func(idx int) int64 {
  var sum int64
  for idx > 0 {
   sum += int64(bit[idx])
   idx -= idx & -idx
  }
  return sum
 }

 var ans int64

 for _, x := range pref {

  // Compressed index
  idx := sort.SearchInts(values, x) + 1

  // Count smaller prefix sums
  ans += query(idx - 1)

  // Insert current prefix sum
  update(idx)
 }

 return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in every language, so I think about the solution in one way and then adapt the syntax.

First, I transform the array. Every `target` becomes `+1`, and every other value becomes `-1`. This is the most important step because it changes the original majority check into a simple sum check.

Next, I build prefix sums. This helps me measure the total value of any subarray using two prefix values instead of scanning the subarray again and again.

Then I realize that a valid subarray is one where the later prefix sum is greater than the earlier prefix sum. So the problem becomes: for each prefix sum, how many previous prefix sums are smaller?

That is why I use coordinate compression. Prefix sums may be negative, and they may spread out a lot. Compression lets me map them to small consecutive indices without changing their order.

After that, I use a Fenwick Tree. It keeps track of how many prefix sums I have already seen. A query tells me how many are smaller than the current one. An update adds the current prefix sum to the structure.

The order matters. I query first and update second. That avoids counting the current prefix sum against itself.

If `target` never appears, then every transformed value becomes `-1`, so no subarray can have a positive sum. The answer naturally becomes `0`.

## Examples

### Example 1

Input:
`nums = [1, 2, 2, 3]`, `target = 2`

Output:
`5`

Brief trace:

* Transform the array into `[-1, +1, +1, -1]`
* Build prefix sums
* Count all pairs where the later prefix sum is bigger than the earlier one
* The valid subarrays are exactly the ones where `2` is the majority element

### Example 2

Input:
`nums = [1, 1, 1, 1]`, `target = 1`

Output:
`10`

Brief trace:

* Every value becomes `+1`
* Every subarray has a positive sum
* So all `n * (n + 1) / 2 = 10` subarrays are valid

### Example 3

Input:
`nums = [1, 2, 3]`, `target = 4`

Output:
`0`

Brief trace:

* `4` never appears
* Every transformed value becomes `-1`
* No subarray can have a positive sum
* So the answer is `0`

## How to Use / Run Locally

For C++, place the solution in a file like `main.cpp`, then compile and run with:

`g++ -std=c++17 -O2 -o main main.cpp && ./main`

For Java, save the solution in `Solution.java`, then run:

`javac Solution.java && java Solution`

For JavaScript, save the file as `solution.js`, then run:

`node solution.js`

For Python3, save the file as `solution.py`, then run:

`python3 solution.py`

For Go, save the file as `main.go`, then run:

`go run main.go`

## Notes & Optimizations

A direct brute force solution would be too slow because there can be `O(n^2)` subarrays.

The prefix sum plus Fenwick Tree approach is the main optimization here.

Another possible way is to use an ordered map or balanced tree style structure, but the Fenwick Tree is clean and fast when coordinate compression is available.

The biggest trick in this problem is the transformation from majority checking to prefix sum comparison. Once that idea clicks, the rest becomes a standard counting problem.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
