import java.util.PriorityQueue;

// The KthLargest class is designed to find the k-th largest element
// in a stream of numbers. It uses a min-heap (a PriorityQueue in Java)
// to efficiently track the k largest elements seen so far.
class KthLargest {

    // Instance variables to store the value of k and the min-heap.
    private int k;
    private PriorityQueue<Integer> minHeap;

    // Constructor to initialize the KthLargest object with the
    // desired k value and an initial array of numbers.
    public KthLargest(int k, int[] nums) {
        this.k = k; // Store the value of k
        // Initialize the min-heap with a capacity of k elements.
        // The min-heap will always contain the k largest elements seen so far.
        this.minHeap = new PriorityQueue<>(k);

        // Add each number from the initial array to the min-heap.
        for (int num : nums) {
            add(num); // Call the add method to maintain the k largest elements.
        }
    }

    // The add method adds a new value to the stream of numbers and returns
    // the k-th largest element after the addition.
    public int add(int val) {
        // If the heap has less than k elements, simply add the new value.
        // Since the heap size is less than k, we haven't yet reached the point
        // where we can determine the k-th largest element.
        if (minHeap.size() < k) {
            minHeap.offer(val); // Add the new value to the heap.
        } else if (val > minHeap.peek()) {
            // If the new value is larger than the smallest value in the heap (root),
            // it means this new value should be part of the k largest elements.
            // Remove the smallest value (which is at the root) from the heap
            // and add the new value.
            minHeap.poll(); // Remove the smallest element.
            minHeap.offer(val); // Add the new value to the heap.
        }
        // The root of the min-heap now represents the k-th largest element,
        // because the heap always maintains the k largest elements, with
        // the smallest of these k elements at the root.
        return minHeap.peek();
    }
}
