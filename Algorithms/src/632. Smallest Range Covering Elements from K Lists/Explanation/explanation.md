# Smallest Range Covering Elements from K Lists - Step-by-Step Explanation

This repository contains solutions to the **Smallest Range Covering Elements from K Lists** problem in **C++**, **Java**, **JavaScript**, **Python**, and **Go**, with a detailed step-by-step breakdown of the approach used in each language.

## Problem Overview

You are given `k` lists of sorted integers. The task is to find the smallest range that includes at least one number from each list. We utilize an optimized approach with a min-heap to solve this problem efficiently across multiple programming languages.

---

## C++ Code Explanation

1. **Min-Heap Initialization**:  
   - We create a priority queue to store the smallest elements from each list.
   - The heap stores a tuple containing the current value, the index of the list, and the index within the list.

2. **Track Maximum Value**:  
   - We initialize a variable to track the maximum value encountered during the heap operations, as this helps in calculating the range.

3. **Insert First Elements**:  
   - Insert the first element from each list into the heap while updating the maximum value.

4. **Extract-Min and Update Range**:  
   - Extract the smallest element from the heap and calculate the current range using the max value.
   - Compare the current range with the previous smallest range and update if it’s better.

5. **Move to Next Element**:  
   - Move to the next element in the list from which the minimum was extracted and push it into the heap.
   - If no more elements exist in that list, stop the process.

6. **Return Smallest Range**:  
   - Finally, return the smallest range found during the process.

---

## Java Code Explanation

1. **Heap Construction with Custom Comparator**:  
   - We utilize a `PriorityQueue` to keep track of the smallest elements from each list, stored in a tuple (value, list index, element index).

2. **Track Maximum Value**:  
   - As we insert the first element from each list into the heap, we simultaneously update the maximum value.

3. **Processing Heap**:  
   - While the heap is not empty, we extract the minimum element and calculate the current range with the tracked maximum value.

4. **Update Range**:  
   - If the new range is smaller than the previous range, we update our result.

5. **Move to Next Element**:  
   - After processing the smallest element, we attempt to insert the next element from the same list into the heap, adjusting the maximum value as needed.

6. **Termination Condition**:  
   - The process stops when one of the lists is exhausted.

---

## JavaScript Code Explanation

1. **Priority Queue Setup**:  
   - We use a custom priority queue implementation (or a library) to maintain the smallest element at the top.

2. **Tracking Maximum Value**:  
   - Keep a variable to track the current maximum value, as we extract and process elements from the queue.

3. **Insert First Elements**:  
   - Add the first element from each sorted list to the priority queue, keeping track of the maximum value encountered.

4. **Range Calculation**:  
   - For each extracted element, compute the current range and compare it with the smallest range found so far.

5. **Update the Queue**:  
   - Push the next element from the list of the extracted element into the queue and update the maximum value if necessary.

6. **Return Smallest Range**:  
   - Once a list is exhausted, stop and return the smallest range found during the process.

---

## Python Code Explanation

1. **Min-Heap Setup**:  
   - Use Python’s `heapq` module to implement the min-heap. Store tuples representing the value, list index, and index within the list.

2. **Track Maximum Value**:  
   - Keep track of the largest number encountered while processing the heap to calculate the range effectively.

3. **Insert Initial Elements**:  
   - Insert the first element from each list into the heap, ensuring the maximum value is updated accordingly.

4. **Extract Minimum and Compute Range**:  
   - Extract the minimum element, calculate the current range using the tracked maximum, and update the result if the current range is smaller.

5. **Move to the Next Element**:  
   - Push the next element from the list of the extracted element into the heap, updating the maximum value.

6. **Exit Condition**:  
   - Once one of the lists is exhausted, the loop terminates, and the smallest range is returned.

---

## Go Code Explanation

1. **Priority Queue Construction**:  
   - In Go, we manually implement a min-heap using the `container/heap` package. Each heap element contains the value, list index, and index within the list.

2. **Track Maximum Value**:  
   - We maintain a `maxValue` variable that keeps track of the maximum value in the current set of elements extracted from the heap.

3. **Push Initial Elements**:  
   - Insert the first element from each list into the priority queue and update the `maxValue`.

4. **Range Calculation**:  
   - Extract the minimum element, calculate the current range using the `maxValue`, and update the result if a smaller range is found.

5. **Push Next Element**:  
   - Insert the next element from the same list into the heap. If that element exceeds `maxValue`, update the `maxValue`.

6. **Termination Condition**:  
   - The process stops when a list runs out of elements, ensuring the smallest range has been found.

---

## Conclusion

By following these steps in any of the five programming languages, we efficiently solve the problem of finding the smallest range that covers at least one number from each of the given `k` lists. Each solution uses the same core logic of a min-heap to track the smallest element while ensuring that the range is minimized by adjusting the largest element dynamically.

For a detailed look at the full code, please refer to the respective files in this repository!
