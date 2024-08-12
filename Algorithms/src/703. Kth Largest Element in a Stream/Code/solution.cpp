#include <queue>
#include <vector>

class KthLargest
{
public:
    // Constructor to initialize the KthLargest object
    // Takes an integer k and a vector of integers nums as parameters
    KthLargest(int k, std::vector<int> &nums) : k(k)
    {
        // Iterate over each element in nums and add it to the minHeap using the add() function
        for (int num : nums)
        {
            add(num);
        }
    }

    // Function to add a new value to the Kth largest tracker
    // Returns the k-th largest element after adding the new value
    int add(int val)
    {
        // If the size of the minHeap is less than k, simply add the value to the heap
        if (minHeap.size() < k)
        {
            minHeap.push(val);
        }
        // If the size of the minHeap is equal to k and the new value is greater than
        // the smallest element in the heap (the top of the heap), replace the smallest element
        // with the new value to maintain the k largest elements in the heap
        else if (val > minHeap.top())
        {
            minHeap.pop();     // Remove the smallest element from the heap
            minHeap.push(val); // Add the new value to the heap
        }
        // Return the k-th largest element, which is the smallest element in the heap
        return minHeap.top();
    }

private:
    int k; // The k-th largest element to track
    // A min-heap (priority_queue with greater<int> comparator) to store the k largest elements
    // The smallest element in the heap (top) represents the k-th largest element in the stream
    std::priority_queue<int, std::vector<int>, std::greater<int>> minHeap;
};
