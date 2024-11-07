# Largest Combination With Bitwise AND Greater Than Zero

This repository contains solutions to solve the problem **"Largest Combination With Bitwise AND Greater Than Zero"**. The goal is to find the size of the largest subset from an array of integers such that the bitwise AND of the subset is greater than zero. Below is a step-by-step breakdown of the solution logic used in each language.

---

### Problem Approach

1. **Intuition**: The problem requires us to identify the largest group of numbers where, if bitwise AND is applied across all elements in the subset, the result remains greater than zero.
2. **Bit Manipulation Insight**: If a certain bit position has many numbers with a `1` at that position, we can create a large subset with a non-zero bitwise AND.
3. **Optimization Strategy**: Instead of calculating bitwise ANDs for each subset, we count how many numbers have a `1` in each bit position (0 to 30), then select the maximum count among these.

---

## C++ Code Explanation

1. **Initialize Count Array**: Create an array `bitCount` to hold counts of `1`s for each of the 31 bit positions (since integers up to \(10^7\) have up to 31 bits).
2. **Counting Set Bits**:
   - For each number in the input array:
     - Loop through each bit position from 0 to 30.
     - Check if the bit at that position is `1` using bitwise AND (`num & (1 << i)`).
     - If `1`, increment the count in `bitCount` for that bit position.
3. **Get Maximum Count**:
   - After counting all bit positions, find the highest value in `bitCount`.
   - This value represents the maximum subset size where bitwise AND will be greater than zero.
4. **Return Result**: The highest count is returned as the solution.

---

## Java Code Explanation

1. **Initialize Count Array**: Create an array `bitCount` to store the count of `1`s for each bit position from 0 to 30.
2. **Counting Set Bits**:
   - For each integer in `candidates`:
     - Loop through bit positions from 0 to 30.
     - Use `num & (1 << i)` to check if the bit at that position is `1`.
     - If `1`, increment the corresponding index in `bitCount`.
3. **Get Maximum Count**:
   - Iterate through `bitCount` to find the highest count.
   - This maximum count represents the largest subset size that can achieve a non-zero bitwise AND.
4. **Return Result**: Return this maximum value as the result.

---

## JavaScript Code Explanation

1. **Initialize Count Array**: Use an array `bitCount` initialized to 31 zeros, each representing a bit position's count of `1`s.
2. **Counting Set Bits**:
   - For each number in `candidates`:
     - Loop over each bit position from 0 to 30.
     - Use `num & (1 << i)` to check if that bit position is `1`.
     - If so, increment `bitCount` at that position.
3. **Get Maximum Count**:
   - Use `Math.max(...bitCount)` to find the maximum count in `bitCount`.
   - This maximum value gives the size of the largest subset with a bitwise AND greater than zero.
4. **Return Result**: Return this maximum value as the solution.

---

## Python Code Explanation

1. **Initialize Count Array**: Create a list `bit_count` with 31 zeros to store counts for each bit position.
2. **Counting Set Bits**:
   - For each number in `candidates`:
     - Iterate over bit positions from 0 to 30.
     - Use `num & (1 << i)` to check if the `i`-th bit is `1`.
     - If true, increment `bit_count[i]`.
3. **Get Maximum Count**:
   - Use `max(bit_count)` to find the highest count in `bit_count`.
   - This highest count represents the maximum subset size with a non-zero AND.
4. **Return Result**: Return this maximum value.

---

## Go Code Explanation

1. **Initialize Count Array**: Declare a fixed-size array `bitCount` of 31 elements, initialized to zero, for each bit position count.
2. **Counting Set Bits**:
   - For each element in `candidates`:
     - Loop through each bit position from 0 to 30.
     - Check if the bit is set using `num & (1 << i)`.
     - If set, increment `bitCount[i]`.
3. **Get Maximum Count**:
   - Find the highest value in `bitCount` by iterating through it.
   - This maximum value is the largest subset size that can achieve a non-zero AND.
4. **Return Result**: Return this maximum count.
