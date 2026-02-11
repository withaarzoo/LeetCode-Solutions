# Longest Balanced Subarray II

---

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given an integer array `nums`.

A subarray is called **balanced** if:

* The number of **distinct even numbers**
* Equals the number of **distinct odd numbers**

inside that subarray.

We need to return the **length of the longest balanced subarray**.

Important: We are counting **distinct numbers**, not total occurrences.

---

## Constraints

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^5`

So we must design an **O(n log n)** or better solution.

---

## Intuition

When I first read the problem, I noticed something important:

We don’t care about frequency.

We only care about whether a number appears at least once inside the subarray.

So I thought:

Let’s define:

```bash
diff = (# distinct odd) - (# distinct even)
```

If `diff == 0`, then the subarray is balanced.

Now the problem becomes:

For each starting index `l`, I want to find the **largest `r`** such that:

```bash
diff(l, r) == 0
```

But recalculating distinct counts for every subarray would be too slow.

So I transformed the problem into:

* Each distinct value contributes:

  * `+1` if odd
  * `-1` if even

And I use a **Segment Tree with Lazy Propagation** to maintain these contributions efficiently.

---

## Approach

### Step 1: Store all occurrences

For every number `v`, store all indices where it appears:

```bash
pos[v] = [i1, i2, i3...]
```

---

### Step 2: Initialize contributions

For each distinct value:

* Let its first occurrence be `p`
* If value is odd → add `+1`
* If value is even → add `-1`

Apply this contribution to range:

```bash
[p, n-1]
```

Now the segment tree at index `r` represents:

```bash
diff(0, r)
```

---

### Step 3: Slide left pointer

For each `l` from `0` to `n-1`:

1. Query segment tree to find the **rightmost index r ≥ l where value = 0**
2. Update answer
3. Move the first occurrence pointer of `nums[l]`
4. Remove its old contribution

This removal becomes:

```bash
add -sign to range [l, nextOccurrence - 1]
```

Everything is done in O(log n).

---

## Data Structures Used

* HashMap → store occurrences
* Lazy Segment Tree → range add + find zero
* Pointers → track current first occurrence

---

## Operations & Behavior Summary

| Operation           | Purpose                           |
| ------------------- | --------------------------------- |
| Build pos map       | Track occurrences                 |
| Range add           | Apply contribution                |
| Lazy propagation    | Optimize range updates            |
| Find rightmost zero | Detect balanced subarray          |
| Pointer shift       | Maintain correctness when sliding |

---

## Complexity

### Time Complexity

`O(n log n)`

* n updates
* n queries
* Each takes O(log n)

Where:

* n = nums.length

---

### Space Complexity

`O(n)`

* Segment tree
* Position map
* Pointer map

---

# Multi-language Solutions

---

## C++

```cpp
/* C++ (optimized, commented) */
#include <bits/stdc++.h>
using namespace std;

struct SegTree
{
    int n;
    vector<int> mn, mx, lazy;
    SegTree(int _n) : n(_n), mn(4 * n, 0), mx(4 * n, 0), lazy(4 * n, 0) {}
    void apply(int idx, int v)
    {
        mn[idx] += v;
        mx[idx] += v;
        lazy[idx] += v;
    }
    void push(int idx)
    {
        if (lazy[idx] != 0)
        {
            apply(idx << 1, lazy[idx]);
            apply(idx << 1 | 1, lazy[idx]);
            lazy[idx] = 0;
        }
    }
    void pull(int idx)
    {
        mn[idx] = min(mn[idx << 1], mn[idx << 1 | 1]);
        mx[idx] = max(mx[idx << 1], mx[idx << 1 | 1]);
    }
    void add_range(int idx, int l, int r, int ql, int qr, int val)
    {
        if (ql > qr)
            return;
        if (ql <= l && r <= qr)
        {
            apply(idx, val);
            return;
        }
        push(idx);
        int mid = (l + r) >> 1;
        if (ql <= mid)
            add_range(idx << 1, l, mid, ql, min(qr, mid), val);
        if (qr > mid)
            add_range(idx << 1 | 1, mid + 1, r, max(ql, mid + 1), qr, val);
        pull(idx);
    }
    // public wrapper
    void add_range(int l, int r, int val)
    {
        if (l > r)
            return;
        add_range(1, 0, n - 1, l, r, val);
    }
    // find rightmost index in [ql, qr] with value == 0, or -1 if none
    int find_rightmost_zero(int idx, int l, int r, int ql, int qr)
    {
        if (ql > qr || qr < l || ql > r)
            return -1;
        if (mn[idx] > 0 || mx[idx] < 0)
            return -1; // no zero inside
        if (l == r)
        {
            if (mn[idx] == 0)
                return l;
            return -1;
        }
        push(idx);
        int mid = (l + r) >> 1;
        // try right child first to get rightmost
        if (qr > mid)
        {
            int res = find_rightmost_zero(idx << 1 | 1, mid + 1, r, max(ql, mid + 1), qr);
            if (res != -1)
                return res;
        }
        if (ql <= mid)
        {
            return find_rightmost_zero(idx << 1, l, mid, ql, min(qr, mid));
        }
        return -1;
    }
    int find_rightmost_zero(int ql, int qr)
    {
        if (ql > qr)
            return -1;
        return find_rightmost_zero(1, 0, n - 1, ql, qr);
    }
};

class Solution
{
public:
    int longestBalanced(vector<int> &nums)
    {
        int n = nums.size();
        unordered_map<int, vector<int>> pos;
        pos.reserve(n * 2);
        for (int i = 0; i < n; ++i)
            pos[nums[i]].push_back(i);

        SegTree st(n);
        // initial: for each value, add sign to [firstPos, n-1]
        for (auto &kv : pos)
        {
            int val = kv.first;
            int sign = (val & 1) ? 1 : -1;
            int p = kv.second[0];
            st.add_range(p, n - 1, sign);
        }

        // pointers to current first occurrence for each value
        unordered_map<int, int> ptr;
        ptr.reserve(pos.size() * 2);
        for (auto &kv : pos)
            ptr[kv.first] = 0;

        int ans = 0;
        for (int l = 0; l < n; ++l)
        {
            int r = st.find_rightmost_zero(l, n - 1);
            if (r != -1)
                ans = max(ans, r - l + 1);

            int x = nums[l];
            int pIndex = ptr[x]; // should point to l
            // move pointer forward
            ptr[x] = pIndex + 1;
            int nextPos = (ptr[x] < (int)pos[x].size()) ? pos[x][ptr[x]] : n;
            int sign = (x & 1) ? 1 : -1;
            // net effect: apply -sign to range [l, nextPos-1]
            int L = l, R = nextPos - 1;
            if (L <= R)
                st.add_range(L, R, -sign);
        }
        return ans;
    }
};
```

---

## Java

```java

// Java (clear, commented)
import java.util.*;

class Solution {
    static class SegTree {
        int n;
        int[] mn, mx, lazy;

        SegTree(int n) {
            this.n = n;
            mn = new int[4 * n];
            mx = new int[4 * n];
            lazy = new int[4 * n];
            // arrays default 0
        }

        void apply(int idx, int v) {
            mn[idx] += v;
            mx[idx] += v;
            lazy[idx] += v;
        }

        void push(int idx) {
            int z = lazy[idx];
            if (z != 0) {
                apply(idx << 1, z);
                apply(idx << 1 | 1, z);
                lazy[idx] = 0;
            }
        }

        void pull(int idx) {
            mn[idx] = Math.min(mn[idx << 1], mn[idx << 1 | 1]);
            mx[idx] = Math.max(mx[idx << 1], mx[idx << 1 | 1]);
        }

        void addRange(int idx, int l, int r, int ql, int qr, int val) {
            if (ql > qr)
                return;
            if (ql <= l && r <= qr) {
                apply(idx, val);
                return;
            }
            push(idx);
            int mid = (l + r) >> 1;
            if (ql <= mid)
                addRange(idx << 1, l, mid, ql, Math.min(qr, mid), val);
            if (qr > mid)
                addRange(idx << 1 | 1, mid + 1, r, Math.max(ql, mid + 1), qr, val);
            pull(idx);
        }

        void addRange(int l, int r, int v) {
            if (l > r)
                return;
            addRange(1, 0, n - 1, l, r, v);
        }

        int findRightmostZero(int idx, int l, int r, int ql, int qr) {
            if (ql > qr || qr < l || ql > r)
                return -1;
            if (mn[idx] > 0 || mx[idx] < 0)
                return -1;
            if (l == r) {
                return mn[idx] == 0 ? l : -1;
            }
            push(idx);
            int mid = (l + r) >> 1;
            if (qr > mid) {
                int res = findRightmostZero(idx << 1 | 1, mid + 1, r, Math.max(ql, mid + 1), qr);
                if (res != -1)
                    return res;
            }
            if (ql <= mid) {
                return findRightmostZero(idx << 1, l, mid, ql, Math.min(qr, mid));
            }
            return -1;
        }

        int findRightmostZero(int l, int r) {
            if (l > r)
                return -1;
            return findRightmostZero(1, 0, n - 1, l, r);
        }
    }

    public int longestBalanced(int[] nums) {
        int n = nums.length;
        HashMap<Integer, ArrayList<Integer>> pos = new HashMap<>();
        for (int i = 0; i < n; ++i) {
            pos.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }
        SegTree st = new SegTree(n);
        for (Map.Entry<Integer, ArrayList<Integer>> e : pos.entrySet()) {
            int val = e.getKey();
            int sign = (val % 2 == 1) ? 1 : -1;
            int p = e.getValue().get(0);
            st.addRange(p, n - 1, sign);
        }
        HashMap<Integer, Integer> ptr = new HashMap<>();
        for (int k : pos.keySet())
            ptr.put(k, 0);

        int ans = 0;
        for (int l = 0; l < n; ++l) {
            int r = st.findRightmostZero(l, n - 1);
            if (r != -1)
                ans = Math.max(ans, r - l + 1);

            int x = nums[l];
            int pIndex = ptr.get(x);
            ptr.put(x, pIndex + 1);
            ArrayList<Integer> lst = pos.get(x);
            int nextPos = (pIndex + 1 < lst.size()) ? lst.get(pIndex + 1) : n;
            int sign = (x % 2 == 1) ? 1 : -1;
            int L = l, R = nextPos - 1;
            if (L <= R)
                st.addRange(L, R, -sign);
        }
        return ans;
    }
}
```

---

## JavaScript

```js
/**
 * JavaScript (Node / leetcode style)
 * @param {number[]} nums
 * @return {number}
 */
var longestBalanced = function (nums) {
  const n = nums.length;
  const pos = new Map();
  for (let i = 0; i < n; ++i) {
    if (!pos.has(nums[i])) pos.set(nums[i], []);
    pos.get(nums[i]).push(i);
  }

  // Segment tree with mn, mx, lazy
  class SegTree {
    constructor(n) {
      this.n = n;
      this.mn = new Array(4 * n).fill(0);
      this.mx = new Array(4 * n).fill(0);
      this.lazy = new Array(4 * n).fill(0);
    }
    apply(idx, v) {
      this.mn[idx] += v;
      this.mx[idx] += v;
      this.lazy[idx] += v;
    }
    push(idx) {
      const z = this.lazy[idx];
      if (z !== 0) {
        this.apply(idx << 1, z);
        this.apply((idx << 1) | 1, z);
        this.lazy[idx] = 0;
      }
    }
    pull(idx) {
      this.mn[idx] = Math.min(this.mn[idx << 1], this.mn[(idx << 1) | 1]);
      this.mx[idx] = Math.max(this.mx[idx << 1], this.mx[(idx << 1) | 1]);
    }
    addRange(idx, l, r, ql, qr, val) {
      if (ql > qr) return;
      if (ql <= l && r <= qr) {
        this.apply(idx, val);
        return;
      }
      this.push(idx);
      const mid = (l + r) >> 1;
      if (ql <= mid)
        this.addRange(idx << 1, l, mid, ql, Math.min(qr, mid), val);
      if (qr > mid)
        this.addRange(
          (idx << 1) | 1,
          mid + 1,
          r,
          Math.max(ql, mid + 1),
          qr,
          val,
        );
      this.pull(idx);
    }
    add(l, r, v) {
      if (l > r) return;
      this.addRange(1, 0, this.n - 1, l, r, v);
    }
    findRightmostZero(idx, l, r, ql, qr) {
      if (ql > qr || qr < l || ql > r) return -1;
      if (this.mn[idx] > 0 || this.mx[idx] < 0) return -1;
      if (l === r) {
        return this.mn[idx] === 0 ? l : -1;
      }
      this.push(idx);
      const mid = (l + r) >> 1;
      if (qr > mid) {
        const res = this.findRightmostZero(
          (idx << 1) | 1,
          mid + 1,
          r,
          Math.max(ql, mid + 1),
          qr,
        );
        if (res !== -1) return res;
      }
      if (ql <= mid) {
        return this.findRightmostZero(idx << 1, l, mid, ql, Math.min(qr, mid));
      }
      return -1;
    }
    findRightmost(l, r) {
      if (l > r) return -1;
      return this.findRightmostZero(1, 0, this.n - 1, l, r);
    }
  }

  const st = new SegTree(n);
  for (let [val, arr] of pos) {
    const sign = val & 1 ? 1 : -1;
    st.add(arr[0], n - 1, sign);
  }
  const ptr = new Map();
  for (let k of pos.keys()) ptr.set(k, 0);

  let ans = 0;
  for (let l = 0; l < n; ++l) {
    const r = st.findRightmost(l, n - 1);
    if (r !== -1) ans = Math.max(ans, r - l + 1);

    const x = nums[l];
    const pi = ptr.get(x);
    ptr.set(x, pi + 1);
    const arr = pos.get(x);
    const nextPos = pi + 1 < arr.length ? arr[pi + 1] : n;
    const sign = x & 1 ? 1 : -1;
    const L = l,
      R = nextPos - 1;
    if (L <= R) st.add(L, R, -sign);
  }
  return ans;
};
```

---

## Python3

```py
# Python3 (concise and commented)
from typing import List
class SegTree:
    def __init__(self, n):
        self.n = n
        self.mn = [0] * (4*n)
        self.mx = [0] * (4*n)
        self.lazy = [0] * (4*n)
    def apply(self, idx, v):
        self.mn[idx] += v
        self.mx[idx] += v
        self.lazy[idx] += v
    def push(self, idx):
        z = self.lazy[idx]
        if z:
            self.apply(idx<<1, z)
            self.apply(idx<<1|1, z)
            self.lazy[idx] = 0
    def pull(self, idx):
        self.mn[idx] = min(self.mn[idx<<1], self.mn[idx<<1|1])
        self.mx[idx] = max(self.mx[idx<<1], self.mx[idx<<1|1])
    def add_range(self, idx, l, r, ql, qr, val):
        if ql > qr: return
        if ql <= l and r <= qr:
            self.apply(idx, val); return
        self.push(idx)
        mid = (l + r) >> 1
        if ql <= mid: self.add_range(idx<<1, l, mid, ql, min(qr, mid), val)
        if qr > mid:  self.add_range(idx<<1|1, mid+1, r, max(ql, mid+1), qr, val)
        self.pull(idx)
    def add(self, l, r, v):
        if l > r: return
        self.add_range(1, 0, self.n-1, l, r, v)
    def find_rightmost_zero(self, idx, l, r, ql, qr):
        if ql > qr or qr < l or ql > r: return -1
        if self.mn[idx] > 0 or self.mx[idx] < 0: return -1
        if l == r:
            return l if self.mn[idx] == 0 else -1
        self.push(idx)
        mid = (l + r) >> 1
        if qr > mid:
            res = self.find_rightmost_zero(idx<<1|1, mid+1, r, max(ql, mid+1), qr)
            if res != -1: return res
        if ql <= mid:
            return self.find_rightmost_zero(idx<<1, l, mid, ql, min(qr, mid))
        return -1
    def find(self, l, r):
        if l > r: return -1
        return self.find_rightmost_zero(1, 0, self.n-1, l, r)

class Solution:
    def longestBalanced(self, nums: List[int]) -> int:
        n = len(nums)
        pos = {}
        for i, v in enumerate(nums):
            pos.setdefault(v, []).append(i)
        st = SegTree(n)
        for v, lst in pos.items():
            sign = 1 if (v & 1) else -1
            st.add(lst[0], n-1, sign)
        ptr = {v:0 for v in pos}
        ans = 0
        for l in range(n):
            r = st.find(l, n-1)
            if r != -1:
                ans = max(ans, r - l + 1)
            x = nums[l]
            pIndex = ptr[x]; ptr[x] = pIndex + 1
            lst = pos[x]
            nextPos = lst[ptr[x]] if ptr[x] < len(lst) else n
            sign = 1 if (x & 1) else -1
            L, R = l, nextPos - 1
            if L <= R:
                st.add(L, R, -sign)
        return ans
```

---

## Go

```go
// Go (clear and efficient)
package main
import (
    "fmt"
)

type SegTree struct {
    n int
    mn, mx, lazy []int
}

func NewSegTree(n int) *SegTree {
    return &SegTree{
        n: n,
        mn: make([]int, 4*n),
        mx: make([]int, 4*n),
        lazy: make([]int, 4*n),
    }
}
func (st *SegTree) apply(idx, v int) {
    st.mn[idx] += v
    st.mx[idx] += v
    st.lazy[idx] += v
}
func (st *SegTree) push(idx int) {
    z := st.lazy[idx]
    if z != 0 {
        st.apply(idx<<1, z)
        st.apply(idx<<1|1, z)
        st.lazy[idx] = 0
    }
}
func (st *SegTree) pull(idx int) {
    if st.mn[idx<<1] < st.mn[idx<<1|1] {
        st.mn[idx] = st.mn[idx<<1]
    } else {
        st.mn[idx] = st.mn[idx<<1|1]
    }
    if st.mx[idx<<1] > st.mx[idx<<1|1] {
        st.mx[idx] = st.mx[idx<<1]
    } else {
        st.mx[idx] = st.mx[idx<<1|1]
    }
}
func (st *SegTree) addRange(idx, l, r, ql, qr, val int) {
    if ql > qr { return }
    if ql <= l && r <= qr {
        st.apply(idx, val); return
    }
    st.push(idx)
    mid := (l + r) >> 1
    if ql <= mid { st.addRange(idx<<1, l, mid, ql, min(qr, mid), val) }
    if qr > mid  { st.addRange(idx<<1|1, mid+1, r, max(ql, mid+1), qr, val) }
    st.pull(idx)
}
func (st *SegTree) Add(l, r, v int) {
    if l > r { return }
    st.addRange(1, 0, st.n-1, l, r, v)
}
func (st *SegTree) findRightmostZero(idx, l, r, ql, qr int) int {
    if ql > qr || qr < l || ql > r { return -1 }
    if st.mn[idx] > 0 || st.mx[idx] < 0 { return -1 }
    if l == r {
        if st.mn[idx] == 0 { return l }
        return -1
    }
    st.push(idx)
    mid := (l + r) >> 1
    if qr > mid {
        res := st.findRightmostZero(idx<<1|1, mid+1, r, max(ql, mid+1), qr)
        if res != -1 { return res }
    }
    if ql <= mid {
        return st.findRightmostZero(idx<<1, l, mid, ql, min(qr, mid))
    }
    return -1
}
func (st *SegTree) FindRightmost(l, r int) int {
    if l > r { return -1 }
    return st.findRightmostZero(1, 0, st.n-1, l, r)
}
func min(a,b int) int { if a<b { return a }; return b }
func max(a,b int) int { if a>b { return a }; return b }

// Example wrapper function (LeetCode expects a function)
func longestBalanced(nums []int) int {
    n := len(nums)
    pos := map[int][]int{}
    for i, v := range nums {
        pos[v] = append(pos[v], i)
    }
    st := NewSegTree(n)
    for v, arr := range pos {
        sign := -1
        if v&1 == 1 { sign = 1 }
        st.Add(arr[0], n-1, sign)
    }
    ptr := map[int]int{}
    for k := range pos { ptr[k] = 0 }
    ans := 0
    for l := 0; l < n; l++ {
        r := st.FindRightmost(l, n-1)
        if r != -1 && r-l+1 > ans { ans = r-l+1 }
        x := nums[l]
        pIndex := ptr[x]; ptr[x] = pIndex+1
        arr := pos[x]
        nextPos := n
        if ptr[x] < len(arr) { nextPos = arr[ptr[x]] }
        sign := -1
        if x&1 == 1 { sign = 1 }
        L, R := l, nextPos-1
        if L <= R { st.Add(L, R, -sign) }
    }
    return ans
}

// main is just for quick local test (remove on leetcode)
func main() {
    fmt.Println(longestBalanced([]int{2,5,4,3})) // expect 4
}
```

---

# Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core logic is identical in all languages:

1. Build map of positions
2. Initialize segment tree
3. Add contribution of first occurrences
4. Slide left pointer
5. Remove old contribution
6. Query for rightmost zero
7. Update answer

Each language only differs syntactically.

---

# Examples

### Example 1

```bash
Input: [2,5,4,3]
Output: 4
```

Distinct evens = {2,4}
Distinct odds = {5,3}
Balanced length = 4

---

### Example 2

```bash
Input: [3,2,2,5,4]
Output: 5
```

Distinct evens = {2,4}
Distinct odds = {3,5}

---

# How to use / Run locally

### C++

```bash
g++ solution.cpp -std=c++17
./a.out
```

### Java

```bash
javac Solution.java
java Solution
```

### Python

```bash
python3 solution.py
```

### Go

```bash
go run solution.go
```

---

# Notes & Optimizations

* We must use Lazy Propagation.
* Without segment tree, solution becomes O(n²).
* We search **rightmost zero** for maximum length.
* HashMap usage ensures O(1) average lookup.

---

# Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
