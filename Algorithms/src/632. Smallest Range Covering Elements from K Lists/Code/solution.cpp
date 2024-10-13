#include <vector>
#include <queue>
#include <climits>

using namespace std;

class Solution
{
public:
    vector<int> smallestRange(vector<vector<int>> &nums)
    {
        int k = nums.size();
        // Min-heap to store (value, list index, element index)
        priority_queue<pair<int, pair<int, int>>, vector<pair<int, pair<int, int>>>, greater<pair<int, pair<int, int>>>> minHeap;
        int maxValue = INT_MIN;

        // Initialize the heap with the first element of each list
        for (int i = 0; i < k; ++i)
        {
            minHeap.push({nums[i][0], {i, 0}});
            maxValue = max(maxValue, nums[i][0]);
        }

        // Initialize result range
        int rangeStart = 0, rangeEnd = INT_MAX;

        // Process until one of the lists is exhausted
        while (!minHeap.empty())
        {
            auto [minValue, pos] = minHeap.top();
            minHeap.pop();
            int row = pos.first, col = pos.second;

            // Update the smallest range
            if (maxValue - minValue < rangeEnd - rangeStart)
            {
                rangeStart = minValue;
                rangeEnd = maxValue;
            }

            // Move to the next element in the current list
            if (col + 1 < nums[row].size())
            {
                minHeap.push({nums[row][col + 1], {row, col + 1}});
                maxValue = max(maxValue, nums[row][col + 1]);
            }
            else
            {
                // One list is exhausted, we stop
                break;
            }
        }

        return {rangeStart, rangeEnd};
    }
};