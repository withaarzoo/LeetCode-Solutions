# First Completely Painted Row or Column

## Problem Overview

You are tasked with determining the first index in the array `arr` where either an entire row or an entire column of the matrix `mat` becomes fully painted. The painting process follows the order given in `arr`.  

---

## Solution Explanation (Step-by-Step)

### **C++ Code**

1. **Initialize Data Structures**:
   - Create a mapping to store the position (row and column) of each value in the matrix. This allows for constant-time lookups.
   - Use two arrays, `rowCount` and `colCount`, to track the number of painted cells in each row and column.

2. **Populate the Mapping**:
   - Iterate through the matrix and record the row and column of every value in the mapping.

3. **Simulate the Painting Process**:
   - Loop through the array `arr`. For each value:
     - Retrieve its row and column using the mapping.
     - Increment the count for the corresponding row and column.

4. **Check Completion**:
   - After updating the counts, check if the row or column has been fully painted. If so, return the current index.

5. **Edge Case**:
   - If no row or column becomes fully painted after processing all elements of `arr`, return `-1`.

---

### **Java Code**

1. **Define Position Mapping**:
   - Use a `HashMap` to map each value in the matrix to its row and column. This avoids searching through the matrix repeatedly.

2. **Track Painting Progress**:
   - Use two arrays, `rowCount` and `colCount`, to count how many cells have been painted for each row and column.

3. **Iterate Through `arr`**:
   - For each number in `arr`, find its position using the mapping and update the painting counts.

4. **Check Completion**:
   - After updating the row and column counts, check if either the row or the column is fully painted. If yes, return the current index.

5. **Handle Edge Cases**:
   - If no row or column is fully painted after processing all values in `arr`, return `-1`.

---

### **JavaScript Code**

1. **Setup Position Map**:
   - Use a `Map` to store the position (row and column) of each value in the matrix for fast lookups.

2. **Track Painted Cells**:
   - Create two arrays, `rowCount` and `colCount`, to track the progress of painting for rows and columns.

3. **Loop Through `arr`**:
   - For each value in `arr`, get its position using the mapping.
   - Increment the corresponding row and column counts.

4. **Check for Completion**:
   - After incrementing, check if the row or column has reached its maximum count. If so, return the current index.

5. **Return Result**:
   - If no row or column is completely painted after processing all elements, return `-1`.

---

### **Python Code**

1. **Create Position Mapping**:
   - Use a dictionary to map each value in the matrix to its row and column. This avoids unnecessary iterations.

2. **Initialize Counters**:
   - Use two lists, `rowCount` and `colCount`, to track how many cells of each row and column are painted.

3. **Simulate Painting**:
   - For every value in `arr`, look up its position in the mapping and update the corresponding counters.

4. **Completion Check**:
   - After updating, check if the row or column has reached its maximum size. If yes, return the current index.

5. **Handle Remaining Cases**:
   - If no row or column is painted fully after all iterations, return `-1`.

---

### **Go Code**

1. **Build a Position Map**:
   - Use a `map` data structure to store the position of each value in the matrix for efficient lookups.

2. **Initialize Row and Column Counters**:
   - Create two slices, `rowCount` and `colCount`, to track the progress of painting.

3. **Process the Array**:
   - For each value in `arr`, retrieve its position using the mapping and increment the respective row and column counters.

4. **Check Painting Status**:
   - After updating the counters, check if the row or column is fully painted. If so, return the index.

5. **Final Result**:
   - If no row or column becomes completely painted, return `-1`.

---

## Notes for All Implementations

- **Mapping Values**:
  - The matrix mapping significantly reduces the computational complexity by avoiding repetitive searches.

- **Optimized Updates**:
  - Instead of repeatedly scanning rows or columns, the approach focuses on maintaining a simple count for each row and column.

- **Edge Cases**:
  - Handle cases where `arr` does not result in any fully painted row or column.
