# Explanation of DFS-Based Approach for Removing Stones

## Overview

This repository contains the implementation of a Depth-First Search (DFS) approach to solve the problem of removing the maximum number of stones from a 2D plane while ensuring that at least one stone remains from each connected component.

The problem is solved in five different programming languages: C++, Java, JavaScript, Python, and Go. Below, you'll find a step-by-step explanation of the logic used in each implementation.

## Step-by-Step Explanation

### 1. **Understanding the Problem**

- We are given `n` stones on a 2D plane, each located at integer coordinates.
- Two stones are connected if they share the same row or the same column.
- We want to remove as many stones as possible such that at least one stone remains from each connected component.

### 2. **Graph Representation**

- The problem can be modeled as a graph:
  - Each stone is represented as a node.
  - An edge is drawn between two nodes if the corresponding stones share the same row or column.
- The goal is to count the number of connected components in this graph. Once we have the connected components, the maximum number of stones that can be removed is `n - numComponents`, where `numComponents` is the number of connected components.

### 3. **Depth-First Search (DFS) Approach**

- **Step 1**: **Initialize the Graph Structure**
  - Create an adjacency list to represent the graph where each stone is connected to other stones that share the same row or column.

- **Step 2**: **Build the Graph**
  - Iterate over all pairs of stones.
  - If two stones share the same row or column, add an edge between them in the adjacency list.

- **Step 3**: **DFS to Find Connected Components**
  - Initialize a data structure (`set`, `visited map`, or `boolean array`) to keep track of visited nodes (stones).
  - Iterate through each stone:
    - If the stone hasn't been visited, start a DFS from that stone to explore all stones in its connected component.
    - After completing the DFS for a component, increment the component count.

- **Step 4**: **Calculate the Result**
  - The maximum number of stones that can be removed is the total number of stones minus the number of connected components.

### 4. **Implementation Details**

#### C++ Implementation

- **Language Constructs**:
  - `vector` for adjacency list.
  - `unordered_set` for tracking visited nodes.
- **DFS Function**:
  - Recursively visits all neighbors of a node, marking them as visited.
- **Graph Building**:
  - Nested loops to find and connect stones in the same row/column.

#### Java Implementation

- **Language Constructs**:
  - `List<List<Integer>>` for adjacency list.
  - `Set<Integer>` for visited nodes.
- **DFS Function**:
  - Similar recursive approach to traverse all connected stones.
- **Graph Building**:
  - Uses `ArrayList` to maintain adjacency lists for each stone.

#### JavaScript Implementation

- **Language Constructs**:
  - `Array` for adjacency list.
  - `Set` for visited nodes.
- **DFS Function**:
  - Recursive function to explore connected stones.
- **Graph Building**:
  - Uses arrays to represent the graph and connects stones based on shared row/column.

#### Python Implementation

- **Language Constructs**:
  - `List` for adjacency list.
  - `Set` for visited nodes.
- **DFS Function**:
  - Recursive DFS to visit connected components.
- **Graph Building**:
  - Iterates over pairs of stones and builds the adjacency list.

#### Go Implementation

- **Language Constructs**:
  - `[][]int` for adjacency list.
  - `map[int]bool` for tracking visited nodes.
- **DFS Function**:
  - DFS implemented as a separate function, marking nodes as visited.
- **Graph Building**:
  - Constructs the graph by connecting stones with shared rows/columns.

### 5. **Final Notes**

- The DFS approach is efficient and works well with the problem constraints.
- The main challenge lies in correctly identifying and connecting all stones that belong to the same connected component.
- This method ensures that the maximum number of stones are removed while keeping the graph connected.
