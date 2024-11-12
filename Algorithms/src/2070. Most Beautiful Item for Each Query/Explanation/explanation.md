# LeetCode Problem 2070: Most Beautiful Item for Each Query

This repository provides solutions in **C++**, **Java**, **JavaScript**, **Python**, and **Go** to the LeetCode problem "Most Beautiful Item for Each Query."

---

## Problem Overview
Given a list of items, each with a price and beauty value, we need to determine the maximum beauty achievable for each query based on the price limit set by that query. If no item meets the price requirement, the answer should be 0.

---

## Solution Explanation

### Step 1: Understanding and Sorting the Items
1. First, we **sort the items by their price** in ascending order. This makes it easier to evaluate items for each price threshold defined by the queries.
2. As we sort, we also **keep track of the maximum beauty** value seen so far for items up to each price point. This allows us to quickly access the maximum beauty within any given price limit.

### Step 2: Optimizing for Maximum Beauty
1. Using the sorted items, we create a **cumulative maximum beauty array**. For each item, if its beauty is higher than the previous maximum, we update it. This array will store the maximum beauty achievable at or below each price point.

### Step 3: Handling Queries Efficiently
1. **Sort the queries** in ascending order while keeping track of their original indexes. Sorting helps in using a two-pointer technique to evaluate each query efficiently.
2. For each query:
   - Use a **binary search** (or two-pointer) technique to find the highest price in the items that is still within the query's limit.
   - Retrieve the corresponding maximum beauty value from the cumulative array.

### Step 4: Mapping Results Back to Original Query Order
1. After processing each query in sorted order, map the results back to the original query order to get the correct answers for each query.

---

## Code Explanation (Language-Specific Walkthroughs)

### C++ Code

- **Step 1**: Define a `vector` of pairs for items where each item has a price and beauty.
- **Step 2**: Sort the items by price in ascending order.
- **Step 3**: Build a cumulative array for maximum beauty values by iterating through the sorted items.
- **Step 4**: Sort the queries along with their original indexes to ensure the correct final ordering.
- **Step 5**: Use binary search on the cumulative beauty array to find the maximum beauty for each query price limit.
- **Step 6**: Store results and remap them to the original query order.

### Java Code

- **Step 1**: Use a list of integer arrays to store items, each with price and beauty.
- **Step 2**: Sort the items by price. Use `Arrays.sort` with a custom comparator for this purpose.
- **Step 3**: Create a cumulative array to store the maximum beauty at each price point.
- **Step 4**: Sort the queries with their original indexes to retain the initial query order after processing.
- **Step 5**: Use binary search (`Arrays.binarySearch` or custom implementation) to find the maximum beauty for each query’s price.
- **Step 6**: Store and map results back to the original order.

### JavaScript Code

- **Step 1**: Represent items as arrays, each containing price and beauty values.
- **Step 2**: Sort the items array based on price using `sort`.
- **Step 3**: Create a cumulative beauty array to store the highest beauty achievable at each price.
- **Step 4**: Sort queries along with their original indices so we can map results back to their initial order.
- **Step 5**: Use binary search (e.g., using a helper function) to find the largest possible beauty for each query.
- **Step 6**: Construct an answer array, placing each result back in the original query order.

### Python Code

- **Step 1**: Use a list of lists to represent items where each sublist has price and beauty.
- **Step 2**: Sort items by price using the `sorted` function.
- **Step 3**: Build a cumulative maximum beauty array for quick access to the max beauty at any price point.
- **Step 4**: Sort the queries with their original indices so that after sorting, results can be mapped back correctly.
- **Step 5**: Use `bisect_left` (from `bisect` module) for binary search to find the highest beauty value that meets the price limit of each query.
- **Step 6**: Create the answer array and remap the results back to the original query order.

### Go Code

- **Step 1**: Create a slice of structs to represent items, each with a price and beauty.
- **Step 2**: Sort the items by price using `sort.Slice`.
- **Step 3**: Construct a cumulative maximum beauty slice to record max beauty up to each price.
- **Step 4**: Sort the queries along with their original indices, allowing remapping to original order after processing.
- **Step 5**: Use binary search (`sort.Search` function) to find the highest beauty achievable for each query.
- **Step 6**: Place results in an answer slice and reorder based on original query indices.

---

## Complexity Analysis

- **Time Complexity**: Sorting items and queries each take \(O(n \log n + m \log m)\), where \(n\) is the number of items and \(m\) is the number of queries. For each query, binary search takes \(O(\log n)\), making the overall time complexity \(O(n \log n + m \log n)\).
  
- **Space Complexity**: \(O(n + m)\) for storing cumulative beauty values and the query results.

---

## Conclusion

The solutions in C++, Java, JavaScript, Python, and Go demonstrate efficient handling of item sorting, cumulative maximum tracking, and binary search for query processing. Each approach ensures that we get the maximum beauty for each query's price limit effectively.