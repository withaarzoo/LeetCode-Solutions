# Smallest Distance Pair - Step-by-Step Explanation

This section provides a detailed explanation of how to find the k-th smallest distance pair in a given array using C++, Java, JavaScript, Python, and Go. The problem involves using a combination of sorting, two-pointer technique, and binary search to achieve the desired result.

---

## C++ Implementation

1. **Sorting the Array:**
   - The first step involves sorting the array to facilitate the two-pointer technique for finding pairs with the smallest distances.

2. **Binary Search Setup:**
   - The search range for the smallest distance is initialized, with `low` as the smallest possible distance (0) and `high` as the largest possible distance (difference between the maximum and minimum elements in the array).

3. **Counting Valid Pairs:**
   - For a given midpoint (`mid`) in the binary search, count how many pairs in the array have a distance less than or equal to `mid`.
   - This involves iterating through the array and using a second pointer to find valid pairs that meet the condition.

4. **Binary Search Execution:**
   - Adjust the search range based on the count of valid pairs found.
   - If the count is greater than or equal to `k`, search the lower half; otherwise, search the upper half.

5. **Result:**
   - After binary search completes, the `low` value will represent the k-th smallest distance.

---

## Java Implementation

1. **Sorting the Array:**
   - The array is sorted to allow efficient pair counting with the two-pointer technique.

2. **Binary Search Initialization:**
   - Set up the binary search range with `low` as 0 and `high` as the difference between the maximum and minimum elements in the array.

3. **Counting Pairs with Two Pointers:**
   - For each midpoint (`mid`), count the number of pairs with a difference less than or equal to `mid`.
   - Use two pointers to iterate through the array, counting valid pairs.

4. **Binary Search Logic:**
   - Adjust the binary search range based on whether the count of valid pairs is greater than or equal to `k`.

5. **Final Result:**
   - The value of `low` after binary search will be the k-th smallest distance.

---

## JavaScript Implementation

1. **Sorting the Array:**
   - Begin by sorting the array to facilitate pair counting with two pointers.

2. **Binary Search Preparation:**
   - Initialize the search range with `low` as the smallest possible distance (0) and `high` as the maximum possible distance.

3. **Counting Pairs:**
   - For a given midpoint (`mid`), count the number of pairs with a difference less than or equal to `mid` using two pointers.

4. **Binary Search Execution:**
   - Based on the count of valid pairs, adjust the search range: search the lower half if the count is greater than or equal to `k`, otherwise search the upper half.

5. **Final Value:**
   - After completing the binary search, `low` will hold the k-th smallest distance.

---

## Python Implementation

1. **Sorting the Array:**
   - The array is sorted to enable efficient pair counting.

2. **Binary Search Setup:**
   - Set up the search range with `low` as 0 and `high` as the difference between the maximum and minimum elements in the array.

3. **Counting Valid Pairs:**
   - For each midpoint (`mid`), use a two-pointer technique to count pairs with a difference less than or equal to `mid`.

4. **Binary Search Progression:**
   - Adjust the search range based on the count of valid pairs. If the count is greater than or equal to `k`, search the lower half; otherwise, search the upper half.

5. **Conclusion:**
   - The value of `low` after binary search will be the k-th smallest distance.

---

## Go Implementation

1. **Array Sorting:**
   - Sort the array to prepare for pair counting using two pointers.

2. **Binary Search Initialization:**
   - Set up the binary search range, with `low` as 0 and `high` as the difference between the maximum and minimum elements in the array.

3. **Counting Pairs:**
   - For each midpoint (`mid`), count the number of pairs with a distance less than or equal to `mid`.

4. **Binary Search Execution:**
   - Adjust the binary search range based on the count of valid pairs. If the count is sufficient, search the lower half; otherwise, search the upper half.

5. **Final Output:**
   - After binary search, `low` will contain the k-th smallest distance.
