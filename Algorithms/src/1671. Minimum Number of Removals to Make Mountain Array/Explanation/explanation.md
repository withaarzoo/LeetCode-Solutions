# Minimum Number of Removals to Make Mountain Array

This README explains the step-by-step approach to solving the "Minimum Number of Removals to Make Mountain Array" problem. We'll walk through each language solution—C++, Java, JavaScript, Python, and Go—and detail the logic and methodology for creating an optimized solution.

---

## Problem Description

Given an integer array, the task is to transform it into a "mountain array" by removing the fewest elements possible. A mountain array has a sequence that strictly increases to a peak and then strictly decreases after the peak.

## Approach

### Step 1: Understanding Mountain Array Requirements

- A valid mountain array has:
  1. **An increasing sequence** before a peak element.
  2. **A decreasing sequence** after the peak.
- For any given element in the array to serve as a "peak," it should have:
  - At least one element on its left in an increasing order.
  - At least one element on its right in a decreasing order.

### Step 2: Define LIS and LDS Arrays

- **LIS** (Longest Increasing Subsequence) array:
  - For each element, calculate the length of the longest increasing subsequence ending at that element.
- **LDS** (Longest Decreasing Subsequence) array:
  - For each element, calculate the length of the longest decreasing subsequence starting at that element.

### Step 3: Calculate LIS for Each Element

- **For each element** in the array:
  - Iterate through all previous elements.
  - If the current element is greater than a previous element, update its LIS value to `LIS[i] = max(LIS[i], LIS[j] + 1)`.

### Step 4: Calculate LDS for Each Element

- **For each element**, starting from the end of the array:
  - Iterate through all subsequent elements.
  - If the current element is greater than the next element, update its LDS value to `LDS[i] = max(LDS[i], LDS[j] + 1)`.

### Step 5: Find the Maximum Length of a Mountain Array

- **For each element**:
  - Check if it can serve as a valid peak by ensuring both `LIS[i] > 1` and `LDS[i] > 1`.
  - For valid peaks, calculate the mountain length as `LIS[i] + LDS[i] - 1`.
  - Track the maximum mountain length found.

### Step 6: Calculate Minimum Removals

- The minimum number of removals required to form a mountain array is `n - maxMountainLength`, where `n` is the size of the input array.

---

## Step-by-Step Code Explanation

### C++ Code

1. **Initialize LIS and LDS arrays** with default values of 1 for each element.
2. **Calculate the LIS values** for each element by comparing with all previous elements.
3. **Calculate the LDS values** for each element by comparing with all following elements.
4. **Identify valid peaks** where both LIS and LDS values are greater than 1.
5. **Calculate the maximum mountain length** by adding the LIS and LDS values for each valid peak.
6. **Return the minimum removals** needed by subtracting the maximum mountain length from the total array size.

---

### Java Code

1. **Initialize LIS and LDS arrays** with default values of 1 for each element.
2. **Calculate the LIS for each element** by iterating through all previous elements to find the longest increasing subsequence.
3. **Calculate the LDS for each element** by iterating through all elements following it to find the longest decreasing subsequence.
4. **Identify valid peaks** based on the criteria that both LIS and LDS are greater than 1.
5. **Find the maximum mountain length** for valid peaks.
6. **Return the minimum removals required** by calculating the difference between the array size and the maximum mountain length.

---

### JavaScript Code

1. **Initialize LIS and LDS arrays** with 1 as the default value for each index.
2. **Compute the LIS array** for each index by comparing it with all previous elements.
3. **Compute the LDS array** for each index by comparing it with all elements that follow.
4. **Identify valid peaks** where `LIS > 1` and `LDS > 1`.
5. **Determine the maximum mountain length** by summing LIS and LDS for each valid peak.
6. **Return the minimum number of elements to remove** by calculating the difference between the array length and the maximum mountain length.

---

### Python Code

1. **Initialize LIS and LDS arrays** with 1 as default for each index.
2. **Calculate the LIS**:
   - For each element, iterate over all prior elements.
   - Update the LIS value to maintain the longest increasing subsequence up to that element.
3. **Calculate the LDS**:
   - For each element, iterate over all elements after it.
   - Update the LDS value for the longest decreasing subsequence starting from that element.
4. **Identify valid peaks** by ensuring both `LIS[i] > 1` and `LDS[i] > 1`.
5. **Calculate maximum mountain length** by summing `LIS` and `LDS` values for valid peaks and tracking the maximum.
6. **Return the result** by subtracting the maximum mountain length from the total array length.

---

### Go Code

1. **Initialize LIS and LDS arrays** with values of 1 for each element in the array.
2. **Compute LIS for each element**:
   - Iterate through all previous elements.
   - Update LIS values to hold the longest increasing subsequence length up to each element.
3. **Compute LDS for each element**:
   - Iterate from the end of the array backward.
   - Update LDS values for the longest decreasing subsequence length from each element.
4. **Identify valid peaks** where LIS and LDS are both greater than 1.
5. **Find the maximum mountain length** by summing the LIS and LDS values for each valid peak.
6. **Calculate and return the minimum removals** as the difference between the array length and the maximum mountain length.
