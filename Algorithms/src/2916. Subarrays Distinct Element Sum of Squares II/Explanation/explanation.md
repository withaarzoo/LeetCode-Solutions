# LeetCode 2916. Subarrays Distinct Element Sum of Squares II – Segment Tree Solution

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

You are given a 0-indexed integer array called `nums`.  

For every possible non-empty contiguous subarray, count how many different numbers appear inside it. Then square that count and add all those squared values together.  

Because the final number can get extremely large, return the answer modulo \(10^9 + 7\).

In short: sum of (number of distinct elements)² over all subarrays, taken modulo \(10^9 + 7\).

## Constraints

- \(1 \leq\) nums.length \(\leq 10^5\)
- \(1 \leq\) nums[i] \(\leq 10^5\)

These limits rule out any \(O(n^2)\) solution. We need an approach that stays close to \(O(n \log n)\).

## Intuition

When the array is long (up to \(10^5\)), looping over every subarray is impossible.  

I noticed that once I know the distinct counts for all subarrays ending at index \(i\), I can quickly turn them into the counts for subarrays ending at \(i+1\).  

The only subarrays whose distinct count increases are those that start after the previous occurrence of the new element. Everything before that previous occurrence already contains the value, so the count stays the same.  

That observation turns the problem into a series of range updates on the left endpoints, which a segment tree can handle efficiently.

## Approach

I maintain an imaginary array \(d[0 \dots n-1]\) where \(d[L]\) stores the number of distinct values in the subarray that starts at \(L\) and ends at the current right endpoint \(r\).

I process the array from left to right. For each new right endpoint \(r\):

1. Look up the last index where `nums[r]` appeared (or –1 if it never appeared).
2. Add 1 to every \(d[L]\) for \(L\) in the range (last + 1) \ldots (r – 1). Those are exactly the subarrays that receive a brand-new value.
3. Set \(d[r] = 1\) (the single-element subarray).
4. Query the sum of squares of all \(d[0 \dots r]\) and add it to the answer.
5. Record the current position of `nums[r]`.

Because a whole range of values is increased by the same constant, the sum of squares on that range can be updated with a simple closed-form formula:

\[
\text{newSumSq} = \text{oldSumSq} + 2 \cdot \text{add} \cdot \text{oldSum} + \text{add}^2 \cdot \text{length}
\]

A segment tree that stores both the sum and the sum of squares, together with a lazy tag for pending additions, makes every range update and range query run in \(O(\log n)\) time.

## Data Structures Used

- **Segment Tree with Lazy Propagation**  
  Each node keeps the current sum of \(d\) values, the sum of their squares, and a lazy value that remembers a pending uniform addition. This lets me add 1 to any contiguous range of left endpoints and instantly get the new sum of squares.

- **Last-occurrence array**  
  A simple array of size \(10^5 + 1\) that stores the most recent index of every possible value. Looking up the previous occurrence of the current number is therefore \(O(1)\).

## Operations & Behavior Summary

1. Initialise the segment tree so every \(d[L]\) starts at 0.  
2. For each right endpoint \(r\) from 0 to \(n-1\):  
   - Retrieve the previous index of `nums[r]`.  
   - Range-add +1 on the interval (previous + 1) \ldots (r – 1).  
   - Point-update position \(r\) to 1.  
   - Query the sum of squares on the prefix [0 \ldots r] and accumulate it.  
   - Update the last-occurrence table.  
3. After the loop finishes, the accumulated value (already taken modulo \(10^9 + 7\)) is the answer.

## Complexity

| Type              | Value          | Explanation                                                                 |
|-------------------|----------------|-----------------------------------------------------------------------------|
| Time Complexity   | \(O(n \log n)\) | One constant number of range updates and one range query per position; each costs \(O(\log n)\). |
| Space Complexity  | \(O(n)\)        | The segment tree needs a linear number of nodes; the last-occurrence array is fixed size \(10^5\). |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    static const int MOD = 1e9 + 7;
    struct Node {
        long long sum = 0, sumsq = 0, lazy = 0;
    };
    vector<Node> tree;
    int n;
    void push(int v, int tl, int tr) {
        if (tree[v].lazy && tl != tr) {
            int mid = (tl + tr) >> 1;
            int llen = mid - tl + 1, rlen = tr - mid;
            long long add = tree[v].lazy;
            tree[v << 1].sumsq = (tree[v << 1].sumsq + 2 * add * tree[v << 1].sum + add * add % MOD * llen) % MOD;
            tree[v << 1].sum = (tree[v << 1].sum + add * llen) % MOD;
            tree[v << 1].lazy = (tree[v << 1].lazy + add) % MOD;
            tree[v << 1 | 1].sumsq = (tree[v << 1 | 1].sumsq + 2 * add * tree[v << 1 | 1].sum + add * add % MOD * rlen) % MOD;
            tree[v << 1 | 1].sum = (tree[v << 1 | 1].sum + add * rlen) % MOD;
            tree[v << 1 | 1].lazy = (tree[v << 1 | 1].lazy + add) % MOD;
            tree[v].lazy = 0;
        }
    }
    void update(int v, int tl, int tr, int l, int r, long long add) {
        if (l > r) return;
        if (l == tl && r == tr) {
            int len = tr - tl + 1;
            tree[v].sumsq = (tree[v].sumsq + 2 * add * tree[v].sum + add * add % MOD * len) % MOD;
            tree[v].sum = (tree[v].sum + add * len) % MOD;
            tree[v].lazy = (tree[v].lazy + add) % MOD;
            return;
        }
        push(v, tl, tr);
        int mid = (tl + tr) >> 1;
        update(v << 1, tl, mid, l, min(r, mid), add);
        update(v << 1 | 1, mid + 1, tr, max(l, mid + 1), r, add);
        tree[v].sum = (tree[v << 1].sum + tree[v << 1 | 1].sum) % MOD;
        tree[v].sumsq = (tree[v << 1].sumsq + tree[v << 1 | 1].sumsq) % MOD;
    }
    long long query(int v, int tl, int tr, int l, int r) {
        if (l > r) return 0;
        if (l == tl && r == tr) return tree[v].sumsq;
        push(v, tl, tr);
        int mid = (tl + tr) >> 1;
        return (query(v << 1, tl, mid, l, min(r, mid)) + query(v << 1 | 1, mid + 1, tr, max(l, mid + 1), r)) % MOD;
    }
    int sumCounts(vector<int>& nums) {
        n = nums.size();
        tree.assign(4 * n, Node());
        vector<int> last(100001, -1);
        long long ans = 0;
        for (int r = 0; r < n; ++r) {
            int prv = last[nums[r]];
            if (prv + 1 <= r - 1) update(1, 0, n - 1, prv + 1, r - 1, 1);
            update(1, 0, n - 1, r, r, 1);
            ans = (ans + query(1, 0, n - 1, 0, r)) % MOD;
            last[nums[r]] = r;
        }
        return (int)ans;
    }
};
```

### Java
```java
class Solution {
    static final int MOD = 1000000007;
    static class Node {
        long sum, sumsq, lazy;
    }
    Node[] tree;
    int n;
    void push(int v, int tl, int tr) {
        if (tree[v].lazy != 0 && tl != tr) {
            int mid = (tl + tr) >> 1;
            int llen = mid - tl + 1, rlen = tr - mid;
            long add = tree[v].lazy;
            tree[v << 1].sumsq = (tree[v << 1].sumsq + 2 * add * tree[v << 1].sum + add * add % MOD * llen) % MOD;
            tree[v << 1].sum = (tree[v << 1].sum + add * llen) % MOD;
            tree[v << 1].lazy = (tree[v << 1].lazy + add) % MOD;
            tree[v << 1 | 1].sumsq = (tree[v << 1 | 1].sumsq + 2 * add * tree[v << 1 | 1].sum + add * add % MOD * rlen) % MOD;
            tree[v << 1 | 1].sum = (tree[v << 1 | 1].sum + add * rlen) % MOD;
            tree[v << 1 | 1].lazy = (tree[v << 1 | 1].lazy + add) % MOD;
            tree[v].lazy = 0;
        }
    }
    void update(int v, int tl, int tr, int l, int r, long add) {
        if (l > r) return;
        if (l == tl && r == tr) {
            int len = tr - tl + 1;
            tree[v].sumsq = (tree[v].sumsq + 2 * add * tree[v].sum + add * add % MOD * len) % MOD;
            tree[v].sum = (tree[v].sum + add * len) % MOD;
            tree[v].lazy = (tree[v].lazy + add) % MOD;
            return;
        }
        push(v, tl, tr);
        int mid = (tl + tr) >> 1;
        update(v << 1, tl, mid, l, Math.min(r, mid), add);
        update(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r, add);
        tree[v].sum = (tree[v << 1].sum + tree[v << 1 | 1].sum) % MOD;
        tree[v].sumsq = (tree[v << 1].sumsq + tree[v << 1 | 1].sumsq) % MOD;
    }
    long query(int v, int tl, int tr, int l, int r) {
        if (l > r) return 0;
        if (l == tl && r == tr) return tree[v].sumsq;
        push(v, tl, tr);
        int mid = (tl + tr) >> 1;
        return (query(v << 1, tl, mid, l, Math.min(r, mid)) + query(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r)) % MOD;
    }
    public int sumCounts(int[] nums) {
        n = nums.length;
        tree = new Node[4 * n];
        for (int i = 0; i < tree.length; i++) tree[i] = new Node();
        int[] last = new int[100001];
        Arrays.fill(last, -1);
        long ans = 0;
        for (int r = 0; r < n; r++) {
            int prv = last[nums[r]];
            if (prv + 1 <= r - 1) update(1, 0, n - 1, prv + 1, r - 1, 1);
            update(1, 0, n - 1, r, r, 1);
            ans = (ans + query(1, 0, n - 1, 0, r)) % MOD;
            last[nums[r]] = r;
        }
        return (int) ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var sumCounts = function(nums) {
    const MOD = 1000000007;
    const n = nums.length;
    const sum = new Array(4 * n).fill(0);
    const sumsq = new Array(4 * n).fill(0);
    const lazy = new Array(4 * n).fill(0);

    const push = (v, tl, tr) => {
        if (lazy[v] !== 0 && tl !== tr) {
            const mid = (tl + tr) >> 1;
            const llen = mid - tl + 1;
            const rlen = tr - mid;
            const add = lazy[v];

            sumsq[v << 1] = (sumsq[v << 1] + 2 * add * sum[v << 1] + add * add % MOD * llen) % MOD;
            sum[v << 1] = (sum[v << 1] + add * llen) % MOD;
            lazy[v << 1] = (lazy[v << 1] + add) % MOD;

            sumsq[v << 1 | 1] = (sumsq[v << 1 | 1] + 2 * add * sum[v << 1 | 1] + add * add % MOD * rlen) % MOD;
            sum[v << 1 | 1] = (sum[v << 1 | 1] + add * rlen) % MOD;
            lazy[v << 1 | 1] = (lazy[v << 1 | 1] + add) % MOD;

            lazy[v] = 0;
        }
    };

    const update = (v, tl, tr, l, r, add) => {
        if (l > r) return;
        if (l === tl && r === tr) {
            const len = tr - tl + 1;
            sumsq[v] = (sumsq[v] + 2 * add * sum[v] + add * add % MOD * len) % MOD;
            sum[v] = (sum[v] + add * len) % MOD;
            lazy[v] = (lazy[v] + add) % MOD;
            return;
        }
        push(v, tl, tr);
        const mid = (tl + tr) >> 1;
        update(v << 1, tl, mid, l, Math.min(r, mid), add);
        update(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r, add);
        sum[v] = (sum[v << 1] + sum[v << 1 | 1]) % MOD;
        sumsq[v] = (sumsq[v << 1] + sumsq[v << 1 | 1]) % MOD;
    };

    const query = (v, tl, tr, l, r) => {
        if (l > r) return 0;
        if (l === tl && r === tr) return sumsq[v];
        push(v, tl, tr);
        const mid = (tl + tr) >> 1;
        return (query(v << 1, tl, mid, l, Math.min(r, mid)) +
                query(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r)) % MOD;
    };

    const last = new Array(100001).fill(-1);
    let ans = 0;
    for (let r = 0; r < n; r++) {
        const prv = last[nums[r]];
        if (prv + 1 <= r - 1) update(1, 0, n - 1, prv + 1, r - 1, 1);
        update(1, 0, n - 1, r, r, 1);
        ans = (ans + query(1, 0, n - 1, 0, r)) % MOD;
        last[nums[r]] = r;
    }
    return ans;
};
```

### TypeScript
```typescript
function sumCounts(nums: number[]): number {
    const MOD = 1000000007;
    const n = nums.length;
    const sum = new Array(4 * n).fill(0);
    const sumsq = new Array(4 * n).fill(0);
    const lazy = new Array(4 * n).fill(0);

    const push = (v: number, tl: number, tr: number) => {
        if (lazy[v] !== 0 && tl !== tr) {
            const mid = (tl + tr) >> 1;
            const llen = mid - tl + 1;
            const rlen = tr - mid;
            const add = lazy[v];

            sumsq[v << 1] = (sumsq[v << 1] + 2 * add * sum[v << 1] + add * add % MOD * llen) % MOD;
            sum[v << 1] = (sum[v << 1] + add * llen) % MOD;
            lazy[v << 1] = (lazy[v << 1] + add) % MOD;

            sumsq[v << 1 | 1] = (sumsq[v << 1 | 1] + 2 * add * sum[v << 1 | 1] + add * add % MOD * rlen) % MOD;
            sum[v << 1 | 1] = (sum[v << 1 | 1] + add * rlen) % MOD;
            lazy[v << 1 | 1] = (lazy[v << 1 | 1] + add) % MOD;

            lazy[v] = 0;
        }
    };

    const update = (v: number, tl: number, tr: number, l: number, r: number, add: number) => {
        if (l > r) return;
        if (l === tl && r === tr) {
            const len = tr - tl + 1;
            sumsq[v] = (sumsq[v] + 2 * add * sum[v] + add * add % MOD * len) % MOD;
            sum[v] = (sum[v] + add * len) % MOD;
            lazy[v] = (lazy[v] + add) % MOD;
            return;
        }
        push(v, tl, tr);
        const mid = (tl + tr) >> 1;
        update(v << 1, tl, mid, l, Math.min(r, mid), add);
        update(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r, add);
        sum[v] = (sum[v << 1] + sum[v << 1 | 1]) % MOD;
        sumsq[v] = (sumsq[v << 1] + sumsq[v << 1 | 1]) % MOD;
    };

    const query = (v: number, tl: number, tr: number, l: number, r: number): number => {
        if (l > r) return 0;
        if (l === tl && r === tr) return sumsq[v];
        push(v, tl, tr);
        const mid = (tl + tr) >> 1;
        return (query(v << 1, tl, mid, l, Math.min(r, mid)) +
                query(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r)) % MOD;
    };

    const last = new Array(100001).fill(-1);
    let ans = 0;
    for (let r = 0; r < n; r++) {
        const prv = last[nums[r]];
        if (prv + 1 <= r - 1) update(1, 0, n - 1, prv + 1, r - 1, 1);
        update(1, 0, n - 1, r, r, 1);
        ans = (ans + query(1, 0, n - 1, 0, r)) % MOD;
        last[nums[r]] = r;
    }
    return ans;
}
```

### Python3
```python
class Solution:
    def sumCounts(self, nums: List[int]) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        sum_ = [0] * (4 * n)
        sumsq = [0] * (4 * n)
        lazy = [0] * (4 * n)

        def push(v, tl, tr):
            if lazy[v] != 0 and tl != tr:
                mid = (tl + tr) // 2
                llen = mid - tl + 1
                rlen = tr - mid
                add = lazy[v]

                sumsq[v * 2] = (sumsq[v * 2] + 2 * add * sum_[v * 2] + add * add % MOD * llen) % MOD
                sum_[v * 2] = (sum_[v * 2] + add * llen) % MOD
                lazy[v * 2] = (lazy[v * 2] + add) % MOD

                sumsq[v * 2 + 1] = (sumsq[v * 2 + 1] + 2 * add * sum_[v * 2 + 1] + add * add % MOD * rlen) % MOD
                sum_[v * 2 + 1] = (sum_[v * 2 + 1] + add * rlen) % MOD
                lazy[v * 2 + 1] = (lazy[v * 2 + 1] + add) % MOD

                lazy[v] = 0

        def update(v, tl, tr, l, r, add):
            if l > r:
                return
            if l == tl and r == tr:
                length = tr - tl + 1
                sumsq[v] = (sumsq[v] + 2 * add * sum_[v] + add * add % MOD * length) % MOD
                sum_[v] = (sum_[v] + add * length) % MOD
                lazy[v] = (lazy[v] + add) % MOD
                return
            push(v, tl, tr)
            mid = (tl + tr) // 2
            update(v * 2, tl, mid, l, min(r, mid), add)
            update(v * 2 + 1, mid + 1, tr, max(l, mid + 1), r, add)
            sum_[v] = (sum_[v * 2] + sum_[v * 2 + 1]) % MOD
            sumsq[v] = (sumsq[v * 2] + sumsq[v * 2 + 1]) % MOD

        def query(v, tl, tr, l, r):
            if l > r:
                return 0
            if l == tl and r == tr:
                return sumsq[v]
            push(v, tl, tr)
            mid = (tl + tr) // 2
            return (query(v * 2, tl, mid, l, min(r, mid)) +
                    query(v * 2 + 1, mid + 1, tr, max(l, mid + 1), r)) % MOD

        last = [-1] * 100001
        ans = 0
        for r in range(n):
            prv = last[nums[r]]
            if prv + 1 <= r - 1:
                update(1, 0, n - 1, prv + 1, r - 1, 1)
            update(1, 0, n - 1, r, r, 1)
            ans = (ans + query(1, 0, n - 1, 0, r)) % MOD
            last[nums[r]] = r
        return ans
```

### Go
```go
func sumCounts(nums []int) int {
    const MOD = 1000000007
    n := len(nums)
    type Node struct {
        sum, sumsq, lazy int64
    }
    tree := make([]Node, 4*n)
    var push func(v, tl, tr int)
    push = func(v, tl, tr int) {
        if tree[v].lazy != 0 && tl != tr {
            mid := (tl + tr) >> 1
            llen := int64(mid - tl + 1)
            rlen := int64(tr - mid)
            add := tree[v].lazy
            tree[v<<1].sumsq = (tree[v<<1].sumsq + 2*add*tree[v<<1].sum + add*add%MOD*llen) % MOD
            tree[v<<1].sum = (tree[v<<1].sum + add*llen) % MOD
            tree[v<<1].lazy = (tree[v<<1].lazy + add) % MOD
            tree[v<<1|1].sumsq = (tree[v<<1|1].sumsq + 2*add*tree[v<<1|1].sum + add*add%MOD*rlen) % MOD
            tree[v<<1|1].sum = (tree[v<<1|1].sum + add*rlen) % MOD
            tree[v<<1|1].lazy = (tree[v<<1|1].lazy + add) % MOD
            tree[v].lazy = 0
        }
    }
    var update func(v, tl, tr, l, r int, add int64)
    update = func(v, tl, tr, l, r int, add int64) {
        if l > r {
            return
        }
        if l == tl && r == tr {
            length := int64(tr - tl + 1)
            tree[v].sumsq = (tree[v].sumsq + 2*add*tree[v].sum + add*add%MOD*length) % MOD
            tree[v].sum = (tree[v].sum + add*length) % MOD
            tree[v].lazy = (tree[v].lazy + add) % MOD
            return
        }
        push(v, tl, tr)
        mid := (tl + tr) >> 1
        update(v<<1, tl, mid, l, min(r, mid), add)
        update(v<<1|1, mid+1, tr, max(l, mid+1), r, add)
        tree[v].sum = (tree[v<<1].sum + tree[v<<1|1].sum) % MOD
        tree[v].sumsq = (tree[v<<1].sumsq + tree[v<<1|1].sumsq) % MOD
    }
    var query func(v, tl, tr, l, r int) int64
    query = func(v, tl, tr, l, r int) int64 {
        if l > r {
            return 0
        }
        if l == tl && r == tr {
            return tree[v].sumsq
        }
        push(v, tl, tr)
        mid := (tl + tr) >> 1
        return (query(v<<1, tl, mid, l, min(r, mid)) + query(v<<1|1, mid+1, tr, max(l, mid+1), r)) % MOD
    }
    last := make([]int, 100001)
    for i := range last {
        last[i] = -1
    }
    var ans int64
    for r := 0; r < n; r++ {
        prv := last[nums[r]]
        if prv+1 <= r-1 {
            update(1, 0, n-1, prv+1, r-1, 1)
        }
        update(1, 0, n-1, r, r, 1)
        ans = (ans + query(1, 0, n-1, 0, r)) % MOD
        last[nums[r]] = r
    }
    return int(ans)
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea is identical in every language; only the syntax changes.

I allocate a segment tree large enough to cover the interval [0, n – 1]. Each node stores three numbers: the sum of the current \(d\) values, the sum of their squares, and a lazy tag.

When a range needs to be increased by a constant `add`, I apply the three arithmetic updates shown earlier and also accumulate `add` into the lazy field. All arithmetic is performed modulo \(10^9 + 7\). Because the update formula is linear, working under the modulus from the beginning stays mathematically correct and prevents overflow.

The push routine is the classic lazy-propagation step: if a node still has a non-zero lazy value and it is not a leaf, the value is pushed onto its two children and then cleared.

In the main loop I walk from left to right. For the current right endpoint I fetch the previous occurrence of the number. All left endpoints that lie strictly after that previous occurrence receive +1; that is one range-update call. I then force the single-element subarray at the current position to have distinct count 1. At that moment the tree already holds the correct distinct counts for every subarray ending at the current index, so a single prefix query of the sum of squares gives me the contribution that must be added to the answer.

Finally I write the current index into the last-occurrence table so future appearances of the same value know where the previous one was.

Edge-case handling is automatic: when a value appears for the first time the previous index is –1, so the whole prefix [0 \ldots r – 1] receives +1 and the new position is set to 1, which is exactly what we want.

## Examples

**Example 1**  
Input: `nums = [1, 2, 1]`  
Output: `15`

- After index 0: only subarray [1] → \(1^2 = 1\)  
- After index 1: subarrays [1], [2], [1,2] → \(1^2 + 1^2 + 2^2 = 6\) (running total 7, wait – the final sum is 15)  
- After index 2: the six subarrays contribute \(1 + 1 + 1 + 4 + 4 + 4 = 15\)

**Example 2**  
Input: `nums = [2, 2]`  
Output: `3`

- After index 0: [2] → \(1^2 = 1\)  
- After index 1: [2], [2], [2,2] → \(1 + 1 + 1 = 3\)

**Example 3**  
Input: `nums = [1]`  
Output: `1`  
(The single subarray has one distinct value, \(1^2 = 1\).)

## How to Use / Run Locally

1. Copy the code for the language you want into a file (`solution.cpp`, `Solution.java`, `solution.js`, etc.).  
2. Make sure the function signature matches the LeetCode template for that language.  
3. Compile / run with the usual commands:  
   - C++: `g++ -std=c++17 solution.cpp -o solution && ./solution`  
   - Java: `javac Solution.java && java Solution`  
   - JavaScript: `node solution.js`  
   - TypeScript: `tsc solution.ts && node solution.js`  
   - Python3: `python3 solution.py`  
   - Go: `go run solution.go`  
4. Feed the sample inputs through standard input or hard-code them inside a small driver `main` if you want to test offline.

## Notes & Optimizations

- The segment tree stores modular values from the start. This is safe because the update formula is linear; any multiple of the modulus disappears.  
- A Fenwick tree alone cannot maintain both sum and sum of squares under range updates, so the segment tree is the natural choice.  
- If the value range were larger than \(10^5\) a hash map could replace the last-occurrence array, but the given constraints make a plain array faster and simpler.  
- The same technique works for many “sum of f(distinct count)” problems; only the closed-form update of the maintained aggregate changes.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)