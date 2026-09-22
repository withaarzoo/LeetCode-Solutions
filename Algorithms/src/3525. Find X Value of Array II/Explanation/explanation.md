# 3525. Find X Value of Array II – LeetCode Solution with Segment Tree

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

LeetCode 3525, Find X Value of Array II, asks us to handle a series of updates and range queries on an array of positive integers.  

You are given an array `nums` and a positive integer `k`. You also receive a list of queries. Each query has four values: an index, a new value, a start position, and a target remainder `x`.  

For every query you must:  
1. Permanently change `nums[index]` to the new value.  
2. Look only at the suffix that starts at position `start`.  
3. Count how many non-empty prefixes of that suffix have a product that leaves remainder `x` when divided by `k`.  

The answer for each query is that count. The final output is an array containing the answer for every query in order.

This problem mixes point updates with range product-modulo queries, which makes it a classic candidate for a segment tree that stores extra modular information.

## Constraints

- `1 <= nums[i] <= 10^9`
- `1 <= nums.length <= 10^5`
- `1 <= k <= 5`
- `1 <= queries.length <= 2 * 10^4`
- Each query is of the form `[index_i, value_i, start_i, x_i]`
- `0 <= index_i < nums.length`
- `1 <= value_i <= 10^9`
- `0 <= start_i < nums.length`
- `0 <= x_i < k`

The small value of `k` is the key that lets us keep frequency arrays of size at most 5 inside every segment-tree node.

## Intuition

The first thing I noticed was that `k` never exceeds 5. That means any information about products modulo `k` can be stored in a tiny array of size 5.  

What I really need for each query is the number of prefixes of a suffix whose product is congruent to a given remainder. If I can maintain, for every contiguous segment, both the total product of the segment and the frequency of every possible prefix-product remainder, then a segment tree can answer the required range queries after each update in logarithmic time.

Merging two adjacent segments is straightforward: the prefixes that stay inside the left segment keep their old remainders, and the prefixes that cross into the right segment simply multiply the left segment’s total product by each right-side prefix remainder.

## Approach

I build a segment tree over the entire array.  

Every node stores:  
- the product of all elements in its range, taken modulo `k`  
- an array of size `k` that counts how many prefixes (starting from the left end of the node) produce each possible remainder  

A leaf that holds a single number `v` sets its product to `v % k` and puts a count of 1 in the corresponding frequency bucket.  

When two children are merged, the parent’s product becomes the product of the two children. The parent’s frequency array is formed by copying the left child’s frequencies and then adding the right child’s frequencies after they have been multiplied by the left child’s total product.  

An update replaces the value at a leaf and rebuilds every ancestor by re-merging.  

A query asks the tree for the combined node that covers the suffix `[start \ldots n-1]`. The answer is simply the frequency of the requested remainder `x` stored in that node.

Because every merge costs only a few operations on arrays of length 5, the whole solution stays fast enough for the given constraints.

## Data Structures Used

- **Segment Tree** – the main structure. It supports point updates and range queries while maintaining both the modular product and the prefix-frequency array of every segment.  
- **Node object / struct** – holds the product and the frequency array of size `k`.  
- **Frequency array of size k** – records how many prefixes of a segment produce each remainder. Because `k ≤ 5` this array is extremely cheap to store and merge.

No extra heavy data structures are required; the segment tree alone is sufficient.

## Operations & Behavior Summary

1. Build the segment tree from the initial `nums` array. Each leaf records its value modulo `k` and sets a single frequency entry.  
2. For every query:  
   - Perform a point update at the given index with the new value.  
   - Query the range that starts at the given `start` and ends at the last index.  
   - Read the frequency of the target remainder `x` from the returned node and store it as the answer for that query.  
3. Return the collected answers.

The merge operation is the only non-trivial part: it correctly combines prefix frequencies across the boundary of two adjacent segments by using the left segment’s total product as a multiplier.

## Complexity

| Type              | Complexity                  | Explanation |
|-------------------|-----------------------------|-------------|
| Time Complexity   | O((n + q) · log n · k²)    | Building the tree costs O(n log n · k²). Each of the q updates and range queries costs O(log n · k²). With k ≤ 5 the quadratic factor is a tiny constant. |
| Space Complexity  | O(n · k)                   | The segment tree contains O(n) nodes and each node stores an array of length k. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    struct Node {
        int prod;
        array<int, 5> freq;
        Node() : prod(1) { freq.fill(0); }
    };
    
    int k;
    vector<Node> tree;
    vector<int> nums;
    
    Node merge(const Node& L, const Node& R) {
        Node res;
        res.prod = (1LL * L.prod * R.prod) % k;
        res.freq = L.freq;
        for (int r = 0; r < k; ++r) {
            if (R.freq[r]) {
                int nr = (1LL * L.prod * r) % k;
                res.freq[nr] += R.freq[r];
            }
        }
        return res;
    }
    
    void build(int v, int tl, int tr) {
        if (tl == tr) {
            tree[v].prod = nums[tl] % k;
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        int tm = (tl + tr) / 2;
        build(v * 2, tl, tm);
        build(v * 2 + 1, tm + 1, tr);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    void update(int v, int tl, int tr, int pos, int val) {
        if (tl == tr) {
            tree[v].prod = val % k;
            tree[v].freq.fill(0);
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        int tm = (tl + tr) / 2;
        if (pos <= tm) update(v * 2, tl, tm, pos, val);
        else update(v * 2 + 1, tm + 1, tr, pos, val);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    Node query(int v, int tl, int tr, int l, int r) {
        if (l > r) return Node();
        if (l == tl && r == tr) return tree[v];
        int tm = (tl + tr) / 2;
        return merge(query(v * 2, tl, tm, l, min(r, tm)),
                     query(v * 2 + 1, tm + 1, tr, max(l, tm + 1), r));
    }
    
    vector<int> resultArray(vector<int>& nums, int k, vector<vector<int>>& queries) {
        this->k = k;
        this->nums = nums;
        int n = nums.size();
        tree.assign(4 * n, Node());
        build(1, 0, n - 1);
        vector<int> ans;
        for (auto& q : queries) {
            int idx = q[0], val = q[1], start = q[2], x = q[3];
            update(1, 0, n - 1, idx, val);
            Node res = query(1, 0, n - 1, start, n - 1);
            ans.push_back(res.freq[x]);
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    static class Node {
        int prod;
        int[] freq;
        Node(int k) {
            prod = 1;
            freq = new int[k];
        }
    }
    
    int k;
    Node[] tree;
    int[] nums;
    
    Node merge(Node L, Node R) {
        Node res = new Node(k);
        res.prod = (int)((1L * L.prod * R.prod) % k);
        System.arraycopy(L.freq, 0, res.freq, 0, k);
        for (int r = 0; r < k; r++) {
            if (R.freq[r] != 0) {
                int nr = (int)((1L * L.prod * r) % k);
                res.freq[nr] += R.freq[r];
            }
        }
        return res;
    }
    
    void build(int v, int tl, int tr) {
        if (tl == tr) {
            tree[v].prod = nums[tl] % k;
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        int tm = (tl + tr) / 2;
        build(v * 2, tl, tm);
        build(v * 2 + 1, tm + 1, tr);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    void update(int v, int tl, int tr, int pos, int val) {
        if (tl == tr) {
            tree[v].prod = val % k;
            Arrays.fill(tree[v].freq, 0);
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        int tm = (tl + tr) / 2;
        if (pos <= tm) update(v * 2, tl, tm, pos, val);
        else update(v * 2 + 1, tm + 1, tr, pos, val);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    Node query(int v, int tl, int tr, int l, int r) {
        if (l > r) return new Node(k);
        if (l == tl && r == tr) return tree[v];
        int tm = (tl + tr) / 2;
        return merge(query(v * 2, tl, tm, l, Math.min(r, tm)),
                     query(v * 2 + 1, tm + 1, tr, Math.max(l, tm + 1), r));
    }
    
    public int[] resultArray(int[] nums, int k, int[][] queries) {
        this.k = k;
        this.nums = nums;
        int n = nums.length;
        tree = new Node[4 * n];
        for (int i = 0; i < tree.length; i++) tree[i] = new Node(k);
        build(1, 0, n - 1);
        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int idx = queries[i][0], val = queries[i][1], start = queries[i][2], x = queries[i][3];
            update(1, 0, n - 1, idx, val);
            Node res = query(1, 0, n - 1, start, n - 1);
            ans[i] = res.freq[x];
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
 * @param {number[][]} queries
 * @return {number[]}
 */
var resultArray = function(nums, k, queries) {
    const n = nums.length;
    const treeProd = new Int32Array(4 * n);
    const treeFreq = Array.from({ length: 4 * n }, () => new Int32Array(k));

    function merge(pL, fL, pR, fR, pRes, fRes) {
        pRes[0] = (pL[0] * pR[0]) % k;
        for (let i = 0; i < k; i++) fRes[i] = fL[i];
        for (let r = 0; r < k; r++) {
            if (fR[r] !== 0) {
                const nr = (pL[0] * r) % k;
                fRes[nr] += fR[r];
            }
        }
    }

    function build(v, tl, tr) {
        if (tl === tr) {
            treeProd[v] = nums[tl] % k;
            treeFreq[v][treeProd[v]] = 1;
            return;
        }
        const tm = (tl + tr) >> 1;
        build(v << 1, tl, tm);
        build(v << 1 | 1, tm + 1, tr);
        const pL = treeProd[v << 1], fL = treeFreq[v << 1];
        const pR = treeProd[v << 1 | 1], fR = treeFreq[v << 1 | 1];
        treeProd[v] = (pL * pR) % k;
        const fRes = treeFreq[v];
        fRes.fill(0);
        for (let i = 0; i < k; i++) fRes[i] = fL[i];
        for (let r = 0; r < k; r++) {
            if (fR[r]) {
                const nr = (pL * r) % k;
                fRes[nr] += fR[r];
            }
        }
    }

    function update(v, tl, tr, pos, val) {
        if (tl === tr) {
            treeProd[v] = val % k;
            treeFreq[v].fill(0);
            treeFreq[v][treeProd[v]] = 1;
            return;
        }
        const tm = (tl + tr) >> 1;
        if (pos <= tm) update(v << 1, tl, tm, pos, val);
        else update(v << 1 | 1, tm + 1, tr, pos, val);

        const pL = treeProd[v << 1], fL = treeFreq[v << 1];
        const pR = treeProd[v << 1 | 1], fR = treeFreq[v << 1 | 1];
        treeProd[v] = (pL * pR) % k;
        const fRes = treeFreq[v];
        fRes.fill(0);
        for (let i = 0; i < k; i++) fRes[i] = fL[i];
        for (let r = 0; r < k; r++) {
            if (fR[r]) {
                const nr = (pL * r) % k;
                fRes[nr] += fR[r];
            }
        }
    }

    function query(v, tl, tr, l, r, outFreq, outProd) {
        if (l > r) {
            outProd[0] = 1;
            outFreq.fill(0);
            return;
        }
        if (l === tl && r === tr) {
            outProd[0] = treeProd[v];
            for (let i = 0; i < k; i++) outFreq[i] = treeFreq[v][i];
            return;
        }
        const tm = (tl + tr) >> 1;
        const leftFreq = new Int32Array(k);
        const rightFreq = new Int32Array(k);
        const leftProd = new Int32Array(1);
        const rightProd = new Int32Array(1);

        query(v << 1, tl, tm, l, Math.min(r, tm), leftFreq, leftProd);
        query(v << 1 | 1, tm + 1, tr, Math.max(l, tm + 1), r, rightFreq, rightProd);

        outProd[0] = (leftProd[0] * rightProd[0]) % k;
        outFreq.fill(0);
        for (let i = 0; i < k; i++) outFreq[i] = leftFreq[i];
        for (let r = 0; r < k; r++) {
            if (rightFreq[r]) {
                const nr = (leftProd[0] * r) % k;
                outFreq[nr] += rightFreq[r];
            }
        }
    }

    build(1, 0, n - 1);

    const ans = new Array(queries.length);
    const tempFreq = new Int32Array(k);
    const tempProd = new Int32Array(1);

    for (let i = 0; i < queries.length; i++) {
        const idx = queries[i][0];
        const val = queries[i][1];
        const start = queries[i][2];
        const x = queries[i][3];

        update(1, 0, n - 1, idx, val);
        query(1, 0, n - 1, start, n - 1, tempFreq, tempProd);
        ans[i] = tempFreq[x];
    }
    return ans;
};
```

### TypeScript
```typescript
function resultArray(nums: number[], k: number, queries: number[][]): number[] {
    const n = nums.length;
    const treeProd = new Int32Array(4 * n);
    const treeFreq: Int32Array[] = Array.from({ length: 4 * n }, () => new Int32Array(k));

    function build(v: number, tl: number, tr: number): void {
        if (tl === tr) {
            treeProd[v] = nums[tl] % k;
            treeFreq[v][treeProd[v]] = 1;
            return;
        }
        const tm = (tl + tr) >> 1;
        build(v << 1, tl, tm);
        build(v << 1 | 1, tm + 1, tr);

        const pL = treeProd[v << 1], fL = treeFreq[v << 1];
        const pR = treeProd[v << 1 | 1], fR = treeFreq[v << 1 | 1];
        treeProd[v] = (pL * pR) % k;
        const fRes = treeFreq[v];
        fRes.fill(0);
        for (let i = 0; i < k; i++) fRes[i] = fL[i];
        for (let r = 0; r < k; r++) {
            if (fR[r]) {
                const nr = (pL * r) % k;
                fRes[nr] += fR[r];
            }
        }
    }

    function update(v: number, tl: number, tr: number, pos: number, val: number): void {
        if (tl === tr) {
            treeProd[v] = val % k;
            treeFreq[v].fill(0);
            treeFreq[v][treeProd[v]] = 1;
            return;
        }
        const tm = (tl + tr) >> 1;
        if (pos <= tm) update(v << 1, tl, tm, pos, val);
        else update(v << 1 | 1, tm + 1, tr, pos, val);

        const pL = treeProd[v << 1], fL = treeFreq[v << 1];
        const pR = treeProd[v << 1 | 1], fR = treeFreq[v << 1 | 1];
        treeProd[v] = (pL * pR) % k;
        const fRes = treeFreq[v];
        fRes.fill(0);
        for (let i = 0; i < k; i++) fRes[i] = fL[i];
        for (let r = 0; r < k; r++) {
            if (fR[r]) {
                const nr = (pL * r) % k;
                fRes[nr] += fR[r];
            }
        }
    }

    function query(v: number, tl: number, tr: number, l: number, r: number, outFreq: Int32Array, outProd: Int32Array): void {
        if (l > r) {
            outProd[0] = 1;
            outFreq.fill(0);
            return;
        }
        if (l === tl && r === tr) {
            outProd[0] = treeProd[v];
            outFreq.set(treeFreq[v]);
            return;
        }
        const tm = (tl + tr) >> 1;
        const leftFreq = new Int32Array(k);
        const rightFreq = new Int32Array(k);
        const leftProd = new Int32Array(1);
        const rightProd = new Int32Array(1);

        query(v << 1, tl, tm, l, Math.min(r, tm), leftFreq, leftProd);
        query(v << 1 | 1, tm + 1, tr, Math.max(l, tm + 1), r, rightFreq, rightProd);

        outProd[0] = (leftProd[0] * rightProd[0]) % k;
        outFreq.fill(0);
        for (let i = 0; i < k; i++) outFreq[i] = leftFreq[i];
        for (let rr = 0; rr < k; rr++) {
            if (rightFreq[rr]) {
                const nr = (leftProd[0] * rr) % k;
                outFreq[nr] += rightFreq[rr];
            }
        }
    }

    build(1, 0, n - 1);

    const ans: number[] = new Array(queries.length);
    const tempFreq = new Int32Array(k);
    const tempProd = new Int32Array(1);

    for (let i = 0; i < queries.length; i++) {
        const [idx, val, start, x] = queries[i];
        update(1, 0, n - 1, idx, val);
        query(1, 0, n - 1, start, n - 1, tempFreq, tempProd);
        ans[i] = tempFreq[x];
    }
    return ans;
}
```

### Python3
```python
class Solution:
    def resultArray(self, nums: List[int], k: int, queries: List[List[int]]) -> List[int]:
        class Node:
            __slots__ = ('prod', 'freq')
            def __init__(self):
                self.prod = 1
                self.freq = [0] * k
        
        n = len(nums)
        tree = [Node() for _ in range(4 * n)]
        
        def merge(L: Node, R: Node) -> Node:
            res = Node()
            res.prod = (L.prod * R.prod) % k
            res.freq = L.freq[:]
            for r in range(k):
                if R.freq[r]:
                    nr = (L.prod * r) % k
                    res.freq[nr] += R.freq[r]
            return res
        
        def build(v: int, tl: int, tr: int) -> None:
            if tl == tr:
                tree[v].prod = nums[tl] % k
                tree[v].freq[tree[v].prod] = 1
                return
            tm = (tl + tr) // 2
            build(v * 2, tl, tm)
            build(v * 2 + 1, tm + 1, tr)
            tree[v] = merge(tree[v * 2], tree[v * 2 + 1])
        
        def update(v: int, tl: int, tr: int, pos: int, val: int) -> None:
            if tl == tr:
                tree[v].prod = val % k
                tree[v].freq = [0] * k
                tree[v].freq[tree[v].prod] = 1
                return
            tm = (tl + tr) // 2
            if pos <= tm:
                update(v * 2, tl, tm, pos, val)
            else:
                update(v * 2 + 1, tm + 1, tr, pos, val)
            tree[v] = merge(tree[v * 2], tree[v * 2 + 1])
        
        def query(v: int, tl: int, tr: int, l: int, r: int) -> Node:
            if l > r:
                return Node()
            if l == tl and r == tr:
                return tree[v]
            tm = (tl + tr) // 2
            return merge(query(v * 2, tl, tm, l, min(r, tm)),
                         query(v * 2 + 1, tm + 1, tr, max(l, tm + 1), r))
        
        build(1, 0, n - 1)
        ans = []
        for idx, val, start, x in queries:
            update(1, 0, n - 1, idx, val)
            res = query(1, 0, n - 1, start, n - 1)
            ans.append(res.freq[x])
        return ans
```

### Go
```go
func resultArray(nums []int, k int, queries [][]int) []int {
    type Node struct {
        prod int
        freq []int
    }
    
    n := len(nums)
    tree := make([]Node, 4*n)
    for i := range tree {
        tree[i] = Node{prod: 1, freq: make([]int, k)}
    }
    
    merge := func(L, R Node) Node {
        res := Node{prod: 1, freq: make([]int, k)}
        res.prod = (L.prod * R.prod) % k
        copy(res.freq, L.freq)
        for r := 0; r < k; r++ {
            if R.freq[r] != 0 {
                nr := (L.prod * r) % k
                res.freq[nr] += R.freq[r]
            }
        }
        return res
    }
    
    var build func(v, tl, tr int)
    build = func(v, tl, tr int) {
        if tl == tr {
            tree[v].prod = nums[tl] % k
            tree[v].freq[tree[v].prod] = 1
            return
        }
        tm := (tl + tr) / 2
        build(v*2, tl, tm)
        build(v*2+1, tm+1, tr)
        tree[v] = merge(tree[v*2], tree[v*2+1])
    }
    
    var update func(v, tl, tr, pos, val int)
    update = func(v, tl, tr, pos, val int) {
        if tl == tr {
            tree[v].prod = val % k
            for i := range tree[v].freq {
                tree[v].freq[i] = 0
            }
            tree[v].freq[tree[v].prod] = 1
            return
        }
        tm := (tl + tr) / 2
        if pos <= tm {
            update(v*2, tl, tm, pos, val)
        } else {
            update(v*2+1, tm+1, tr, pos, val)
        }
        tree[v] = merge(tree[v*2], tree[v*2+1])
    }
    
    var query func(v, tl, tr, l, r int) Node
    query = func(v, tl, tr, l, r int) Node {
        if l > r {
            return Node{prod: 1, freq: make([]int, k)}
        }
        if l == tl && r == tr {
            return tree[v]
        }
        tm := (tl + tr) / 2
        return merge(query(v*2, tl, tm, l, min(r, tm)),
                     query(v*2+1, tm+1, tr, max(l, tm+1), r))
    }
    
    build(1, 0, n-1)
    ans := make([]int, 0, len(queries))
    for _, q := range queries {
        idx, val, start, x := q[0], q[1], q[2], q[3]
        update(1, 0, n-1, idx, val)
        res := query(1, 0, n-1, start, n-1)
        ans = append(ans, res.freq[x])
    }
    return ans
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

All six implementations follow the same high-level design; only the syntax for defining a node, allocating arrays, and writing recursive functions changes.

I first declare a Node that stores the modular product and a frequency array of length k. In C++ I use `std::array`, in Java and the JavaScript-family languages I use ordinary arrays, and in Go I use a slice. The default state of a node is product = 1 and all frequencies zero, which correctly represents an empty segment.

The merge function is identical across languages. It multiplies the two products, copies the left frequencies, then walks through every remainder that appears on the right and adds the corresponding count to the bucket `(left.product * right_remainder) % k`.

Build is a standard bottom-up recursion. When a leaf is reached it records `nums[i] % k` and sets the matching frequency to 1. On the way back up, each parent is formed by merging its two children.

Update descends to the correct leaf, rewrites the product and the frequency array, then re-merges every ancestor on the path to the root. This keeps the tree consistent after every permanent assignment required by a query.

Query returns a fresh Node that represents the exact merge of all complete segments covering the requested suffix. Because the merge operation is associative, the frequency array inside that Node is exactly the information needed for the current query. I simply read the entry for the target remainder and push it into the answer list.

The main driver function builds the tree once, then processes the queries one after another. Each query first updates the array and then asks for the frequency on the current suffix. Later queries automatically see all previous updates because the changes stay inside the segment tree.

Edge cases such as a single-element suffix or a start index that points to the last element are handled naturally by the leaf logic and by the empty-node base case of the query recursion.

## Examples

**Example 1**  
Input: `nums = [1,2,3,4,5]`, `k = 3`, `queries = [[2,2,0,2],[3,3,3,0],[0,1,0,1]]`  
Output: `[2,2,2]`  

- After the first update the array becomes `[1,2,2,4,5]`. The whole array (start = 0) has two prefixes whose product is 2 mod 3.  
- After the second update the array is `[1,2,2,3,5]`. Looking only at the suffix that starts at index 3 gives two valid prefixes.  
- After the third update the array is `[1,2,2,3,5]` again and the whole array still yields two prefixes congruent to 1 mod 3.

**Example 2**  
Input: `nums = [1,2,4,8,16,32]`, `k = 4`, `queries = [[0,2,0,2],[0,2,0,1]]`  
Output: `[1,0]`  

Only the full array after the first update has a product congruent to 2 mod 4; after the second update no prefix satisfies the new target remainder.

**Example 3**  
Input: `nums = [1,1,2,1,1]`, `k = 2`, `queries = [[2,1,0,1]]`  
Output: `[5]`  

After changing the middle element to 1 every one of the five possible prefixes has an odd product, so the count for remainder 1 is 5.

## How to Use / Run Locally

1. Copy the code for the language you prefer into a file (`main.cpp`, `Main.java`, `solution.js`, etc.).  
2. Make sure the language runtime or compiler is installed.  
3. Compile (if needed) and run. Most online judges accept the solution class directly; for local testing you can wrap the `resultArray` function with a simple `main` that reads input and prints the returned array.  

Example compile commands:  
- C++: `g++ -std=c++17 main.cpp -o main && ./main`  
- Java: `javac Main.java && java Main`  
- Python: `python3 solution.py`  
- Go: `go run main.go`  
- JavaScript / TypeScript: use Node.js (`node solution.js` or compile TypeScript first).

## Notes & Optimizations

- Because `k` is at most 5 the extra factor of `k²` never becomes a problem; the solution easily fits inside typical time limits.  
- If `k` were larger a different technique (for example sparse tables or heavy-light decomposition with modular products) would be required, but that is unnecessary here.  
- Empty prefixes are never counted because every frequency array only records non-empty prefixes, satisfying the problem statement.  
- The same segment-tree idea can be reused for any problem that needs prefix-product frequencies under updates, making the code a useful template for similar modular range-query tasks.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)