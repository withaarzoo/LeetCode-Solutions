# Find Subtree Sizes After Changes - LeetCode 3331 Solution

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

You are given a tree with n nodes rooted at node 0. The tree is described by a parent array where parent[i] tells the parent of node i. You also get a string s of length n, and s[i] is the character assigned to node i.

For every node x from 1 to n-1 we do the following change at the same time:

- Look for the closest ancestor y of x such that s[x] equals s[y].
- If such a y exists, we cut the edge from x to its current parent and attach x under y instead.

After all these simultaneous reparenting operations finish, we need to return an array answer of size n where answer[i] is the size of the subtree rooted at node i in the final tree.

This is a classic tree reparenting problem that appears in many competitive programming contests and coding interviews. The goal is to compute subtree sizes after the closest-same-character ancestor changes.

## Constraints

- 1 <= n <= 10^5
- parent.length == n
- s.length == n
- parent[0] == -1
- 0 <= parent[i] <= n-1 for all i >= 1
- parent forms a valid tree
- s contains only lowercase English letters

## Intuition

The first thing I noticed is that every node wants to jump to its nearest ancestor that carries the same letter. Because all jumps happen simultaneously on the original tree, I cannot simply walk up the parent chain for each node one by one. That would be too slow for n = 10^5 and would also break the simultaneous requirement.

The key observation is that a depth-first search already visits nodes in an order that respects the ancestor relationship. If I keep track of the most recent node I have seen for every character while I go down the tree, then the moment I arrive at a node the value stored for its character is exactly the closest matching ancestor. That single idea turns an expensive per-node search into a linear-time pass.

## Approach

I start by building the original adjacency list from the parent array so I can traverse the tree easily.

I create a newParent array that initially copies the given parent array. This array will store the final parent of every node after the reparenting step.

I run a DFS from the root. During the DFS I maintain a small array of size 26 that remembers the deepest node seen so far for each letter. When I enter a node I check the current entry for its letter. If it is not -1 I write that value into newParent. Then I update the entry to the current node, recurse on the children, and finally restore the previous value so that siblings still see the correct ancestor.

After the DFS finishes I have the correct new parents. I build a second adjacency list from these new parents and run one more DFS that simply computes the size of every subtree and stores the results in an answer array.

The whole process stays linear because every node is visited a constant number of times and the character-tracking array never grows beyond 26 slots.

## Data Structures Used

- Adjacency lists (vector of vectors / list of lists) – to store both the original tree and the new tree after reparenting. They give O(1) access to children and support efficient DFS.
- Array of size 26 – to remember the most recent node for each lowercase letter. Fixed size makes it very fast and cache-friendly.
- newParent array – holds the final parent of every node after the simultaneous moves.
- Answer array – stores the computed subtree sizes.

## Operations & Behavior Summary

1. Build the original tree from the parent array.
2. Copy the parent array into newParent.
3. Perform a DFS while maintaining the last-seen node for every character.
4. For each node, if a matching ancestor exists, update its entry in newParent.
5. After the first DFS, construct the new tree from newParent.
6. Run a second DFS on the new tree to calculate and store every subtree size.
7. Return the size array.

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time       | O(n)  | Each node is visited a constant number of times in two DFS passes. The character array is fixed at size 26. |
| Space      | O(n)  | Adjacency lists, newParent array, recursion stack and answer array all take linear space. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<int> findSubtreeSizes(vector<int>& parent, string s) {
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) g[parent[i]].push_back(i);
        vector<int> np(parent.begin(), parent.end());
        vector<int> last(26, -1);
        function<void(int)> dfs = [&](int u) {
            int c = s[u] - 'a';
            int prev = last[c];
            if (prev != -1) np[u] = prev;
            last[c] = u;
            for (int v : g[u]) dfs(v);
            last[c] = prev;
        };
        dfs(0);
        vector<vector<int>> ng(n);
        for (int i = 1; i < n; i++) ng[np[i]].push_back(i);
        vector<int> ans(n);
        function<int(int)> size = [&](int u) {
            int sz = 1;
            for (int v : ng[u]) sz += size(v);
            return ans[u] = sz;
        };
        size(0);
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int[] findSubtreeSizes(int[] parent, String s) {
        int n = parent.length;
        List<Integer>[] g = new ArrayList[n];
        for (int i = 0; i < n; i++) g[i] = new ArrayList<>();
        for (int i = 1; i < n; i++) g[parent[i]].add(i);
        int[] np = parent.clone();
        int[] last = new int[26];
        Arrays.fill(last, -1);
        dfs(0, s, g, np, last);
        List<Integer>[] ng = new ArrayList[n];
        for (int i = 0; i < n; i++) ng[i] = new ArrayList<>();
        for (int i = 1; i < n; i++) ng[np[i]].add(i);
        int[] ans = new int[n];
        size(0, ng, ans);
        return ans;
    }
    void dfs(int u, String s, List<Integer>[] g, int[] np, int[] last) {
        int c = s.charAt(u) - 'a';
        int prev = last[c];
        if (prev != -1) np[u] = prev;
        last[c] = u;
        for (int v : g[u]) dfs(v, s, g, np, last);
        last[c] = prev;
    }
    int size(int u, List<Integer>[] ng, int[] ans) {
        int sz = 1;
        for (int v : ng[u]) sz += size(v, ng, ans);
        return ans[u] = sz;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} parent
 * @param {string} s
 * @return {number[]}
 */
var findSubtreeSizes = function(parent, s) {
    const n = parent.length;
    const g = Array.from({length: n}, () => []);
    for (let i = 1; i < n; i++) g[parent[i]].push(i);
    const np = parent.slice();
    const last = new Array(26).fill(-1);
    const dfs = (u) => {
        const c = s.charCodeAt(u) - 97;
        const prev = last[c];
        if (prev !== -1) np[u] = prev;
        last[c] = u;
        for (const v of g[u]) dfs(v);
        last[c] = prev;
    };
    dfs(0);
    const ng = Array.from({length: n}, () => []);
    for (let i = 1; i < n; i++) ng[np[i]].push(i);
    const ans = new Array(n);
    const size = (u) => {
        let sz = 1;
        for (const v of ng[u]) sz += size(v);
        return ans[u] = sz;
    };
    size(0);
    return ans;
};
```

### TypeScript
```typescript
function findSubtreeSizes(parent: number[], s: string): number[] {
    const n = parent.length;
    const g: number[][] = Array.from({length: n}, () => []);
    for (let i = 1; i < n; i++) g[parent[i]].push(i);
    const np = parent.slice();
    const last = new Array(26).fill(-1);
    const dfs = (u: number) => {
        const c = s.charCodeAt(u) - 97;
        const prev = last[c];
        if (prev !== -1) np[u] = prev;
        last[c] = u;
        for (const v of g[u]) dfs(v);
        last[c] = prev;
    };
    dfs(0);
    const ng: number[][] = Array.from({length: n}, () => []);
    for (let i = 1; i < n; i++) ng[np[i]].push(i);
    const ans = new Array(n);
    const size = (u: number): number => {
        let sz = 1;
        for (const v of ng[u]) sz += size(v);
        return ans[u] = sz;
    };
    size(0);
    return ans;
};
```

### Python3
```python
class Solution:
    def findSubtreeSizes(self, parent: List[int], s: str) -> List[int]:
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parent[i]].append(i)
        np = parent[:]
        last = [-1] * 26
        def dfs(u):
            c = ord(s[u]) - 97
            prev = last[c]
            if prev != -1:
                np[u] = prev
            last[c] = u
            for v in g[u]:
                dfs(v)
            last[c] = prev
        dfs(0)
        ng = [[] for _ in range(n)]
        for i in range(1, n):
            ng[np[i]].append(i)
        ans = [0] * n
        def size(u):
            sz = 1
            for v in ng[u]:
                sz += size(v)
            ans[u] = sz
            return sz
        size(0)
        return ans
```

### Go
```go
func findSubtreeSizes(parent []int, s string) []int {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        g[parent[i]] = append(g[parent[i]], i)
    }
    np := make([]int, n)
    copy(np, parent)
    last := make([]int, 26)
    for i := range last {
        last[i] = -1
    }
    var dfs func(int)
    dfs = func(u int) {
        c := int(s[u] - 'a')
        prev := last[c]
        if prev != -1 {
            np[u] = prev
        }
        last[c] = u
        for _, v := range g[u] {
            dfs(v)
        }
        last[c] = prev
    }
    dfs(0)
    ng := make([][]int, n)
    for i := 1; i < n; i++ {
        ng[np[i]] = append(ng[np[i]], i)
    }
    ans := make([]int, n)
    var size func(int) int
    size = func(u int) int {
        sz := 1
        for _, v := range ng[u] {
            sz += size(v)
        }
        ans[u] = sz
        return sz
    }
    size(0)
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax changes.

I first allocate an adjacency list of size n and fill it by walking through the parent array from index 1 to n-1. This gives me a ready-to-traverse original tree.

I then create a newParent array that starts as an exact copy of the input parent array. At the same time I prepare a last-seen array of length 26 filled with -1.

The first DFS is the heart of the solution. When I enter a node u I compute its character index (s[u] - 'a'). I look up the value currently stored for that character. If it is not -1 it means I have already seen a matching ancestor, so I write that ancestor into newParent[u]. Next I overwrite the last-seen slot with u itself so that my descendants will treat me as their candidate. After finishing the recursive calls on all children I restore the previous value. Restoring is essential; without it a sibling would incorrectly think the current node is still its ancestor.

Once the first DFS returns, newParent contains the correct final parent for every node. I build a second adjacency list from this array exactly the same way I built the first one.

The second DFS is a classic subtree-size computation. For each node I start with size 1 (myself) and add the sizes returned by every child. I store the result in an answer array as I go. When the call on the root finishes, the answer array is complete and can be returned.

Edge cases are handled naturally: the root never changes parent, nodes without a matching ancestor keep their original parent, and the simultaneous nature of the moves is respected because the first DFS works only on the original tree.

## Examples

**Example 1**

Input: parent = [-1,0,0,1,1,1], s = "abaabc"  
Output: [6,3,1,1,1,1]

Trace:  
- Node 3 has character 'a'. Its closest matching ancestor is node 0.  
- All other nodes either have no matching ancestor or already point to the correct one.  
- After reparenting, the final tree has root 0 with children 1 and 2; node 1 keeps children 4 and 5 and also receives node 3.  
- Subtree sizes become 6 (root), 3 (node 1), 1 (node 2), 1 (node 3), 1 (node 4), 1 (node 5).

**Example 2**

Input: parent = [-1,0,4,0,1], s = "abbba"  
Output: [5,2,1,1,1]

Trace:  
- Node 4 has character 'b' and finds node 0 as its closest matching ancestor, so it moves under 0.  
- Node 2 has character 'b' and finds node 1 as its closest matching ancestor, so it moves under 1.  
- The final tree has root 0 with children 1 and 3; node 1 has child 2; node 4 becomes a direct child of 0.  
- Subtree sizes evaluate to [5,2,1,1,1].

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `main.cpp`.  
2. Compile with `g++ -std=c++17 main.cpp -o main`.  
3. Run with `./main`.

**Java**  
1. Save the code in a file named `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run with `java Solution`.

**JavaScript**  
1. Save the code in a file named `solution.js`.  
2. Run with `node solution.js`.

**TypeScript**  
1. Save the code in a file named `solution.ts`.  
2. Compile with `tsc solution.ts`.  
3. Run the generated JavaScript with `node solution.js`.

**Python3**  
1. Save the code in a file named `solution.py`.  
2. Run with `python3 solution.py`.

**Go**  
1. Save the code in a file named `main.go`.  
2. Run with `go run main.go`.

In each case you will need to add a small driver that reads the parent array and the string, calls the function, and prints the result. The core solution function is ready to drop into any online judge.

## Notes & Optimizations

- The solution runs in strict linear time and uses only a constant amount of extra memory beyond the input and output arrays, which is optimal for this problem.
- Because the character set is limited to 26 letters, the last-seen array never becomes a bottleneck.
- An alternative approach would be to climb the parent chain for every node while caching results, but that can degrade to quadratic time on skewed trees and is therefore unsafe under the given constraints.
- The two-DFS design keeps the code simple and easy to verify. Merging the size calculation into the first DFS is possible but makes the logic harder to follow and does not improve asymptotic complexity.
- The simultaneous reparenting requirement is satisfied automatically because all decisions are made on the original tree before any edges are changed.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
