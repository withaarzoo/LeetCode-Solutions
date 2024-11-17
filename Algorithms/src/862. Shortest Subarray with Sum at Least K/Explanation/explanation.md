# Shortest Subarray with Sum at Least `K` - Step-by-Step Explanation  

This README provides a **step-by-step explanation** of solving the problem of finding the shortest subarray with a sum of at least `k` in **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Each step corresponds to a key part of the algorithm, with an explanation for its role in the solution.

---

## **C++ Code Explanation**  

### **Step 1: Initialize Variables**

- Create a `prefix` array to store the cumulative sums of the input array.
- Use a `deque` to keep track of indices in a monotonic order.
- Set `minLength` to a very large value (`INT_MAX`) to store the result.

### **Step 2: Compute Prefix Sums**

- Iterate through the input array to compute the cumulative sum of elements up to each index. Store these sums in the `prefix` array.

### **Step 3: Check for Valid Subarray**

- Traverse the `prefix` array:
  - Check if subtracting the smallest prefix sum from the current prefix sum yields a value greater than or equal to `k`.
  - If so, calculate the length of this subarray and update `minLength`.

### **Step 4: Maintain Monotonicity**

- Remove indices from the back of the deque if the current prefix sum is smaller than or equal to the last element in the deque.
- This ensures that the deque contains indices in increasing order of prefix sums.

### **Step 5: Return the Result**

- If no valid subarray is found, return `-1`.
- Otherwise, return `minLength`.

---

## **Java Code Explanation**  

### **Step 1: Initialize Variables**

- Use a `long[]` array called `prefix` to store cumulative sums.
- Initialize a `Deque` to maintain indices in a monotonic order.
- Set `minLength` to `Integer.MAX_VALUE` to store the result.

### **Step 2: Compute Prefix Sums**

- Loop through the input array and calculate prefix sums. Each prefix sum represents the sum of elements up to that index.

### **Step 3: Validate Subarrays**

- For each index in the `prefix` array:
  - Check if the difference between the current prefix sum and the smallest prefix sum in the deque is greater than or equal to `k`.
  - If true, update the result with the length of the subarray and remove the front of the deque.

### **Step 4: Update the Deque**

- Remove elements from the back of the deque if the current prefix sum is smaller or equal to the last prefix sum in the deque. This maintains monotonicity.

### **Step 5: Return the Final Result**

- Return `-1` if no valid subarray is found; otherwise, return `minLength`.

---

## **JavaScript Code Explanation**  

### **Step 1: Setup Variables**

- Create an array `prefix` to store cumulative sums.  
- Use an empty array as a deque to maintain indices in increasing order of prefix sums.
- Set `minLength` to `Infinity` as the initial result.

### **Step 2: Calculate Prefix Sums**

- Iterate through the input array, updating the `prefix` array to reflect the cumulative sum up to each index.

### **Step 3: Evaluate Subarrays**

- For each index in the `prefix` array:
  - Check if the difference between the current prefix sum and the smallest prefix sum in the deque is at least `k`.
  - If so, update `minLength` and remove the front element from the deque.

### **Step 4: Maintain Monotonic Order**

- Remove elements from the back of the deque if the current prefix sum is smaller than or equal to the last element in the deque. This ensures that the deque maintains increasing order.

### **Step 5: Return the Shortest Subarray**

- If no valid subarray is found, return `-1`. Otherwise, return `minLength`.

---

## **Python Code Explanation**  

### **Step 1: Initialize Structures**

- Create a `prefix` array to hold cumulative sums of the input array.
- Use a `deque` to maintain indices in increasing order of prefix sums.
- Initialize `min_length` with infinity (`float('inf')`).

### **Step 2: Compute Prefix Sums**

- Compute prefix sums using a loop. Each element in the `prefix` array represents the sum of elements up to that index.

### **Step 3: Find Valid Subarrays**

- Traverse the `prefix` array:
  - Check if subtracting the smallest prefix sum (at the front of the deque) from the current prefix sum yields a value greater than or equal to `k`.
  - If true, calculate the length of this subarray and update `min_length`.

### **Step 4: Update the Deque**

- Remove indices from the back of the deque if the current prefix sum is smaller or equal to the last element in the deque. This ensures monotonicity.

### **Step 5: Return the Result**

- Return `-1` if no valid subarray is found. Otherwise, return `min_length`.

---

## **Go Code Explanation**  

### **Step 1: Initialize Variables**

- Create a `prefix` array of type `int64` to store cumulative sums.
- Use a slice as a deque to maintain indices in a monotonic order.
- Set `minLength` to a large value (e.g., `n+1`) to store the shortest subarray length.

### **Step 2: Compute Prefix Sums**

- Iterate through the input array to compute the prefix sums. Store these sums in the `prefix` array.

### **Step 3: Check Valid Subarrays**

- Traverse the `prefix` array:
  - Check if the difference between the current prefix sum and the smallest prefix sum in the deque is at least `k`.
  - If true, calculate the length of this subarray and update `minLength`.

### **Step 4: Maintain Monotonic Order**

- Remove indices from the back of the deque if the current prefix sum is smaller or equal to the last prefix sum in the deque. This ensures that the deque only contains useful indices.

### **Step 5: Return the Result**

- If no valid subarray is found, return `-1`.
- Otherwise, return `minLength`.
