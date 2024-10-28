# Longest Square Streak Solution Explanation

This repository provides solutions to the "Longest Square Streak" problem implemented in **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Below is a step-by-step explanation of the approach and logic for each language solution.

---

## Problem Description

Given an array of integers, the task is to find the longest sequence in which each element is the square of the previous one. For example, in the array `[2, 4, 16]`, the sequence `2 -> 4 -> 16` has a length of 3. The solution should return the maximum length of such a streak, or `-1` if no streak of at least length 2 exists.

---

## Solution Approach (General)

1. **Sort the Array**: Sorting helps to evaluate each number starting from the smallest, which simplifies streak building.
2. **Use a Set for Quick Lookup**: Convert the array to a set, allowing constant-time lookups to determine if a squared number exists in the array.
3. **Build Streaks for Each Number**:
   - For each number, attempt to build a streak by repeatedly squaring the current number and checking if it exists in the set.
   - Stop the streak if the squared value exceeds the maximum limit (to avoid overflow).
4. **Track Longest Streak**: For each streak that has at least 2 elements, compare its length to the maximum streak length found so far and update if necessary.
5. **Return Result**: The result is the longest streak length found, or `-1` if no streak of length 2 or more was found.

---

### C++ Solution Explanation

1. **Sort the Input Array**: The array is sorted so that we start from the smallest number, making it easier to build streaks in ascending order.
2. **Initialize a Set**: Create a set from the array to quickly check if a squared number exists.
3. **Loop Through Each Number**:
   - For each number, initialize `length` to zero.
   - While the current number exists in the set:
     - Increase the length by one.
     - Square the current number to attempt extending the streak.
     - If the squared value exceeds a threshold (like \(10^9\)), stop to prevent overflow.
4. **Update Maximum Length**: If the streak length is at least 2, update `maxLength` if this streak is the longest found so far.
5. **Return Result**: After all numbers are evaluated, return the longest streak length.

---

### Java Solution Explanation

1. **Sort the Array**: Sorting helps to evaluate the numbers from smallest to largest, building streaks in ascending order.
2. **Convert to Set**: Use a HashSet to store the array elements, allowing for quick existence checks.
3. **Evaluate Each Number for Streak**:
   - For each number, initialize `length` to zero.
   - Check if the current number exists in the set:
     - Increment `length` if it does.
     - Square the number to extend the streak.
     - Stop if squaring the number results in a value beyond `Integer.MAX_VALUE` to prevent overflow.
4. **Track Longest Streak**: Update `maxLength` if the streak length meets the minimum requirement (2) and is the longest found so far.
5. **Return the Maximum Streak Length**: Return `maxLength` after all elements are processed.

---

### JavaScript Solution Explanation

1. **Sort the Array**: Sort the array to start building streaks from the smallest element.
2. **Create a Set for Quick Lookup**: Use a Set to store the array elements, ensuring constant-time lookups for each number.
3. **Loop Through Each Element**:
   - For each element, initialize `length` to zero.
   - Check if the current element exists in the set:
     - Increase the `length` counter.
     - Square the current value to attempt extending the streak.
     - Stop if the squared value exceeds a set threshold (e.g., \(10^9\)) to prevent large values.
4. **Update the Maximum Length**: Update `maxLength` only if the streak length is at least 2 and is the longest found.
5. **Return the Result**: After processing all numbers, return the longest streak length.

---

### Python Solution Explanation

1. **Sort the Array**: Sorting is done to start streaks from the smallest element and build them in ascending order.
2. **Use a Set for Quick Lookup**: Create a set from the array, allowing constant-time checks for existence of each squared number.
3. **Check Each Number for Streak**:
   - Initialize `length` to zero for each number.
   - While the current value exists in the set:
     - Increase the `length` counter.
     - Square the current number to continue the streak.
     - Break out of the loop if squaring results in a number larger than a set limit (like \(10^9\)) to prevent excessive values.
4. **Track the Maximum Streak Length**: Update `maxLength` if this streak is the longest so far and meets the minimum length requirement (2).
5. **Return the Longest Streak**: After evaluating all numbers, return the longest streak found.

---

### Go Solution Explanation

1. **Sort the Input Array**: Sorting allows the streaks to start from the smallest number, which is more efficient.
2. **Convert to Map for Fast Lookup**: Use a map to store the array values, enabling quick existence checks for squared values.
3. **Loop Through Each Element**:
   - For each element, initialize `length` to zero.
   - Check if the current number exists in the map:
     - Increment the streak `length`.
     - Square the current number and continue the streak.
     - Exit the loop if squaring leads to an excessively large number.
4. **Update Maximum Streak Length**: If a streak length is at least 2, check if it is the longest streak found so far, and update `maxLength` accordingly.
5. **Return Final Result**: Return `maxLength` after all numbers have been evaluated.
