# Eventual Safe States - Step-by-Step Explanation for Multiple Languages

This repository provides solutions to the **"Eventual Safe States"** problem implemented in **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Below is a step-by-step explanation for each language, with a friendly and approachable tone. Let's dive into how we solved the problem and broke it into understandable steps!

---

## C++ Code Explanation

1. **Graph Representation**:  
   Represent the input graph as an adjacency list where each node points to the nodes it can visit.

2. **Tracking Node States**:  
   Use an array to track the states of each node:
   - 0: Node is unvisited.
   - 1: Node is being visited (part of the current path).
   - 2: Node is already determined to be safe.

3. **DFS Traversal**:  
   Perform a depth-first search to determine if a node leads to a cycle:
   - Mark the node as being visited.
   - For each connected node, recursively check if it leads to a cycle.
   - If any neighbor leads to a cycle, the current node is unsafe.

4. **Marking Safe Nodes**:  
   If all neighbors are safe, mark the current node as safe and continue.

5. **Result Compilation**:  
   After processing all nodes, return the list of nodes marked as safe.

---

## Java Code Explanation

1. **Graph Representation**:  
   Convert the input into an adjacency list to represent the graph efficiently.

2. **State Tracking**:  
   Use an array to keep track of whether a node is safe:
   - 0: Node is not yet processed.
   - 1: Node is safe.
   - 2: Node is unsafe or part of a cycle.

3. **DFS Logic**:  
   Recursively check each node:
   - If a node is already marked, return its state.
   - If not, explore its neighbors.
   - If any neighbor leads to a cycle, the current node is unsafe.

4. **Storing Safe Nodes**:  
   Collect all nodes that are eventually determined to be safe.

5. **Return Result**:  
   Sort the list of safe nodes and return as the result.

---

## JavaScript Code Explanation

1. **Graph Conversion**:  
   Parse the input graph into a usable format, typically an adjacency list.

2. **State Management**:  
   Create a state array to track each node's status:
   - `0`: Not visited.
   - `1`: Safe.
   - `2`: Unsafe or cyclic.

3. **Recursive DFS Function**:  
   Write a helper function to explore each node recursively:
   - If a node is already marked, return its state.
   - Otherwise, check all its neighbors.
   - If any neighbor is part of a cycle, mark the current node as unsafe.

4. **Identifying Safe Nodes**:  
   Add nodes identified as safe to the result array.

5. **Final Output**:  
   Return the sorted list of safe nodes.

---

## Python Code Explanation

1. **Graph Input**:  
   Read and transform the input graph into an adjacency list for easier traversal.

2. **State Array**:  
   Use an array to track whether a node is:
   - `0`: Not yet processed.
   - `1`: Safe.
   - `2`: Unsafe.

3. **DFS Implementation**:  
   Define a recursive function to explore nodes:
   - Check if a node has been visited before.
   - Explore its neighbors recursively.
   - If all neighbors are safe, mark the node as safe.

4. **Collecting Results**:  
   Iterate through all nodes and call the DFS function to determine their safety.

5. **Output**:  
   Return a sorted list of nodes that are safe.

---

## Go Code Explanation

1. **Graph Structure**:  
   Represent the input graph as an adjacency list using slices for efficient traversal.

2. **State Management**:  
   Use a slice to track node states:
   - `0`: Not processed.
   - `1`: Safe.
   - `2`: Unsafe or cyclic.

3. **Recursive DFS Function**:  
   Implement a helper function to determine if a node leads to a cycle:
   - Mark nodes during traversal.
   - Check neighbors recursively.
   - If any neighbor is unsafe, mark the current node as unsafe.

4. **Identifying Safe Nodes**:  
   Iterate through the graph, marking nodes as safe if they meet the conditions.

5. **Result Compilation**:  
   Collect all safe nodes, sort them, and return the final list.

---

## Conclusion

By following the above steps, you can confidently understand how the solutions work across different programming languages. Each implementation uses depth-first search (DFS) and state tracking to determine whether nodes are eventually safe. Feel free to explore the respective code files for a more detailed view of the implementation!
