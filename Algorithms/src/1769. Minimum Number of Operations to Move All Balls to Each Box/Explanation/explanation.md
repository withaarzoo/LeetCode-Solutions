# Minimum Operations to Move Balls to Each Box

This GitHub README provides a **step-by-step explanation** of how the solution works for the problem "Minimum Operations to Move Balls to Each Box." The explanation is language-specific (C++, Java, JavaScript, Python, and Go). Each step corresponds to how the logic is applied without explicitly showing the code but by describing the process in detail.

---

## C++ Code Explanation

1. **Initialize Result Array:**  
   Create a result array `answer` initialized with zeros, one for each box.

2. **First Pass (Left-to-Right):**  
   - Use variables `balls` (to count the number of balls encountered so far) and `operations` (to track cumulative operations).  
   - Traverse the string of boxes from left to right.  
   - For each box:
     - Add the cumulative operations (`operations`) to the result array.  
     - Check if the current box has a ball (`boxes[i] == '1'`); if yes, increment the `balls` counter.  
     - Add the current count of `balls` to `operations` for the next box.

3. **Second Pass (Right-to-Left):**  
   - Reset `balls` and `operations` variables to zero.  
   - Traverse the string of boxes from right to left.  
   - For each box:
     - Add the cumulative operations (`operations`) to the result array.  
     - Check if the current box has a ball (`boxes[i] == '1'`); if yes, increment the `balls` counter.  
     - Add the current count of `balls` to `operations` for the next box.

4. **Return Result:**  
   The result array now contains the minimum operations required for each box.  

---

## Java Code Explanation

1. **Initialize Result Array:**  
   Create an integer array `answer` with a size equal to the number of boxes, initialized with zeros.

2. **First Pass (Left-to-Right):**  
   - Define variables `balls` and `operations` for tracking the number of balls and cumulative operations.  
   - Traverse the `boxes` string from left to right.  
   - For each box:
     - Add the cumulative `operations` to the current index in the result array.  
     - If the current box contains a ball (`boxes.charAt(i) == '1'`), increment the `balls` count.  
     - Update `operations` by adding the `balls` count for the next box.

3. **Second Pass (Right-to-Left):**  
   - Reset `balls` and `operations` to zero.  
   - Traverse the string from right to left.  
   - For each box:
     - Add the cumulative `operations` to the current index in the result array.  
     - If the current box contains a ball, increment the `balls` count.  
     - Update `operations` by adding the `balls` count for the next box.

4. **Return Result:**  
   The result array now holds the minimum operations for each box.

---

## JavaScript Code Explanation

1. **Initialize Result Array:**  
   Create an array `answer` of size `n` (length of `boxes` string) filled with zeros.

2. **First Pass (Left-to-Right):**  
   - Use two variables: `balls` to count the number of balls encountered and `operations` for cumulative operations.  
   - Traverse the string of boxes from left to right.  
   - For each box:
     - Update the `answer[i]` by adding the current cumulative `operations`.  
     - If the current box contains a ball (`boxes[i] === '1'`), increment the `balls`.  
     - Add the `balls` count to `operations` for the next box.

3. **Second Pass (Right-to-Left):**  
   - Reset `balls` and `operations` to zero.  
   - Traverse the string from right to left.  
   - For each box:
     - Add the cumulative `operations` to `answer[i]`.  
     - If the current box contains a ball, increment `balls`.  
     - Update `operations` by adding `balls` for the next box.

4. **Return Result:**  
   The `answer` array now contains the minimum operations needed for each box.

---

## Python Code Explanation

1. **Initialize Result Array:**  
   Create a list `answer` of size `n` (length of `boxes` string) initialized to zeros.

2. **First Pass (Left-to-Right):**  
   - Initialize variables `balls` and `operations` to zero.  
   - Traverse the string from left to right.  
   - For each box:
     - Add the cumulative `operations` to the current index in the `answer` list.  
     - If the current box contains a ball (`boxes[i] == '1'`), increment the `balls` counter.  
     - Update `operations` by adding the `balls` count for the next box.

3. **Second Pass (Right-to-Left):**  
   - Reset `balls` and `operations` to zero.  
   - Traverse the string from right to left.  
   - For each box:
     - Add the cumulative `operations` to the current index in the `answer` list.  
     - If the current box contains a ball, increment the `balls` counter.  
     - Update `operations` by adding the `balls` count for the next box.

4. **Return Result:**  
   The `answer` list now contains the minimum operations required for each box.

---

## Go Code Explanation

1. **Initialize Result Array:**  
   Create a slice `answer` of size `n` (length of `boxes` string) initialized with zeros.

2. **First Pass (Left-to-Right):**  
   - Use two variables: `balls` to count the number of balls and `operations` to track cumulative operations.  
   - Traverse the string of boxes from left to right.  
   - For each box:
     - Add the cumulative `operations` to `answer[i]`.  
     - If the current box contains a ball (`boxes[i] == '1'`), increment the `balls` counter.  
     - Update `operations` by adding the `balls` count for the next box.

3. **Second Pass (Right-to-Left):**  
   - Reset `balls` and `operations` to zero.  
   - Traverse the string from right to left.  
   - For each box:
     - Add the cumulative `operations` to `answer[i]`.  
     - If the current box contains a ball, increment the `balls` counter.  
     - Update `operations` by adding the `balls` count for the next box.

4. **Return Result:**  
   The `answer` slice now holds the minimum operations required for each box.

---

## Summary

In all solutions (C++, Java, JavaScript, Python, Go):

- The problem is solved using two linear sweeps: left-to-right and right-to-left.
- The result array is updated incrementally with the number of operations from both passes.
- The time complexity is **O(n)**, and the space complexity is **O(1)** (excluding the result array).

Feel free to explore any of the language-specific implementations!
