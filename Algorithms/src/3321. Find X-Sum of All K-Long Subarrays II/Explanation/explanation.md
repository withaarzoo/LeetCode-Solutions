# Problem Title

**LeetCode 3321 – Find X-Sum of All K-Long Subarrays II**

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

For each subarray of length `k` in `nums`, compute its **x-sum**:

1. Count the frequency of each distinct value inside the subarray.
2. Keep only the **top `x` pairs** according to:

   * Higher frequency first.
   * If frequencies tie, higher value first.
3. The x-sum is the sum of `value * frequency` over the chosen pairs.

Return an array of these x-sums for all sliding windows of size `k`.

---

## Constraints

* `1 ≤ n = nums.length`
* `1 ≤ k ≤ n`
* `1 ≤ nums[i] ≤ 1e9` (values can be large; only ordering matters)
* `1 ≤ x ≤ number of distinct elements in window` (if fewer distinct elements exist, we just take all of them)

---

## Intuition

I wanted a sliding window that can **instantly** tell me the x-sum.
If I always know the **best `x` (freq, value)** pairs in the current window, the answer is just:

```
Σ (value * frequency)   over those best x pairs
```

So I keep the window’s elements split into two groups:

* **TOP**: exactly the best `x` pairs.
* **REST**: everything else.

Then, when I slide the window (one element enters and one leaves), only a **few pairs change frequency by ±1**. That means I just need tiny local fixes to maintain which pairs belong to TOP.

---

## Approach

**Invariant:**
`TOP` contains the top `x` pairs (by frequency desc, value desc).
`REST` contains the rest.
I maintain `topSum = Σ value * freq` over `TOP`.

**Add a value `v`:**

1. Remove old `(f, v)` from its current set (TOP or REST). If it was in TOP, subtract `v*f` from `topSum`.
2. Increase frequency to `f+1`.
3. Insert `(f+1, v)` directly into **TOP** and add `v*(f+1)` to `topSum`.
4. If `TOP.size() > x`, demote its **smallest** (worst among best) to REST and subtract that contribution.

**Remove a value `v`:**

1. Remove old `(f, v)` from its set; subtract from `topSum` if it was in TOP.
2. Decrease frequency to `f-1`.

   * If `f-1 > 0` → insert `(f-1, v)` into **REST**.
3. If `TOP.size() < x` and REST isn’t empty, **promote** REST’s **largest** (best among rest) to TOP and add its contribution.

**Sliding window:**

* Build the first window using the “add” operation.
* For each shift: erase left element, insert right element, record `topSum`.

---

## Data Structures Used

* `unordered_map` / `HashMap` / `Map<number, number>`: frequency of each value in the current window.
* Two ordered containers of pairs `(freq, value)` using ordering `(freq desc, value desc)`:

  * `TOP` (size ≤ `x`)
  * `REST`
* In languages without a balanced tree that supports erase by key easily (JS/Python/Go), I use heaps with **lazy deletion** plus small helper sets/indices.

---

## Operations & Behavior Summary

* **insert(value)**
  Bump frequency, move pair to TOP, demote one if TOP overflow.
* **erase(value)**
  Drop old pair, decrease frequency, push to REST if still positive, promote one if TOP underfilled.
* Every step touches O(log U) elements where `U` is current distinct count (≤ `k`).

---

## Complexity

* **Time:** `O(n log U)`
  `n` = length of `nums`, `U ≤ k` is the number of distinct elements inside the window.
  Each add/remove is a balanced-tree or heap operation.
* **Space:** `O(U)`
  For frequency map + the two structures.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    using P = pair<int,int>; // (freq, value)
    struct Cmp {
        bool operator()(const P& a, const P& b) const {
            if (a.first != b.first) return a.first > b.first; // freq desc
            return a.second > b.second;                        // value desc
        }
    };

    vector<long long> findXSum(vector<int>& nums, int k, int x) {
        int n = (int)nums.size();
        vector<long long> ans(n - k + 1);

        unordered_map<int,int> cnt; cnt.reserve(n * 2);
        set<P, Cmp> top, rest; // top = best x, rest = others
        long long topSum = 0;

        auto pull = [&](int v, int f){
            P key{f, v};
            auto it = top.find(key);
            if (it != top.end()){
                topSum -= 1LL * v * f;
                top.erase(it);
            } else {
                auto jt = rest.find(key);
                if (jt != rest.end()) rest.erase(jt);
            }
        };
        auto pushToTop = [&](int v, int f){
            top.insert({f, v});
            topSum += 1LL * v * f;
        };

        auto addVal = [&](int v){
            int f = cnt[v];
            if (f) pull(v, f);
            cnt[v] = ++f;
            pushToTop(v, f);
            if ((int)top.size() > x){
                auto it = prev(top.end()); // smallest in top
                topSum -= 1LL * it->first * it->second;
                rest.insert(*it);
                top.erase(it);
            }
        };

        auto removeVal = [&](int v){
            auto itc = cnt.find(v);
            if (itc == cnt.end()) return;
            int f = itc->second;
            pull(v, f);
            if (--f == 0) {
                cnt.erase(itc);
            } else {
                cnt[v] = f;
                rest.insert({f, v});
            }
            if ((int)top.size() < x && !rest.empty()){
                auto best = rest.begin(); // largest in rest
                topSum += 1LL * best->first * best->second;
                top.insert(*best);
                rest.erase(best);
            }
        };

        for (int i = 0; i < k; ++i) addVal(nums[i]);
        ans[0] = topSum;
        for (int i = k; i < n; ++i){
            removeVal(nums[i-k]);
            addVal(nums[i]);
            ans[i-k+1] = topSum;
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    private static final class Pair {
        final int f, v;
        Pair(int f, int v){ this.f = f; this.v = v; }
        @Override public boolean equals(Object o){
            if (this == o) return true;
            if (!(o instanceof Pair)) return false;
            Pair p = (Pair)o; return f == p.f && v == p.v;
        }
        @Override public int hashCode(){ return Objects.hash(f, v); }
    }
    private static final Comparator<Pair> DESC = (a, b) -> {
        if (a.f != b.f) return Integer.compare(b.f, a.f);
        return Integer.compare(b.v, a.v);
    };

    private Map<Integer,Integer> cnt;
    private TreeSet<Pair> top, rest;
    private long topSum;

    private void pull(int v, int f){
        Pair key = new Pair(f, v);
        if (top.remove(key)) topSum -= 1L * f * v;
        else rest.remove(key);
    }
    private void pushToTop(int v, int f){
        top.add(new Pair(f, v));
        topSum += 1L * f * v;
    }
    private void addVal(int v, int x){
        int f = cnt.getOrDefault(v, 0);
        if (f > 0) pull(v, f);
        f += 1; cnt.put(v, f);
        pushToTop(v, f);
        if (top.size() > x){
            Pair worst = top.last();
            top.remove(worst);
            topSum -= 1L * worst.f * worst.v;
            rest.add(worst);
        }
    }
    private void removeVal(int v, int x){
        Integer F = cnt.get(v);
        if (F == null || F == 0) return;
        int f = F;
        pull(v, f);
        f -= 1;
        if (f == 0) cnt.remove(v);
        else { cnt.put(v, f); rest.add(new Pair(f, v)); }
        if (top.size() < x && !rest.isEmpty()){
            Pair best = rest.first();
            rest.remove(best);
            top.add(best);
            topSum += 1L * best.f * best.v;
        }
    }

    public long[] findXSum(int[] nums, int k, int x) {
        int n = nums.length;
        long[] ans = new long[n - k + 1];

        cnt = new HashMap<>(Math.max(16, n * 2));
        top = new TreeSet<>(DESC);
        rest = new TreeSet<>(DESC);
        topSum = 0;

        for (int i = 0; i < k; ++i) addVal(nums[i], x);
        ans[0] = topSum;
        for (int i = k; i < n; ++i){
            removeVal(nums[i - k], x);
            addVal(nums[i], x);
            ans[i - k + 1] = topSum;
        }
        return ans;
    }
}
```

### JavaScript

```javascript
// Heap-based with lazy deletion (keeps chosen values in TOP)
var findXSum = function(nums, k, x) {
  class PQ {
    constructor(cmp){ this.a = []; this.cmp = cmp; }
    size(){ return this.a.length; }
    peek(){ return this.a[0]; }
    push(v){ this.a.push(v); this._up(this.a.length - 1); }
    pop(){ const n = this.a.length; if (!n) return;
      [this.a[0], this.a[n-1]] = [this.a[n-1], this.a[0]];
      const v = this.a.pop(); this._down(0); return v; }
    _up(i){ while (i){
      const p = (i-1)>>1;
      if (this.cmp(this.a[p], this.a[i]) <= 0) break;
      [this.a[p], this.a[i]] = [this.a[i], this.a[p]]; i = p; } }
    _down(i){ const n = this.a.length; for(;;){
      let l=i*2+1, r=l+1, b=i;
      if (l<n && this.cmp(this.a[b], this.a[l])>0) b=l;
      if (r<n && this.cmp(this.a[b], this.a[r])>0) b=r;
      if (b===i) break; [this.a[b], this.a[i]]=[this.a[i], this.a[b]]; i=b; } }
  }
  const minCmp = (a,b)=> a[0]!==b[0]? a[0]-b[0] : a[1]-b[1]; // (f asc, v asc)
  const maxCmp = (a,b)=> a[0]!==b[0]? b[0]-a[0] : b[1]-a[1]; // (f desc, v desc)

  const n = nums.length, ans = new Array(n - k + 1);
  const freq = new Map();
  const chosen = new Set();
  const hot = new PQ(minCmp);   // worst in TOP
  const pool = new PQ(maxCmp);  // best in REST
  let sum = 0n;

  const clean = () => {
    while (hot.size()){
      const [f, v] = hot.peek();
      if (chosen.has(v) && (freq.get(v)||0) === f) break;
      hot.pop();
    }
    while (pool.size()){
      const [f, v] = pool.peek();
      if (!chosen.has(v) && (freq.get(v)||0) === f && f > 0) break;
      pool.pop();
    }
  };
  const demoteIfChosen = (v) => {
    if (chosen.has(v)) {
      chosen.delete(v);
      const f = freq.get(v) || 0;
      sum -= BigInt(v) * BigInt(f);
    }
  };
  const promoteIfNeeded = () => {
    clean();
    while (chosen.size < x && pool.size()){
      const [f, v] = pool.pop();
      if ((freq.get(v)||0) !== f || chosen.has(v) || f === 0) continue;
      chosen.add(v);
      sum += BigInt(v) * BigInt(f);
      hot.push([f, v]);
      clean();
    }
  };

  const addOne = (v) => {
    demoteIfChosen(v);
    const f = (freq.get(v)||0) + 1;
    freq.set(v, f);
    pool.push([f, v]);
    if (chosen.size < x) {
      promoteIfNeeded();
    } else {
      clean();
      if (pool.size() && hot.size()){
        const [bf, bv] = pool.peek();
        const [wf, wv] = hot.peek();
        if (bf > wf || (bf === wf && bv > wv)) {
          pool.pop();
          chosen.add(bv);
          sum += BigInt(bv) * BigInt(bf);
          hot.push([bf, bv]);
          clean();
          const [df, dv] = hot.pop(); // demote worst
          if (chosen.has(dv) && (freq.get(dv)||0) === df) {
            chosen.delete(dv);
            sum -= BigInt(dv) * BigInt(df);
            pool.push([df, dv]);
          }
          clean();
        }
      }
    }
  };
  const removeOne = (v) => {
    demoteIfChosen(v);
    const f = (freq.get(v)||0) - 1;
    if (f <= 0) freq.delete(v); else { freq.set(v, f); pool.push([f, v]); }
    promoteIfNeeded();
  };

  for (let i = 0; i < k; ++i) addOne(nums[i]);
  ans[0] = Number(sum);
  for (let i = k; i < n; ++i) {
    removeOne(nums[i - k]);
    addOne(nums[i]);
    ans[i - k + 1] = Number(sum);
  }
  return ans;
};
```

### Python3

```python
from typing import List
import heapq

class Solution:
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        n = len(nums)
        ans = [0]*(n-k+1)

        cnt = {}
        chosen = set()             # values currently in TOP

        hot: list[tuple[int,int]] = []   # min-heap (freq, value): worst in TOP
        pool: list[tuple[int,int]] = []  # max-heap via negatives (-freq, -value)
        total = 0

        def clean():
            while hot and (hot[0][1] not in chosen or cnt.get(hot[0][1], 0) != hot[0][0]):
                heapq.heappop(hot)
            while pool and ((-pool[0][1]) in chosen or cnt.get(-pool[0][1], 0) != -pool[0][0] or -pool[0][0] == 0):
                heapq.heappop(pool)

        def demote_if_chosen(v: int):
            nonlocal total
            if v in chosen:
                chosen.remove(v)
                total -= v * cnt.get(v, 0)

        def promote_if_needed():
            nonlocal total
            clean()
            while len(chosen) < x and pool:
                f, v = -pool[0][0], -pool[0][1]
                if cnt.get(v, 0) != f or v in chosen or f == 0:
                    heapq.heappop(pool)
                    continue
                heapq.heappop(pool)
                chosen.add(v)
                total += v * f
                heapq.heappush(hot, (f, v))
            clean()

        def add_one(v: int):
            nonlocal total
            demote_if_chosen(v)
            f = cnt.get(v, 0) + 1
            cnt[v] = f
            heapq.heappush(pool, (-f, -v))
            if len(chosen) < x:
                promote_if_needed()
            else:
                clean()
                if pool and hot:
                    bf, bv = -pool[0][0], -pool[0][1]
                    wf, wv = hot[0]
                    if bf > wf or (bf == wf and bv > wv):
                        heapq.heappop(pool)
                        chosen.add(bv)
                        total += bv * bf
                        heapq.heappush(hot, (bf, bv))
                        heapq.heappop(hot)
                        if wv in chosen:
                            chosen.remove(wv)
                            total -= wv * wf
                        heapq.heappush(pool, (-wf, -wv))
                clean()

        def remove_one(v: int):
            nonlocal total
            demote_if_chosen(v)
            f = cnt.get(v, 0) - 1
            if f <= 0:
                cnt.pop(v, None)
            else:
                cnt[v] = f
                heapq.heappush(pool, (-f, -v))
            promote_if_needed()

        for i in range(k):
            add_one(nums[i])
        ans[0] = total
        for i in range(k, n):
            remove_one(nums[i-k])
            add_one(nums[i])
            ans[i-k+1] = total
        return ans
```

### Go

```go
package main

import "container/heap"

type item struct {
 val   int
 freq  int
 idx   int
 inTop bool
}

type hotHeap []*item // min-heap by (freq asc, value asc)
func (h hotHeap) Len() int { return len(h) }
func (h hotHeap) Less(i, j int) bool {
 if h[i].freq != h[j].freq { return h[i].freq < h[j].freq }
 return h[i].val < h[j].val
}
func (h hotHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i]; h[i].idx, h[j].idx = i, j }
func (h *hotHeap) Push(x interface{}){ it := x.(*item); it.idx = len(*h); *h = append(*h, it) }
func (h *hotHeap) Pop() interface{} { old := *h; it := old[len(old)-1]; *h = old[:len(old)-1]; return it }

type restHeap []*item // max-heap by (freq desc, value desc)
func (h restHeap) Len() int { return len(h) }
func (h restHeap) Less(i, j int) bool {
 if h[i].freq != h[j].freq { return h[i].freq > h[j].freq }
 return h[i].val > h[j].val
}
func (h restHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i]; h[i].idx, h[j].idx = i, j }
func (h *restHeap) Push(x interface{}){ it := x.(*item); it.idx = len(*h); *h = append(*h, it) }
func (h *restHeap) Pop() interface{} { old := *h; it := old[len(old)-1]; *h = old[:len(old)-1]; return it }

func findXSum(nums []int, k int, x int) []int64 {
 n := len(nums)
 ans := make([]int64, n-k+1)

 freq := map[int]*item{}
 hot := &hotHeap{}; rest := &restHeap{}
 heap.Init(hot); heap.Init(rest)

 var sum int64

 removeFromCurrent := func(it *item){
  if it.inTop {
   sum -= int64(it.val) * int64(it.freq)
   heap.Remove(hot, it.idx)
   it.inTop = false
  } else if it.freq > 0 && it.idx >= 0 && it.idx < rest.Len() && (*rest)[it.idx] == it {
   heap.Remove(rest, it.idx)
  }
 }

 promoteIfPossible := func(){
  for hot.Len() < x && rest.Len() > 0 {
   best := heap.Pop(rest).(*item)
   best.inTop = true
   sum += int64(best.val) * int64(best.freq)
   heap.Push(hot, best)
  }
 }

 addVal := func(v int){
  it, ok := freq[v]
  if !ok {
   it = &item{val: v, idx: -1}
   freq[v] = it
  } else {
   removeFromCurrent(it)
  }
  it.freq++
  it.inTop = true
  sum += int64(it.val) * int64(it.freq)
  heap.Push(hot, it)
  if hot.Len() > x {
   worst := heap.Pop(hot).(*item)
   sum -= int64(worst.val) * int64(worst.freq)
   worst.inTop = false
   heap.Push(rest, worst)
  }
 }

 removeVal := func(v int){
  it := freq[v]
  removeFromCurrent(it)
  it.freq--
  if it.freq == 0 {
   delete(freq, v)
   it.idx, it.inTop = -1, false
  } else {
   heap.Push(rest, it)
  }
  promoteIfPossible()
 }

 for i := 0; i < k; i++ { addVal(nums[i]) }
 ans[0] = sum
 for i := k; i < n; i++ {
  removeVal(nums[i-k])
  addVal(nums[i])
  ans[i-k+1] = sum
 }
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. **Frequency Map (`cnt`)**
   Tracks current counts in the sliding window. On every add/remove we update `cnt[v]`.

2. **Two Containers – `TOP` and `REST`**

   * Ordered by `(freq desc, value desc)`.
   * In C++/Java: `set` / `TreeSet` supports erase-by-key directly.
   * In JS/Python/Go: two heaps with **lazy deletion** are used:

     * `hot` = min-heap of `(freq, value)` keeps the *worst* of TOP at the root.
     * `pool` = max-heap of candidates for TOP (REST’s best).

3. **Adding a value `v`**

   * Remove `(oldFreq, v)` from whichever structure it sits in. Subtract `v*oldFreq` from the running sum if it was in TOP.
   * Increase to `newFreq`. Insert `(newFreq, v)` into **TOP**. Add `v*newFreq` to the sum.
   * If `TOP` overflows size `x`, demote its smallest to `REST` and subtract its contribution.

4. **Removing a value `v`**

   * Remove its current pair; subtract if it was in TOP.
   * Decrease to `newFreq`. If `newFreq > 0`, put `(newFreq, v)` into REST.
   * If `TOP` has fewer than `x` items, promote the best from REST to TOP and add its contribution.

5. **Answer**

   * After building the first window, the current `topSum` is the first output.
   * Then slide one step at a time: remove left, add right, and append the updated `topSum`.

---

## Examples

**Example 1**

```
nums = [1,1,2,2,3,4,2,3], k = 6, x = 2
Output: [6, 10, 12]
```

* Window [1..6]: most frequent are 1(2),2(2) → 1*2 + 2*2 = 6
* Window [2..7]: keep 2(3),4(1) vs 1(1),3(1) → best two: 2(3),4(1) → 2*3 + 4*1 = 10
* Window [3..8]: best two: 2(3),3(2) → 2*3 + 3*2 = 12

**Example 2**

```
nums = [3,8,7,8,7,5], k = 2, x = 2
Each k-window has ≤ 2 distinct → x-sum = sum of the subarray.
Output: [11, 15, 15, 12, 12]
```

---

## How to use / Run locally

* **C++**

  ```bash
  g++ -std=c++17 -O2 main.cpp && ./a.out
  ```

* **Java**

  ```bash
  javac Solution.java && java Solution
  ```

* **JavaScript (Node)**

  ```bash
  node solution.js
  ```

* **Python3**

  ```bash
  python3 solution.py
  ```

* **Go**

  ```bash
  go run main.go
  ```

Each sample `main` can instantiate the `Solution` class and call `findXSum(nums, k, x)`.

---

## Notes & Optimizations

* Pair ordering is **critical**: `(freq desc, value desc)`.
  This guarantees that the *smallest* in `TOP` and the *largest* in `REST` define the boundary.
* Using sets (C++/Java) avoids lazy deletion and is the cleanest/fastest.
  In JS/Python/Go, heaps with lazy deletion are practical and efficient enough.
* Values may be up to `1e9`, but we never sort by values alone—only inside `(freq, value)` pairs.
* If `x` ≥ number of distinct values in the window, `TOP` will simply hold all distincts — the x-sum equals `sum(value * freq)` which equals `sum(window)`.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
