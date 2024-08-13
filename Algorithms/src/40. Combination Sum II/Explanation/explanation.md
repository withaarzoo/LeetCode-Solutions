# Combination Sum II - Step-by-Step Explanation

This repository contains implementations of the `combinationSum2` algorithm in multiple programming languages: C++, Java, JavaScript, Python, and Go. The algorithm finds all unique combinations of numbers from a list that sum up to a given target. Each combination is unique, and the same number cannot be used multiple times.

---

## Common Steps Across All Implementations

1. **Sorting the Input Array**:  
   The candidates array is sorted to handle duplicates easily and optimize the backtracking process. This ensures that the algorithm can skip over duplicate elements and stop early when it’s clear that no further combination will meet the target.

2. **Backtracking Setup**:  
   Backtracking is used to explore all possible combinations. A helper function is defined that recursively tries to build combinations by adding elements from the sorted array.

3. **Base Case - Target Reached**:  
   If the target becomes zero, it indicates that the current combination sums up to the desired target. This combination is then added to the result list.

4. **Skipping Duplicates**:  
   To avoid generating duplicate combinations, the algorithm skips over elements that are the same as the previous one when iterating through the array.

5. **Early Stopping**:  
   If an element in the sorted array exceeds the current target, the loop breaks because any subsequent elements will also be too large.

6. **Backtracking**:  
   After exploring a combination, the last element added is removed (backtracking) to try other possible combinations.

---

## C++ Implementation

1. **Sort the Input Array**:  
   The candidates array is sorted to make duplicate handling easier.

2. **Initialize Result and Current Combination Vectors**:  
   Two vectors are initialized, one for storing valid combinations and one for tracking the current combination being explored.

3. **Backtracking Function**:  
   The main function calls a helper function `backtrack`, passing the current index, target, current combination, and result.

4. **Base Case - Target Reached**:  
   If the target is zero, the current combination is added to the result.

5. **Iterate Through Candidates**:  
   The loop iterates through the candidates, skipping duplicates and stopping early if the candidate exceeds the target.

6. **Recursively Backtrack**:  
   The `backtrack` function is called recursively with the updated target and the next index.

7. **Backtrack - Remove Last Element**:  
   After exploring a path, the last element is removed to explore other combinations.

## Java Implementation

1. **Sort the Array**:  
   The array is sorted for easier handling of duplicates.

2. **Initialize Result and Temporary List**:  
   Two lists are created: one for storing results and one for the current combination.

3. **Backtracking Method**:  
   A helper method `backtrack` is defined to handle the recursive exploration of combinations.

4. **Base Case - Target Zero**:  
   If the target is zero, a copy of the current list is added to the result.

5. **Loop Through Candidates**:  
   The loop checks for duplicates and whether the current candidate can be part of a valid combination.

6. **Recursive Exploration**:  
   The `backtrack` method is called with an updated target and index.

7. **Remove Last Element**:  
   Backtrack by removing the last added element before exploring the next candidate.

## JavaScript Implementation

1. **Sort the Candidates Array**:  
   The array is sorted to manage duplicates and allow early stopping.

2. **Initialize Result and Current Arrays**:  
   Arrays are initialized to store results and the current combination.

3. **Define Backtracking Function**:  
   A function `backtrack` is defined to explore combinations.

4. **Check for Target Zero**:  
   When the target is zero, the current combination is added to the result.

5. **Iterate Through Candidates**:  
   The loop skips duplicates and breaks early if a candidate is too large.

6. **Recursive Backtracking**:  
   Call `backtrack` recursively with the remaining target and next index.

7. **Backtrack by Popping Last Element**:  
   Remove the last element to try other combinations.

## Python Implementation

1. **Sort the Candidates List**:  
   The list is sorted to facilitate duplicate handling and early stopping.

2. **Initialize Result List**:  
   A list is created to store the valid combinations.

3. **Backtracking Function**:  
   A nested function `backtrack` is defined to handle recursive exploration.

4. **Check for Target Zero**:  
   When the target reaches zero, the current combination is added to the result.

5. **Loop Through Candidates**:  
   The loop skips duplicates and stops early if a candidate is greater than the target.

6. **Recursively Call Backtrack**:  
   Call `backtrack` recursively with the reduced target and updated index.

7. **Backtrack by Removing Last Element**:  
   After exploring a combination, the last element is removed to try another path.

## Go Implementation

1. **Sort the Input Slice**:  
   The slice is sorted for managing duplicates and stopping early.

2. **Initialize Result and Current Slices**:  
   Slices are initialized to store results and the current combination.

3. **Define Backtracking Function**:  
   A function `backtrack` is defined to recursively explore combinations.

4. **Check for Target Zero**:  
   When the target is zero, the current combination is added to the result.

5. **Iterate Over Candidates**:  
   The loop checks for duplicates and breaks if a candidate exceeds the target.

6. **Recursive Backtracking**:  
   The `backtrack` function is called with the updated target and index.

7. **Backtrack by Popping Last Element**:  
   The last added element is removed to explore other possibilities.
