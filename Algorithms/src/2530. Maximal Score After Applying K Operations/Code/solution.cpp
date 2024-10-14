#include <queue>
#include <cmath>
#include <vector>

class Solution
{
public:
    long long maxKelements(std::vector<int> &nums, int k)
    {
        // Max-heap to store the numbers (priority_queue is by default a max-heap)
        std::priority_queue<long long> maxHeap;

        // Push all elements into the heap
        for (int num : nums)
        {
            maxHeap.push(num);
        }

        long long score = 0;

        // Perform k operations
        for (int i = 0; i < k; ++i)
        {
            long long maxVal = maxHeap.top();
            maxHeap.pop();

            // Add the largest value to the score
            score += maxVal;

            // Replace the number with ceil(maxVal / 3)
            maxHeap.push((maxVal + 2) / 3); // Using (maxVal + 2) / 3 to simulate ceil
        }

        return score;
    }
};