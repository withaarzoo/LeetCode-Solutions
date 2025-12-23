class Solution
{
public:
    int maxTwoEvents(vector<vector<int>> &events)
    {
        // Sort events by start time
        sort(events.begin(), events.end());

        // Create another array sorted by end time
        vector<vector<int>> endSorted = events;
        sort(endSorted.begin(), endSorted.end(),
             [](auto &a, auto &b)
             {
                 return a[1] < b[1];
             });

        int n = events.size();
        vector<int> maxValueTill(n);

        // Prefix max of values based on end time
        maxValueTill[0] = endSorted[0][2];
        for (int i = 1; i < n; i++)
        {
            maxValueTill[i] = max(maxValueTill[i - 1], endSorted[i][2]);
        }

        int ans = 0;
        int j = 0;

        // Traverse events by start time
        for (int i = 0; i < n; i++)
        {
            int start = events[i][0];
            int value = events[i][2];

            // Move pointer for events that end before current start
            while (j < n && endSorted[j][1] < start)
            {
                j++;
            }

            // Option 1: take only this event
            ans = max(ans, value);

            // Option 2: take this + best previous
            if (j > 0)
            {
                ans = max(ans, value + maxValueTill[j - 1]);
            }
        }

        return ans;
    }
};
