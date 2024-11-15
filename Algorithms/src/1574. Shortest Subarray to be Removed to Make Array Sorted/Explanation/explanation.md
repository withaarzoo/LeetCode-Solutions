# Shortest Subarray to be Removed to Make Array Sorted

## Problem Overview

The goal is to find the minimum number of elements to remove from a given array in order to make the entire array sorted in non-decreasing order. The solution needs to identify and utilize already sorted segments of the array to minimize the number of elements that must be removed.

---

## Approach

### General Steps

1. **Identify the Longest Sorted Prefix**:
   - Start from the beginning of the array and move forward until you find a break in the non-decreasing order. This part of the array is the longest "prefix" that is already sorted.

2. **Identify the Longest Sorted Suffix**:
   - Starting from the end of the array, move backward until you find a break in the non-decreasing order. This part of the array is the longest "suffix" that is already sorted.

3. **Check if Entire Array is Sorted**:
   - If the longest sorted prefix already covers the entire array, no removal is needed, so the result is `0`.

4. **Calculate Initial Minimum Removals**:
   - If the entire array is not sorted, calculate the initial minimum number of elements to remove by either:
     - Removing the entire suffix (from the end of the prefix to the array’s end).
     - Removing the entire prefix (from the start to the beginning of the suffix).

5. **Optimize Removal with Two Pointers**:
   - Use a two-pointer technique to attempt merging parts of the prefix and suffix.
   - Adjust pointers to find if a smaller section in between can be removed to achieve the sorted order, updating the minimum elements to remove if a smaller solution is found.

---

## Solutions in Multiple Languages

### C++ Code - Step-by-Step Explanation

1. **Find the Longest Sorted Prefix**: Initialize a pointer from the start and increment it while the next element is larger than or equal to the current. Stop when the order breaks.

2. **Find the Longest Sorted Suffix**: Initialize a pointer from the end and decrement it while the previous element is smaller than or equal to the current. Stop when the order breaks.

3. **Check Full Array Condition**: If the prefix pointer reaches the end, the array is already sorted. Return `0`.

4. **Calculate Initial Minimum Removals**: Compute the minimum removals by comparing two options:
   - Removing all elements from the end of the prefix to the array's end.
   - Removing all elements from the start to the beginning of the suffix.

5. **Merge with Two-Pointer Technique**: Set two pointers on the prefix and suffix. Adjust them while maintaining sorted order, minimizing the number of elements to remove for a sorted result.

### Java Code - Step-by-Step Explanation

1. **Identify Longest Sorted Prefix**: Use a loop to traverse from the beginning, stopping when the non-decreasing order is broken.

2. **Identify Longest Sorted Suffix**: Traverse from the end towards the beginning, stopping when the order breaks.

3. **Full Array Check**: If the prefix pointer reaches the last index, the array is already sorted, so return `0`.

4. **Calculate Minimum Removals**:
   - Calculate the number of elements to remove if only keeping the prefix.
   - Calculate the number of elements to remove if only keeping the suffix.
   - Set the minimum of these values as the initial result.

5. **Optimize with Two Pointers**: Using pointers on both the prefix and suffix, try merging sections and update the minimum removals if a smaller solution is possible.

### JavaScript Code - Step-by-Step Explanation

1. **Longest Sorted Prefix**: Start from index `0`, increment until reaching a break in the order.

2. **Longest Sorted Suffix**: Start from the last index, decrement until reaching a break in the order.

3. **Check if Sorted**: If the entire array is already sorted (i.e., the prefix reaches the end), return `0`.

4. **Calculate Initial Minimum**: Evaluate the removals required by only keeping the prefix or only keeping the suffix and set the minimum.

5. **Two-Pointer Optimization**: Use two pointers on the prefix and suffix to find a merged solution, minimizing the required removals by adjusting the pointers to maintain order.

### Python Code - Step-by-Step Explanation

1. **Find the Longest Non-Decreasing Prefix**: Initialize a pointer and move it forward while each element is less than or equal to the next.

2. **Find the Longest Non-Decreasing Suffix**: Initialize a pointer at the end and move it backward while each element is greater than or equal to the previous.

3. **Full Array Check**: If the prefix reaches the last element, the array is fully sorted, and the answer is `0`.

4. **Initial Minimum Removals**: Calculate removals by keeping either the prefix or the suffix, whichever results in fewer removals.

5. **Two-Pointer Technique for Optimal Merging**: Attempt to merge the prefix and suffix while keeping them in order. Adjust pointers to find the smallest number of elements that need to be removed.

### Go Code - Step-by-Step Explanation

1. **Identify Longest Sorted Prefix**: Set a pointer to start and increment as long as each element is in non-decreasing order with the next.

2. **Identify Longest Sorted Suffix**: Set a pointer to the end and decrement while each element is in non-decreasing order with the previous.

3. **Entire Array Check**: If the prefix covers the whole array, it’s already sorted, so return `0`.

4. **Compute Initial Minimum Removals**: Calculate the required removals if keeping only the prefix or the suffix. Set the result to the minimum of these two values.

5. **Two-Pointer Optimization for Minimal Removals**: Using two pointers on the prefix and suffix, move them towards each other while keeping the array sorted to find the minimal number of elements to remove.

---

## Complexity Analysis

- **Time Complexity**: \(O(n)\) - Each solution involves linear scans to identify the sorted prefix and suffix and a two-pointer pass, making the solution efficient.
- **Space Complexity**: \(O(1)\) - Only a few variables are used for indices and result calculation, leading to constant extra space usage.

---

Each solution leverages the initial sorted segments of the array (prefix and suffix) to minimize the removal count, making it a highly efficient approach to solve this problem.
