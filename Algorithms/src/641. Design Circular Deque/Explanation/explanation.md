# MyCircularDeque Implementation in Multiple Languages

This README provides a detailed step-by-step explanation of how a **Circular Deque** is implemented across different programming languages including C++, Java, JavaScript, and Python. We'll break down the logic used in each method and how the circular nature of the deque is handled. Each section below focuses on the core concepts applied in the implementation.

## C++ Implementation

### Key Concepts

1. **Vector Representation**:
    - The deque is implemented using a vector (`deque`) to store elements. An extra space (`k + 1`) is allocated to differentiate between the full and empty states of the deque.

2. **Pointers for Front and Rear**:
    - `front` and `rear` pointers keep track of the positions at the front and rear of the deque. Initially, both are set to `0`.

3. **Circular Movement**:
    - Both the front and rear pointers move in a circular manner using modulo operations (`%`) to wrap around when they reach the end or beginning of the deque.

### Methods

#### Constructor

- Initializes the deque with `k + 1` size to handle the wrap-around logic. The `front` and `rear` pointers are both set to 0.

#### `insertFront`

- Checks if the deque is full using the `isFull()` method.
- Moves the `front` pointer backward using a circular operation and inserts the value.

#### `insertLast`

- Checks if the deque is full.
- Inserts the value at the `rear` pointer's current position and moves the `rear` pointer forward in a circular manner.

#### `deleteFront`

- Checks if the deque is empty.
- Moves the `front` pointer forward to remove the element at the front.

#### `deleteLast`

- Checks if the deque is empty.
- Moves the `rear` pointer backward to remove the last element.

#### `getFront` and `getRear`

- Returns the front or rear element. In the case of `getRear`, the `rear` pointer points to the next position, so we subtract one and wrap it around.

#### `isEmpty` and `isFull`

- `isEmpty`: The deque is empty when `front == rear`.
- `isFull`: The deque is full when the next position of `rear` equals `front`.

---

## Java Implementation

### Key Concepts

1. **Array Representation**:
    - The deque is represented as an array (`deque[]`) with an extra space (`k + 1`) to distinguish between full and empty states.

2. **Pointers for Front and Rear**:
    - Similar to C++, the `front` and `rear` pointers track the positions at the front and rear of the deque. The pointers move in a circular manner using modulo operations.

### Methods

#### Constructor

- Initializes an array of size `k + 1` with `front` and `rear` pointers set to 0.

#### `insertFront`

- Checks if the deque is full.
- Moves the `front` pointer backward and inserts the value at the new position.

#### `insertLast`

- Inserts a value at the `rear` position and moves the `rear` pointer forward.

#### `deleteFront`

- Moves the `front` pointer forward, effectively removing the front element.

#### `deleteLast`

- Moves the `rear` pointer backward to remove the last element.

#### `getFront` and `getRear`

- Retrieves the front or rear element. For the rear, the `rear - 1` is used, wrapped around with modulo.

#### `isEmpty` and `isFull`

- Similar to C++, the deque is empty when `front == rear` and full when `(rear + 1) % size == front`.

---

## JavaScript Implementation

### Key Concepts

1. **Array Representation**:
    - An array (`deque[]`) is used to represent the deque with an extra space (`k + 1`).

2. **Circular Pointer Movement**:
    - The `front` and `rear` pointers are managed using modulo operations for circular movement.

### Methods

#### Constructor

- Initializes the array with size `k + 1` and both `front` and `rear` pointers at 0.

#### `insertFront`

- Checks if the deque is full. Moves the `front` pointer backward in a circular manner and inserts the value.

#### `insertLast`

- Inserts the value at the current `rear` position and moves the `rear` pointer forward.

#### `deleteFront`

- Moves the `front` pointer forward to remove the front element.

#### `deleteLast`

- Moves the `rear` pointer backward to remove the last element.

#### `getFront` and `getRear`

- Retrieves the front or rear element, with circular management of the `rear - 1` for `getRear`.

#### `isEmpty` and `isFull`

- Similar logic as before, empty when `front == rear` and full when the next position of `rear` matches `front`.

---

## Python Implementation

### Key Concepts

1. **Array Representation**:
    - A list (`deque[]`) of size `k + 1` is used to represent the deque.

2. **Pointer Management**:
    - `front` and `rear` pointers are used to track the start and end of the deque, and they move circularly using modulo operations.

### Methods

#### Constructor

- Initializes the deque with size `k + 1` and sets `front` and `rear` pointers to 0.

#### `insertFront`

- Moves the `front` pointer backward and inserts the value.

#### `insertLast`

- Inserts a value at the current `rear` position and moves the `rear` pointer forward.

#### `deleteFront`

- Moves the `front` pointer forward to remove the front element.

#### `deleteLast`

- Moves the `rear` pointer backward to remove the last element.

#### `getFront` and `getRear`

- Retrieves the front or rear element, adjusting the `rear - 1` using modulo for `getRear`.

#### `isEmpty` and `isFull`

- Same logic as other languages, empty when `front == rear`, and full when `(rear + 1) % size == front`.

---

## Conclusion

The Circular Deque implementation across C++, Java, JavaScript, and Python shares the same key concepts but is expressed differently based on the syntax and data structures available in each language. By utilizing circular pointers and extra space, these implementations ensure efficient handling of the deque operations.
