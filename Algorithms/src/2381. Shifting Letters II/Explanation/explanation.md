# Shifting Letters with Range Updates  

This repository contains the implementation of the **Shifting Letters** problem in multiple languages (**C++**, **Java**, **JavaScript**, **Python**, and **Go**). Below is the detailed step-by-step explanation for each implementation.

---

## C++ Code Explanation

1. **Initialize Variables**:
   - Create a difference array `diff` with size `n + 1` (where \(n\) is the length of the string).
   - Initialize all elements in the array to `0`.

2. **Build the Difference Array**:
   - Iterate through the `shifts` array.
   - For each shift, calculate the range `[start, end]` and direction (forward or backward).
   - Update the `start` and `end + 1` indices in the difference array based on the direction of the shift.

3. **Calculate Cumulative Shifts**:
   - Iterate through the difference array to calculate the cumulative sum.
   - Normalize the cumulative shift to always stay within the range `[0, 25]` using modular arithmetic.

4. **Apply the Shifts**:
   - Iterate through the string.
   - Modify each character based on the cumulative shift value, ensuring wrap-around in the alphabet.

5. **Return Result**:
   - Construct and return the modified string after applying all the shifts.

---

## Java Code Explanation

1. **Initialize Variables**:
   - Create an array `diff` with size `n + 1` to represent the difference array.
   - Initialize all elements to `0`.

2. **Build the Difference Array**:
   - Loop through the `shifts` array.
   - For each operation, adjust the difference array based on the range `[start, end]` and direction.

3. **Calculate Cumulative Shifts**:
   - Use a running sum to calculate the total shift for each character.
   - Normalize the shift using modular arithmetic to handle wrap-around.

4. **Apply the Shifts**:
   - Convert the string to a character array.
   - Modify each character based on the cumulative shift, and update the array.

5. **Return Result**:
   - Convert the character array back to a string and return it.

---

## JavaScript Code Explanation

1. **Initialize Variables**:
   - Create an array `diff` of size `n + 1` initialized with zeros.

2. **Build the Difference Array**:
   - Iterate through the `shifts` array.
   - For each shift operation, adjust the start and end indices of the difference array.

3. **Calculate Cumulative Shifts**:
   - Compute the running sum for the difference array.
   - Normalize the shift values using modular arithmetic to stay within the alphabet range.

4. **Apply the Shifts**:
   - Convert the string into an array of characters for easier manipulation.
   - Apply the cumulative shift to each character, wrapping around within the alphabet.

5. **Return Result**:
   - Join the modified character array back into a string and return it.

---

## Python Code Explanation

1. **Initialize Variables**:
   - Create a difference array `diff` of size `n + 1`, initialized to zeros.

2. **Build the Difference Array**:
   - Loop through the `shifts` array.
   - For each shift, update the difference array based on the range `[start, end]` and direction.

3. **Calculate Cumulative Shifts**:
   - Use a running sum to calculate the total shift for each character.
   - Normalize the shifts to ensure they stay within `[0, 25]` using modular arithmetic.

4. **Apply the Shifts**:
   - Convert the string into a list of characters.
   - Apply the cumulative shift to each character, wrapping around as needed.

5. **Return Result**:
   - Join the modified list of characters into a string and return it.

---

## Go Code Explanation

1. **Initialize Variables**:
   - Create a slice `diff` of size `n + 1` initialized to zeros.

2. **Build the Difference Array**:
   - Iterate through the `shifts` slice.
   - For each operation, adjust the `start` and `end + 1` indices in the difference slice.

3. **Calculate Cumulative Shifts**:
   - Use a running sum to compute the cumulative shifts for each index.
   - Normalize the shift values using modular arithmetic to wrap around within the alphabet.

4. **Apply the Shifts**:
   - Convert the string to a slice of bytes for easier manipulation.
   - Modify each byte based on the cumulative shift, ensuring wrap-around within the alphabet.

5. **Return Result**:
   - Convert the modified byte slice back to a string and return it.

---

## Additional Notes

- **Efficiency**: All implementations use the difference array technique to aggregate shifts efficiently, achieving a time complexity of \(O(n + m)\), where \(n\) is the length of the string and \(m\) is the number of shifts.
- **Normalization**: Modular arithmetic ensures all calculations stay within the bounds of the alphabet.
