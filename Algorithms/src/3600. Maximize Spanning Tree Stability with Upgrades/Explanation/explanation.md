# Maximize Spanning Tree Stability with Upgrades

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

# Problem Summary

You are given:

* An integer `n` representing `n` nodes labeled from `0` to `n-1`.
* A list of edges where each edge is represented as:

```
[u, v, s, must]
```

Where:

* `u` and `v` are nodes connected by the edge
* `s` is the strength of the edge
* `must` is either `0` or `1`

Rules:

* If `must = 1`, the edge **must be included** in the spanning tree
* Mandatory edges **cannot be upgraded**

You are also given:

* An integer `k` representing the maximum number of upgrades allowed

Upgrade rule:

* An upgrade **doubles the strength of an edge**
* Each optional edge (`must = 0`) can be upgraded **at most once**

The **stability of a spanning tree** is defined as:

```
minimum edge strength among all edges in the tree
```

Goal:

Return the **maximum possible stability** of a valid spanning tree.

If it is impossible to connect all nodes, return `-1`.

---

## Constraints

```
2 <= n <= 10^5
1 <= edges.length <= 10^5
edges[i] = [u_i, v_i, s_i, must_i]
0 <= u_i, v_i < n
u_i != v_i
1 <= s_i <= 10^5
must_i is either 0 or 1
0 <= k <= n
No duplicate edges
```

---

# Intuition

When I first read the problem, I noticed that the stability of the tree depends on the **minimum edge strength** inside the tree.

So the goal is not to maximize the sum of weights, but to **maximize the smallest edge weight** in the spanning tree.

This immediately reminded me of **Maximum Spanning Tree logic**.

However, the problem adds two complications:

1. Some edges are mandatory and must be included.
2. Optional edges can be upgraded up to `k` times.

My idea was:

1. First include all mandatory edges.
2. If mandatory edges create a cycle, the answer is impossible.
3. Then connect remaining components using optional edges with the largest strength.
4. After building the spanning tree, upgrade the weakest edges to increase the final stability.

To efficiently manage connected components, I used **Union-Find (Disjoint Set Union)**.

---

# Approach

Step 1 — Initialize Union-Find

Each node initially belongs to its own component.

Step 2 — Add mandatory edges

For every edge where `must = 1`:

* If it connects two already connected nodes, it forms a cycle → return `-1`
* Otherwise union them
* Track the minimum strength among mandatory edges

Step 3 — Store optional edges

All edges with `must = 0` are stored for later processing.

Step 4 — Build a Maximum Spanning Tree

Sort optional edges in **descending order of strength**.

Add edges greedily if they connect two different components.

Step 5 — Check connectivity

If after processing we still have more than one component, it means we cannot connect the graph.

Return `-1`.

Step 6 — Apply upgrades

The weakest edges determine the stability.

So we sort selected optional edges in ascending order and upgrade the smallest ones first.

Step 7 — Compute final stability

The answer is the minimum strength among all edges in the final spanning tree.

---

# Data Structures Used

Union-Find (Disjoint Set Union)

Purpose:

* Detect cycles
* Track connected components efficiently

Other structures:

* Vector/List for optional edges
* Vector/List for selected edges
* Sorting utilities

---

# Operations & Behavior Summary

Mandatory edge processing

* Must always be included
* Cannot be upgraded
* Cycle detection required

Optional edge processing

* Used to connect remaining components
* Can be upgraded once

Upgrade strategy

* Upgrade smallest edges first
* At most `k` upgrades

---

# Complexity

Time Complexity

```
O(E log E)
```

Where:

* `E` = number of edges

Sorting optional edges dominates the runtime.

Union-Find operations are almost constant time.

Space Complexity

```
O(E + N)
```

Used for:

* Union-Find arrays
* Optional edge storage
* Selected edge storage

---

# Multi-language Solutions

## C++

```cpp
class Solution {
public:

    struct DSU {
        vector<int> parent, rank;

        DSU(int n) {
            parent.resize(n);
            rank.resize(n);

            for(int i=0;i<n;i++)
                parent[i]=i;
        }

        int find(int x){
            if(parent[x]==x)
                return x;

            return parent[x]=find(parent[x]);
        }

        bool unite(int a,int b){
            a=find(a);
            b=find(b);

            if(a==b)
                return false;

            if(rank[a]<rank[b])
                swap(a,b);

            parent[b]=a;

            if(rank[a]==rank[b])
                rank[a]++;

            return true;
        }
    };


    int maxStability(int n, vector<vector<int>>& edges, int k) {

        DSU dsu(n);

        int components=n;
        int mandatoryMin=INT_MAX;

        vector<vector<int>> optional;

        for(auto &e:edges){

            int u=e[0];
            int v=e[1];
            int s=e[2];
            int must=e[3];

            if(must==1){

                if(!dsu.unite(u,v))
                    return -1;

                components--;
                mandatoryMin=min(mandatoryMin,s);

            }else{

                optional.push_back(e);
            }
        }


        sort(optional.begin(),optional.end(),
        [](auto &a,auto &b){
            return a[2]>b[2];
        });


        vector<int> used;

        for(auto &e:optional){

            if(dsu.unite(e[0],e[1])){

                used.push_back(e[2]);
                components--;

                if(components==1)
                    break;
            }
        }


        if(components>1)
            return -1;


        sort(used.begin(),used.end());


        int ans=mandatoryMin;

        for(int w:used){

            int val=w;

            if(k>0){
                val*=2;
                k--;
            }

            if(ans==INT_MAX)
                ans=val;
            else
                ans=min(ans,val);
        }

        return ans;
    }
};
```

---

## Java

```java
class Solution {

    class DSU{

        int[] parent;
        int[] rank;

        DSU(int n){

            parent=new int[n];
            rank=new int[n];

            for(int i=0;i<n;i++)
                parent[i]=i;
        }


        int find(int x){

            if(parent[x]==x)
                return x;

            return parent[x]=find(parent[x]);
        }


        boolean union(int a,int b){

            a=find(a);
            b=find(b);

            if(a==b)
                return false;


            if(rank[a]<rank[b]){
                int t=a;
                a=b;
                b=t;
            }

            parent[b]=a;

            if(rank[a]==rank[b])
                rank[a]++;

            return true;
        }
    }


    public int maxStability(int n, int[][] edges, int k) {

        DSU dsu=new DSU(n);

        int comp=n;
        int mandatoryMin=Integer.MAX_VALUE;

        List<int[]> optional=new ArrayList<>();


        for(int[] e:edges){

            if(e[3]==1){

                if(!dsu.union(e[0],e[1]))
                    return -1;

                comp--;
                mandatoryMin=Math.min(mandatoryMin,e[2]);

            }else{

                optional.add(e);
            }
        }


        optional.sort((a,b)->b[2]-a[2]);


        List<Integer> used=new ArrayList<>();

        for(int[] e:optional){

            if(dsu.union(e[0],e[1])){

                used.add(e[2]);
                comp--;

                if(comp==1)
                    break;
            }
        }


        if(comp>1)
            return -1;


        Collections.sort(used);

        int ans=mandatoryMin;

        for(int w:used){

            int val=w;

            if(k>0){
                val*=2;
                k--;
            }

            if(ans==Integer.MAX_VALUE)
                ans=val;
            else
                ans=Math.min(ans,val);
        }

        return ans;
    }
}
```

---

## JavaScript

```javascript
var maxStability = function(n, edges, k) {

    const parent=Array(n).fill(0).map((_,i)=>i);
    const rank=Array(n).fill(0);


    function find(x){
        if(parent[x]!==x)
            parent[x]=find(parent[x]);

        return parent[x];
    }


    function union(a,b){

        a=find(a);
        b=find(b);

        if(a===b)
            return false;


        if(rank[a]<rank[b]){
            let temp=a;
            a=b;
            b=temp;
        }

        parent[b]=a;

        if(rank[a]===rank[b])
            rank[a]++;

        return true;
    }


    let comp=n;
    let mandatoryMin=Infinity;

    let optional=[];


    for(const e of edges){

        if(e[3]===1){

            if(!union(e[0],e[1]))
                return -1;

            comp--;
            mandatoryMin=Math.min(mandatoryMin,e[2]);

        }else{

            optional.push(e);
        }
    }


    optional.sort((a,b)=>b[2]-a[2]);


    let used=[];

    for(const e of optional){

        if(union(e[0],e[1])){

            used.push(e[2]);
            comp--;

            if(comp===1)
                break;
        }
    }


    if(comp>1)
        return -1;


    used.sort((a,b)=>a-b);


    let ans=mandatoryMin;

    for(let w of used){

        let val=w;

        if(k>0){
            val*=2;
            k--;
        }

        if(ans===Infinity)
            ans=val;
        else
            ans=Math.min(ans,val);
    }


    return ans;
};
```

---

## Python3

```python
class Solution:

    def maxStability(self, n: int, edges: List[List[int]], k: int) -> int:

        parent=list(range(n))
        rank=[0]*n

        def find(x):
            if parent[x]!=x:
                parent[x]=find(parent[x])
            return parent[x]


        def union(a,b):

            ra=find(a)
            rb=find(b)

            if ra==rb:
                return False

            if rank[ra]<rank[rb]:
                ra,rb=rb,ra

            parent[rb]=ra

            if rank[ra]==rank[rb]:
                rank[ra]+=1

            return True


        comp=n
        mandatory_min=float('inf')

        optional=[]


        for u,v,s,m in edges:

            if m==1:

                if not union(u,v):
                    return -1

                comp-=1
                mandatory_min=min(mandatory_min,s)

            else:

                optional.append((u,v,s))


        optional.sort(key=lambda x:-x[2])


        used=[]

        for u,v,s in optional:

            if union(u,v):

                used.append(s)
                comp-=1

                if comp==1:
                    break


        if comp>1:
            return -1


        used.sort()


        ans=mandatory_min

        for w in used:

            val=w

            if k>0:
                val*=2
                k-=1

            if ans==float('inf'):
                ans=val
            else:
                ans=min(ans,val)

        return ans
```

---

## Go

```go
func maxStability(n int, edges [][]int, k int) int {

    parent:=make([]int,n)
    rank:=make([]int,n)

    for i:=0;i<n;i++{
        parent[i]=i
    }


    var find func(int) int

    find = func(x int) int{
        if parent[x]!=x{
            parent[x]=find(parent[x])
        }
        return parent[x]
    }


    union := func(a,b int) bool{

        ra:=find(a)
        rb:=find(b)

        if ra==rb{
            return false
        }

        if rank[ra]<rank[rb]{
            ra,rb=rb,ra
        }

        parent[rb]=ra

        if rank[ra]==rank[rb]{
            rank[ra]++
        }

        return true
    }


    comp:=n
    mandatoryMin:=int(^uint(0)>>1)

    optional:=[][]int{}


    for _,e:=range edges{

        u:=e[0]
        v:=e[1]
        s:=e[2]
        m:=e[3]

        if m==1{

            if !union(u,v){
                return -1
            }

            comp--

            if s<mandatoryMin{
                mandatoryMin=s
            }

        }else{

            optional=append(optional,e)
        }
    }


    sort.Slice(optional,func(i,j int) bool{
        return optional[i][2] > optional[j][2]
    })


    used:=[]int{}

    for _,e:=range optional{

        if union(e[0],e[1]){

            used=append(used,e[2])
            comp--

            if comp==1{
                break
            }
        }
    }


    if comp>1{
        return -1
    }


    sort.Ints(used)


    ans:=mandatoryMin

    for _,w:=range used{

        val:=w

        if k>0{
            val*=2
            k--
        }

        if ans==int(^uint(0)>>1){
            ans=val
        }else if val<ans{
            ans=val
        }
    }


    return ans
}
```

---

# Step-by-step Detailed Explanation

1. Initialize DSU

Each node starts as its own parent.

This helps us detect cycles and connect components efficiently.

1. Process mandatory edges

Mandatory edges must be added first.

If adding one forms a cycle, a valid spanning tree cannot exist.

We also record the minimum strength among mandatory edges.

1. Store optional edges

Optional edges are saved for later because they may or may not be used.

1. Build the spanning tree

Sort optional edges by strength in descending order.

Greedily add edges that connect different components.

This builds a Maximum Spanning Tree.

1. Check connectivity

If the graph is still disconnected, return `-1`.

1. Optimize stability with upgrades

Sort the selected optional edges in ascending order.

Upgrade the smallest edges first because they affect the minimum value.

1. Compute final stability

The stability is the minimum strength among all edges in the tree.

Return this value.

---

# Examples

Example 1

```
Input
n = 3
edges = [[0,1,2,1],[1,2,3,0]]
k = 1

Output
2
```

Example 2

```
Input
n = 3
edges = [[0,1,4,0],[1,2,3,0],[0,2,1,0]]
k = 2

Output
6
```

Example 3

```
Input
n = 3
edges = [[0,1,1,1],[1,2,1,1],[2,0,1,1]]
k = 0

Output
-1
```

---

# How to use / Run locally

Clone repository

```
git clone https://github.com/your-username/your-repo.git
```

Run C++ solution

```
g++ solution.cpp
./a.out
```

Run Python solution

```
python solution.py
```

Run Java solution

```
javac Solution.java
java Solution
```

Run JavaScript

```
node solution.js
```

Run Go

```
go run solution.go
```

---

# Notes & Optimizations

Union-Find uses path compression and union by rank.

This keeps operations nearly constant time.

Sorting edges ensures we always choose the strongest edges first.

Upgrading the smallest selected edges maximizes the minimum value of the tree.

This greedy strategy ensures optimal stability.

---

# Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
