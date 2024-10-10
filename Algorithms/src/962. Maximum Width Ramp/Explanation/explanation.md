# Maximum Width Ramp - Step-by-Step Explanation

## Problem Overview

The task is to find the maximum width of a ramp in an array `nums`. A ramp is defined as a pair `(i, j)` where `i < j` and `nums[i] <= nums[j]`. The width of the ramp is `j - i`.

Below are step-by-step explanations for the solution implemented in multiple languages.

---

## Step-by-Step Solution (C++ Code Explanation)

1. **Initialization**:
   - Store the length of the array `nums` in `n`.
   - Create an empty stack `s` to keep track of indices in a decreasing order based on `nums` values.

2. **Building the Decreasing Stack**:
   - Traverse through the array `nums` from the start.
   - For each element, push its index into the stack if either the stack is empty or the current element is smaller than the top element of the stack. This ensures that the stack contains indices of elements in decreasing order.

3. **Finding the Maximum Width Ramp**:
   - Initialize `maxWidth` to 0.
   - Traverse the array from the end to the beginning. For each element, pop from the stack as long as the current element is greater than or equal to the element at the index stored in the top of the stack.
   - For each valid pop, update the `maxWidth` by calculating the difference between the current index and the popped index.

4. **Return Result**:
   - After the second traversal, return the `maxWidth` as the result.

---

## Step-by-Step Solution (Java Code Explanation)

1. **Initialization**:
   - Store the length of the array `nums` in `n`.
   - Create an empty `Stack` to store the indices of elements in decreasing order.

2. **Building the Decreasing Stack**:
   - Loop through the array from the start.
   - Push the index onto the stack if the stack is empty or the current element is smaller than the element at the index at the top of the stack.

3. **Finding the Maximum Width Ramp**:
   - Initialize `maxWidth` as 0.
   - Traverse the array from the last index to the first.
   - For each element, pop from the stack while the current element is greater than or equal to the element at the index stored in the stack.
   - Calculate and update the `maxWidth` for every valid pop operation.

4. **Return the Result**:
   - Return the `maxWidth` after the entire process.

---

## Step-by-Step Solution (JavaScript Code Explanation)

1. **Initialization**:
   - Store the length of the `nums` array in a variable `n`.
   - Create an empty `stack` to store the indices in decreasing order based on the values in `nums`.

2. **Building the Decreasing Stack**:
   - Iterate through `nums` from the start.
   - Push the current index onto the stack if the stack is empty or the current element is smaller than the element at the index on the top of the stack.

3. **Finding the Maximum Width Ramp**:
   - Initialize `maxWidth` to 0.
   - Traverse `nums` from the last index to the first.
   - For each element, while the stack is not empty and the current element is greater than or equal to the element at the index stored in the stack, pop the stack and update `maxWidth`.

4. **Return Result**:
   - Return the final `maxWidth`.

---

## Step-by-Step Solution (Python Code Explanation)

1. **Initialization**:
   - Store the length of `nums` in `n`.
   - Create an empty `stack` to store indices in decreasing order based on the values in `nums`.

2. **Building the Decreasing Stack**:
   - Iterate through the array `nums` from left to right.
   - For each element, append its index to the stack if the stack is empty or the element is smaller than the element at the index stored at the top of the stack.

3. **Finding the Maximum Width Ramp**:
   - Initialize `maxWidth` to 0.
   - Traverse the array `nums` from right to left.
   - For each element, pop from the stack while the current element is greater than or equal to the element at the index stored at the top of the stack.
   - For each valid pop, update the `maxWidth`.

4. **Return the Result**:
   - Return the computed `maxWidth`.

---

## Step-by-Step Solution (Go Code Explanation)

1. **Initialization**:
   - Store the length of the `nums` array in `n`.
   - Create an empty `stack` to store indices in decreasing order based on the values in `nums`.

2. **Building the Decreasing Stack**:
   - Traverse through the array from the start.
   - Append the index of the element to the stack if it is empty or the current element is smaller than the element at the top of the stack.

3. **Finding the Maximum Width Ramp**:
   - Initialize `maxWidth` to 0.
   - Traverse through the array from the end.
   - While the stack is not empty and the current element is greater than or equal to the element at the top index in the stack, pop the index from the stack and update `maxWidth`.

4. **Return Result**:
   - Return the final `maxWidth`.

---

### Conclusion

This problem can be solved using a **monotonic stack** approach, which ensures efficient calculation of the maximum width ramp in linear time. The approach remains consistent across all languages by maintaining a stack of decreasing indices and then checking the largest valid ramp widths from the end of the array.
