class MyCircularDeque {
    // Array to hold the elements of the deque
    private int[] deque;
    // Pointers to the front and rear of the deque, and size to manage capacity
    private int front, rear, size;

    /**
     * Constructor to initialize the deque with a given capacity.
     * 
     * @param k The maximum number of elements the deque can hold.
     */
    public MyCircularDeque(int k) {
        // We allocate size k+1 to manage the circular nature (1 extra space for
        // distinction between full and empty)
        deque = new int[k + 1];
        front = 0; // Front pointer starts at the 0th position
        rear = 0; // Rear pointer starts at the 0th position
        size = k + 1; // Total size is k+1 to differentiate between full and empty
    }

    /**
     * Inserts an element at the front of the deque.
     * 
     * @param value The value to be inserted.
     * @return true if insertion is successful, false if deque is full.
     */
    public boolean insertFront(int value) {
        // Check if deque is full
        if (isFull())
            return false;
        // Move the front pointer backward in a circular manner
        front = (front - 1 + size) % size;
        // Place the value at the new front position
        deque[front] = value;
        return true;
    }

    /**
     * Inserts an element at the rear of the deque.
     * 
     * @param value The value to be inserted.
     * @return true if insertion is successful, false if deque is full.
     */
    public boolean insertLast(int value) {
        // Check if deque is full
        if (isFull())
            return false;
        // Insert the value at the current rear position
        deque[rear] = value;
        // Move the rear pointer forward in a circular manner
        rear = (rear + 1) % size;
        return true;
    }

    /**
     * Deletes an element from the front of the deque.
     * 
     * @return true if deletion is successful, false if deque is empty.
     */
    public boolean deleteFront() {
        // Check if deque is empty
        if (isEmpty())
            return false;
        // Move the front pointer forward in a circular manner
        front = (front + 1) % size;
        return true;
    }

    /**
     * Deletes an element from the rear of the deque.
     * 
     * @return true if deletion is successful, false if deque is empty.
     */
    public boolean deleteLast() {
        // Check if deque is empty
        if (isEmpty())
            return false;
        // Move the rear pointer backward in a circular manner
        rear = (rear - 1 + size) % size;
        return true;
    }

    /**
     * Gets the front element of the deque.
     * 
     * @return The front element or -1 if the deque is empty.
     */
    public int getFront() {
        // Check if deque is empty
        if (isEmpty())
            return -1;
        // Return the element at the front pointer
        return deque[front];
    }

    /**
     * Gets the rear element of the deque.
     * 
     * @return The rear element or -1 if the deque is empty.
     */
    public int getRear() {
        // Check if deque is empty
        if (isEmpty())
            return -1;
        // Return the element just before the rear pointer (rear - 1) in a circular
        // manner
        return deque[(rear - 1 + size) % size];
    }

    /**
     * Checks if the deque is empty.
     * 
     * @return true if deque is empty, false otherwise.
     */
    public boolean isEmpty() {
        // Deque is empty when the front and rear pointers are equal
        return front == rear;
    }

    /**
     * Checks if the deque is full.
     * 
     * @return true if deque is full, false otherwise.
     */
    public boolean isFull() {
        // Deque is full when moving the rear pointer forward brings it to the front
        // position
        return (rear + 1) % size == front;
    }
}
