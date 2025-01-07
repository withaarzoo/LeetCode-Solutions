# String Matching in an Array  

This repository contains solutions to the "String Matching in an Array" problem in multiple programming languages: **C++**, **Java**, **JavaScript**, **Python**, and **Go**.

---

## Problem Description  

Given an array of strings `words`, find all strings in the array that are substrings of another string in the array. Return the list of all such strings.  

---

## Steps to Solve  

Below, you'll find a step-by-step explanation of how the problem is solved in each language.

---

### 🚀 C++ Solution  

#### Step 1: Sort the Strings  

- Use a custom comparator to sort the strings in ascending order by length.  
- This ensures shorter strings are checked first as they are more likely to be substrings of longer strings.  

#### Step 2: Use a Nested Loop  

- The outer loop iterates through each string in the sorted array.  
- For every string in the outer loop, the inner loop checks if it is a substring of any string that appears after it in the sorted array.  

#### Step 3: Substring Check  

- Use the `find` method to check if the current string from the outer loop is a substring of the string in the inner loop.  
- If a match is found, add the string to the result list and stop further checks for that string.  

#### Step 4: Return Results  

- Once all iterations are complete, return the list of strings that are substrings of others.  

---

### ☕ Java Solution  

#### Step 1: Sort the Strings  

- Use `Arrays.sort` with a custom comparator to sort the strings by their lengths in ascending order.  
- This ensures shorter strings are processed first.  

#### Step 2: Use a Nested Loop  

- The outer loop iterates over each string in the sorted array.  
- The inner loop checks if the string in the outer loop is a substring of any string that comes after it in the sorted array.  

#### Step 3: Substring Check  

- Use the `contains` method of the `String` class to check if the current string is a substring of another.  
- If a match is found, add the string to the result list and break the inner loop to move to the next string.  

#### Step 4: Return Results  

- Return the list of strings that were found to be substrings of others.  

---

### 🖥️ JavaScript Solution  

#### Step 1: Sort the Strings  

- Use the `sort` function with a comparator to arrange the strings in ascending order based on their lengths.  

#### Step 2: Use a Nested Loop  

- Loop through each string in the sorted array with an outer loop.  
- Use an inner loop to compare the current string with the strings that come after it in the sorted array.  

#### Step 3: Substring Check  

- Use the `includes` method to check if the current string is a substring of any longer string.  
- If a match is found, add the string to the result array and stop further checks for that string.  

#### Step 4: Return Results  

- After all iterations are complete, return the result array containing all substrings.  

---

### 🐍 Python Solution  

#### Step 1: Sort the Strings  

- Use the `sort` method with the `key=len` parameter to sort the strings by their lengths in ascending order.  

#### Step 2: Use a Nested Loop  

- The outer loop iterates through each string in the sorted list.  
- The inner loop compares the current string with all longer strings that come after it in the sorted list.  

#### Step 3: Substring Check  

- Use the `in` operator to check if the current string is a substring of any other string.  
- If a match is found, add the string to the result list and stop further checks for that string.  

#### Step 4: Return Results  

- Return the result list containing all the substrings.  

---

### 🛠️ Go Solution  

#### Step 1: Sort the Strings  

- Use `sort.Slice` to sort the strings in ascending order by length.  

#### Step 2: Use a Nested Loop  

- Iterate through the sorted slice of strings with an outer loop.  
- Use an inner loop to compare the current string with all strings that appear later in the sorted slice.  

#### Step 3: Substring Check  

- Use `strings.Contains` to check if the current string is a substring of another.  
- If a match is found, add the string to the result slice and stop further checks for that string.  

#### Step 4: Return Results  

- Return the result slice containing all the substrings.  

---

## Complexity Analysis  

### Time Complexity  

- **Sorting:** Sorting the array takes \(O(n \log n)\), where \(n\) is the number of strings.  
- **Substring Checks:** Nested loops result in \(O(n^2 \cdot m)\), where \(m\) is the average length of the strings.  
- **Overall:** \(O(n^2 \cdot m)\).  

### Space Complexity  

- **Sorting:** \(O(1)\) additional space.  
- **Result Storage:** \(O(k)\), where \(k\) is the number of substrings.  
- **Overall:** \(O(k)\).  

---

## Supported Languages  

- **C++**  
- **Java**  
- **JavaScript**  
- **Python**  
- **Go**  

Feel free to explore each language's solution and choose the one that fits your needs! 😊  
