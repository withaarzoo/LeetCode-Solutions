# Find the Maximum Fish in a Grid

## Step-by-Step Explanation

This problem involves finding the maximum number of fish that can be collected in a grid by exploring connected components. Below is a detailed breakdown of the approach and logic used in the implementation across **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Each step explains the logic for all the solutions without directly revealing the code.

---

### General Approach

1. **Identify Connected Components**  
   - The problem can be visualized as identifying connected groups of grid cells containing fish (non-zero values).
   - Use Depth First Search (DFS) or Breadth First Search (BFS) to traverse connected cells starting from any cell with a positive value.

2. **Collect Fish from Connected Cells**  
   - As you traverse each connected component, sum up the values (number of fish) in that component.

3. **Track the Maximum Fish Count**  
   - Compare the total fish collected from each connected component to keep track of the maximum.

---

### C++ Code: Step-by-Step

1. **Grid Traversal**  
   - Loop through each cell in the grid. If the cell contains fish (non-zero value) and hasn't been visited yet, initiate a DFS traversal.

2. **DFS Implementation**  
   - Use a recursive function to traverse all connected cells, summing up the fish values.
   - Mark visited cells to avoid processing them again.

3. **Track Maximum Fish**  
   - After completing the DFS for each component, compare the total fish count with the current maximum and update accordingly.

---

### Java Code: Step-by-Step

1. **Grid Traversal**  
   - Use nested loops to go through all the grid cells. Start DFS whenever a fish-filled cell is found.

2. **DFS as a Helper Function**  
   - Implement DFS recursively or with a stack. For each connected cell, add the fish value to a running total and mark it as visited.

3. **Compare and Update Maximum**  
   - At the end of the traversal for each component, compare its total fish count with the global maximum.

---

### JavaScript Code: Step-by-Step

1. **Iterate Through the Grid**  
   - Use nested `for` loops to inspect each cell in the grid. When a cell with fish is found, initiate a DFS.

2. **Handle Connected Components**  
   - Implement a DFS function using recursion or an iterative approach (using a stack) to explore all connected cells. Add fish values and mark cells as visited.

3. **Update Maximum Fish Count**  
   - After completing the DFS for one component, compare its total fish count with the global maximum and store the higher value.

---

### Python Code: Step-by-Step

1. **Loop Over the Grid**  
   - Iterate through every cell using nested loops. When a cell with fish is encountered, start a DFS.

2. **Recursive DFS Implementation**  
   - Use a recursive function to traverse all connected cells. Add the value of each cell to the total count and mark it as visited.

3. **Global Maximum Update**  
   - At the end of each DFS, compare the total fish collected in that component with the maximum recorded so far and update as needed.

---

### Go Code: Step-by-Step

1. **Traverse Each Cell**  
   - Loop over all the cells in the grid. When a cell containing fish is found, trigger a DFS.

2. **DFS with a Helper Function**  
   - Implement DFS to explore connected cells. Maintain a running total of fish and ensure cells are marked as visited.

3. **Update Maximum Fish Count**  
   - After completing the DFS for a component, compare its total fish count with the global maximum and update.

---

### Complexity Analysis

- **Time Complexity**:  
  $$O(n \times m)$$ for all solutions, where \(n\) and \(m\) are the grid dimensions. This is because each cell is processed once during the traversal.

- **Space Complexity**:  
  - For DFS: \(O(\text{stack depth})\), which is proportional to the size of the grid in the worst case.  
  - For BFS: \(O(n \times m)\) for the queue used to store cells during traversal.
