# CustomStack Implementation - Step by Step Explanation

In this `README`, we will walk through the implementation of the `CustomStack` class in five different programming languages: C++, Java, JavaScript, Python, and Go. The `CustomStack` class is designed to maintain a stack with additional functionality to increment the bottom k elements. Each section will describe how the core methods (`push`, `pop`, and `increment`) work step by step without revealing the actual code.

---

## C++ Code Explanation

1. **Initialization**:
   - The constructor accepts `maxSize` and initializes two vectors: one for the stack elements and another for tracking increments.

2. **Push Operation**:
   - If the current size of the stack is less than `maxSize`, the new element is added to the stack.
   - Along with the element, an increment placeholder (initialized to 0) is added to a separate vector to track future increments for each stack element.

3. **Pop Operation**:
   - If the stack is empty, return `-1`.
   - The value to be popped is the sum of the top element and any accumulated increments.
   - If the stack has more than one element, the increment of the current element is propagated to the next lower element.
   - Both the element and its corresponding increment placeholder are removed.

4. **Increment Operation**:
   - The bottom `k` elements of the stack are incremented by a given value.
   - This is done by adding the increment to the corresponding position in the `inc` vector.

---

## Java Code Explanation

1. **Initialization**:
   - The constructor accepts `maxSize` and initializes two lists: one for the stack and another for the increments.

2. **Push Operation**:
   - If the stack size is smaller than `maxSize`, the element is added.
   - A 0 is appended to the `inc` list, serving as a placeholder for future increments.

3. **Pop Operation**:
   - Return `-1` if the stack is empty.
   - Otherwise, pop the top element and add any accumulated increment to it.
   - If there are more elements below, propagate the current increment to the element below it.
   - Remove the element and the corresponding increment entry.

4. **Increment Operation**:
   - Increment the bottom `k` elements by a given value.
   - This is done by modifying the `inc` list at the appropriate position.

---

## JavaScript Code Explanation

1. **Initialization**:
   - The constructor takes `maxSize` and initializes two arrays: one for stack elements and one for increments.

2. **Push Operation**:
   - Add an element to the stack if the current size is less than `maxSize`.
   - Also, append a 0 to the `inc` array for future increment tracking.

3. **Pop Operation**:
   - If the stack is empty, return `-1`.
   - Otherwise, retrieve the top element and add any pending increment to it.
   - If there are elements below, propagate the current increment to the next lower element.
   - Remove the element and the corresponding increment value.

4. **Increment Operation**:
   - Add a value to the bottom `k` elements by updating the `inc` array.
   - The increment is applied to the appropriate element in the `inc` array.

---

## Python Code Explanation

1. **Initialization**:
   - The constructor accepts `maxSize` and initializes two lists: one for the stack and another for the increments.

2. **Push Operation**:
   - If the stack has fewer elements than `maxSize`, push the new element to the stack.
   - Append `0` to the `inc` list, representing the placeholder for future increments.

3. **Pop Operation**:
   - If the stack is empty, return `-1`.
   - Otherwise, pop the top element and apply any accumulated increment.
   - If there are other elements in the stack, propagate the increment to the next element down.
   - Remove both the top element and its corresponding increment.

4. **Increment Operation**:
   - Increment the bottom `k` elements by updating the value in the `inc` list.
   - The increment is applied to the lowest `k` elements via the `inc` array.

---

## Go Code Explanation

1. **Initialization**:
   - The constructor takes `maxSize` and initializes two slices: one for the stack and another for the increments.

2. **Push Operation**:
   - Add the element to the stack if its size is less than `maxSize`.
   - A zero is appended to the `inc` slice for future increments.

3. **Pop Operation**:
   - If the stack is empty, return `-1`.
   - Otherwise, the top element is popped and any accumulated increment is applied.
   - If the stack has more than one element, propagate the increment down to the next element.
   - Remove both the element and its increment value.

4. **Increment Operation**:
   - Increment the bottom `k` elements by updating the `inc` slice.
   - This involves adding the given value to the appropriate position in the `inc` slice.

---

### Conclusion

In all five implementations, the `CustomStack` class performs three core operations: `push`, `pop`, and `increment`. While the syntax differs across languages, the underlying logic remains consistent:

- The `push` method adds elements up to the maximum size.
- The `pop` method removes the top element while applying any pending increments.
- The `increment` method allows modifying the bottom `k` elements efficiently.

Each implementation uses an additional array or list to track the increments, allowing for deferred and efficient handling of bulk updates during the `pop` operation.
