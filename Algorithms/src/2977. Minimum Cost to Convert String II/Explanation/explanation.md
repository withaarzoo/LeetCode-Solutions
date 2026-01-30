# 🧠 Minimum Cost to Convert String II

**LeetCode – Problem 2977**

---

## 📑 Table of Contents

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

## 🧩 Problem Summary

I am given two strings:

* `source`
* `target`

Both have the **same length**.

I am also given multiple **substring conversion rules**, where:

* A substring from `original[i]` can be converted to `changed[i]`
* Each conversion has a specific `cost[i]`

### Rules

* I can apply **any number of operations**
* Operations must be either:

  * **Disjoint**, or
  * **Exactly the same range**
* I can chain multiple operations on the **same substring range**

### Goal

👉 Find the **minimum total cost** to convert `source` into `target`
👉 If it is **not possible**, return `-1`

---

## 📏 Constraints

* `1 ≤ source.length == target.length ≤ 1000`
* `1 ≤ original.length == changed.length ≤ 100`
* Substring lengths `≤ source.length`
* All strings contain only lowercase English letters
* `1 ≤ cost[i] ≤ 10⁶`

---

## 💡 Intuition

When I read the problem carefully, I realized something important:

❌ This is **not a character-by-character problem**
✅ This is a **substring transformation problem**

Also:

* I can apply **multiple operations on the same substring**
* That means substring conversions can be **chained**

So I split the problem into **two layers**:

1. **What is the minimum cost to convert one substring into another substring?**
   → This looks like a **graph shortest path problem**

2. **How do I convert the full string from left to right efficiently?**
   → This is a **Dynamic Programming** problem

That’s why the final solution is:

> **Floyd-Warshall + Dynamic Programming**

---

## 🚀 Approach

### Step 1: Treat substrings as graph nodes

* Every unique string from `original` and `changed` becomes a **node**
* Each conversion rule becomes a **directed edge with cost**

---

### Step 2: Use Floyd-Warshall

* I compute the **minimum cost** to convert:

  ```
  substring A → substring B
  ```

* This allows chaining multiple operations **on the same index range**

---

### Step 3: Dynamic Programming on the main string

Let:

```
dp[i] = minimum cost to convert source[0...i-1] → target[0...i-1]
```

Transitions:

1. If `source[i] == target[i]`

   * I move forward with **zero cost**
2. Try all valid substring lengths:

   * Convert `source[i...i+len-1]` → `target[i...i+len-1]`
   * Use precomputed Floyd-Warshall cost

---

### Step 4: Final Answer

* If `dp[n]` is unreachable → return `-1`
* Else → return `dp[n]`

---

## 🧰 Data Structures Used

* `unordered_map / HashMap` → substring to ID mapping
* `2D matrix` → Floyd-Warshall distance table
* `dp[] array` → dynamic programming
* `set / unordered_set` → valid substring lengths

---

## 🔄 Operations & Behavior Summary

| Operation            | Behavior                       |
| -------------------- | ------------------------------ |
| Same substring       | Can chain multiple conversions |
| Different substrings | Must be disjoint               |
| Substring match      | Zero cost                      |
| Missing path         | Conversion impossible          |

---

## ⏱️ Complexity

### Time Complexity

* Floyd-Warshall: **O(k³)**
  where `k ≤ 200` (unique substrings)
* DP traversal: **O(n × m)**

✅ Total: **O(k³ + n·m)**

---

### Space Complexity

* Distance matrix: **O(k²)**
* DP array: **O(n)**

---

## 🌍 Multi-language Solutions

### C++

```cpp
class Solution {
public:
    static const unsigned long long INF = ULLONG_MAX;
    unsigned long long dist[201][201];
    unsigned long long dp[1001];

    unordered_map<string, int> id;
    unordered_set<int> lens;

    long long minimumCost(string source, string target,
                          vector<string>& original,
                          vector<string>& changed,
                          vector<int>& cost) {

        int m = original.size();
        int n = source.size();

        id.clear();
        lens.clear();

        int sz = 0;
        memset(dist, 255, sizeof(dist));

        for (int i = 0; i < m; i++) {
            if (!id.count(original[i])) {
                id[original[i]] = sz++;
                lens.insert(original[i].size());
            }
            if (!id.count(changed[i])) {
                id[changed[i]] = sz++;
            }
            int u = id[original[i]];
            int v = id[changed[i]];
            dist[u][v] = min(dist[u][v], (unsigned long long)cost[i]);
        }

        for (int i = 0; i < sz; i++) dist[i][i] = 0;

        for (int k = 0; k < sz; k++)
            for (int i = 0; i < sz; i++)
                if (dist[i][k] != INF)
                    for (int j = 0; j < sz; j++)
                        if (dist[k][j] != INF)
                            dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j]);

        memset(dp, 255, sizeof(dp));
        dp[0] = 0;

        for (int i = 0; i < n; i++) {
            if (dp[i] == INF) continue;

            if (source[i] == target[i])
                dp[i + 1] = min(dp[i + 1], dp[i]);

            for (int L : lens) {
                if (i + L > n) continue;

                string s = source.substr(i, L);
                string t = target.substr(i, L);

                if (id.count(s) && id.count(t)) {
                    unsigned long long d = dist[id[s]][id[t]];
                    if (d != INF)
                        dp[i + L] = min(dp[i + L], dp[i] + d);
                }
            }
        }

        return dp[n] == INF ? -1 : dp[n];
    }
};
```

### Java

```java
class Solution {
    static final long INF = Long.MAX_VALUE;
    public long minimumCost(String source, String target,
                            String[] original, String[] changed, int[] cost) {

        Map<String, Integer> id = new HashMap<>();
        Set<Integer> lens = new HashSet<>();

        int sz = 0;
        int m = original.length;
        int n = source.length();

        long[][] dist = new long[201][201];
        for (long[] row : dist) Arrays.fill(row, INF);

        for (int i = 0; i < m; i++) {
            if (!id.containsKey(original[i])) {
                id.put(original[i], sz++);
                lens.add(original[i].length());
            }
            if (!id.containsKey(changed[i])) {
                id.put(changed[i], sz++);
            }
            int u = id.get(original[i]);
            int v = id.get(changed[i]);
            dist[u][v] = Math.min(dist[u][v], cost[i]);
        }

        for (int i = 0; i < sz; i++) dist[i][i] = 0;

        for (int k = 0; k < sz; k++)
            for (int i = 0; i < sz; i++)
                if (dist[i][k] != INF)
                    for (int j = 0; j < sz; j++)
                        if (dist[k][j] != INF)
                            dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);

        long[] dp = new long[n + 1];
        Arrays.fill(dp, INF);
        dp[0] = 0;

        for (int i = 0; i < n; i++) {
            if (dp[i] == INF) continue;

            if (source.charAt(i) == target.charAt(i))
                dp[i + 1] = Math.min(dp[i + 1], dp[i]);

            for (int L : lens) {
                if (i + L > n) continue;

                String s = source.substring(i, i + L);
                String t = target.substring(i, i + L);

                if (id.containsKey(s) && id.containsKey(t)) {
                    long d = dist[id.get(s)][id.get(t)];
                    if (d != INF)
                        dp[i + L] = Math.min(dp[i + L], dp[i] + d);
                }
            }
        }

        return dp[n] == INF ? -1 : dp[n];
    }
}
```

### JavaScript

```javascript
var minimumCost = function(source, target, original, changed, cost) {
    const INF = BigInt("18446744073709551615");
    const id = new Map();
    const lens = new Set();
    let sz = 0;

    const dist = Array.from({ length: 201 }, () =>
        Array(201).fill(INF)
    );

    for (let i = 0; i < original.length; i++) {
        if (!id.has(original[i])) {
            id.set(original[i], sz++);
            lens.add(original[i].length);
        }
        if (!id.has(changed[i])) {
            id.set(changed[i], sz++);
        }
        const u = id.get(original[i]);
        const v = id.get(changed[i]);
        dist[u][v] = BigInt(Math.min(Number(dist[u][v]), cost[i]));
    }

    for (let i = 0; i < sz; i++) dist[i][i] = 0n;

    for (let k = 0; k < sz; k++)
        for (let i = 0; i < sz; i++)
            if (dist[i][k] !== INF)
                for (let j = 0; j < sz; j++)
                    if (dist[k][j] !== INF)
                        dist[i][j] = BigInt(Math.min(
                            Number(dist[i][j]),
                            Number(dist[i][k] + dist[k][j])
                        ));

    const n = source.length;
    const dp = Array(n + 1).fill(INF);
    dp[0] = 0n;

    for (let i = 0; i < n; i++) {
        if (dp[i] === INF) continue;

        if (source[i] === target[i])
            dp[i + 1] = dp[i + 1] < dp[i] ? dp[i + 1] : dp[i];

        for (const L of lens) {
            if (i + L > n) continue;
            const s = source.substr(i, L);
            const t = target.substr(i, L);

            if (id.has(s) && id.has(t)) {
                const d = dist[id.get(s)][id.get(t)];
                if (d !== INF)
                    dp[i + L] = dp[i + L] < dp[i] + d ? dp[i + L] : dp[i] + d;
            }
        }
    }

    return dp[n] === INF ? -1 : Number(dp[n]);
};
```

### Python3

```python
class Solution:
    def minimumCost(self, source, target, original, changed, cost):
        INF = 10**30
        id = {}
        lens = set()
        sz = 0

        dist = [[INF]*201 for _ in range(201)]

        for s, t, c in zip(original, changed, cost):
            if s not in id:
                id[s] = sz
                lens.add(len(s))
                sz += 1
            if t not in id:
                id[t] = sz
                sz += 1
            dist[id[s]][id[t]] = min(dist[id[s]][id[t]], c)

        for i in range(sz):
            dist[i][i] = 0

        for k in range(sz):
            for i in range(sz):
                if dist[i][k] < INF:
                    for j in range(sz):
                        if dist[k][j] < INF:
                            dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

        n = len(source)
        dp = [INF] * (n + 1)
        dp[0] = 0

        for i in range(n):
            if dp[i] == INF:
                continue

            if source[i] == target[i]:
                dp[i + 1] = min(dp[i + 1], dp[i])

            for L in lens:
                if i + L > n:
                    continue
                s = source[i:i+L]
                t = target[i:i+L]
                if s in id and t in id:
                    dp[i + L] = min(dp[i + L], dp[i] + dist[id[s]][id[t]])

        return -1 if dp[n] == INF else dp[n]
```

### Go

```go
func minimumCost(source string, target string,
    original []string, changed []string, cost []int) int64 {

    const INF int64 = 1<<62
    id := map[string]int{}
    lens := map[int]bool{}
    sz := 0

    dist := make([][]int64, 201)
    for i := range dist {
        dist[i] = make([]int64, 201)
        for j := range dist[i] {
            dist[i][j] = INF
        }
    }

    for i := 0; i < len(original); i++ {
        if _, ok := id[original[i]]; !ok {
            id[original[i]] = sz
            lens[len(original[i])] = true
            sz++
        }
        if _, ok := id[changed[i]]; !ok {
            id[changed[i]] = sz
            sz++
        }
        u := id[original[i]]
        v := id[changed[i]]
        if int64(cost[i]) < dist[u][v] {
            dist[u][v] = int64(cost[i])
        }
    }

    for i := 0; i < sz; i++ {
        dist[i][i] = 0
    }

    for k := 0; k < sz; k++ {
        for i := 0; i < sz; i++ {
            if dist[i][k] < INF {
                for j := 0; j < sz; j++ {
                    if dist[k][j] < INF {
                        if dist[i][k]+dist[k][j] < dist[i][j] {
                            dist[i][j] = dist[i][k] + dist[k][j]
                        }
                    }
                }
            }
        }
    }

    n := len(source)
    dp := make([]int64, n+1)
    for i := range dp {
        dp[i] = INF
    }
    dp[0] = 0

    for i := 0; i < n; i++ {
        if dp[i] == INF {
            continue
        }
        if source[i] == target[i] {
            if dp[i] < dp[i+1] {
                dp[i+1] = dp[i]
            }
        }
        for L := range lens {
            if i+L > n {
                continue
            }
            s := source[i : i+L]
            t := target[i : i+L]
            if x, ok1 := id[s]; ok1 {
                if y, ok2 := id[t]; ok2 {
                    if dist[x][y] < INF {
                        if dp[i]+dist[x][y] < dp[i+L] {
                            dp[i+L] = dp[i] + dist[x][y]
                        }
                    }
                }
            }
        }
    }

    if dp[n] == INF {
        return -1
    }
    return dp[n]
}
```

> 💡 All implementations follow the **same algorithm**, only syntax differs.

---

## 🧠 Step-by-step Detailed Explanation

1. Assign every unique substring a numeric ID
2. Build a directed graph using conversion rules
3. Run Floyd-Warshall to compute all-pair minimum costs
4. Initialize `dp[0] = 0`
5. Traverse source string left to right
6. Try:

   * Single character match
   * Valid substring conversions
7. Store minimum cost at each index
8. Return final DP result

---

## 🧪 Examples

### Example 1

```
source = "abcd"
target = "acbe"
Output = 28
```

✔ Multiple substring conversions
✔ Chained operations
✔ DP + shortest paths

---

### Example 2

```
source = "abcdefgh"
target = "addeeghh"
Output = 9
```

---

### Example 3

```
source = "abcdefgh"
target = "addddddd"
Output = -1
```

❌ No valid non-overlapping conversion sequence

---

## ▶️ How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
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

### JavaScript

```bash
node solution.js
```

### Go

```bash
go run solution.go
```

---

## 📝 Notes & Optimizations

* Only **substring lengths that exist** are checked
* Floyd-Warshall is safe due to small constraints
* This approach is **interview-ready**
* Avoids incorrect greedy or per-character logic

---

## 👤 Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
