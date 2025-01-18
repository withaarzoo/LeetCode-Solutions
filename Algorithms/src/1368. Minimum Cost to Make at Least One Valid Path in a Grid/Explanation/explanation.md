# Minimum Cost to Make Grid Traversable 🚀

This repository contains solutions in multiple languages for the problem of finding the **minimum cost** to traverse a grid based on directional costs.

## Intuition 🧠

When approaching this problem, the goal is to minimize the total cost incurred while navigating from the top-left to the bottom-right corner of the grid. Each cell offers a preferred direction with a cost of 0; deviating from this incurs an additional cost.

The best approach involves **graph traversal**, treating the grid as a weighted graph and leveraging algorithms like **0-1 BFS** to ensure efficiency.

---

## Approach 🛠️

1. **Model the Grid as a Weighted Graph**:  
   Treat each cell as a node. Moving to the preferred direction has a cost of 0, while deviating adds a cost of 1.

2. **Use 0-1 BFS**:  
   Utilize a deque for efficient processing. Push nodes with cost 0 to the front and nodes with cost 1 to the back. This ensures the shortest path is always evaluated first.

3. **Direction Mapping**:  
   Map directions (right, left, down, up) to specific grid movements, allowing straightforward navigation.

4. **Handle Boundary Conditions**:  
   Ensure moves stay within the grid boundaries and prevent revisiting processed cells.

---

## Complexity 📊

- **Time Complexity**:  
  $$O(n \times m)$$, where \(n\) and \(m\) are the dimensions of the grid. Each cell is processed at most once.
  
- **Space Complexity**:  
  $$O(n \times m)$$, for the visited array and deque.

---

## Step-by-Step Explanation 📝

### C++ Solution 🔵

1. **Initialize Deque**:  
   Start with a deque and push the top-left cell with a cost of 0.

2. **Direction Mapping**:  
   Create an array to represent directions for right, left, down, and up.

3. **Grid Traversal**:  
   For each cell, calculate its cost based on the direction. Push cells with a cost of 0 to the front of the deque and others to the back.

4. **Boundary Check**:  
   Ensure the new position is within the grid. Avoid revisiting cells.

5. **Return Cost**:  
   Once the bottom-right cell is reached, return the accumulated cost.

---

### Java Solution 🟡

1. **Queue Implementation**:  
   Use a `Deque` to implement 0-1 BFS, starting with the top-left cell.

2. **Direction Mapping**:  
   Use arrays for directions and map them to movements (e.g., right = (0, 1)).

3. **Traverse the Grid**:  
   For each cell, calculate the cost for moving in all four directions. Prioritize moves in the preferred direction.

4. **Boundary Conditions**:  
   Validate positions before processing them. Skip cells already visited.

5. **End Condition**:  
   Return the cost when the bottom-right cell is reached.

---

### JavaScript Solution 🟢

1. **Deque Initialization**:  
   Use a `double-ended queue` (e.g., `deque` library) to process cells based on their costs.

2. **Direction Mapping**:  
   Define an array of direction offsets for moving right, left, down, and up.

3. **Traverse and Process**:  
   For each cell, evaluate the cost to move in each direction. Push zero-cost moves to the front and higher-cost moves to the back.

4. **Boundary Handling**:  
   Check for out-of-bound indices and mark visited cells.

5. **Terminate on Completion**:  
   Stop traversal when reaching the bottom-right corner and return the total cost.

---

### Python Solution 🟣

1. **Deque Setup**:  
   Use `collections.deque` to implement 0-1 BFS for efficient grid traversal.

2. **Direction Mapping**:  
   Store direction offsets for right, left, down, and up in a list for easy iteration.

3. **BFS Traversal**:  
   For each cell, calculate costs for all possible moves. Append moves with zero cost to the front and others to the back of the deque.

4. **Boundary and Visited Check**:  
   Ensure moves are within bounds and avoid revisiting cells.

5. **Return Minimum Cost**:  
   Once the bottom-right cell is reached, output the accumulated cost.

---

### Go Solution 🔵

1. **Deque Initialization**:  
   Use a `doubly linked list` from the `container` package to implement the deque.

2. **Directional Offsets**:  
   Use arrays to define the movement offsets for right, left, down, and up.

3. **Process Grid Cells**:  
   For each cell, compute the cost of moving in all directions. Add zero-cost moves to the front and others to the back.

4. **Validate Moves**:  
   Ensure the new cell positions are within bounds and avoid revisiting them.

5. **Finish and Return Cost**:  
   Once the destination is reached, return the minimum cost.

---

## Languages Included 🌐

- **C++**  
- **Java**  
- **JavaScript**  
- **Python**  
- **Go**

Each implementation follows the same logic, optimized for the respective language's syntax and data structures.
