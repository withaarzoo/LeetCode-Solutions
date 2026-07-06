class Solution {
public:
    int removeCoveredIntervals(vector<vector<int>>& intervals) {
        // Sort by start in ascending order.
        // For the same start, place the larger end first
        // so the covering interval is processed before the covered one.
        sort(intervals.begin(), intervals.end(),
             [](const vector<int>& a, const vector<int>& b) {
                 if (a[0] == b[0]) {
                     return a[1] > b[1];
                 }
                 return a[0] < b[0];
             });

        // Every valid end is greater than 0, so -1 is a safe initial value.
        int maxEnd = -1;

        // This stores how many intervals are not covered.
        int remaining = 0;

        // Check every interval in the sorted order.
        for (const auto& interval : intervals) {
            // A larger end means no earlier interval can cover this one.
            if (interval[1] > maxEnd) {
                remaining++;

                // Remember the farthest ending point seen so far.
                maxEnd = interval[1];
            }
            // If interval[1] <= maxEnd, an earlier interval covers it,
            // so I do not count it.
        }

        // Return the number of intervals left after removing covered ones.
        return remaining;
    }
}; 