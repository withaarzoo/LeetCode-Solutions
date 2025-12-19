# Find All People With Secret (LeetCode 2092)

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given:

* `n` people numbered from `0` to `n - 1`
* A list of meetings where each meeting is `[x, y, time]`
* A `firstPerson` who receives the secret from person `0` at time `0`

Rules:

* If two people meet at a time and **one knows the secret**, the other learns it
* The secret spreads **instantly within the same time**
* Meetings at different times do not affect each other

Your task is to return **all people who know the secret after all meetings**.

---

## Constraints

* `2 ≤ n ≤ 100000`
* `1 ≤ meetings.length ≤ 100000`
* `meetings[i].length == 3`
* `0 ≤ x, y < n`
* `x != y`
* `1 ≤ time ≤ 100000`
* `1 ≤ firstPerson < n`

---

## Intuition

When I first looked at the problem, I noticed something very important:

* Meetings are **time-based**
* Secrets spread **instantly within the same time**
* Connections should **not carry over** to the next time

So I thought:

* I must process meetings **time by time**
* People meeting at the same time form a **temporary group**
* Inside that group, if **anyone knows the secret**, it spreads to everyone connected

This is a perfect use case for **Union-Find (Disjoint Set Union)**.

---

## Approach

1. Sort all meetings by time
2. Mark person `0` and `firstPerson` as knowing the secret
3. Process meetings grouped by the same time
4. Union people who meet at the same time
5. Check which groups contain at least one person who knows the secret
6. Allow the secret to spread only inside those groups
7. Reset connections for people who did not receive the secret
8. Collect all people who know the secret at the end

---

## Data Structures Used

* **Union-Find (Disjoint Set Union)**
  Used to group people meeting at the same time
* **Boolean array / Set**
  To track who knows the secret
* **Sorting**
  To process meetings in correct time order

---

## Operations & Behavior Summary

* Meetings at the same time → temporary graph
* Secret spreads instantly inside that graph
* Graph resets after each time block
* No secret leakage across different times

---

## Complexity

* **Time Complexity:** `O(m log m + m α(n))`

  * `m` = number of meetings
  * Sorting dominates the runtime
* **Space Complexity:** `O(n)`

  * Parent array + secret tracking

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> parent;

    int find(int x) {
        if (parent[x] == x) return x;
        return parent[x] = find(parent[x]);
    }

    void unite(int x, int y) {
        x = find(x);
        y = find(y);
        if (x != y) parent[y] = x;
    }

    vector<int> findAllPeople(int n, vector<vector<int>>& meetings, int firstPerson) {
        sort(meetings.begin(), meetings.end(),
             [](auto &a, auto &b) { return a[2] < b[2]; });

        parent.resize(n);
        for (int i = 0; i < n; i++) parent[i] = i;

        vector<bool> knows(n, false);
        knows[0] = knows[firstPerson] = true;

        int i = 0;
        while (i < meetings.size()) {
            int time = meetings[i][2];
            vector<int> people;

            int j = i;
            while (j < meetings.size() && meetings[j][2] == time) {
                unite(meetings[j][0], meetings[j][1]);
                people.push_back(meetings[j][0]);
                people.push_back(meetings[j][1]);
                j++;
            }

            unordered_map<int, bool> good;
            for (int p : people)
                if (knows[p]) good[find(p)] = true;

            for (int p : people) {
                if (good[find(p)]) knows[p] = true;
                else parent[p] = p;
            }
            i = j;
        }

        vector<int> ans;
        for (int i = 0; i < n; i++)
            if (knows[i]) ans.push_back(i);

        return ans;
    }
};
```

---

### Java

```java
class Solution {
    int[] parent;

    int find(int x) {
        if (parent[x] == x) return x;
        return parent[x] = find(parent[x]);
    }

    void union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x != y) parent[y] = x;
    }

    public List<Integer> findAllPeople(int n, int[][] meetings, int firstPerson) {
        Arrays.sort(meetings, (a, b) -> a[2] - b[2]);

        parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;

        boolean[] knows = new boolean[n];
        knows[0] = knows[firstPerson] = true;

        int i = 0;
        while (i < meetings.length) {
            int time = meetings[i][2];
            List<Integer> people = new ArrayList<>();

            int j = i;
            while (j < meetings.length && meetings[j][2] == time) {
                union(meetings[j][0], meetings[j][1]);
                people.add(meetings[j][0]);
                people.add(meetings[j][1]);
                j++;
            }

            Set<Integer> good = new HashSet<>();
            for (int p : people)
                if (knows[p]) good.add(find(p));

            for (int p : people) {
                if (good.contains(find(p))) knows[p] = true;
                else parent[p] = p;
            }
            i = j;
        }

        List<Integer> ans = new ArrayList<>();
        for (i = 0; i < n; i++)
            if (knows[i]) ans.add(i);

        return ans;
    }
}
```

---

### JavaScript

```javascript
var findAllPeople = function(n, meetings, firstPerson) {
    meetings.sort((a, b) => a[2] - b[2]);

    const parent = Array.from({ length: n }, (_, i) => i);
    const knows = Array(n).fill(false);
    knows[0] = knows[firstPerson] = true;

    const find = (x) => parent[x] === x ? x : (parent[x] = find(parent[x]));
    const union = (x, y) => {
        x = find(x); y = find(y);
        if (x !== y) parent[y] = x;
    };

    let i = 0;
    while (i < meetings.length) {
        let time = meetings[i][2];
        let people = [];

        let j = i;
        while (j < meetings.length && meetings[j][2] === time) {
            union(meetings[j][0], meetings[j][1]);
            people.push(meetings[j][0], meetings[j][1]);
            j++;
        }

        const good = new Set();
        for (let p of people)
            if (knows[p]) good.add(find(p));

        for (let p of people) {
            if (good.has(find(p))) knows[p] = true;
            else parent[p] = p;
        }
        i = j;
    }

    return knows.map((v, i) => v ? i : -1).filter(v => v !== -1);
};
```

---

### Python3

```python
class Solution:
    def findAllPeople(self, n, meetings, firstPerson):
        meetings.sort(key=lambda x: x[2])

        parent = list(range(n))
        knows = [False] * n
        knows[0] = knows[firstPerson] = True

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(x, y):
            x, y = find(x), find(y)
            if x != y:
                parent[y] = x

        i = 0
        while i < len(meetings):
            time = meetings[i][2]
            people = []

            j = i
            while j < len(meetings) and meetings[j][2] == time:
                union(meetings[j][0], meetings[j][1])
                people.extend([meetings[j][0], meetings[j][1]])
                j += 1

            good = set(find(p) for p in people if knows[p])

            for p in people:
                if find(p) in good:
                    knows[p] = True
                else:
                    parent[p] = p

            i = j

        return [i for i in range(n) if knows[i]]
```

---

## Step-by-step Detailed Explanation

1. Sort meetings by time
2. Union people meeting at the same time
3. Identify which connected groups already know the secret
4. Spread secret only inside those groups
5. Reset connections for others
6. Repeat for next time block
7. Collect final result

---

## Examples

**Input**

```
n = 5
meetings = [[3,4,2],[1,2,1],[2,3,1]]
firstPerson = 1
```

**Output**

```
[0,1,2,3,4]
```

---

## How to Use / Run Locally

1. Clone the repository
2. Choose your language file
3. Run with standard compiler/interpreter

   * C++: `g++ solution.cpp && ./a.out`
   * Java: `javac Solution.java && java Solution`
   * JS: `node solution.js`
   * Python: `python solution.py`

---

## Notes & Optimizations

* Union-Find reset is the key optimization
* Prevents incorrect secret spread across time
* Works efficiently for large inputs
* Interview-safe and production-ready logic

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
