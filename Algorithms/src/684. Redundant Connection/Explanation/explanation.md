# **Intuition**

When I first saw this problem, it reminded me of a cycle detection problem in an undirected graph. The problem states that we have a connected graph with **n nodes** and **n edges** (i.e., a tree with an extra edge).  

A **tree** with `n` nodes should have exactly `n-1` edges. Since we have `n` edges, this means there's **one extra edge** forming a cycle, and we need to **find and remove that extra edge** while keeping the graph connected.  

A common way to handle **cycle detection in an undirected graph** is by using **Union-Find (Disjoint Set Union, DSU)**. So, my first instinct was to solve this using **DSU with path compression and union by rank**.

---

# **Approach**

To find the **redundant edge** efficiently, we use the **Union-Find (DSU) data structure**:

1. **Initialize DSU:**  
   - We maintain a parent array where each node is its own parent initially.  
   - We also maintain a rank array for optimization.

2. **Process each edge one by one:**  
   - If two nodes in an edge already belong to the same connected component (i.e., they have the same root parent), then adding this edge **forms a cycle**.  
   - This means the current edge is **redundant**, so we return it as the answer.  

3. **Union-Find operations:**  
   - **Find function:** Uses path compression to keep the tree flat.  
   - **Union function:** Uses union by rank to attach smaller trees under larger ones, keeping the structure balanced.  

By iterating over all edges and performing these operations, we can efficiently detect the first edge that forms a cycle.

---

# **Complexity Analysis**

- **Time Complexity:** \(O(n \log n)\)  
  - The `find` and `union` operations have an **amortized** complexity of **O(α(n))**, which is nearly constant.  
  - Since we process `n` edges, the overall complexity is **O(n α(n))**, which simplifies to **O(n log n)** in the worst case.

- **Space Complexity:** **O(n)**  
  - We maintain a `parent` and `rank` array, both of size **O(n)**.
