# 1722. Minimize Hamming Distance After Swap Operations

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I am given two integer arrays:

* `source`
* `target`

Both arrays have the same length.

I am also given `allowedSwaps`, where each pair `[a, b]` means I can swap `source[a]` and `source[b]` any number of times.

My task is to return the minimum possible Hamming Distance between `source` and `target` after performing any number of swaps.

The Hamming Distance is the number of positions where:

```text
source[i] != target[i]
```

---

## Constraints

```text
1 <= source.length, target.length <= 10^5
1 <= source[i], target[i] <= 10^5
0 <= allowedSwaps.length <= 10^5
allowedSwaps[i].length == 2
0 <= ai, bi < source.length
ai != bi
```

---

## Intuition

I thought about what the allowed swaps really mean.

If index `0` can swap with `1`, and `1` can swap with `2`, then I can also rearrange values among `0`, `1`, and `2` freely.

That means all connected indices form one group.

Inside one connected group:

* I can rearrange the `source` values however I want
* So I only need to check how many values can be matched with the `target` array

To efficiently find connected groups, I use Disjoint Set Union (DSU) / Union Find.

---

## Approach

1. Create a DSU for all indices.
2. Connect indices using `allowedSwaps`.
3. Find all connected components.
4. Group indices by their parent/root.
5. For each connected component:

   * Count the frequency of values in `source`
   * Try to match values in `target`
6. If a target value is available in the frequency map, use it.
7. Otherwise, increase the answer because that value cannot be matched.

---

## Data Structures Used

* Disjoint Set Union (DSU) / Union Find
* Hash Map / Dictionary
* Vector / ArrayList / List
* Counter / Frequency Map

---

## Operations & Behavior Summary

| Operation          | Purpose                                                  |
| ------------------ | -------------------------------------------------------- |
| Union              | Connect two swappable indices                            |
| Find               | Get the root parent of an index                          |
| Grouping           | Collect indices belonging to the same component          |
| Frequency Counting | Store how many times a value appears in source           |
| Matching           | Try to match target values using available source values |

---

## Complexity

* Time Complexity: `O(n * α(n))`

  * `n` is the size of the array
  * `α(n)` is inverse Ackermann function, almost constant

* Space Complexity: `O(n)`

  * Used for parent array, groups, and frequency maps

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> parent, rankArr;

    int find(int x) {
        if (parent[x] == x)
            return x;
        return parent[x] = find(parent[x]);
    }

    void unite(int a, int b) {
        int pa = find(a);
        int pb = find(b);

        if (pa == pb)
            return;

        if (rankArr[pa] < rankArr[pb]) {
            parent[pa] = pb;
        } else if (rankArr[pb] < rankArr[pa]) {
            parent[pb] = pa;
        } else {
            parent[pb] = pa;
            rankArr[pa]++;
        }
    }

    int minimumHammingDistance(vector<int>& source, vector<int>& target, vector<vector<int>>& allowedSwaps) {
        int n = source.size();

        parent.resize(n);
        rankArr.resize(n, 0);

        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }

        for (auto& swap : allowedSwaps) {
            unite(swap[0], swap[1]);
        }

        unordered_map<int, vector<int>> groups;

        for (int i = 0; i < n; i++) {
            groups[find(i)].push_back(i);
        }

        int answer = 0;

        for (auto& entry : groups) {
            unordered_map<int, int> freq;

            for (int idx : entry.second) {
                freq[source[idx]]++;
            }

            for (int idx : entry.second) {
                if (freq[target[idx]] > 0) {
                    freq[target[idx]]--;
                } else {
                    answer++;
                }
            }
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    int[] parent;
    int[] rank;

    private int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }

    private void union(int a, int b) {
        int pa = find(a);
        int pb = find(b);

        if (pa == pb) return;

        if (rank[pa] < rank[pb]) {
            parent[pa] = pb;
        } else if (rank[pb] < rank[pa]) {
            parent[pb] = pa;
        } else {
            parent[pb] = pa;
            rank[pa]++;
        }
    }

    public int minimumHammingDistance(int[] source, int[] target, int[][] allowedSwaps) {
        int n = source.length;

        parent = new int[n];
        rank = new int[n];

        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }

        for (int[] swap : allowedSwaps) {
            union(swap[0], swap[1]);
        }

        Map<Integer, List<Integer>> groups = new HashMap<>();

        for (int i = 0; i < n; i++) {
            int root = find(i);
            groups.putIfAbsent(root, new ArrayList<>());
            groups.get(root).add(i);
        }

        int answer = 0;

        for (List<Integer> indices : groups.values()) {
            Map<Integer, Integer> freq = new HashMap<>();

            for (int idx : indices) {
                freq.put(source[idx], freq.getOrDefault(source[idx], 0) + 1);
            }

            for (int idx : indices) {
                if (freq.getOrDefault(target[idx], 0) > 0) {
                    freq.put(target[idx], freq.get(target[idx]) - 1);
                } else {
                    answer++;
                }
            }
        }

        return answer;
    }
}
```

### JavaScript

```javascript
var minimumHammingDistance = function(source, target, allowedSwaps) {
    const n = source.length;

    const parent = Array.from({ length: n }, (_, i) => i);
    const rank = Array(n).fill(0);

    function find(x) {
        if (parent[x] !== x) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }

    function union(a, b) {
        let pa = find(a);
        let pb = find(b);

        if (pa === pb) return;

        if (rank[pa] < rank[pb]) {
            parent[pa] = pb;
        } else if (rank[pb] < rank[pa]) {
            parent[pb] = pa;
        } else {
            parent[pb] = pa;
            rank[pa]++;
        }
    }

    for (const [u, v] of allowedSwaps) {
        union(u, v);
    }

    const groups = new Map();

    for (let i = 0; i < n; i++) {
        const root = find(i);

        if (!groups.has(root)) {
            groups.set(root, []);
        }

        groups.get(root).push(i);
    }

    let answer = 0;

    for (const indices of groups.values()) {
        const freq = new Map();

        for (const idx of indices) {
            freq.set(source[idx], (freq.get(source[idx]) || 0) + 1);
        }

        for (const idx of indices) {
            const val = target[idx];

            if ((freq.get(val) || 0) > 0) {
                freq.set(val, freq.get(val) - 1);
            } else {
                answer++;
            }
        }
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def minimumHammingDistance(self, source: List[int], target: List[int], allowedSwaps: List[List[int]]) -> int:
        n = len(source)

        parent = list(range(n))
        rank = [0] * n

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(a, b):
            pa = find(a)
            pb = find(b)

            if pa == pb:
                return

            if rank[pa] < rank[pb]:
                parent[pa] = pb
            elif rank[pb] < rank[pa]:
                parent[pb] = pa
            else:
                parent[pb] = pa
                rank[pa] += 1

        for u, v in allowedSwaps:
            union(u, v)

        from collections import defaultdict, Counter

        groups = defaultdict(list)

        for i in range(n):
            groups[find(i)].append(i)

        answer = 0

        for indices in groups.values():
            freq = Counter()

            for idx in indices:
                freq[source[idx]] += 1

            for idx in indices:
                if freq[target[idx]] > 0:
                    freq[target[idx]] -= 1
                else:
                    answer += 1

        return answer
```

### Go

```go
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
    n := len(source)

    parent := make([]int, n)
    rank := make([]int, n)

    for i := 0; i < n; i++ {
        parent[i] = i
    }

    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }

    union := func(a, b int) {
        pa := find(a)
        pb := find(b)

        if pa == pb {
            return
        }

        if rank[pa] < rank[pb] {
            parent[pa] = pb
        } else if rank[pb] < rank[pa] {
            parent[pb] = pa
        } else {
            parent[pb] = pa
            rank[pa]++
        }
    }

    for _, swap := range allowedSwaps {
        union(swap[0], swap[1])
    }

    groups := make(map[int][]int)

    for i := 0; i < n; i++ {
        root := find(i)
        groups[root] = append(groups[root], i)
    }

    answer := 0

    for _, indices := range groups {
        freq := make(map[int]int)

        for _, idx := range indices {
            freq[source[idx]]++
        }

        for _, idx := range indices {
            if freq[target[idx]] > 0 {
                freq[target[idx]]--
            } else {
                answer++
            }
        }
    }

    return answer
}
```

## Step-by-step Detailed Explanation

### Step 1: Build connected components

Using DSU, I connect all indices that can be swapped.

```text
allowedSwaps = [[0,1],[1,2]]
```

This means indices `0`, `1`, and `2` all belong to the same component.

---

### Step 2: Group indices by parent

After DSU is complete, every connected component gets grouped together.

Example:

```text
Component 1 -> [0, 1, 3]
Component 2 -> [2, 4]
```

---

### Step 3: Count source frequencies

For every component, I count how many times each value appears.

```text
source values = [1, 2, 2, 3]

Frequency Map:
1 -> 1
2 -> 2
3 -> 1
```

---

### Step 4: Match target values

Now I try to match every target value with the frequency map.

```text
Target values = [2, 1, 3]
```

* `2` exists
* `1` exists
* `3` does not exist

So one mismatch remains.

---

### Step 5: Add mismatches to answer

Whenever a target value is unavailable in the component, I increase the answer.

That final answer becomes the minimum Hamming Distance.

---

## Examples

### Example 1

```text
Input:
source = [1,2,3,4]
target = [2,1,4,5]
allowedSwaps = [[0,1],[2,3]]

Output:
1
```

### Example 2

```text
Input:
source = [1,2,3,4]
target = [1,3,2,4]
allowedSwaps = []

Output:
2
```

### Example 3

```text
Input:
source = [5,1,2,4,3]
target = [1,5,4,2,3]
allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]

Output:
0
```

---

## How to use / Run locally

```bash
g++ filename.cpp -o output
./output
```

```bash
javac Solution.java
java Solution
```

```bash
node solution.js
```

```bash
python solution.py
```

```bash
go run solution.go
```

---

## Notes & Optimizations

* DSU is used because it efficiently finds connected groups.
* Path Compression makes `find()` very fast.
* Union by Rank helps keep the DSU tree shallow.
* Frequency maps avoid unnecessary sorting.
* Overall solution works efficiently for large constraints up to `10^5`.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
