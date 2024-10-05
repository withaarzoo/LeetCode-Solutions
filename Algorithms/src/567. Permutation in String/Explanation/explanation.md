# Check Inclusion of Permutation - Step by Step Explanation

This repository contains solutions to the "Check Inclusion of Permutation" problem implemented in multiple languages: C++, Java, JavaScript, Python, and Go. Each language implements a sliding window approach to solve the problem efficiently.

## Problem Overview

Given two strings `s1` and `s2`, the task is to check if a permutation of `s1` is present as a substring in `s2`. A sliding window technique is used to compare character frequency counts of `s1` with substrings of `s2`.

### Steps to Solution

### 1. **Initial Checks**

- **Condition**: If the length of `s1` is greater than the length of `s2`, immediately return `false`. This is because `s2` cannot contain any permutation of `s1` if it's shorter.

### 2. **Frequency Count Setup**

- **Count Frequency**: Initialize two frequency count arrays (or lists) of size 26 (for 26 letters in the alphabet). One array counts the frequency of characters in `s1`, and the other does the same for the first window (substring of the same length as `s1`) in `s2`.

   Example:

- For each character in `s1`, update its count in `s1Count`.
- For each character in the first window of `s2`, update its count in `s2Count`.

### 3. **Sliding the Window**

- **Slide Window**: Slide the window across `s2` and update the character frequency counts as you go. At each step, compare the counts of the current window in `s2` with the counts of `s1`.
- **Compare Counts**: If the two frequency arrays match, return `true` because a permutation of `s1` is found in `s2`.
- **Window Update**:
  - Decrement the count of the character that is sliding out of the window (the leftmost character).
  - Increment the count of the new character that is entering the window (the rightmost character).

### 4. **Final Check**

- **Last Window**: After sliding through the entire `s2`, compare the frequency counts one last time for the final window in `s2`. If they match, return `true`; otherwise, return `false`.

---

## C++ Code Explanation

1. **Initial Checks**: Compare the length of `s1` and `s2`.
2. **Frequency Count Setup**: Use two `vector<int>` arrays to track the frequency of each character in `s1` and the first window of `s2`.
3. **Sliding the Window**: Slide the window across `s2` by adjusting the frequency counts of characters entering and leaving the window.
4. **Final Check**: Compare the counts of the final window with `s1Count`.

---

## Java Code Explanation

1. **Initial Checks**: Check if `s1` is longer than `s2`. Return `false` if it is.
2. **Frequency Count Setup**: Initialize two `int[]` arrays for frequency counts.
3. **Sliding the Window**: Move the window across `s2` and update counts using the helper method `matches` to compare `s1Count` and `s2Count`.
4. **Final Check**: After sliding through all possible windows, compare the last window's count.

---

## JavaScript Code Explanation

1. **Initial Checks**: If the length of `s1` is larger than `s2`, return `false`.
2. **Frequency Count Setup**: Create two arrays of size 26 for `s1Count` and `s2Count`.
3. **Sliding the Window**: Update the frequency counts as the window moves across `s2`. Use a helper function `matches` to compare the two arrays.
4. **Final Check**: After sliding through all windows, ensure the last window is checked.

---

## Python Code Explanation

1. **Initial Checks**: Verify if `s1` is longer than `s2`. If so, return `false`.
2. **Frequency Count Setup**: Create two lists of size 26 for character frequency counts.
3. **Sliding the Window**: Adjust the counts for each new character entering and leaving the sliding window.
4. **Final Check**: Perform the last check for the final window after sliding across `s2`.

---

## Go Code Explanation

1. **Initial Checks**: If `s1` is longer than `s2`, return `false`.
2. **Frequency Count Setup**: Create two slices (arrays) of size 26 for `s1Count` and `s2Count`.
3. **Sliding the Window**: Update character counts by decrementing the count for the character that leaves the window and incrementing for the one that enters.
4. **Final Check**: After sliding through `s2`, compare the frequency counts for the last window.

---

This step-by-step explanation follows a uniform structure in all the languages, highlighting how the sliding window technique is applied to solve the problem efficiently.
