#include <vector>
#include <queue>
#include <algorithm>
using namespace std;

class Solution
{
public:
    int minGroups(vector<vector<int>> &intervals)
    {
        // Sort intervals by their start time
        sort(intervals.begin(), intervals.end());

        // Min-heap to track the end times of intervals in groups
        priority_queue<int, vector<int>, greater<int>> pq;

        // Traverse through all intervals
        for (const auto &interval : intervals)
        {
            int start = interval[0], end = interval[1];

            // If the earliest end time is less than the current start, reuse that group
            if (!pq.empty() && pq.top() < start)
            {
                pq.pop();
            }

            // Push the end time of the current interval into the heap
            pq.push(end);
        }

        // The size of the heap at the end is the number of groups
        return pq.size();
    }
};
