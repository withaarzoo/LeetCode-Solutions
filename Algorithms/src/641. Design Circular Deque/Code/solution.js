class MyCircularDeque {
  // Constructor that initializes a deque with a capacity of k
  constructor(k) {
    // Create an array of size k+1 to differentiate between full and empty cases
    this.deque = new Array(k + 1); // The extra 1 is used for distinguishing full from empty
    this.front = 0; // Front pointer to track the beginning of the deque
    this.rear = 0; // Rear pointer to track the end of the deque
    this.size = k + 1; // Actual size of the circular deque (k + 1 to account for the extra space)
  }

  // Insert an element at the front of the deque
  insertFront(value) {
    // Check if the deque is full, return false if it is
    if (this.isFull()) return false;
    // Move the front pointer back in a circular manner
    // (this.front - 1 + this.size) % this.size ensures it wraps around correctly
    this.front = (this.front - 1 + this.size) % this.size;
    // Place the value at the new front position
    this.deque[this.front] = value;
    return true; // Successfully inserted
  }

  // Insert an element at the rear of the deque
  insertLast(value) {
    // Check if the deque is full, return false if it is
    if (this.isFull()) return false;
    // Place the value at the current rear position
    this.deque[this.rear] = value;
    // Move the rear pointer forward in a circular manner
    // This ensures that the rear wraps around correctly when reaching the end
    this.rear = (this.rear + 1) % this.size;
    return true; // Successfully inserted
  }

  // Delete an element from the front of the deque
  deleteFront() {
    // Check if the deque is empty, return false if it is
    if (this.isEmpty()) return false;
    // Move the front pointer forward in a circular manner
    // This essentially "removes" the front element by moving the pointer
    this.front = (this.front + 1) % this.size;
    return true; // Successfully deleted
  }

  // Delete an element from the rear of the deque
  deleteLast() {
    // Check if the deque is empty, return false if it is
    if (this.isEmpty()) return false;
    // Move the rear pointer back in a circular manner
    // This removes the rear element by moving the pointer
    this.rear = (this.rear - 1 + this.size) % this.size;
    return true; // Successfully deleted
  }

  // Get the front element of the deque
  getFront() {
    // Check if the deque is empty, return -1 if it is
    if (this.isEmpty()) return -1;
    // Return the value at the front pointer
    return this.deque[this.front];
  }

  // Get the rear element of the deque
  getRear() {
    // Check if the deque is empty, return -1 if it is
    if (this.isEmpty()) return -1;
    // Return the value at the rear pointer minus one (since rear points to the next empty space)
    // (this.rear - 1 + this.size) % this.size handles the circular nature of the deque
    return this.deque[(this.rear - 1 + this.size) % this.size];
  }

  // Check if the deque is empty
  isEmpty() {
    // The deque is empty when the front and rear pointers are equal
    return this.front === this.rear;
  }

  // Check if the deque is full
  isFull() {
    // The deque is full when the next position of rear (rear + 1) equals the front
    return (this.rear + 1) % this.size === this.front;
  }
}
