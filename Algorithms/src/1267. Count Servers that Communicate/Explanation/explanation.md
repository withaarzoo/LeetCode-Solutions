# Count Servers That Communicate  

This repository contains solutions to the problem "Count Servers That Communicate" in **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Below is the **step-by-step explanation** of the logic and approach for each language.

---

## Problem Overview  

You are given a grid where:

- `1` represents a server.
- `0` represents an empty space.

The goal is to count the servers that can communicate. Two servers can communicate if:

- They are in the same row **or** the same column.

---

## Approach (All Languages)  

We use a **two-pass approach**:

1. First pass: Count the servers in each row and column.
2. Second pass: Identify servers that can communicate based on the counts from the first pass.

### Step-by-Step Breakdown  

### C++ Code

1. **Initialize Row and Column Count**: Create two arrays to store the number of servers in each row and column.  
2. **First Pass**: Traverse the grid and update the row and column count for each server found (`1`).  
3. **Second Pass**: Traverse the grid again. For each server, check if its row or column count is greater than `1`. If so, it can communicate. Increment the count.  
4. **Return the Result**: Finally, return the total count of servers that can communicate.

---

### Java Code

1. **Initialize Row and Column Count**: Use two arrays to store the number of servers in each row and column.  
2. **First Pass**: Iterate through the grid, and for every server (`1`), increment the respective row and column counts.  
3. **Second Pass**: Iterate through the grid again. For every server (`1`), check if its row count or column count is greater than `1`. If true, it can communicate, so add to the total count.  
4. **Return the Result**: Return the count of all communicable servers.

---

### JavaScript Code

1. **Initialize Row and Column Count**: Create two arrays, one for rows and one for columns, and initialize them to zero.  
2. **First Pass**: Use nested loops to traverse the grid. For every server (`1`), increment the row and column counters at the respective indices.  
3. **Second Pass**: Loop through the grid again. Check for each server (`1`) if the row or column count at its position is greater than `1`. If so, increment the result count.  
4. **Return the Result**: Return the total count of servers that can communicate.

---

### Python Code

1. **Initialize Row and Column Count**: Use two lists to keep track of the number of servers in each row and column.  
2. **First Pass**: Iterate through the grid. For every server (`1`), increment the respective row and column counts.  
3. **Second Pass**: Iterate through the grid again. For every server (`1`), check if the row count or column count is greater than `1`. If true, add it to the total count.  
4. **Return the Result**: Finally, return the count of servers that can communicate.

---

### Go Code

1. **Initialize Row and Column Count**: Create slices to store the count of servers in each row and column.  
2. **First Pass**: Loop through the grid using nested loops. For every server (`1`), increment the corresponding row and column counts.  
3. **Second Pass**: Loop through the grid again. For each server (`1`), check if the count in its row or column is greater than `1`. If true, increment the total count.  
4. **Return the Result**: Return the total number of servers that can communicate.

---

## Notes  

- All implementations follow the same **two-pass approach** to solve the problem.  
- The time complexity of all solutions is **O(m × n)**, where `m` is the number of rows and `n` is the number of columns.  
- The space complexity of all solutions is **O(m + n)** for storing row and column counts.
