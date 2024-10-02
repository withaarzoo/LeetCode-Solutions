# Array Rank Transform - Step by Step Explanation

This guide will provide step-by-step explanations for solving the "Array Rank Transform" problem in multiple languages: C++, Java, JavaScript, Python, and Go. The solution involves transforming the elements of an array into their ranks based on their sorted order. We will break down each step for all implementations.

## Approach Overview

1. **Check for Empty Array**: If the array is empty, return an empty array.
2. **Sort the Array**: Create a copy of the array and sort it in ascending order.
3. **Assign Ranks**: Create a mapping of each unique element to its rank.
4. **Transform the Original Array**: Replace each element in the original array with its rank based on the sorted array.

---

## C++ Code Explanation

1. **Check for Empty Array**:  
   - If the input array is empty, return an empty vector.

2. **Sort the Array**:  
   - Copy the original array into a new vector.
   - Sort the copied array to arrange the elements in ascending order.

3. **Assign Ranks**:  
   - Use an unordered map to store the rank of each unique element in the sorted array.
   - Loop through the sorted array and assign ranks, ensuring each unique element receives a rank only once.

4. **Transform the Original Array**:  
   - Iterate through the original array and replace each element with its corresponding rank from the rank map.

---

## Java Code Explanation

1. **Check for Empty Array**:  
   - If the input array length is 0, return a new empty array.

2. **Sort the Array**:  
   - Clone the original array.
   - Sort the cloned array to get the elements in ascending order.

3. **Assign Ranks**:  
   - Use a HashMap to map each unique element to its rank.
   - Iterate through the sorted array and assign ranks only to elements that have not been assigned a rank yet.

4. **Transform the Original Array**:  
   - Iterate through the original array and replace each element with its rank based on the rank map.

---

## JavaScript Code Explanation

1. **Check for Empty Array**:  
   - If the array length is 0, return an empty array.

2. **Sort the Array**:  
   - Create a copy of the original array using the spread operator.
   - Sort the copied array in ascending order using a custom comparator for numerical sorting.

3. **Assign Ranks**:  
   - Use a Map to store the rank of each unique element.
   - Loop through the sorted array and assign ranks only to elements that haven't been added to the map yet.

4. **Transform the Original Array**:  
   - Loop through the original array and replace each element with its rank from the map.

---

## Python Code Explanation

1. **Check for Empty Array**:  
   - If the array is empty, return an empty list.

2. **Sort the Array**:  
   - Sort the original array and store it in a new list.

3. **Assign Ranks**:  
   - Use a dictionary to store the rank of each unique element.
   - Iterate over the sorted array and assign ranks only to elements that are not already in the dictionary.

4. **Transform the Original Array**:  
   - Use a list comprehension to replace each element in the original array with its corresponding rank from the dictionary.

---

## Go Code Explanation

1. **Check for Empty Array**:  
   - If the array is empty, return an empty slice.

2. **Sort the Array**:  
   - Create a copy of the original array using the `copy` function.
   - Sort the copied array using the `sort.Ints` function.

3. **Assign Ranks**:  
   - Use a map to store the rank of each unique element.
   - Loop through the sorted array and assign ranks to elements that haven't been added to the map yet.

4. **Transform the Original Array**:  
   - Iterate over the original array and replace each element with its rank using the map.

---

## Conclusion

In each implementation, the process follows a similar pattern:

1. **Check for an empty array**.
2. **Sort a copy** of the original array.
3. **Map each unique element** to a rank.
4. **Transform the original array** by replacing each element with its rank.

This approach ensures that the elements are ranked based on their sorted order, and the time complexity is dominated by the sorting operation, making it **O(n log n)** in all implementations.
