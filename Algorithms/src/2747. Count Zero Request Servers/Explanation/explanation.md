# Count Zero Request Servers - LeetCode 2747 Solution

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

You are given the total number of servers `n` and a list of request logs. Each log entry tells you which server received a request and at what time.

You also receive a list of queries and a fixed window length `x`. For every query time `q`, you must find how many servers received zero requests in the closed time interval `[q - x, q]`.

The output is an array of the same length as the queries. Each position stores the count of servers that stayed completely idle during that specific window.

This problem is a classic combination of sorting, sliding window, and frequency tracking. The goal is to answer many range queries efficiently without scanning the entire log list for every query.

## Constraints

- 1 <= n <= 10^5
- 1 <= logs.length <= 10^5
- 1 <= queries.length <= 10^5
- logs[i].length == 2
- 1 <= logs[i][0] <= n
- 1 <= logs[i][1] <= 10^9
- 1 <= x <= 10^5
- x < queries[i] <= 10^9

These limits force us to design an O((L + Q) log (L + Q)) solution. A nested loop would be too slow.

## Intuition

The first thing I noticed is that every query asks about a sliding time window of fixed length `x`. If I process the queries in increasing order of time, the windows only move forward.

That observation immediately suggests a two-pointer sliding window on the logs after they are sorted by time. I only need to keep track of how many unique servers are currently active inside the window. The answer for each query is simply `n` minus that active count.

Sorting the logs once and sorting the queries once lets me move both pointers only to the right, so every log is examined a constant number of times.

## Approach

1. Sort the logs by their time values so earlier requests come first.
2. Create a list of (query time, original index) pairs and sort that list by query time. This lets me answer queries in sorted order and still restore the original order at the end.
3. Maintain a frequency map that counts how many times each server appears inside the current window.
4. Keep an integer that stores the number of servers with a positive frequency (the active servers).
5. Use two pointers `left` and `right` on the sorted logs.
6. For every query in sorted order:
   - Move `right` forward while the next log time is still <= the current query time. Update frequencies and the active count.
   - Move `left` forward while the log at `left` has a time strictly less than `query - x`. Decrease frequencies and the active count when a server drops to zero.
7. Store `n - active` at the original index of the query.
8. After processing all queries, return the answer array.

The same logic works in every language; only the syntax for maps, pairs, and sorting changes.

## Data Structures Used

- Sorted list / array of logs – allows linear two-pointer movement after the initial sort.
- Array of (time, index) pairs – preserves original query order after sorting.
- Hash map (frequency map) – tracks how many times each server appears in the current window. Chosen because server IDs are up to 10^5 and we need O(1) average updates.
- Integer counter for active servers – avoids scanning the map on every query.

## Operations & Behavior Summary

- Sort logs by time.
- Sort queries while remembering their original positions.
- Initialize an empty frequency map and set active = 0, left = 0, right = 0.
- For each query in increasing time order:
  - Expand the right side of the window to include every log whose time is <= query.
  - Shrink the left side of the window to remove every log whose time is < query - x.
  - Record n - active as the answer for that query.
- Place the answers back into the original query order and return the result.

## Complexity

| Complexity       | Value                          | Explanation                                                                 |
|------------------|--------------------------------|-----------------------------------------------------------------------------|
| Time Complexity  | O((L + Q) log (L + Q))         | Sorting logs and queries dominates. The two-pointer scan is linear.         |
| Space Complexity | O(n + Q)                       | Frequency map can hold up to n servers. Extra array of size Q stores indices. |

L = number of logs, Q = number of queries.

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<int> countServers(int n, vector<vector<int>>& logs, int x, vector<int>& queries) {
        sort(logs.begin(), logs.end(), [](const vector<int>& a, const vector<int>& b) {
            return a[1] < b[1];
        });
        int m = queries.size();
        vector<pair<int, int>> q;
        for (int i = 0; i < m; i++) q.emplace_back(queries[i], i);
        sort(q.begin(), q.end());
        vector<int> ans(m);
        unordered_map<int, int> freq;
        int active = 0;
        int left = 0, right = 0;
        int L = logs.size();
        for (auto& [time, idx] : q) {
            while (right < L && logs[right][1] <= time) {
                int s = logs[right][0];
                if (++freq[s] == 1) active++;
                right++;
            }
            while (left < right && logs[left][1] < time - x) {
                int s = logs[left][0];
                if (--freq[s] == 0) active--;
                left++;
            }
            ans[idx] = n - active;
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int[] countServers(int n, int[][] logs, int x, int[] queries) {
        Arrays.sort(logs, (a, b) -> Integer.compare(a[1], b[1]));
        int m = queries.length;
        int[][] q = new int[m][2];
        for (int i = 0; i < m; i++) {
            q[i][0] = queries[i];
            q[i][1] = i;
        }
        Arrays.sort(q, (a, b) -> Integer.compare(a[0], b[0]));
        int[] ans = new int[m];
        Map<Integer, Integer> freq = new HashMap<>();
        int active = 0;
        int left = 0, right = 0;
        int L = logs.length;
        for (int[] pair : q) {
            int time = pair[0];
            int idx = pair[1];
            while (right < L && logs[right][1] <= time) {
                int s = logs[right][0];
                freq.put(s, freq.getOrDefault(s, 0) + 1);
                if (freq.get(s) == 1) active++;
                right++;
            }
            while (left < right && logs[left][1] < time - x) {
                int s = logs[left][0];
                freq.put(s, freq.get(s) - 1);
                if (freq.get(s) == 0) active--;
                left++;
            }
            ans[idx] = n - active;
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number} n
 * @param {number[][]} logs
 * @param {number} x
 * @param {number[]} queries
 * @return {number[]}
 */
var countServers = function(n, logs, x, queries) {
    logs.sort((a, b) => a[1] - b[1]);
    const m = queries.length;
    const q = queries.map((t, i) => [t, i]);
    q.sort((a, b) => a[0] - b[0]);
    const ans = new Array(m);
    const freq = new Map();
    let active = 0;
    let left = 0, right = 0;
    const L = logs.length;
    for (const [time, idx] of q) {
        while (right < L && logs[right][1] <= time) {
            const s = logs[right][0];
            freq.set(s, (freq.get(s) || 0) + 1);
            if (freq.get(s) === 1) active++;
            right++;
        }
        while (left < right && logs[left][1] < time - x) {
            const s = logs[left][0];
            freq.set(s, freq.get(s) - 1);
            if (freq.get(s) === 0) active--;
            left++;
        }
        ans[idx] = n - active;
    }
    return ans;
};
```

### TypeScript
```typescript
function countServers(n: number, logs: number[][], x: number, queries: number[]): number[] {
    logs.sort((a, b) => a[1] - b[1]);
    const m = queries.length;
    const q: [number, number][] = queries.map((t, i) => [t, i]);
    q.sort((a, b) => a[0] - b[0]);
    const ans = new Array(m);
    const freq = new Map<number, number>();
    let active = 0;
    let left = 0, right = 0;
    const L = logs.length;
    for (const [time, idx] of q) {
        while (right < L && logs[right][1] <= time) {
            const s = logs[right][0];
            freq.set(s, (freq.get(s) || 0) + 1);
            if (freq.get(s) === 1) active++;
            right++;
        }
        while (left < right && logs[left][1] < time - x) {
            const s = logs[left][0];
            freq.set(s, freq.get(s)! - 1);
            if (freq.get(s) === 0) active--;
            left++;
        }
        ans[idx] = n - active;
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def countServers(self, n: int, logs: List[List[int]], x: int, queries: List[int]) -> List[int]:
        logs.sort(key=lambda a: a[1])
        m = len(queries)
        q = [(queries[i], i) for i in range(m)]
        q.sort()
        ans = [0] * m
        freq = {}
        active = 0
        left = right = 0
        L = len(logs)
        for time, idx in q:
            while right < L and logs[right][1] <= time:
                s = logs[right][0]
                freq[s] = freq.get(s, 0) + 1
                if freq[s] == 1:
                    active += 1
                right += 1
            while left < right and logs[left][1] < time - x:
                s = logs[left][0]
                freq[s] -= 1
                if freq[s] == 0:
                    active -= 1
                left += 1
            ans[idx] = n - active
        return ans
```

### Go
```go
func countServers(n int, logs [][]int, x int, queries []int) []int {
    sort.Slice(logs, func(i, j int) bool {
        return logs[i][1] < logs[j][1]
    })
    m := len(queries)
    type pair struct{ time, idx int }
    q := make([]pair, m)
    for i := 0; i < m; i++ {
        q[i] = pair{queries[i], i}
    }
    sort.Slice(q, func(i, j int) bool {
        return q[i].time < q[j].time
    })
    ans := make([]int, m)
    freq := make(map[int]int)
    active := 0
    left, right := 0, 0
    L := len(logs)
    for _, p := range q {
        time, idx := p.time, p.idx
        for right < L && logs[right][1] <= time {
            s := logs[right][0]
            freq[s]++
            if freq[s] == 1 {
                active++
            }
            right++
        }
        for left < right && logs[left][1] < time-x {
            s := logs[left][0]
            freq[s]--
            if freq[s] == 0 {
                active--
            }
            left++
        }
        ans[idx] = n - active
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea is identical in every language. The only differences are how we declare maps, pairs, and how we sort.

I first sort the logs so that time values increase. This is necessary because the two pointers can only move forward safely when the array is ordered.

Next I create a temporary list that pairs every query with its original index. After sorting this list by time, I can still write the final answers back to the correct positions.

I keep a frequency map. In C++ and Java I use unordered_map / HashMap. In JavaScript, TypeScript, Python and Go I use the built-in map or dictionary. Every time a server’s count goes from 0 to 1 I increase the active counter. When the count falls back to 0 I decrease it.

The right pointer advances as long as the next log still belongs to the current window (time <= query). The left pointer advances as long as the log at the left is too old (time < query - x).

Because queries are processed in sorted order, both pointers only ever increase. No log is ever revisited, which keeps the whole scan linear after the sorting step.

Edge cases are handled naturally:
- If a window contains no logs, active stays 0 and the answer is n.
- If every server appears at least once, active becomes n and the answer is 0.
- Duplicate requests from the same server in the same window are correctly counted only once for the active total because we only care about presence, not the exact number of requests.

The final answer array is filled using the stored original indices so the output matches the order of the input queries.

## Examples

**Example 1**

Input:  
n = 3  
logs = [[1,3],[2,6],[1,5]]  
x = 5  
queries = [10,11]

Output: [1,2]

Trace for query 10 (window [5,10]):
- Logs at times 5 and 6 fall inside the window.
- Servers 1 and 2 are active.
- Server 3 is idle → answer 1.

Trace for query 11 (window [6,11]):
- Only the log at time 6 stays inside.
- Only server 2 is active.
- Servers 1 and 3 are idle → answer 2.

**Example 2**

Input:  
n = 3  
logs = [[2,4],[2,1],[1,2],[3,1]]  
x = 2  
queries = [3,4]

Output: [0,1]

Trace for query 3 (window [1,3]):
- All four logs fall inside or on the boundary.
- Servers 1, 2 and 3 all appear → answer 0.

Trace for query 4 (window [2,4]):
- Logs at times 2 and 4 remain.
- Servers 1 and 2 are active.
- Server 3 is idle → answer 1.

## How to Use / Run Locally

1. Copy the solution for the language you want into a file named according to the language (Solution.cpp, Solution.java, etc.).
2. Make sure the function signature matches the LeetCode template for that language.
3. Compile and run with a small test driver that creates the input arrays and prints the returned vector / array.

C++  
g++ -std=c++17 Solution.cpp -o solution && ./solution

Java  
javac Solution.java && java Solution

JavaScript  
node solution.js

TypeScript  
tsc solution.ts && node solution.js

Python3  
python3 solution.py

Go  
go run solution.go

Replace the driver code with your own test cases or feed the function the examples shown above.

## Notes & Optimizations

- The frequency map is the main extra space. If memory is extremely tight you could use a fixed-size array of size n+1 instead of a hash map, since server IDs are in the range 1..n.
- Sorting both arrays is unavoidable for the two-pointer technique. Any attempt to avoid sorting would force a heavier data structure such as a Fenwick tree or segment tree on the time axis, which is more complicated and usually slower in practice for these constraints.
- The solution correctly handles duplicate timestamps and multiple requests from the same server inside one window.
- Because the left and right pointers only move forward, the algorithm is optimal for the given constraints.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
