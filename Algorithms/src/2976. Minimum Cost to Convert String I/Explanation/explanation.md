# Minimum Cost to Convert String I (LeetCode 2976)

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I am given two strings `source` and `target` of the same length.
Both strings contain only lowercase English letters.

I am also given:

* `original[]` characters
* `changed[]` characters
* `cost[]` where `cost[i]` is the cost to convert `original[i]` → `changed[i]`

In **one operation**, I can change a character into another character **only if a rule exists**.

My task is to find the **minimum total cost** to convert `source` into `target`.

If conversion is **not possible**, I must return `-1`.

---

## Constraints

* `1 ≤ source.length ≤ 10⁵`
* `source.length == target.length`
* All characters are lowercase English letters
* `1 ≤ original.length ≤ 2000`
* `1 ≤ cost[i] ≤ 10⁶`
* `original[i] != changed[i]`

---

## Intuition

When I read the problem, I realized one key thing:

There are **only 26 lowercase letters**.

So instead of thinking in terms of strings, I thought in terms of **characters as nodes in a graph**.

Each conversion rule:

```bash
a → b with cost c
```

is like a **directed edge** in a graph.

Since I can apply multiple operations, I need the **shortest path** between characters.

Once I know the minimum cost to convert **any letter to any other letter**, converting the whole string becomes easy.

---

## Approach

I solved this in **three main steps**.

### Step 1: Create a graph

* Each character (`a` to `z`) is a node
* Store the minimum direct cost between characters

### Step 2: Precompute shortest paths

* Since we only have 26 nodes, I used **Floyd–Warshall**
* This gives me the minimum cost between **every pair of characters**

### Step 3: Convert the string

* For each index:

  * If `source[i] == target[i]`, cost is `0`
  * Otherwise, add the precomputed shortest cost
  * If no path exists, return `-1`

---

## Data Structures Used

* 2D array `dist[26][26]` to store minimum conversion costs
* Simple loops, no heavy data structures needed

---

## Operations & Behavior Summary

* Direct character conversions are stored first
* Indirect conversions are handled using shortest paths
* The solution works efficiently even for large strings
* Impossible conversions are detected early

---

## Complexity

### Time Complexity

* Floyd–Warshall: `O(26³)` → constant time
* String traversal: `O(n)`

**Total Time:** `O(n)`

### Space Complexity

* `26 × 26` matrix

**Total Space:** `O(1)`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long minimumCost(string source, string target,
                          vector<char>& original,
                          vector<char>& changed,
                          vector<int>& cost) {

        const long long INF = 1e18;
        vector<vector<long long>> dist(26, vector<long long>(26, INF));

        for (int i = 0; i < 26; i++) dist[i][i] = 0;

        for (int i = 0; i < original.size(); i++) {
            int u = original[i] - 'a';
            int v = changed[i] - 'a';
            dist[u][v] = min(dist[u][v], (long long)cost[i]);
        }

        for (int k = 0; k < 26; k++)
            for (int i = 0; i < 26; i++)
                for (int j = 0; j < 26; j++)
                    dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j]);

        long long ans = 0;
        for (int i = 0; i < source.size(); i++) {
            int s = source[i] - 'a';
            int t = target[i] - 'a';
            if (dist[s][t] == INF) return -1;
            ans += dist[s][t];
        }

        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public long minimumCost(String source, String target,
                            char[] original, char[] changed, int[] cost) {

        long INF = (long)1e18;
        long[][] dist = new long[26][26];

        for (int i = 0; i < 26; i++)
            for (int j = 0; j < 26; j++)
                dist[i][j] = (i == j) ? 0 : INF;

        for (int i = 0; i < original.length; i++) {
            int u = original[i] - 'a';
            int v = changed[i] - 'a';
            dist[u][v] = Math.min(dist[u][v], cost[i]);
        }

        for (int k = 0; k < 26; k++)
            for (int i = 0; i < 26; i++)
                for (int j = 0; j < 26; j++)
                    dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);

        long ans = 0;
        for (int i = 0; i < source.length(); i++) {
            int s = source.charAt(i) - 'a';
            int t = target.charAt(i) - 'a';
            if (dist[s][t] == INF) return -1;
            ans += dist[s][t];
        }

        return ans;
    }
}
```

---

### JavaScript

```javascript
var minimumCost = function(source, target, original, changed, cost) {
    const INF = 1e18;
    const dist = Array.from({ length: 26 }, () => Array(26).fill(INF));

    for (let i = 0; i < 26; i++) dist[i][i] = 0;

    for (let i = 0; i < original.length; i++) {
        const u = original[i].charCodeAt(0) - 97;
        const v = changed[i].charCodeAt(0) - 97;
        dist[u][v] = Math.min(dist[u][v], cost[i]);
    }

    for (let k = 0; k < 26; k++)
        for (let i = 0; i < 26; i++)
            for (let j = 0; j < 26; j++)
                dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);

    let ans = 0;
    for (let i = 0; i < source.length; i++) {
        const s = source.charCodeAt(i) - 97;
        const t = target.charCodeAt(i) - 97;
        if (dist[s][t] === INF) return -1;
        ans += dist[s][t];
    }

    return ans;
};
```

---

### Python3

```python
class Solution:
    def minimumCost(self, source, target, original, changed, cost):
        INF = 10**18
        dist = [[INF]*26 for _ in range(26)]

        for i in range(26):
            dist[i][i] = 0

        for o, c, w in zip(original, changed, cost):
            u = ord(o) - 97
            v = ord(c) - 97
            dist[u][v] = min(dist[u][v], w)

        for k in range(26):
            for i in range(26):
                for j in range(26):
                    dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

        ans = 0
        for s, t in zip(source, target):
            u = ord(s) - 97
            v = ord(t) - 97
            if dist[u][v] == INF:
                return -1
            ans += dist[u][v]

        return ans
```

---

### Go

```go
func minimumCost(source string, target string,
 original []byte, changed []byte, cost []int) int64 {

 const INF int64 = 1e18
 dist := make([][]int64, 26)

 for i := 0; i < 26; i++ {
  dist[i] = make([]int64, 26)
  for j := 0; j < 26; j++ {
   if i == j {
    dist[i][j] = 0
   } else {
    dist[i][j] = INF
   }
  }
 }

 for i := 0; i < len(original); i++ {
  u := original[i] - 'a'
  v := changed[i] - 'a'
  if int64(cost[i]) < dist[u][v] {
   dist[u][v] = int64(cost[i])
  }
 }

 for k := 0; k < 26; k++ {
  for i := 0; i < 26; i++ {
   for j := 0; j < 26; j++ {
    if dist[i][k]+dist[k][j] < dist[i][j] {
     dist[i][j] = dist[i][k] + dist[k][j]
    }
   }
  }
 }

 var ans int64 = 0
 for i := 0; i < len(source); i++ {
  s := source[i] - 'a'
  t := target[i] - 'a'
  if dist[s][t] == INF {
   return -1
  }
  ans += dist[s][t]
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation

1. Initialize a 26×26 matrix with infinity
2. Set same-character conversion cost to 0
3. Fill direct conversion costs
4. Run Floyd–Warshall to allow multi-step conversions
5. Traverse the string and accumulate costs
6. Return `-1` if conversion is impossible

---

## Examples

**Input**

```bash
source = "abcd"
target = "acbe"
```

**Output**

```bash
28
```

**Explanation**
Each character is converted using the cheapest possible path.

---

## How to use / Run locally

1. Copy the solution code
2. Paste it into your local compiler or LeetCode editor
3. Run with sample inputs
4. Submit

---

## Notes & Optimizations

* Floyd–Warshall is perfect because character count is fixed
* No need for Dijkstra or BFS per character
* Works efficiently even for maximum constraints

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
