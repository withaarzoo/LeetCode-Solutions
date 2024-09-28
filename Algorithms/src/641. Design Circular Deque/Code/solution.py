class MyCircularDeque:

    def __init__(self, k: int):
        """
        Initialize the circular deque with a fixed capacity.
        
        Args:
        k (int): The size of the deque.
        
        The deque is implemented using a fixed-size array of length (k + 1) to differentiate
        between an empty and a full deque by leaving one extra space.
        """
        # Initialize the deque with a size of (k + 1) to account for the empty space needed to distinguish full/empty.
        self.deque = [0] * (k + 1)
        # Initialize front and rear pointers to 0, representing an empty deque.
        self.front = 0
        self.rear = 0
        # Store the size (k + 1) of the deque to help with modulo operations.
        self.size = k + 1

    def insertFront(self, value: int) -> bool:
        """
        Inserts an item at the front of the deque.
        
        Args:
        value (int): The value to insert.
        
        Returns:
        bool: True if the insertion is successful, False if the deque is full.
        """
        if self.isFull():
            return False
        # Move the front pointer backwards (circularly), and place the value at the new front.
        self.front = (self.front - 1 + self.size) % self.size
        self.deque[self.front] = value
        return True

    def insertLast(self, value: int) -> bool:
        """
        Inserts an item at the rear of the deque.
        
        Args:
        value (int): The value to insert.
        
        Returns:
        bool: True if the insertion is successful, False if the deque is full.
        """
        if self.isFull():
            return False
        # Place the value at the current rear, then move the rear pointer forward (circularly).
        self.deque[self.rear] = value
        self.rear = (self.rear + 1) % self.size
        return True

    def deleteFront(self) -> bool:
        """
        Deletes an item from the front of the deque.
        
        Returns:
        bool: True if the deletion is successful, False if the deque is empty.
        """
        if self.isEmpty():
            return False
        # Move the front pointer forward (circularly) to effectively remove the front item.
        self.front = (self.front + 1) % self.size
        return True

    def deleteLast(self) -> bool:
        """
        Deletes an item from the rear of the deque.
        
        Returns:
        bool: True if the deletion is successful, False if the deque is empty.
        """
        if self.isEmpty():
            return False
        # Move the rear pointer backward (circularly) to effectively remove the rear item.
        self.rear = (self.rear - 1 + self.size) % self.size
        return True

    def getFront(self) -> int:
        """
        Gets the front item of the deque.
        
        Returns:
        int: The value at the front of the deque, or -1 if the deque is empty.
        """
        if self.isEmpty():
            return -1
        # Return the value at the front pointer.
        return self.deque[self.front]

    def getRear(self) -> int:
        """
        Gets the rear item of the deque.
        
        Returns:
        int: The value at the rear of the deque, or -1 if the deque is empty.
        """
        if self.isEmpty():
            return -1
        # Return the value at the rear pointer. Since rear points to the next insertion index, we use (rear - 1).
        return self.deque[(self.rear - 1 + self.size) % self.size]

    def isEmpty(self) -> bool:
        """
        Checks whether the deque is empty.
        
        Returns:
        bool: True if the deque is empty, False otherwise.
        """
        # Deque is empty if the front and rear pointers are at the same position.
        return self.front == self.rear

    def isFull(self) -> bool:
        """
        Checks whether the deque is full.
        
        Returns:
        bool: True if the deque is full, False otherwise.
        """
        # Deque is full if the next position of rear equals front (circularly).
        return (self.rear + 1) % self.size == self.front
