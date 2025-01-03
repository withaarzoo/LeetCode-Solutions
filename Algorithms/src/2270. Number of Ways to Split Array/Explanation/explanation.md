# Number of Ways to Split Array - Step-by-Step Explanation

This README provides a detailed explanation of the solution for the problem **"Number of Ways to Split Array"**. Below, we go through the thought process and steps for each implementation in **C++**, **Java**, **JavaScript**, **Python**, and **Go**.

---

## 🚀 Problem Description  

Given an array `nums`, count the number of ways to split the array into two non-empty parts such that the sum of the left part is **greater than or equal to** the sum of the right part.

---

## 🧠 Intuition  

The idea is to calculate the sum of the left and right parts of the array for each split point and check if the left sum is greater than or equal to the right sum. By using **prefix sums**, we can avoid recalculating sums repeatedly and ensure the solution is efficient.

---

## 🛠️ Approach  

1. Calculate the **total sum** of the entire array.  
2. Use a **prefix sum** to keep track of the cumulative sum of elements from the start up to the current index.  
3. At each split point (from index `0` to `n-2`), calculate the right sum as the difference between the total sum and the prefix sum.  
4. Check if the left sum (prefix sum) is greater than or equal to the right sum. If yes, increase the count of valid splits.  
5. Return the final count of valid splits.

---

## 📚 Step-by-Step Explanation  

### **C++ Implementation**  

1. **Calculate the total sum**: Start by iterating through the array to compute the total sum of all elements.  
2. **Initialize variables**: Use a `prefix_sum` to store the cumulative sum of the left part and a `count` to keep track of valid splits.  
3. **Iterate through the array**: For each split point, update the prefix sum with the current element.  
4. **Compute the right sum**: Subtract the prefix sum from the total sum to get the right part's sum.  
5. **Check the condition**: Compare the prefix sum with the right sum. If the prefix sum is greater than or equal to the right sum, increment the count.  
6. **Return the result**: Once the loop finishes, return the count of valid splits.

---

### **Java Implementation**  

1. **Calculate the total sum**: Use a loop to compute the total sum of all elements in the array.  
2. **Initialize variables**: Create variables for `prefixSum` (to track the cumulative left sum) and `count` (to count valid splits).  
3. **Iterate through the array**: Loop from the start of the array to the second-to-last element.  
4. **Update the prefix sum**: Add the current element to the `prefixSum` during each iteration.  
5. **Compute the right sum**: Subtract the `prefixSum` from the `totalSum` to get the right part's sum dynamically.  
6. **Check the condition**: If the `prefixSum` is greater than or equal to the right sum, increment the `count`.  
7. **Return the count**: After the loop, return the total number of valid splits.

---

### **JavaScript Implementation**  

1. **Calculate the total sum**: Use the `reduce()` function to compute the sum of all elements in the array.  
2. **Initialize variables**: Create variables `prefixSum` for the cumulative left sum and `count` to track valid splits.  
3. **Iterate through the array**: Use a `for` loop to traverse the array, stopping at the second-to-last element.  
4. **Update the prefix sum**: Add the current element to `prefixSum` in each iteration.  
5. **Compute the right sum**: Subtract `prefixSum` from `totalSum` to dynamically calculate the sum of the right part.  
6. **Check the condition**: Compare `prefixSum` with the right sum. If `prefixSum` is greater than or equal to the right sum, increment `count`.  
7. **Return the count**: At the end of the loop, return the total count of valid splits.

---

### **Python Implementation**  

1. **Calculate the total sum**: Use Python's `sum()` function to find the sum of all elements in the array.  
2. **Initialize variables**: Create `prefix_sum` for the cumulative left sum and `count` to store the number of valid splits.  
3. **Iterate through the array**: Loop through the array, stopping before the last element.  
4. **Update the prefix sum**: Add the current element to `prefix_sum` during each iteration.  
5. **Compute the right sum**: Subtract the `prefix_sum` from the `total_sum` to get the right part's sum.  
6. **Check the condition**: If the `prefix_sum` is greater than or equal to the right sum, increment the `count`.  
7. **Return the count**: After the loop finishes, return the total count of valid splits.

---

### **Go Implementation**  

1. **Calculate the total sum**: Use a `for` loop to compute the sum of all elements in the array.  
2. **Initialize variables**: Create `prefixSum` to store the left cumulative sum and `count` to track valid splits.  
3. **Iterate through the array**: Loop through the array, stopping before the last element.  
4. **Update the prefix sum**: Add the current element to `prefixSum` in each iteration.  
5. **Compute the right sum**: Subtract `prefixSum` from `totalSum` to dynamically compute the right sum.  
6. **Check the condition**: Compare `prefixSum` and the right sum. If `prefixSum` is greater than or equal to the right sum, increment `count`.  
7. **Return the count**: At the end of the loop, return the total number of valid splits.

---

## ⚙️ Complexity Analysis  

- **Time Complexity**:  
  The algorithm runs in \(O(n)\) time, where \(n\) is the size of the array. This is because we iterate through the array once to calculate the total sum and once more to evaluate valid splits.  

- **Space Complexity**:  
  The solution uses \(O(1)\) extra space since we only use a few variables (`total_sum`, `prefix_sum`, `count`).

---

## ✨ Conclusion  

This approach ensures optimal performance with minimal space usage, making it well-suited for large inputs. The use of prefix sums significantly reduces redundant computations, achieving \(O(n)\) efficiency across all implementations.
