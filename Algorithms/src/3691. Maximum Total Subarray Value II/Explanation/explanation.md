# 3691. Maximum Total Subarray Value II

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

LeetCode 3691: Maximum Total Subarray Value II is a challenging array and heap-based problem that asks us to select exactly `k` distinct non-empty subarrays and maximize the total sum of their values.

The value of a subarray is defined as:

`max(subarray) - min(subarray)`

The goal is to find the maximum possible total value after selecting exactly `k` valid subarrays.

At first glance, checking every subarray seems possible, but the number of subarrays can be extremely large. Because of the given constraints, a brute-force solution is not practical. We need an efficient approach that combines range queries and greedy selection.

This problem is a great example of using:

* Sparse Table
* Range Maximum Query (RMQ)
* Range Minimum Query (RMQ)
* Priority Queue (Max Heap)
* Greedy Algorithms

## Constraints

| Constraint                             | Value                        |
| -------------------------------------- | ---------------------------- |
| `1 <= n == nums.length <= 5 * 10^4`    | Array size                   |
| `0 <= nums[i] <= 10^9`                 | Array values                 |
| `1 <= k <= min(10^5, n * (n + 1) / 2)` | Number of selected subarrays |

## Intuition

The first thing I noticed was that the value of a subarray depends only on its maximum and minimum elements.

For a fixed starting index `l`, if I keep moving the right boundary `r` toward the left, the subarray becomes smaller.

When a range becomes smaller:

* The maximum cannot increase.
* The minimum cannot decrease.

Because of this, the sequence of subarray values for a fixed starting index forms a non-increasing sequence.

That observation is the key.

Instead of generating every possible subarray, I can think of the problem as merging multiple sorted sequences and repeatedly taking the largest available value.

To make this efficient, I need fast range maximum and minimum queries. That is where Sparse Tables become useful.

## Approach

1. Build a Sparse Table for range maximum queries.

2. Build another Sparse Table for range minimum queries.

3. Create a helper function that returns:

   `max(nums[l...r]) - min(nums[l...r])`

   in constant time.

4. For every starting position `l`:

   * Compute the value of `(l, n-1)`.
   * Insert it into a max heap.

5. Repeatedly perform exactly `k` operations:

   * Extract the largest value from the heap.
   * Add it to the answer.
   * Generate the next candidate from the same sequence.
   * Push it back into the heap if it still exists.

6. After performing `k` extractions, the accumulated sum is the answer.

This allows us to efficiently find the globally largest `k` subarray values without generating all subarrays.

## Data Structures Used

### Sparse Table

Used to answer:

* Range Maximum Query
* Range Minimum Query

Both operations become `O(1)` after preprocessing.

### Max Heap (Priority Queue)

Used to always retrieve the current largest available subarray value.

This ensures the greedy selection remains efficient.

### Logarithm Lookup Array

Stores precomputed logarithm values.

This avoids repeated logarithm calculations during RMQ queries.

## Operations & Behavior Summary

1. Precompute logarithm values.
2. Build maximum sparse table.
3. Build minimum sparse table.
4. Insert the largest candidate from every starting position into the heap.
5. Extract the largest value.
6. Add it to the answer.
7. Generate the next candidate from the same sequence.
8. Push it into the heap.
9. Continue until exactly `k` values have been selected.
10. Return the final sum.

## Complexity

| Operation                 | Complexity           |
| ------------------------- | -------------------- |
| Sparse Table Construction | O(n log n)           |
| Single Range Query        | O(1)                 |
| Heap Push / Pop           | O(log n)             |
| Total Solution            | O(n log n + k log n) |
| Extra Space               | O(n log n)           |

### Time Complexity

**O(n log n + k log n)**

Where:

* `n` is the size of the array.
* `k` is the number of selected subarrays.

### Space Complexity

**O(n log n)**

The sparse tables require additional memory for storing range maximum and minimum information.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long maxTotalValue(vector<int>& nums, int k) {
        int n = nums.size();

        // Precompute floor(log2(i))
        vector<int> lg(n + 1);
        for (int i = 2; i <= n; i++) {
            lg[i] = lg[i / 2] + 1;
        }

        int K = lg[n] + 1;

        // Sparse table for maximums
        vector<vector<int>> mx(K, vector<int>(n));

        // Sparse table for minimums
        vector<vector<int>> mn(K, vector<int>(n));

        for (int i = 0; i < n; i++) {
            mx[0][i] = nums[i];
            mn[0][i] = nums[i];
        }

        // Build sparse tables
        for (int j = 1; j < K; j++) {
            for (int i = 0; i + (1 << j) <= n; i++) {
                mx[j][i] = max(mx[j - 1][i],
                               mx[j - 1][i + (1 << (j - 1))]);

                mn[j][i] = min(mn[j - 1][i],
                               mn[j - 1][i + (1 << (j - 1))]);
            }
        }

        // Returns value(l, r) = max - min in O(1)
        auto getValue = [&](int l, int r) -> long long {
            int len = r - l + 1;
            int p = lg[len];

            int mxVal = max(mx[p][l],
                            mx[p][r - (1 << p) + 1]);

            int mnVal = min(mn[p][l],
                            mn[p][r - (1 << p) + 1]);

            return 1LL * mxVal - mnVal;
        };

        struct Node {
            long long val;
            int l;
            int r;

            bool operator<(const Node& other) const {
                return val < other.val; // max heap
            }
        };

        priority_queue<Node> pq;

        // First element from every sorted sequence
        for (int l = 0; l < n; l++) {
            pq.push({getValue(l, n - 1), l, n - 1});
        }

        long long ans = 0;

        while (k--) {
            auto cur = pq.top();
            pq.pop();

            ans += cur.val;

            // Move to next value in the same sequence
            if (cur.r > cur.l) {
                pq.push({
                    getValue(cur.l, cur.r - 1),
                    cur.l,
                    cur.r - 1
                });
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    static class Node {
        long val;
        int l;
        int r;

        Node(long val, int l, int r) {
            this.val = val;
            this.l = l;
            this.r = r;
        }
    }

    public long maxTotalValue(int[] nums, int k) {
        int n = nums.length;

        int[] lg = new int[n + 1];
        for (int i = 2; i <= n; i++) {
            lg[i] = lg[i / 2] + 1;
        }

        int K = lg[n] + 1;

        int[][] mx = new int[K][n];
        int[][] mn = new int[K][n];

        for (int i = 0; i < n; i++) {
            mx[0][i] = nums[i];
            mn[0][i] = nums[i];
        }

        for (int j = 1; j < K; j++) {
            for (int i = 0; i + (1 << j) <= n; i++) {
                mx[j][i] = Math.max(
                    mx[j - 1][i],
                    mx[j - 1][i + (1 << (j - 1))]
                );

                mn[j][i] = Math.min(
                    mn[j - 1][i],
                    mn[j - 1][i + (1 << (j - 1))]
                );
            }
        }

        PriorityQueue<Node> pq = new PriorityQueue<>(
            (a, b) -> Long.compare(b.val, a.val)
        );

        for (int l = 0; l < n; l++) {
            pq.offer(new Node(getValue(l, n - 1, mx, mn, lg), l, n - 1));
        }

        long ans = 0;

        while (k-- > 0) {
            Node cur = pq.poll();

            ans += cur.val;

            if (cur.r > cur.l) {
                pq.offer(new Node(
                    getValue(cur.l, cur.r - 1, mx, mn, lg),
                    cur.l,
                    cur.r - 1
                ));
            }
        }

        return ans;
    }

    private long getValue(
        int l,
        int r,
        int[][] mx,
        int[][] mn,
        int[] lg
    ) {
        int len = r - l + 1;
        int p = lg[len];

        int mxVal = Math.max(
            mx[p][l],
            mx[p][r - (1 << p) + 1]
        );

        int mnVal = Math.min(
            mn[p][l],
            mn[p][r - (1 << p) + 1]
        );

        return (long) mxVal - mnVal;
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
var maxTotalValue = function(nums, k) {
    const n = nums.length;

    const lg = new Array(n + 1).fill(0);

    for (let i = 2; i <= n; i++) {
        lg[i] = lg[i >> 1] + 1;
    }

    const K = lg[n] + 1;

    const mx = Array.from({ length: K }, () => Array(n).fill(0));
    const mn = Array.from({ length: K }, () => Array(n).fill(0));

    for (let i = 0; i < n; i++) {
        mx[0][i] = nums[i];
        mn[0][i] = nums[i];
    }

    for (let j = 1; j < K; j++) {
        for (let i = 0; i + (1 << j) <= n; i++) {
            mx[j][i] = Math.max(
                mx[j - 1][i],
                mx[j - 1][i + (1 << (j - 1))]
            );

            mn[j][i] = Math.min(
                mn[j - 1][i],
                mn[j - 1][i + (1 << (j - 1))]
            );
        }
    }

    const getValue = (l, r) => {
        const len = r - l + 1;
        const p = lg[len];

        const mxVal = Math.max(
            mx[p][l],
            mx[p][r - (1 << p) + 1]
        );

        const mnVal = Math.min(
            mn[p][l],
            mn[p][r - (1 << p) + 1]
        );

        return mxVal - mnVal;
    };

    class MaxHeap {
        constructor() {
            this.heap = [];
        }

        push(x) {
            this.heap.push(x);

            let i = this.heap.length - 1;

            while (i > 0) {
                let p = (i - 1) >> 1;

                if (this.heap[p][0] >= this.heap[i][0]) break;

                [this.heap[p], this.heap[i]] =
                    [this.heap[i], this.heap[p]];

                i = p;
            }
        }

        pop() {
            const top = this.heap[0];
            const last = this.heap.pop();

            if (this.heap.length) {
                this.heap[0] = last;

                let i = 0;

                while (true) {
                    let largest = i;
                    let l = i * 2 + 1;
                    let r = i * 2 + 2;

                    if (
                        l < this.heap.length &&
                        this.heap[l][0] > this.heap[largest][0]
                    ) {
                        largest = l;
                    }

                    if (
                        r < this.heap.length &&
                        this.heap[r][0] > this.heap[largest][0]
                    ) {
                        largest = r;
                    }

                    if (largest === i) break;

                    [this.heap[i], this.heap[largest]] =
                        [this.heap[largest], this.heap[i]];

                    i = largest;
                }
            }

            return top;
        }
    }

    const pq = new MaxHeap();

    for (let l = 0; l < n; l++) {
        pq.push([getValue(l, n - 1), l, n - 1]);
    }

    let ans = 0;

    while (k--) {
        const [val, l, r] = pq.pop();

        ans += val;

        if (r > l) {
            pq.push([getValue(l, r - 1), l, r - 1]);
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        n = len(nums)

        # floor(log2(i))
        lg = [0] * (n + 1)
        for i in range(2, n + 1):
            lg[i] = lg[i // 2] + 1

        K = lg[n] + 1

        # Sparse table for maximums
        mx = [[0] * n for _ in range(K)]

        # Sparse table for minimums
        mn = [[0] * n for _ in range(K)]

        for i in range(n):
            mx[0][i] = nums[i]
            mn[0][i] = nums[i]

        # Build sparse tables
        for j in range(1, K):
            length = 1 << j

            for i in range(n - length + 1):
                mx[j][i] = max(
                    mx[j - 1][i],
                    mx[j - 1][i + (length >> 1)]
                )

                mn[j][i] = min(
                    mn[j - 1][i],
                    mn[j - 1][i + (length >> 1)]
                )

        # O(1) range value query
        def get_value(l, r):
            length = r - l + 1
            p = lg[length]

            mx_val = max(
                mx[p][l],
                mx[p][r - (1 << p) + 1]
            )

            mn_val = min(
                mn[p][l],
                mn[p][r - (1 << p) + 1]
            )

            return mx_val - mn_val

        import heapq

        # Python heap is min-heap, so store negative values
        pq = []

        for l in range(n):
            heapq.heappush(
                pq,
                (-get_value(l, n - 1), l, n - 1)
            )

        ans = 0

        for _ in range(k):
            neg_val, l, r = heapq.heappop(pq)

            val = -neg_val
            ans += val

            if r > l:
                heapq.heappush(
                    pq,
                    (-get_value(l, r - 1), l, r - 1)
                )

        return ans
```

### Go

```go
func maxTotalValue(nums []int, k int) int64 {
 n := len(nums)

 // floor(log2(i))
 lg := make([]int, n+1)
 for i := 2; i <= n; i++ {
  lg[i] = lg[i/2] + 1
 }

 K := lg[n] + 1

 // Sparse table for maximums
 mx := make([][]int, K)

 // Sparse table for minimums
 mn := make([][]int, K)

 for i := 0; i < K; i++ {
  mx[i] = make([]int, n)
  mn[i] = make([]int, n)
 }

 for i := 0; i < n; i++ {
  mx[0][i] = nums[i]
  mn[0][i] = nums[i]
 }

 // Build sparse tables
 for j := 1; j < K; j++ {
  for i := 0; i+(1<<j) <= n; i++ {
   a := mx[j-1][i]
   b := mx[j-1][i+(1<<(j-1))]
   if b > a {
    a = b
   }
   mx[j][i] = a

   c := mn[j-1][i]
   d := mn[j-1][i+(1<<(j-1))]
   if d < c {
    c = d
   }
   mn[j][i] = c
  }
 }

 // O(1) range value query
 getValue := func(l, r int) int64 {
  length := r - l + 1
  p := lg[length]

  a := mx[p][l]
  b := mx[p][r-(1<<p)+1]
  if b > a {
   a = b
  }

  c := mn[p][l]
  d := mn[p][r-(1<<p)+1]
  if d < c {
   c = d
  }

  return int64(a - c)
 }

 type Node struct {
  val int64
  l   int
  r   int
 }

 h := &MaxHeap{}
 heap.Init(h)

 for l := 0; l < n; l++ {
  heap.Push(h, Node{
   val: getValue(l, n-1),
   l:   l,
   r:   n - 1,
  })
 }

 var ans int64 = 0

 for k > 0 {
  cur := heap.Pop(h).(Node)

  ans += cur.val

  if cur.r > cur.l {
   heap.Push(h, Node{
    val: getValue(cur.l, cur.r-1),
    l:   cur.l,
    r:   cur.r - 1,
   })
  }

  k--
 }

 return ans
}

type MaxHeap []struct {
 val int64
 l   int
 r   int
}

func (h MaxHeap) Len() int {
 return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
 return h[i].val > h[j].val
}

func (h MaxHeap) Swap(i, j int) {
 h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
 *h = append(*h, x.(struct {
  val int64
  l   int
  r   int
 }))
}

func (h *MaxHeap) Pop() interface{} {
 old := *h
 n := len(old)

 item := old[n-1]
 *h = old[:n-1]

 return item
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains identical across all five languages.

### Step 1: Build Logarithm Table

I first precompute logarithm values for every possible length.

This allows me to quickly determine which sparse table level should be used during a query.

### Step 2: Build Sparse Tables

I maintain two separate sparse tables:

* One stores maximum values.
* One stores minimum values.

Each level represents intervals of length:

* 1
* 2
* 4
* 8
* 16

and so on.

This preprocessing enables constant-time range queries.

### Step 3: Create Range Query Function

Whenever I need the value of a subarray:

`max(nums[l...r]) - min(nums[l...r])`

I perform:

* One maximum query
* One minimum query

Both take constant time.

### Step 4: Initialize Heap

For every starting index:

`l`

I insert:

`(l, n-1)`

into the max heap.

These represent the largest available candidate from every sequence.

### Step 5: Extract Maximum Candidate

The heap always stores the largest currently available subarray value.

When I remove the top element:

* It becomes part of the answer.
* Its value is added to the total.

### Step 6: Generate Next Candidate

Suppose the extracted subarray is:

`(l, r)`

Then the next element in the same sequence is:

`(l, r-1)`

If that subarray still exists, I insert it into the heap.

### Step 7: Repeat Exactly k Times

Each extraction contributes one chosen subarray.

After exactly `k` selections, the answer contains the maximum achievable total value.

## Examples

### Example 1

Input

```text
nums = [1,3,2]
k = 2
```

Output

```text
4
```

Explanation

Selected subarrays:

```text
[1,3]
value = 3 - 1 = 2

[1,3,2]
value = 3 - 1 = 2
```

Total:

```text
2 + 2 = 4
```

---

### Example 2

Input

```text
nums = [4,2,5,1]
k = 3
```

Output

```text
12
```

Explanation

Selected subarrays:

```text
[4,2,5,1] -> 4
[2,5,1]   -> 4
[5,1]     -> 4
```

Total:

```text
4 + 4 + 4 = 12
```

---

### Example 3

Input

```text
nums = [7]
k = 1
```

Output

```text
0
```

Explanation

Only one subarray exists.

```text
max = 7
min = 7
value = 0
```

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -O2 -std=c++17
```

Run

```bash
./a.out
```

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

### JavaScript

Run

```bash
node solution.js
```

### Python3

Run

```bash
python solution.py
```

### Go

Run

```bash
go run solution.go
```

## Notes & Optimizations

* A brute-force solution would require examining all subarrays and would be far too slow.
* Sparse Tables reduce range maximum and minimum queries to constant time.
* The heap ensures we always select the next best available subarray value.
* The solution scales comfortably to the maximum constraints.
* This approach combines greedy selection with efficient range query preprocessing.
* Segment Trees could also be used for range queries, but Sparse Tables provide faster query performance for static arrays.
* All calculations should use 64-bit integers because the final answer can become very large.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
