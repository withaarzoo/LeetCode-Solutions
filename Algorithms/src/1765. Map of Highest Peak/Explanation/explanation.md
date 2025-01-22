# 🚀 Step-by-Step Explanation of the Solution

This README provides a detailed, step-by-step explanation of how the solution for the "Highest Peak" problem works across multiple languages: C++, Java, JavaScript, Python, and Go. Each explanation dives into the logic while keeping the explanation language-agnostic.

---

## **🛠 Key Steps in Solving the Problem**

1. **Input Grid Initialization**  
   - We are given a grid where cells are either `1` (water) or `0` (land).  
   - The goal is to determine the height of each cell such that:  
     - Water cells have a height of `0`.  
     - Land cells have heights calculated based on their distance from the nearest water cell.

2. **Mark Water Cells**  
   - Traverse the grid and identify all cells marked as water (`1`).  
   - These cells are immediately set to a height of `0`.  
   - Land cells are marked as "unvisited" or assigned a default value (e.g., `-1` or `Infinity`).

3. **Breadth-First Search (BFS) Initialization**  
   - Use a queue (or similar data structure) to perform a level-wise traversal (BFS).  
   - Initially, add all water cells into the queue as the starting points.  

4. **BFS Traversal to Calculate Heights**  
   - Process each cell in the queue:  
     - For each cell, check all its neighbors (up, down, left, right).  
     - If the neighbor is unvisited (land), calculate its height as `current_cell_height + 1`.  
     - Add the neighbor to the queue for further exploration.  

5. **Continue Until All Cells Are Processed**  
   - BFS ensures that heights are calculated layer by layer, starting from the water cells.  
   - This guarantees that each cell's height is the shortest distance from the nearest water cell.  

6. **Return the Resulting Grid**  
   - Once BFS completes, the updated grid contains the required heights for all cells.  

---

### **📌 Explanation for Each Language**

#### **C++ Code**  

1. **Define the Input and Output Grid**  
   - Use a 2D vector to represent the grid.  

2. **Initialize the Queue**  
   - Use a queue to store coordinates of all water cells (`{x, y}`).  

3. **BFS Traversal**  
   - Iterate over all four possible directions (up, down, left, right) using a directions array.  
   - For each unvisited neighbor, update its height and push it into the queue.  

4. **Complete the Grid Update**  
   - Continue the BFS until the queue is empty.  
   - The final grid contains heights for all cells.

---

#### **Java Code**  

1. **Prepare the Input Grid**  
   - Use a 2D array for the grid and a `Queue` to store coordinates of water cells.  

2. **Add Water Cells to the Queue**  
   - Iterate through the grid, marking water cells as `0` and land cells as `-1`.  

3. **Perform BFS Traversal**  
   - Use a directions array to move in all four cardinal directions.  
   - For each neighbor, calculate the height (`current_height + 1`) and enqueue it.  

4. **Return the Updated Grid**  
   - BFS ensures that all cells are visited in the shortest possible distance order.  

---

#### **JavaScript Code**  

1. **Grid Initialization**  
   - Represent the grid as a 2D array and use a `queue` to track BFS levels.  

2. **Identify Water Cells**  
   - Mark all water cells (`1`) with height `0` and add them to the queue.  
   - Land cells are initially marked as unvisited (`-1`).  

3. **BFS Traversal**  
   - For each cell, explore its neighbors using predefined directions.  
   - Update the neighbor's height if it is unvisited, then add it to the queue.  

4. **Update the Grid**  
   - Continue processing until all cells are visited and their heights are calculated.

---

#### **Python Code**  

1. **Input Grid and Initialization**  
   - Use a 2D list to represent the grid.  
   - Initialize a `deque` with all water cells and mark land cells as unvisited (`-1`).  

2. **BFS Setup**  
   - Store all water cells in the deque with height `0`.  
   - Use directions to move up, down, left, and right during traversal.  

3. **Calculate Heights Using BFS**  
   - For each cell, check all its neighbors.  
   - If the neighbor is unvisited, assign its height (`current_height + 1`) and add it to the deque.  

4. **Return the Result**  
   - The final grid contains calculated heights for all cells.

---

#### **Go Code**  

1. **Input Grid Representation**  
   - Use a 2D slice to represent the grid.  

2. **Initialize the Queue**  
   - Create a queue and add all water cells as starting points.  

3. **BFS Traversal**  
   - Use a directions array to explore all neighbors of each cell.  
   - Update the height of each unvisited neighbor and push it to the queue.  

4. **Update and Return the Grid**  
   - Continue BFS until all cells are processed.  
   - The resulting grid contains the required heights.

---

### **🌟 Complexity Analysis**

- **Time Complexity**:  
  $$O(n \times m)$$, where \(n\) and \(m\) are the dimensions of the grid. Each cell is processed once.  

- **Space Complexity**:  
  $$O(n \times m)$$ for the BFS queue and auxiliary storage.  

---

### **🔗 Additional Notes**

This problem emphasizes the importance of BFS for multi-source shortest path problems. The key takeaway is that BFS ensures layer-wise traversal, making it ideal for calculating distances in grid-based problems.
