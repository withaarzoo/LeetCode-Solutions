// Define a structure for the circular deque
type MyCircularDeque struct {
    deque []int   // The deque implemented as a slice (array) of integers
    front, rear, size int  // front: index pointing to the front element, rear: index pointing to the rear element, size: total capacity of the deque
}

// Constructor initializes the deque with a size of k + 1 (one extra space to differentiate between full and empty states)
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        deque: make([]int, k + 1),  // Create an array with size k + 1
        front: 0,                   // Initialize front pointer to 0
        rear: 0,                    // Initialize rear pointer to 0
        size: k + 1,                // Set the total size to k + 1
    }
}

// InsertFront inserts an element at the front of the deque
func (this *MyCircularDeque) InsertFront(value int) bool {
    // Check if the deque is full
    if this.IsFull() {
        return false // Return false if full, cannot insert
    }
    // Move front pointer backwards (circularly) by adjusting it within bounds
    this.front = (this.front - 1 + this.size) % this.size
    // Insert the value at the new front index
    this.deque[this.front] = value
    return true // Return true as insertion is successful
}

// InsertLast inserts an element at the rear of the deque
func (this *MyCircularDeque) InsertLast(value int) bool {
    // Check if the deque is full
    if this.IsFull() {
        return false // Return false if full, cannot insert
    }
    // Insert the value at the current rear index
    this.deque[this.rear] = value
    // Move rear pointer forward (circularly) to the next available position
    this.rear = (this.rear + 1) % this.size
    return true // Return true as insertion is successful
}

// DeleteFront removes an element from the front of the deque
func (this *MyCircularDeque) DeleteFront() bool {
    // Check if the deque is empty
    if this.IsEmpty() {
        return false // Return false if empty, nothing to delete
    }
    // Move front pointer forward (circularly) to remove the current front element
    this.front = (this.front + 1) % this.size
    return true // Return true as deletion is successful
}

// DeleteLast removes an element from the rear of the deque
func (this *MyCircularDeque) DeleteLast() bool {
    // Check if the deque is empty
    if this.IsEmpty() {
        return false // Return false if empty, nothing to delete
    }
    // Move rear pointer backward (circularly) to remove the current rear element
    this.rear = (this.rear - 1 + this.size) % this.size
    return true // Return true as deletion is successful
}

// GetFront returns the front element of the deque without removing it
func (this *MyCircularDeque) GetFront() int {
    // Check if the deque is empty
    if this.IsEmpty() {
        return -1 // Return -1 if empty, no front element exists
    }
    // Return the element at the front index
    return this.deque[this.front]
}

// GetRear returns the rear element of the deque without removing it
func (this *MyCircularDeque) GetRear() int {
    // Check if the deque is empty
    if this.IsEmpty() {
        return -1 // Return -1 if empty, no rear element exists
    }
    // Return the element at the index before the rear pointer (circularly adjusted)
    return this.deque[(this.rear - 1 + this.size) % this.size]
}

// IsEmpty checks whether the deque is empty
func (this *MyCircularDeque) IsEmpty() bool {
    // The deque is empty if the front and rear pointers are at the same position
    return this.front == this.rear
}

// IsFull checks whether the deque is full
func (this *MyCircularDeque) IsFull() bool {
    // The deque is full if the next position of the rear pointer equals the front pointer
    return (this.rear + 1) % this.size == this.front
}
