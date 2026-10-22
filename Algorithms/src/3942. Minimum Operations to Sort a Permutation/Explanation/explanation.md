# Minimum Operations to Sort a Permutation – LeetCode 3942 Solution

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

You are given a permutation of the numbers from 0 to n-1. You can only do two things: reverse the whole array, or rotate it left by one position. The goal is to turn the array into the sorted order [0, 1, 2, \ldots, n-1] using the fewest possible operations. If it is impossible, return -1.

This is LeetCode problem 3942 – Minimum Operations to Sort a Permutation. It asks for the shortest sequence of left rotates and full reverses that sorts a given permutation, or reports that no such sequence exists.

## Constraints

- 1 <= n == nums.length <= 10^5
- 0 <= nums[i] <= n-1
- nums is a permutation of the integers from 0 to n-1

## Intuition

The two allowed operations are very limited. Starting from the sorted array, every reachable array is either a left rotation of the sorted sequence or a left rotation of the reversed sequence. That means only 2n possible arrays can ever be sorted with these operations. If the input is not one of those 2n arrays, the answer is immediately -1.

When the input does belong to one of the two families, we still need the exact number of operations. Because a left rotate is not its own inverse, the distance from the sorted array to the input is not the same as the distance from the input back to the sorted array. We therefore model the problem as a graph of 2n states and search from the input’s state toward the sorted state.

## Approach

First we check whether the given array is a cyclic shift of the increasing sequence 0,1,2,\ldots,n-1. This is true when every element is followed by the next number modulo n. If it is, we record the state (type = 0, shift = nums[0]).

We do the same check for the decreasing sequence n-1,n-2,\ldots,0. If every element is followed by the previous number modulo n, we record the state (type = 1, shift = (n-1-nums[0]) mod n).

If neither check succeeds we return -1. (For very small n both checks can succeed; we simply keep both candidate states.)

Each reachable array can be written as a pair (type, shift). The two operations act on these pairs as follows:

- left rotate changes the pair to (type, (shift + 1) mod n)
- reverse changes the pair to (1 - type, (n - shift) mod n)

We run a multi-source BFS that starts from every recorded pair with distance 0. The first time the search reaches the sorted state (0, 0) that distance is the minimum number of operations needed.

## Data Structures Used

- A small list (or vector) that stores the one or two possible starting states of the input.
- A 2-by-n distance table that records the shortest distance to every (type, shift) pair and also serves as a visited array.
- A queue for the BFS. Because there are only 2n states, the queue never grows larger than O(n).

These structures keep the whole algorithm linear in the size of the input.

## Operations & Behavior Summary

1. Scan the array once and decide whether it is a cyclic increasing permutation, a cyclic decreasing permutation, or neither.
2. Build the list of valid starting states from the checks above.
3. If the list is empty, return -1.
4. Initialize a distance table to -1 and put every starting state into the BFS queue with distance 0.
5. While the queue is not empty, take the front state. If it is already the sorted state (0, 0), return its distance.
6. Otherwise generate the two successors (left-rotate and reverse). Any successor that has never been visited receives distance = current distance + 1 and is enqueued.
7. If the queue empties without hitting the sorted state, return -1 (this case does not occur when the earlier checks succeeded).

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(n)  | Two linear scans plus a BFS that visits at most 2n states, each with constant work. |
| Space Complexity | O(n)  | Distance table of size 2 × n and a queue that holds at most O(n) entries. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        vector<pair<int,int>> starts;
        bool inc = true, dec = true;
        for (int i = 0; i < n; ++i) {
            if (nums[(i + 1) % n] != (nums[i] + 1) % n) inc = false;
            if (nums[(i + 1) % n] != (nums[i] - 1 + n) % n) dec = false;
        }
        if (inc) starts.emplace_back(0, nums[0]);
        if (dec) starts.emplace_back(1, (n - 1 - nums[0] + n) % n);
        if (starts.empty()) return -1;

        vector<vector<int>> dist(2, vector<int>(n, -1));
        queue<pair<int,int>> q;
        for (auto [t, s] : starts) {
            dist[t][s] = 0;
            q.emplace(t, s);
        }
        while (!q.empty()) {
            auto [t, s] = q.front(); q.pop();
            if (t == 0 && s == 0) return dist[t][s];
            int d = dist[t][s];
            int nt = t, ns = (s + 1) % n;
            if (dist[nt][ns] == -1) {
                dist[nt][ns] = d + 1;
                q.emplace(nt, ns);
            }
            nt = 1 - t;
            ns = (n - s) % n;
            if (dist[nt][ns] == -1) {
                dist[nt][ns] = d + 1;
                q.emplace(nt, ns);
            }
        }
        return -1;
    }
};
```

### Java
```java
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        List<int[]> starts = new ArrayList<>();
        boolean inc = true, dec = true;
        for (int i = 0; i < n; i++) {
            if (nums[(i + 1) % n] != (nums[i] + 1) % n) inc = false;
            if (nums[(i + 1) % n] != (nums[i] - 1 + n) % n) dec = false;
        }
        if (inc) starts.add(new int[]{0, nums[0]});
        if (dec) starts.add(new int[]{1, (n - 1 - nums[0] + n) % n});
        if (starts.isEmpty()) return -1;

        int[][] dist = new int[2][n];
        for (int[] row : dist) Arrays.fill(row, -1);
        Queue<int[]> q = new ArrayDeque<>();
        for (int[] p : starts) {
            dist[p[0]][p[1]] = 0;
            q.offer(p);
        }
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            int t = cur[0], s = cur[1];
            if (t == 0 && s == 0) return dist[t][s];
            int d = dist[t][s];
            int nt = t, ns = (s + 1) % n;
            if (dist[nt][ns] == -1) {
                dist[nt][ns] = d + 1;
                q.offer(new int[]{nt, ns});
            }
            nt = 1 - t;
            ns = (n - s) % n;
            if (dist[nt][ns] == -1) {
                dist[nt][ns] = d + 1;
                q.offer(new int[]{nt, ns});
            }
        }
        return -1;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minOperations = function(nums) {
    const n = nums.length;
    const starts = [];
    let inc = true, dec = true;
    for (let i = 0; i < n; i++) {
        if (nums[(i + 1) % n] !== (nums[i] + 1) % n) inc = false;
        if (nums[(i + 1) % n] !== (nums[i] - 1 + n) % n) dec = false;
    }
    if (inc) starts.push([0, nums[0]]);
    if (dec) starts.push([1, (n - 1 - nums[0] + n) % n]);
    if (starts.length === 0) return -1;

    const dist = Array.from({length: 2}, () => Array(n).fill(-1));
    const q = [];
    for (const [t, s] of starts) {
        dist[t][s] = 0;
        q.push([t, s]);
    }
    let head = 0;
    while (head < q.length) {
        const [t, s] = q[head++];
        if (t === 0 && s === 0) return dist[t][s];
        const d = dist[t][s];
        let nt = t, ns = (s + 1) % n;
        if (dist[nt][ns] === -1) {
            dist[nt][ns] = d + 1;
            q.push([nt, ns]);
        }
        nt = 1 - t;
        ns = (n - s) % n;
        if (dist[nt][ns] === -1) {
            dist[nt][ns] = d + 1;
            q.push([nt, ns]);
        }
    }
    return -1;
};
```

### TypeScript
```typescript
function minOperations(nums: number[]): number {
    const n = nums.length;
    const starts: [number, number][] = [];
    let inc = true, dec = true;
    for (let i = 0; i < n; i++) {
        if (nums[(i + 1) % n] !== (nums[i] + 1) % n) inc = false;
        if (nums[(i + 1) % n] !== (nums[i] - 1 + n) % n) dec = false;
    }
    if (inc) starts.push([0, nums[0]]);
    if (dec) starts.push([1, (n - 1 - nums[0] + n) % n]);
    if (starts.length === 0) return -1;

    const dist: number[][] = Array.from({length: 2}, () => Array(n).fill(-1));
    const q: [number, number][] = [];
    for (const [t, s] of starts) {
        dist[t][s] = 0;
        q.push([t, s]);
    }
    let head = 0;
    while (head < q.length) {
        const [t, s] = q[head++];
        if (t === 0 && s === 0) return dist[t][s];
        const d = dist[t][s];
        let nt = t, ns = (s + 1) % n;
        if (dist[nt][ns] === -1) {
            dist[nt][ns] = d + 1;
            q.push([nt, ns]);
        }
        nt = 1 - t;
        ns = (n - s) % n;
        if (dist[nt][ns] === -1) {
            dist[nt][ns] = d + 1;
            q.push([nt, ns]);
        }
    }
    return -1;
};
```

### Python3
```python
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        n = len(nums)
        starts = []
        inc = dec = True
        for i in range(n):
            if nums[(i + 1) % n] != (nums[i] + 1) % n:
                inc = False
            if nums[(i + 1) % n] != (nums[i] - 1 + n) % n:
                dec = False
        if inc:
            starts.append((0, nums[0]))
        if dec:
            starts.append((1, (n - 1 - nums[0]) % n))
        if not starts:
            return -1

        dist = [[-1] * n for _ in range(2)]
        q = deque()
        for t, s in starts:
            dist[t][s] = 0
            q.append((t, s))
        while q:
            t, s = q.popleft()
            if t == 0 and s == 0:
                return dist[t][s]
            d = dist[t][s]
            nt, ns = t, (s + 1) % n
            if dist[nt][ns] == -1:
                dist[nt][ns] = d + 1
                q.append((nt, ns))
            nt, ns = 1 - t, (n - s) % n
            if dist[nt][ns] == -1:
                dist[nt][ns] = d + 1
                q.append((nt, ns))
        return -1
```

### Go
```go
func minOperations(nums []int) int {
    n := len(nums)
    starts := make([][2]int, 0, 2)
    inc, dec := true, true
    for i := 0; i < n; i++ {
        if nums[(i+1)%n] != (nums[i]+1)%n {
            inc = false
        }
        if nums[(i+1)%n] != (nums[i]-1+n)%n {
            dec = false
        }
    }
    if inc {
        starts = append(starts, [2]int{0, nums[0]})
    }
    if dec {
        starts = append(starts, [2]int{1, (n - 1 - nums[0] + n) % n})
    }
    if len(starts) == 0 {
        return -1
    }

    dist := make([][]int, 2)
    for i := range dist {
        dist[i] = make([]int, n)
        for j := range dist[i] {
            dist[i][j] = -1
        }
    }
    type state struct{ t, s int }
    q := make([]state, 0, 2)
    for _, p := range starts {
        dist[p[0]][p[1]] = 0
        q = append(q, state{p[0], p[1]})
    }
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        t, s := cur.t, cur.s
        if t == 0 && s == 0 {
            return dist[t][s]
        }
        d := dist[t][s]
        nt, ns := t, (s+1)%n
        if dist[nt][ns] == -1 {
            dist[nt][ns] = d + 1
            q = append(q, state{nt, ns})
        }
        nt, ns = 1-t, (n-s)%n
        if dist[nt][ns] == -1 {
            dist[nt][ns] = d + 1
            q = append(q, state{nt, ns})
        }
    }
    return -1
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the concrete syntax for arrays, queues and modular arithmetic changes.

We start by measuring n, the length of the array. Then we prepare a small collection that will hold every legal starting state of the given permutation.

We walk the array once and test two modular conditions. The first asks whether every element is followed by the next integer modulo n. If the condition holds, the array is a pure left-rotation of the sorted sequence and we record the pair (0, nums[0]). The second condition asks the same question for the previous integer. If it holds we record the pair (1, (n-1-nums[0]) mod n). When the collection stays empty the array can never become sorted, so we return -1 at once.

We allocate a 2-by-n distance table filled with -1 and a queue. For every recorded starting pair we set its distance to 0 and push it into the queue. This is a multi-source BFS: all possible interpretations of the input are explored together.

Inside the loop we dequeue a state (t, s). If that state is already the sorted vertex (0, 0) we return its distance immediately – that is the minimum number of operations. Otherwise we generate the two successors that the allowed operations produce:

- left-rotate keeps the type and advances the shift by one,
- reverse flips the type and replaces the shift by (n-s) mod n.

Any successor that has never been visited receives distance = current distance + 1 and is enqueued. Because the graph contains only 2n vertices and each has out-degree two, the search finishes after a linear number of steps.

If the queue empties without ever hitting (0, 0) we return -1. In practice this never happens when the earlier modular checks succeeded, because every valid cyclic shift is reachable from the sorted array by a combination of left-rotates and reverses.

Edge cases are handled naturally: n = 1 is a trivial sorted array (distance 0), n = 2 may belong to both families, and the multi-source BFS simply picks the shorter of the two possible paths.

## Examples

**Example 1**  
Input: nums = [0, 2, 1]  
Output: 2  
The array is a left rotation of the reversed sequence. One left rotate produces [2, 1, 0]; a subsequent reverse produces the sorted array [0, 1, 2]. No shorter sequence exists.

**Example 2**  
Input: nums = [1, 0, 2]  
Output: 2  
Again the array belongs to the decreasing family. Reverse first, then left-rotate once, reaches the sorted order.

**Example 3**  
Input: nums = [1, 2, 0]  
Output: 2  
This array is a pure left rotation of the sorted sequence (shift = 1). Two successive left rotates bring 0 to the front and produce [0, 1, 2]. A single left rotate or a reverse alone does not sort it, so the answer is 2.

**Example 4**  
Input: nums = [2, 0, 1, 3]  
Output: -1  
The array is neither a cyclic shift of the increasing sequence nor of the decreasing sequence, therefore it is impossible to sort with the given operations.

## How to Use / Run Locally

**C++**  
Save the code as `main.cpp`, then compile and run:  
```bash
g++ -std=c++17 main.cpp -o main
./main
```

**Java**  
Save the code as `Solution.java`, then:  
```bash
javac Solution.java
java Solution
```

**JavaScript**  
Save the code as `solution.js`, then:  
```bash
node solution.js
```

**TypeScript**  
Save the code as `solution.ts`, then:  
```bash
npx ts-node solution.ts
```

**Python3**  
Save the code as `solution.py`, then:  
```bash
python3 solution.py
```

**Go**  
Save the code as `main.go`, then:  
```bash
go run main.go
```

In every language you can hard-code a test array inside the `main` function or read it from standard input according to your preference.

## Notes & Optimizations

The algorithm already runs in linear time and constant extra factors, so it comfortably handles the maximum constraint n = 10^5. No further asymptotic improvement is needed.

An alternative closed-form solution exists that calculates the minimum number of rotates and reverses without an explicit BFS, but the BFS version is easier to verify and still meets the time limits with a large margin.

The only real edge-case subtlety is n ≤ 2, where an array can belong to both the increasing and the decreasing families at the same time. The multi-source BFS automatically considers both possibilities and returns the shorter distance.

Because left-rotate is not an involution, searching from the input toward the sorted state is mandatory; searching in the opposite direction produces incorrect answers on some test cases.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)