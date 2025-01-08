# **Count Prefix-Suffix Pairs**  

This repository provides step-by-step explanations of the solution for the problem "Count Prefix-Suffix Pairs" in **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Below, you'll find a detailed explanation for each language, written as steps for easy understanding.  

---

## **C++ Code Explanation**  

### Step 1: Initialize Variables  

- Create a variable `count` to store the total number of valid prefix-suffix pairs.  
- Use `n` to store the size of the input array `words`.  

### Step 2: Iterate Through All Pairs  

- Use a **nested loop**:  
  - The outer loop selects each string `words[i]` one by one.  
  - The inner loop compares it with every subsequent string `words[j]` (where \(j > i\)).  

### Step 3: Check Prefix and Suffix  

- Use the `substr` function:  
  - Check if `words[j]` starts with `words[i]` by comparing the first `len(words[i])` characters of `words[j]`.  
  - Check if `words[j]` ends with `words[i]` by comparing the last `len(words[i])` characters of `words[j]`.  

### Step 4: Update the Counter  

- If both conditions (prefix and suffix) are satisfied, increment the `count`.  

### Step 5: Return the Result  

- After evaluating all pairs, return the value of `count`.  

---

## **Java Code Explanation**  

### Step 1: Initialize Variables  

- Create a variable `count` to store the total number of valid prefix-suffix pairs.  
- Use `n` to store the length of the input array `words`.  

### Step 2: Iterate Through All Pairs  

- Use two **nested loops**:  
  - The outer loop iterates through each string `words[i]`.  
  - The inner loop evaluates every subsequent string `words[j]` (\(j > i\)).  

### Step 3: Check Prefix and Suffix  

- Use the `startsWith` and `endsWith` methods:  
  - Verify if `words[j]` starts with `words[i]`.  
  - Verify if `words[j]` ends with `words[i]`.  

### Step 4: Update the Counter  

- If both checks pass, increment the `count`.  

### Step 5: Return the Result  

- Return the `count` after evaluating all pairs.  

---

## **JavaScript Code Explanation**  

### Step 1: Initialize Variables  

- Create a variable `count` to store the total number of valid prefix-suffix pairs.  
- Use `n` to store the length of the `words` array.  

### Step 2: Iterate Through All Pairs  

- Use two **nested loops**:  
  - The outer loop selects `words[i]`.  
  - The inner loop compares it with every subsequent `words[j]` (\(j > i\)).  

### Step 3: Check Prefix and Suffix  

- Use `startsWith` and `endsWith`:  
  - Check if `words[j]` starts with `words[i]`.  
  - Check if `words[j]` ends with `words[i]`.  

### Step 4: Update the Counter  

- Increment `count` if both conditions are satisfied.  

### Step 5: Return the Result  

- After completing the iterations, return `count`.  

---

## **Python Code Explanation**  

### Step 1: Initialize Variables  

- Create a variable `count` to store the total number of valid prefix-suffix pairs.  
- Use `n` to store the size of the `words` list.  

### Step 2: Iterate Through All Pairs  

- Use two **nested loops**:  
  - The outer loop iterates over each string `words[i]`.  
  - The inner loop compares it with every subsequent string `words[j]` (\(j > i\)).  

### Step 3: Check Prefix and Suffix  

- Use `startswith` and `endswith`:  
  - Check if `words[j]` starts with `words[i]`.  
  - Check if `words[j]` ends with `words[i]`.  

### Step 4: Update the Counter  

- Increment `count` if both conditions (prefix and suffix) are satisfied.  

### Step 5: Return the Result  

- Return `count` after evaluating all pairs.  

---

## **Go Code Explanation**  

### Step 1: Initialize Variables  

- Create a variable `count` to store the total number of valid prefix-suffix pairs.  
- Use `n` to store the length of the `words` slice.  

### Step 2: Iterate Through All Pairs  

- Use two **nested loops**:  
  - The outer loop iterates over each string `words[i]`.  
  - The inner loop evaluates every subsequent string `words[j]` (\(j > i\)).  

### Step 3: Check Prefix and Suffix  

- Use `strings.HasPrefix` and `strings.HasSuffix`:  
  - Check if `words[j]` starts with `words[i]`.  
  - Check if `words[j]` ends with `words[i]`.  

### Step 4: Update the Counter  

- Increment `count` if both prefix and suffix conditions are met.  

### Step 5: Return the Result  

- Return `count` after iterating through all pairs.  

---

## **Conclusion**  

This problem involves string manipulation and nested iteration to check prefix-suffix relationships. Although the approach has a time complexity of \(O(n^2 \times m)\), it is straightforward and demonstrates basic string operations in multiple programming languages. If you have further optimizations or need help, feel free to open an issue or discuss here! 😊  
