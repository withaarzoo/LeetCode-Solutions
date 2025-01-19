# Trapping Rain Water II - Solution Explanation

This repository provides solutions to the **Trapping Rain Water II** problem implemented in multiple programming languages: C++, Java, JavaScript, Python, and Go. Below is a detailed step-by-step explanation for each language.

---

## C++ Code

1. **Initialize the Data Structures**  
   - Use a priority queue (min-heap) to efficiently track the smallest boundary height.
   - Maintain a visited array to ensure each cell is processed only once.

2. **Push the Boundary Cells into the Min-Heap**  
   - Add all the boundary cells (first row, last row, first column, last column) to the min-heap.
   - Mark these cells as visited.

3. **Process the Cells in the Min-Heap**  
   - Use the min-heap to simulate water filling from the lowest boundary height.
   - For each cell, check its unvisited neighbors and calculate the trapped water for that neighbor.

4. **Calculate Trapped Water**  
   - The trapped water for a neighbor cell is determined by the difference between the current boundary height and the neighbor's height (if positive).

5. **Update the Min-Heap**  
   - Add the neighbor cell to the min-heap with its updated height and continue the process.

6. **Return the Total Trapped Water**  
   - Keep a running sum of all trapped water and return the result.

---

## Java Code

1. **Setup the Min-Heap and Visited Array**  
   - Use a `PriorityQueue` for the min-heap to store cells by their height.
   - Create a `visited` 2D array to track processed cells.

2. **Add Boundary Cells to the Heap**  
   - Iterate over all the boundary cells (edges of the heightMap) and add them to the heap.
   - Mark them as visited to avoid duplicate processing.

3. **Simulate Water Filling**  
   - Process cells in the heap one by one, starting with the smallest height.
   - For each cell, visit its unprocessed neighbors.

4. **Compute Water Volume for Each Neighbor**  
   - Compare the neighbor’s height with the current boundary height to calculate trapped water.

5. **Update the Boundary Height**  
   - Add the neighbor cell to the heap with its updated height (max of its own height or the current boundary).

6. **Aggregate the Results**  
   - Sum up the trapped water for all cells and return the total.

---

## JavaScript Code

1. **Initialize a Min-Heap and a Visited Set**  
   - Use a custom implementation of a priority queue or a library to create a min-heap.
   - Track visited cells using a `Set`.

2. **Push Boundary Cells into the Min-Heap**  
   - Add all boundary cells of the 2D grid to the heap, storing their height and coordinates.
   - Mark these cells as visited to prevent reprocessing.

3. **Iterate Through the Min-Heap**  
   - Process the smallest cell height first to maintain boundary consistency.
   - For each cell, examine its neighbors.

4. **Calculate Water Trapped**  
   - Check if the neighbor cell is lower than the current boundary height.
   - If so, calculate the water trapped and add it to the running total.

5. **Update the Neighbor’s Height**  
   - Add the neighbor cell back into the heap with its updated height, which ensures proper boundary simulation.

6. **Return the Total Water Volume**  
   - Continue until the heap is empty and return the accumulated trapped water.

---

## Python Code

1. **Setup Initial Structures**  
   - Use the `heapq` library for the min-heap and a `visited` matrix to track processed cells.

2. **Add Boundary Cells to the Min-Heap**  
   - Loop through all edges of the heightMap and push boundary cells into the heap.
   - Mark these cells as visited to avoid duplicate computation.

3. **Simulate Water Trapping**  
   - Extract the smallest height cell from the heap.
   - For each unvisited neighbor, determine if water can be trapped.

4. **Compute Trapped Water**  
   - Calculate the trapped water as the difference between the boundary height and the neighbor’s height.

5. **Update and Continue**  
   - Add the neighbor to the heap with its updated boundary height.
   - Repeat the process for all neighbors until the heap is empty.

6. **Return the Result**  
   - Maintain a total water volume counter and return its value after processing all cells.

---

## Go Code

1. **Initialize Data Structures**  
   - Use a `heap` package to create a priority queue for the min-heap.
   - Create a `visited` 2D slice to track processed cells.

2. **Add Boundary Cells**  
   - Push all the cells along the grid boundary into the min-heap.
   - Mark these cells as visited.

3. **Process the Min-Heap**  
   - Continuously extract the smallest height cell from the heap.
   - For each neighbor of the current cell, check if it can trap water.

4. **Calculate Trapped Water**  
   - Determine the trapped water by comparing the boundary height and the neighbor cell height.

5. **Push Neighbor Cells**  
   - Add each unvisited neighbor to the heap with an updated height (either its own height or the boundary height).

6. **Return the Accumulated Water**  
   - Keep a running total of trapped water and return it after processing all cells.

---

### Final Notes

All the implementations follow the same core logic:

1. **Simulate water filling using a priority queue.**
2. **Process cells in the order of increasing height.**
3. **Calculate water trapped for each neighbor.**
4. **Update the boundary dynamically using the min-heap.**

This ensures the algorithm efficiently computes the trapped water for any 2D height map.
