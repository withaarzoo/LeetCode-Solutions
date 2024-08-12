# KthLargest Element Tracker

This guide provides a step-by-step explanation of how to implement a data structure that tracks the k-th largest element in a stream of numbers using different programming languages: C++, Java, JavaScript, Python, and Go.

---

## C++ Implementation

1. **Initialize the KthLargest class**:
   - Create a class `KthLargest` with a constructor that accepts `k` (the position of the k-th largest element) and a vector of integers `nums`.
   - The constructor initializes a min-heap (`priority_queue` with `greater<int>` comparator) to keep track of the k largest elements.

2. **Add elements to the heap**:
   - Iterate over each element in `nums`.
   - For each element, use the `add()` function to determine if it belongs in the k largest elements.
  
3. **Maintain the heap size**:
   - In the `add()` function, if the heap size is less than `k`, simply add the element.
   - If the heap size is equal to `k`, check if the new element is larger than the smallest element in the heap. If it is, replace the smallest element with the new one.

4. **Return the k-th largest element**:
   - The smallest element in the min-heap represents the k-th largest element in the stream. Return this value.

---

## Java Implementation

1. **Define the KthLargest class**:
   - Create a class `KthLargest` with instance variables `k` and a `PriorityQueue` (min-heap) to store the k largest elements.

2. **Constructor and initialization**:
   - Initialize the class with `k` and an array of integers `nums`.
   - Populate the heap by iterating through `nums` and calling the `add()` method for each element.

3. **Add elements and maintain heap size**:
   - The `add()` method checks if the heap size is less than `k`. If so, add the new value.
   - If the heap already contains `k` elements, compare the new value with the smallest element in the heap. If the new value is larger, replace the smallest element.

4. **Retrieve the k-th largest element**:
   - The smallest element in the heap (root of the `PriorityQueue`) is the k-th largest. Return this value.

---

## JavaScript Implementation

1. **Initialize the KthLargest class**:
   - Define a class `KthLargest` with a constructor that accepts `k` and an array of numbers `nums`.
   - Use a minimum priority queue (`MinPriorityQueue`) to track the k largest elements.

2. **Process the initial array**:
   - Loop through `nums` and add each number to the heap using the `add()` method.

3. **Add new values and maintain heap**:
   - The `add()` method first checks if the heap size is less than `k`. If true, add the new value.
   - If the heap is full, compare the new value with the smallest value in the heap. Replace the smallest value if the new value is larger.

4. **Return the k-th largest element**:
   - The smallest element in the heap is the k-th largest element in the current stream. Return this value.

---

## Python Implementation

1. **Create the KthLargest class**:
   - Define a class `KthLargest` with a constructor that takes `k` and a list `nums`.
   - Use a list `minHeap` to implement a min-heap, utilizing Python's `heapq` module.

2. **Initialize and populate the heap**:
   - Iterate over `nums`, adding each element to the heap via the `add()` method.

3. **Add values and manage heap size**:
   - In the `add()` method, if the heap contains fewer than `k` elements, push the new value.
   - If the heap is full and the new value is larger than the smallest in the heap, replace the smallest element.

4. **Output the k-th largest element**:
   - Return the smallest element in the heap, which corresponds to the k-th largest element.

---

## Go Implementation

1. **Define the KthLargest struct**:
   - Create a `KthLargest` struct with `k` and a pointer to an `IntHeap` (min-heap).

2. **Constructor and initialization**:
   - The `Constructor` function initializes the `KthLargest` struct and processes each element in `nums` by adding them to the heap using the `Add()` method.

3. **Adding elements to the heap**:
   - The `Add()` method checks if the heap has fewer than `k` elements. If true, the new value is pushed onto the heap.
   - If the heap is full, compare the new value with the smallest element. Replace the smallest element if the new value is larger.

4. **Return the k-th largest element**:
   - The smallest element in the heap (root of the heap) is the k-th largest. Return this value.
