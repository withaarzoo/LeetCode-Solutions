# Trionic Array II (LeetCode 3640)

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

You are given an integer array `nums`.

A **trionic subarray** is a contiguous subarray `nums[l…r]` such that there exist indices
`l < p < q < r` where:

* `nums[l…p]` is **strictly increasing**
* `nums[p…q]` is **strictly decreasing**
* `nums[q…r]` is **strictly increasing**

Your task is to **return the maximum possible sum** of any trionic subarray.

It is guaranteed that **at least one valid trionic subarray exists**.

---

## Constraints

* `4 ≤ nums.length ≤ 10^5`
* `-10^9 ≤ nums[i] ≤ 10^9`
* At least one trionic subarray always exists

---

## Intuition

When I first looked at the problem, I noticed something very important.

Every trionic subarray has a **clear structure**:

```bash
increasing → decreasing → increasing
```

The middle part must be a **strictly decreasing contiguous segment**.

So instead of trying all subarrays (which is impossible for large `n`),
I decided to **fix the decreasing part first**, and then try to extend it on both sides.

If I know:

* the best increasing sum ending just **before** the decreasing part
* the sum of the decreasing part
* the best increasing sum starting just **after** the decreasing part

Then I can combine them to form a valid trionic subarray.

This observation turns a hard problem into a linear one.

---

## Approach

I solved the problem in **three main steps**.

### Step 1: Precompute increasing sums

I build two helper arrays:

* `maxEndingAt[i]`
  → maximum sum of a **strictly increasing subarray ending at index i**

* `maxStartingAt[i]`
  → maximum sum of a **strictly increasing subarray starting at index i**

This is similar to Kadane’s algorithm, but extension is allowed **only if strictly increasing**.

---

### Step 2: Decompose the array

I scan the array and split it into **maximal strictly decreasing subarrays**.

Each segment gives me:

* `p` → start index of decreasing part
* `q` → end index of decreasing part
* `sum` → sum of that decreasing segment

This segment will act as the **middle (decreasing) phase**.

---

### Step 3: Build trionic subarrays

For each decreasing segment `[p…q]`:

I check:

* Can I extend to the left with an increasing part?
* Can I extend to the right with an increasing part?

If yes, total sum becomes:

```bash
maxEndingAt[p-1] + sum + maxStartingAt[q+1]
```

I take the maximum over all valid cases.

---

## Data Structures Used

* Arrays (`maxEndingAt`, `maxStartingAt`)
* List / Vector of tuples (to store decreasing segments)

No heavy data structures. Everything is linear.

---

## Operations & Behavior Summary

* Strict comparisons (`<`, `>`) enforce correctness
* Only valid transitions are allowed
* Decreasing segments act as anchors
* Increasing segments are greedily maximized
* One-pass logic ensures efficiency

---

## Complexity

* **Time Complexity:** `O(n)`
  (`n` = length of the array)

* **Space Complexity:** `O(n)`
  Two helper arrays for increasing sums

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<tuple<int,int,long long>> decompose(vector<int>& nums){
        int n = nums.size();
        vector<tuple<int,int,long long>> parts;
        int l = 0;
        long long sum = nums[0];

        for(int i = 1; i < n; i++){
            if(nums[i-1] <= nums[i]){
                parts.push_back({l, i-1, sum});
                l = i;
                sum = 0;
            }
            sum += nums[i];
        }
        parts.push_back({l, n-1, sum});
        return parts;
    }

    long long maxSumTrionic(vector<int>& nums) {
        int n = nums.size();
        vector<long long> left(n), right(n);

        for(int i = 0; i < n; i++){
            left[i] = nums[i];
            if(i > 0 && nums[i-1] < nums[i] && left[i-1] > 0)
                left[i] += left[i-1];
        }

        for(int i = n-1; i >= 0; i--){
            right[i] = nums[i];
            if(i < n-1 && nums[i] < nums[i+1] && right[i+1] > 0)
                right[i] += right[i+1];
        }

        auto parts = decompose(nums);
        long long ans = LLONG_MIN;

        for(auto &[p,q,s] : parts){
            if(p > 0 && q < n-1 && nums[p-1] < nums[p] && nums[q] < nums[q+1] && p < q){
                ans = max(ans, left[p-1] + s + right[q+1]);
            }
        }
        return ans;
    }
};
```

---

### Java

```java
class Solution {
    static class Block {
        int l, r;
        long sum;
        Block(int l, int r, long sum){
            this.l = l;
            this.r = r;
            this.sum = sum;
        }
    }

    List<Block> decompose(int[] nums){
        List<Block> list = new ArrayList<>();
        int l = 0;
        long s = nums[0];

        for(int i = 1; i < nums.length; i++){
            if(nums[i-1] <= nums[i]){
                list.add(new Block(l, i-1, s));
                l = i;
                s = 0;
            }
            s += nums[i];
        }
        list.add(new Block(l, nums.length-1, s));
        return list;
    }

    public long maxSumTrionic(int[] nums) {
        int n = nums.length;
        long[] left = new long[n];
        long[] right = new long[n];

        for(int i = 0; i < n; i++){
            left[i] = nums[i];
            if(i > 0 && nums[i-1] < nums[i] && left[i-1] > 0)
                left[i] += left[i-1];
        }

        for(int i = n-1; i >= 0; i--){
            right[i] = nums[i];
            if(i < n-1 && nums[i] < nums[i+1] && right[i+1] > 0)
                right[i] += right[i+1];
        }

        long ans = Long.MIN_VALUE;
        for(Block b : decompose(nums)){
            if(b.l > 0 && b.r < n-1 && nums[b.l-1] < nums[b.l] && nums[b.r] < nums[b.r+1] && b.l < b.r){
                ans = Math.max(ans, left[b.l-1] + b.sum + right[b.r+1]);
            }
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
var maxSumTrionic = function(nums) {
    const n = nums.length;
    const left = Array(n).fill(0);
    const right = Array(n).fill(0);

    for(let i=0;i<n;i++){
        left[i]=nums[i];
        if(i>0 && nums[i-1]<nums[i] && left[i-1]>0)
            left[i]+=left[i-1];
    }

    for(let i=n-1;i>=0;i--){
        right[i]=nums[i];
        if(i<n-1 && nums[i]<nums[i+1] && right[i+1]>0)
            right[i]+=right[i+1];
    }

    let parts=[];
    let l=0,s=nums[0];
    for(let i=1;i<n;i++){
        if(nums[i-1]<=nums[i]){
            parts.push([l,i-1,s]);
            l=i;s=0;
        }
        s+=nums[i];
    }
    parts.push([l,n-1,s]);

    let ans=-1e18;
    for(const [p,q,sum] of parts){
        if(p>0 && q<n-1 && nums[p-1]<nums[p] && nums[q]<nums[q+1] && p<q){
            ans=Math.max(ans,left[p-1]+sum+right[q+1]);
        }
    }
    return ans;
};
```

---

### Python3

```python
class Solution:
    def maxSumTrionic(self, nums):
        n=len(nums)
        left=[0]*n
        right=[0]*n

        for i in range(n):
            left[i]=nums[i]
            if i>0 and nums[i-1]<nums[i] and left[i-1]>0:
                left[i]+=left[i-1]

        for i in range(n-1,-1,-1):
            right[i]=nums[i]
            if i<n-1 and nums[i]<nums[i+1] and right[i+1]>0:
                right[i]+=right[i+1]

        parts=[]
        l,s=0,nums[0]
        for i in range(1,n):
            if nums[i-1]<=nums[i]:
                parts.append((l,i-1,s))
                l,s=i,0
            s+=nums[i]
        parts.append((l,n-1,s))

        ans=-10**30
        for p,q,sumv in parts:
            if p>0 and q<n-1 and nums[p-1]<nums[p] and nums[q]<nums[q+1] and p<q:
                ans=max(ans,left[p-1]+sumv+right[q+1])
        return ans
```

---

### Go

```go
func maxSumTrionic(nums []int) int64 {
    n := len(nums)
    left := make([]int64, n)
    right := make([]int64, n)

    for i := 0; i < n; i++ {
        left[i] = int64(nums[i])
        if i > 0 && nums[i-1] < nums[i] && left[i-1] > 0 {
            left[i] += left[i-1]
        }
    }

    for i := n-1; i >= 0; i-- {
        right[i] = int64(nums[i])
        if i < n-1 && nums[i] < nums[i+1] && right[i+1] > 0 {
            right[i] += right[i+1]
        }
    }

    type Block struct{ l, r int; sum int64 }
    blocks := []Block{}

    l, s := 0, int64(nums[0])
    for i := 1; i < n; i++ {
        if nums[i-1] <= nums[i] {
            blocks = append(blocks, Block{l, i-1, s})
            l, s = i, 0
        }
        s += int64(nums[i])
    }
    blocks = append(blocks, Block{l, n-1, s})

    ans := int64(-1e18)
    for _, b := range blocks {
        if b.l > 0 && b.r < n-1 && nums[b.l-1] < nums[b.l] && nums[b.r] < nums[b.r+1] && b.l < b.r {
            val := left[b.l-1] + b.sum + right[b.r+1]
            if val > ans {
                ans = val
            }
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Compute best increasing sums from left and right
2. Split array into strictly decreasing blocks
3. Validate left and right expansion conditions
4. Combine three parts into a trionic subarray
5. Track maximum sum

Each step is linear and safe.

---

## Examples

**Input:** `[1,4,2,7]`
**Output:** `14`

**Input:** `[0,-2,-1,-3,0,2,-1]`
**Output:** `-4`

---

## How to use / Run locally

```bash
g++ solution.cpp -o run
./run
```

Or paste directly into LeetCode.

---

## Notes & Optimizations

* Strict comparisons are critical
* Decreasing segment must have length ≥ 2
* Works safely with negative numbers
* Linear and interview-optimal

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
