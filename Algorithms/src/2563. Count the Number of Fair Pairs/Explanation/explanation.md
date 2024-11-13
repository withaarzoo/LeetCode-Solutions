# Count the Number of Fair Pairs

This README provides a step-by-step breakdown of the solution to the problem of counting fair pairs in an array. We'll cover the approach for each language: C++, Java, JavaScript, Python, and Go.

## Problem Summary

Given a `0-indexed` integer array `nums` and two integers `lower` and `upper`, we need to count the number of pairs `(i, j)` such that:

- \( 0 \leq i < j < n \) (i.e., \( i \) should be less than \( j \))
- \( \text{lower} \leq \text{nums}[i] + \text{nums}[j] \leq \text{upper} \) (i.e., the sum of the elements should lie between `lower` and `upper`)

## Approach Overview

To solve this problem efficiently:

1. **Sort the Array**: We start by sorting the array `nums` to simplify the process of finding pairs.
2. **Iterate and Search**: For each element `nums[i]`, calculate the range of values that `nums[j]` should lie within to meet the conditions. Use binary search to quickly find the count of valid `j` indices for each `i`.
3. **Binary Search Utility Functions**: Implement functions for binary search that will find the lower and upper bounds of the range that `nums[j]` can fall into for a valid pair.

### Complexity Analysis

- **Time Complexity**: \(O(n \log n)\) due to sorting and binary search for each element.
- **Space Complexity**: \(O(1)\) (or \(O(n)\) if the sorted array copy is counted as extra space).

---

## C++ Solution Explanation

### Step 1: Sort the Array

- Start by sorting `nums`. This enables us to efficiently find the range of valid pairs for each element.

### Step 2: Iterate Over `nums`

- For each element at index `i`, calculate the minimum and maximum values (`minVal` and `maxVal`) that `nums[j]` (where \( j > i \)) must satisfy.
  
### Step 3: Find Lower and Upper Bounds Using Binary Search

- Use `lower_bound` to find the smallest index `j` where `nums[j] >= minVal`.
- Use `upper_bound` to find the smallest index `j` where `nums[j] > maxVal`.
  
### Step 4: Count Valid Pairs

- For each `i`, add the number of valid pairs `(i, j)` by subtracting the lower bound index from the upper bound index.

---

## Java Solution Explanation

### Step 1: Sort the Array

- Start by sorting `nums` to make the search for valid pairs faster.

### Step 2: Define Binary Search Functions

- Implement custom `lowerBound` and `upperBound` methods.
  - `lowerBound` finds the first index `j` where `nums[j] >= minVal`.
  - `upperBound` finds the first index `j` where `nums[j] > maxVal`.

### Step 3: Iterate Over `nums`

- For each index `i`, calculate `minVal` and `maxVal` based on the range `[lower, upper]`.

### Step 4: Use Binary Search to Count Valid Pairs

- For each `i`, find the range of valid indices `j` using the custom binary search functions.
- Add the difference between `upper bound` and `lower bound` indices to the total count.

---

## JavaScript Solution Explanation

### Step 1: Sort the Array

- Sort `nums` to streamline finding valid pairs later.

### Step 2: Define Binary Search Helper Functions

- Create `lowerBound` and `upperBound` helper functions.
  - `lowerBound` finds the first index `j` where `nums[j] >= minVal`.
  - `upperBound` finds the first index `j` where `nums[j] > maxVal`.

### Step 3: Iterate Over the Array

- For each index `i`, determine `minVal` and `maxVal` for the current element as the required bounds for `nums[j]`.

### Step 4: Count Valid Pairs with Binary Search

- For each `i`, use the helper functions to find the valid range of indices for `j`.
- Increment the count based on the difference between the indices found by the helper functions.

---

## Python Solution Explanation

### Step 1: Sort the Array

- Begin by sorting `nums` to enable efficient range searching.

### Step 2: Define Binary Search Functions

- Use custom functions `lower_bound` and `upper_bound`.
  - `lower_bound` finds the first position where `nums[j] >= minVal`.
  - `upper_bound` finds the first position where `nums[j] > maxVal`.

### Step 3: Iterate and Calculate the Range

- For each element `nums[i]`, calculate `minVal` and `maxVal` for valid pairs.

### Step 4: Count Pairs in the Range

- For each `i`, determine the valid range for `j` indices and count the pairs by subtracting the lower bound index from the upper bound index.

---

## Go Solution Explanation

### Step 1: Sort the Array

- Start by sorting `nums` to make it easier to locate valid pairs.

### Step 2: Define Binary Search Functions

- Implement `lowerBound` and `upperBound` functions.
  - `lowerBound` finds the index where `nums[j] >= minVal`.
  - `upperBound` finds the index where `nums[j] > maxVal`.

### Step 3: Iterate and Set Bounds

- For each element in `nums`, calculate `minVal` and `maxVal` based on the `lower` and `upper` constraints.

### Step 4: Count Valid Pairs Using Range

- For each `i`, use `lowerBound` and `upperBound` to get the range of valid `j` indices.
- Add the count of valid pairs to the total based on the difference in indices.

---

## Summary

The main strategy is to:

1. Sort the array.
2. For each index `i`, find the valid range for `j` using binary search functions.
3. Count the number of pairs in this range.

Each language implements the same approach, with minor variations due to syntax and standard library differences.

---

This structured approach ensures that the solution is both efficient and easy to understand across different programming languages. Each step in the implementation is designed to make the most of binary search and sorting for an optimal solution.
