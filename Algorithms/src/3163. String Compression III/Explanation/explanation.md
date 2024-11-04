# String Compression III - Step-by-Step Explanation

This repository provides a detailed explanation of how to solve the **String Compression III** problem across multiple programming languages. Here, you will find explanations for each language in a point-by-point format, guiding you through the problem-solving process step by step without revealing the full code.

---

## Problem Overview

Given a string `word`, compress it by grouping contiguous occurrences of each character up to a maximum of 9. Each group is represented by the count followed by the character. For instance:

- Input: `word = "aaaaaaaaaaaaaabb"`
- Output: `comp = "9a5a2b"`

### Steps for All Languages

The following explanations break down the steps needed to solve this problem in each of the five programming languages.

---

## C++ Code

1. **Initialize a Result String and Counter**:
   - Create an empty string, `comp`, to store the compressed result.
   - Set an integer variable, `count`, to 1 to begin counting the occurrences of each character.

2. **Iterate Over Each Character**:
   - Loop through each character in `word` starting from the second character.
   - For each character, check if it:
     - Differs from the previous character.
     - Reaches a count of 9.
     - Is the last character in the string.
   - In any of these cases, append the `count` and the previous character to `comp`.

3. **Reset or Increment Counter**:
   - If the above conditions are met, reset `count` to 1 to start counting the next character group.
   - Otherwise, increment `count` by 1 if the current character is the same as the previous one.

4. **Return Compressed String**:
   - Return `comp` after the loop finishes.

---

## Java Code

1. **Initialize a StringBuilder and Counter**:
   - Use a `StringBuilder` object `comp` to store the compressed output.
   - Initialize a `count` variable to keep track of consecutive character occurrences.

2. **Loop Through the String**:
   - Iterate through each character in `word` from the second position onward.
   - For each character, check:
     - If it’s different from the previous character.
     - If `count` has reached 9.
     - If it’s the last character in `word`.
   - If any condition is met, append the `count` and the previous character to `comp`.

3. **Update the Counter**:
   - Reset `count` to 1 if any of the above conditions are met.
   - Otherwise, increment `count` for consecutive identical characters.

4. **Output the Result**:
   - After the loop completes, convert `comp` to a string and return it.

---

## JavaScript Code

1. **Initialize an Empty String and Counter**:
   - Create an empty string `comp` to store the compressed result.
   - Initialize a `count` variable to 1 for counting character occurrences.

2. **Traverse the String**:
   - Loop through each character in `word` starting from the second position.
   - For each character, check if:
     - It differs from the previous character.
     - `count` reaches 9.
     - It’s the last character in `word`.
   - If any condition holds true, concatenate `count` and the previous character to `comp`.

3. **Manage the Counter**:
   - Reset `count` to 1 if the conditions are met.
   - Otherwise, increase `count` for each matching character.

4. **Return the Compressed Output**:
   - Once the loop ends, return `comp` containing the compressed string.

---

## Python Code

1. **Create a List for the Result and Initialize a Counter**:
   - Use a list `comp` to accumulate the compressed parts, which will be joined into a string at the end.
   - Set `count` to 1 to count each character’s consecutive occurrences.

2. **Loop Through Each Character**:
   - Iterate through `word` starting from the second position.
   - For each character, check:
     - If it differs from the previous character.
     - If `count` equals 9.
     - If it’s the last character in the string.
   - If any condition is met, add a formatted string with `count` and the previous character to `comp`.

3. **Adjust the Counter**:
   - Reset `count` to 1 if any condition applies.
   - Otherwise, increment `count` for consecutive identical characters.

4. **Join and Return the Result**:
   - Convert `comp` list to a single string using `join()` and return the result.

---

## Go Code

1. **Initialize an Empty String and Counter**:
   - Start with an empty string `comp` to hold the compressed version of `word`.
   - Set `count` to 1 to track character repetitions.

2. **Iterate Through Each Character**:
   - Loop through the `word` string from the second character onward.
   - For each character, check:
     - If it’s different from the previous character.
     - If `count` equals 9.
     - If it’s the final character in the word.
   - If any condition is met, append a formatted string of `count` and the previous character to `comp`.

3. **Update the Counter**:
   - Reset `count` to 1 if the character changes or `count` reaches 9.
   - Otherwise, increment `count` if the character repeats.

4. **Return the Final String**:
   - After the loop, `comp` will contain the compressed result. Return `comp` as the output.

---

## Summary

This approach ensures that all versions of the solution compress the string by tracking consecutive characters in groups of up to 9, adding each group's count and character to the result. By implementing the same logic across multiple languages, this guide offers a clear and consistent way to understand and solve the **String Compression III** problem.
