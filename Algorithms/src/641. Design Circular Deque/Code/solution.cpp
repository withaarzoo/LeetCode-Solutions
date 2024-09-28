class MyCircularDeque
{
private:
    vector<int> deque; // Vector to represent the deque as an array
    int front, rear;   // Pointers to track the front and rear of the deque
    int size;          // Total size of the deque including one extra space for full/empty differentiation

public:
    // Constructor to initialize the deque with size k
    MyCircularDeque(int k)
    {
        // Resize the deque to have k + 1 spaces to handle the "full" state easily
        deque.resize(k + 1);
        front = 0;    // Initially, both front and rear are set to 0
        rear = 0;     // This means the deque is empty
        size = k + 1; // The size of the deque is k+1 to differentiate full vs empty states
    }

    // Insert an item at the front of the deque
    bool insertFront(int value)
    {
        // If the deque is full, we cannot insert at the front
        if (isFull())
            return false;

        // Move the front pointer backwards in a circular manner
        // This ensures that when the front reaches 0, it wraps around to the end
        front = (front - 1 + size) % size;

        // Insert the value at the new front position
        deque[front] = value;

        return true;
    }

    // Insert an item at the rear of the deque
    bool insertLast(int value)
    {
        // If the deque is full, we cannot insert at the rear
        if (isFull())
            return false;

        // Insert the value at the current rear position
        deque[rear] = value;

        // Move the rear pointer forward in a circular manner
        // This ensures that when the rear reaches the end, it wraps around to 0
        rear = (rear + 1) % size;

        return true;
    }

    // Delete an item from the front of the deque
    bool deleteFront()
    {
        // If the deque is empty, there is nothing to delete
        if (isEmpty())
            return false;

        // Move the front pointer forward in a circular manner
        // This removes the front element by advancing the front pointer
        front = (front + 1) % size;

        return true;
    }

    // Delete an item from the rear of the deque
    bool deleteLast()
    {
        // If the deque is empty, there is nothing to delete
        if (isEmpty())
            return false;

        // Move the rear pointer backward in a circular manner
        // This ensures that when the rear reaches 0, it wraps around to the end
        rear = (rear - 1 + size) % size;

        return true;
    }

    // Get the front item of the deque
    int getFront()
    {
        // If the deque is empty, return -1 to indicate no element at the front
        if (isEmpty())
            return -1;

        // Return the element at the front pointer
        return deque[front];
    }

    // Get the rear item of the deque
    int getRear()
    {
        // If the deque is empty, return -1 to indicate no element at the rear
        if (isEmpty())
            return -1;

        // The rear pointer points to the next position, so we subtract 1 to get the actual last element
        // We use (rear - 1 + size) % size to handle the wrap-around in a circular manner
        return deque[(rear - 1 + size) % size];
    }

    // Check if the deque is empty
    bool isEmpty()
    {
        // The deque is empty if the front and rear pointers are at the same position
        return front == rear;
    }

    // Check if the deque is full
    bool isFull()
    {
        // The deque is full if the next position of rear equals front
        // This is why we have an extra space in the deque
        return (rear + 1) % size == front;
    }
};
