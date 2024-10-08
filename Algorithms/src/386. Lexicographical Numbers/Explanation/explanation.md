# Lexicographical Order of Numbers: Step-by-Step Explanation

This guide explains how to solve the problem of generating numbers from 1 to `n` in lexicographical order using different programming languages. The solution involves a **Depth-First Search (DFS)** strategy to explore numbers by appending digits in sequence.

## C++ Code Explanation

1. **Initialization**:
    - Create a `vector<int>` to store the numbers in lexicographical order.
    - Iterate from numbers 1 to 9 (these act as root nodes for DFS exploration).

2. **DFS Implementation**:
    - For each starting number, perform a DFS to explore all possible numbers by appending digits (0-9).
    - Each time a number is generated, add it to the result list.
    - Stop the DFS when the current number exceeds the limit `n`.

3. **Base Case**:
    - If the current number being generated is greater than `n`, terminate the recursion for that branch.

---

## Java Code Explanation

1. **Initialization**:
    - Create a `List<Integer>` to store the lexicographically ordered numbers.
    - Iterate from 1 to 9, using each number as a starting point for DFS.

2. **DFS Implementation**:
    - For each number, add it to the result list and then explore numbers by appending digits (0-9).
    - Recursively call DFS on the new number generated by appending digits.

3. **Base Case**:
    - When a number exceeds `n`, terminate the recursion for that branch to avoid unnecessary exploration.

---

## JavaScript Code Explanation

1. **Initialization**:
    - Initialize an empty array `result` to store the final lexicographical sequence.
    - Iterate from 1 to 9, treating each number as the start of a DFS traversal.

2. **DFS Implementation**:
    - For each starting number, recursively append digits (0-9) to generate new numbers.
    - Each generated number is added to the result array.

3. **Base Case**:
    - Stop the recursion if the current number exceeds `n`.

---

## Python Code Explanation

1. **Initialization**:
    - Use a list `result` to hold the numbers in lexicographical order.
    - Loop through the numbers 1 to 9 and perform DFS starting with each.

2. **DFS Implementation**:
    - Add the current number to the result list.
    - Recursively generate the next number by appending digits (0-9) to the current number.

3. **Base Case**:
    - Stop recursion when the number being generated exceeds `n`.

---

## Go Code Explanation

1. **Initialization**:
    - Create an empty slice `result` to store numbers in lexicographical order.
    - Iterate over numbers 1 to 9 as the starting points for DFS exploration.

2. **DFS Implementation**:
    - Perform DFS by appending digits (0-9) to the current number.
    - Append the valid numbers to the `result` slice.

3. **Base Case**:
    - If the current number exceeds `n`, stop further recursion for that branch.

---

This solution uses **Depth-First Search (DFS)** in all languages to explore numbers in lexicographical order. The main idea is to treat each digit as a root and recursively generate new numbers by appending digits from 0 to 9 until the number exceeds `n`.
