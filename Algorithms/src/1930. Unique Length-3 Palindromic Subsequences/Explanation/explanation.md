# Count Palindromic Subsequences (Length 3)  

This repository provides a solution to the problem of counting unique length-3 palindromic subsequences (`a _ a`) in a string. The explanation is broken down step-by-step for each language: **C++**, **Java**, **JavaScript**, **Python**, and **Go**.

---

## Problem Statement  

Given a string `s`, find the number of unique palindromic subsequences of length 3 in the form `a _ a`, where the first and last characters are the same, and the middle character can be any character.

---

## Approach  

The solution involves:

1. **Tracking first and last occurrences** of each character in the string.
2. **Finding unique middle characters** between the first and last occurrences for every character.
3. **Counting distinct palindromes** based on unique middle characters.

---

## Step-by-Step Explanation for Each Language  

### C++ Code  

1. **Initialize First and Last Occurrence Arrays**:  
   - Use an array of size 26 (for all lowercase letters) to store the first and last occurrence of each character in the string.
   - Iterate over the string to populate these arrays.

2. **Identify Characters with Multiple Occurrences**:  
   - For each character, check if its first occurrence is not `-1` and if its last occurrence is after its first occurrence.

3. **Collect Unique Middle Characters**:  
   - For the characters identified in step 2, iterate over the range between the first and last occurrences.
   - Use a `set` to store unique characters in this range.

4. **Count Unique Palindromes**:  
   - Add the size of the `set` (number of unique middle characters) to the result for each character.

---

### Java Code  

1. **Set Up Arrays for First and Last Occurrences**:  
   - Use two arrays of size 26 to track the first and last indices of each character. Initialize the `first` array with `-1` to indicate unvisited characters.

2. **Populate the Arrays**:  
   - Traverse the string and update the `first` and `last` occurrence of each character.

3. **Check for Eligible Characters**:  
   - For each letter of the alphabet, verify if it appears more than once by comparing its first and last indices.

4. **Extract Unique Middle Characters**:  
   - Use a `Set` to store all distinct middle characters found between the first and last occurrences.

5. **Count Unique Palindromic Subsequences**:  
   - Add the size of the `Set` to the result.

---

### JavaScript Code  

1. **Initialize Arrays**:  
   - Create two arrays (`first` and `last`) of size 26 and set all elements to `-1`.

2. **Track First and Last Indices**:  
   - Loop through the string to determine the first and last occurrence of each character.

3. **Find Valid Characters for Palindromes**:  
   - Check if a character appears more than once by comparing its first and last indices.

4. **Count Unique Middle Characters**:  
   - Iterate over the range between the first and last indices and use a `Set` to store distinct middle characters.

5. **Sum the Results**:  
   - Add the number of unique middle characters (size of the `Set`) to the final count.

---

### Python Code  

1. **Set Up First and Last Occurrence Lists**:  
   - Use two lists of size 26 initialized to `-1` to store the first and last occurrence indices of each character.

2. **Populate First and Last Occurrences**:  
   - Iterate over the string and update the `first` and `last` lists based on the character’s position.

3. **Identify Valid Characters for Palindromes**:  
   - For each character in the alphabet, check if its first and last indices indicate multiple occurrences.

4. **Extract Middle Characters**:  
   - Use a Python `set` to store unique characters between the first and last indices of valid characters.

5. **Add to Result**:  
   - Count the number of unique middle characters for each valid character and add them to the result.

---

### Go Code  

1. **Initialize Arrays for Tracking Indices**:  
   - Create two slices (`first` and `last`) of size 26 and initialize the `first` slice with `-1`.

2. **Populate the Arrays**:  
   - Loop through the string to find the first and last occurrence of each character.

3. **Validate Characters for Palindromes**:  
   - Check if a character has multiple occurrences by comparing its first and last indices.

4. **Store Unique Middle Characters**:  
   - Use a map to store unique middle characters between the first and last indices of valid characters.

5. **Count Palindromic Subsequences**:  
   - Add the number of unique middle characters for each valid character to the total count.

---

## Complexity Analysis  

- **Time Complexity**:  
  - The solution involves two passes through the string: one to calculate first and last occurrences, and another to count unique middle characters.  
  - Overall time complexity: \(O(n)\).

- **Space Complexity**:  
  - We use additional data structures (arrays or sets) to store indices and unique characters.  
  - Overall space complexity: \(O(n)\).

---

## Additional Notes  

- The solution is optimized for large strings due to its linear time complexity.  
- Using sets or maps ensures we only count unique middle characters efficiently.  

Feel free to explore the code in each language and try it with different test cases! 🚀  
