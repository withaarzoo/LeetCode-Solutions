# Pyramid Transition Matrix (LeetCode 756)

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

I am given a **bottom row of blocks** where each block has a color (A–F).
Using a list of **allowed triangular patterns**, I need to check whether I can build a **pyramid all the way to the top**.

Each rule says:

* Two adjacent blocks at the bottom
* Can form **one specific block on top**

I must return `true` if the pyramid can reach **one block at the top**, otherwise `false`.

---

## Constraints

* `2 ≤ bottom.length ≤ 6`
* `0 ≤ allowed.length ≤ 216`
* `allowed[i].length == 3`
* Characters are only from `{A, B, C, D, E, F}`
* All allowed patterns are unique

---

## Intuition

At first, I thought this was just a simple DFS problem.

But then I noticed something important:

If a certain row **fails once**, trying it again is completely useless.

So I asked myself:

> “Why am I rebuilding the same failed rows again and again?”

That’s where the optimization comes in.

The key idea is:

* Use DFS to explore possibilities
* **Cache failed rows**
* Never recompute them again

This single idea removes **Time Limit Exceeded (TLE)** completely.

---

## Approach

1. Convert `allowed` into a fast lookup structure
   `(left + right) → possible top blocks`

2. Use **DFS** to simulate pyramid building.

3. Build the **next row character by character**, not as a full list.

4. If a row reaches length `1`, return `true`.

5. If a row completely fails, store it in a **memo / invalid set**.

6. If the same row appears again, skip it immediately.

---

## Data Structures Used

* **HashMap / Dictionary**

  * Stores valid transitions for `(left, right)`

* **Set**

  * Stores rows that are already proven invalid

* **Recursion (DFS)**

  * To simulate pyramid construction

---

## Operations & Behavior Summary

* Bottom → build next row → build next → repeat
* If stuck at any level → mark row as invalid
* If top reached → stop and return true
* Memoization avoids repeated failures

---

## Complexity

### Time Complexity

**O(3ⁿ)** in the worst case
But due to memoization and pruning, the actual runtime is much faster.

Here:

* `n` = length of bottom row (max 6)

### Space Complexity

**O(number of failed rows)**
Very small due to limited combinations.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    unordered_map<string, unordered_set<char>> rules;
    unordered_set<string> bad;

    bool dfs(const string& row, int idx, string& next) {
        if (row.size() == 1) return true;

        if (idx == row.size() - 1) {
            if (bad.count(next)) return false;
            bool ok = dfs(next, 0, *(new string()));
            if (!ok) bad.insert(next);
            return ok;
        }

        string key = row.substr(idx, 2);
        if (!rules.count(key)) return false;

        for (char c : rules[key]) {
            next.push_back(c);
            if (dfs(row, idx + 1, next)) return true;
            next.pop_back();
        }
        return false;
    }

    bool pyramidTransition(string bottom, vector<string>& allowed) {
        for (auto& s : allowed)
            rules[s.substr(0, 2)].insert(s[2]);

        string next;
        return dfs(bottom, 0, next);
    }
};
```

---

### Java

```java
class Solution {
    Map<String, List<Character>> rules = new HashMap<>();
    Set<String> bad = new HashSet<>();

    public boolean pyramidTransition(String bottom, List<String> allowed) {
        for (String s : allowed) {
            rules.computeIfAbsent(s.substring(0, 2), k -> new ArrayList<>())
                 .add(s.charAt(2));
        }
        return dfs(bottom, 0, new StringBuilder());
    }

    private boolean dfs(String row, int idx, StringBuilder next) {
        if (row.length() == 1) return true;

        if (idx == row.length() - 1) {
            String nextRow = next.toString();
            if (bad.contains(nextRow)) return false;
            boolean ok = dfs(nextRow, 0, new StringBuilder());
            if (!ok) bad.add(nextRow);
            return ok;
        }

        String key = row.substring(idx, idx + 2);
        if (!rules.containsKey(key)) return false;

        for (char c : rules.get(key)) {
            next.append(c);
            if (dfs(row, idx + 1, next)) return true;
            next.deleteCharAt(next.length() - 1);
        }
        return false;
    }
}
```

---

### JavaScript

```javascript
var pyramidTransition = function(bottom, allowed) {
    const rules = {};
    for (const s of allowed) {
        const key = s[0] + s[1];
        if (!rules[key]) rules[key] = new Set();
        rules[key].add(s[2]);
    }

    const bad = new Set();

    function dfs(row, idx, next) {
        if (row.length === 1) return true;

        if (idx === row.length - 1) {
            if (bad.has(next)) return false;
            const ok = dfs(next, 0, '');
            if (!ok) bad.add(next);
            return ok;
        }

        const key = row.substring(idx, idx + 2);
        if (!rules[key]) return false;

        for (const c of rules[key]) {
            if (dfs(row, idx + 1, next + c)) return true;
        }
        return false;
    }

    return dfs(bottom, 0, '');
};
```

---

### Python3

```python
class Solution:
    def pyramidTransition(self, bottom: str, allowed: List[str]) -> bool:
        from collections import defaultdict

        rules = defaultdict(set)
        for a, b, c in allowed:
            rules[a + b].add(c)

        bad = set()

        def dfs(row, idx, nxt):
            if len(row) == 1:
                return True

            if idx == len(row) - 1:
                if nxt in bad:
                    return False
                ok = dfs(nxt, 0, "")
                if not ok:
                    bad.add(nxt)
                return ok

            for c in rules[row[idx:idx + 2]]:
                if dfs(row, idx + 1, nxt + c):
                    return True
            return False

        return dfs(bottom, 0, "")
```

---

### Go

```go
func pyramidTransition(bottom string, allowed []string) bool {
    rules := map[string][]byte{}
    for _, s := range allowed {
        rules[s[:2]] = append(rules[s[:2]], s[2])
    }

    bad := map[string]bool{}

    var dfs func(string, int, []byte) bool
    dfs = func(row string, idx int, next []byte) bool {
        if len(row) == 1 {
            return true
        }

        if idx == len(row)-1 {
            nxt := string(next)
            if bad[nxt] {
                return false
            }
            ok := dfs(nxt, 0, nil)
            if !ok {
                bad[nxt] = true
            }
            return ok
        }

        key := row[idx : idx+2]
        for _, c := range rules[key] {
            if dfs(row, idx+1, append(next, c)) {
                return true
            }
        }
        return false
    }

    return dfs(bottom, 0, nil)
}
```

---

## Step-by-step Detailed Explanation

1. Start from the bottom row.
2. Take two adjacent blocks.
3. Check which blocks can sit on top.
4. Build the next row gradually.
5. If stuck, mark the row as invalid.
6. If top reached, return true.
7. Memoization avoids repeated failures.

---

## Examples

**Input**

```
bottom = "BCD"
allowed = ["BCC","CDE","CEA","FFF"]
```

**Output**

```
true
```

---

**Input**

```
bottom = "AAAA"
allowed = ["AAB","AAC","BCD","BBE","DEF"]
```

**Output**

```
false
```

---

## How to use / Run locally

* Copy the solution of your preferred language
* Paste into LeetCode editor or local compiler
* Run with sample inputs

---

## Notes & Optimizations

* Memoization is the key to avoiding TLE
* Building next row inline reduces memory
* DFS + pruning is optimal for small constraints

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
