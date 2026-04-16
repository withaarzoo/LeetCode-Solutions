class Solution
{
public:
    vector<int> solveQueries(vector<int> &nums, vector<int> &queries)
    {
        int n = nums.size();

        // Store all indices for every value
        unordered_map<int, vector<int>> positions;

        for (int i = 0; i < n; i++)
        {
            positions[nums[i]].push_back(i);
        }

        // answer[i] = minimum circular distance for index i
        vector<int> answer(n, -1);

        // Process each group of equal values
        for (auto &entry : positions)
        {
            vector<int> &pos = entry.second;
            int m = pos.size();

            // If value appears only once, answer remains -1
            if (m == 1)
                continue;

            for (int i = 0; i < m; i++)
            {
                int curr = pos[i];

                // Previous and next occurrence in circular order
                int prev = pos[(i - 1 + m) % m];
                int next = pos[(i + 1) % m];

                // Distance to previous occurrence
                int distPrev = abs(curr - prev);
                distPrev = min(distPrev, n - distPrev);

                // Distance to next occurrence
                int distNext = abs(curr - next);
                distNext = min(distNext, n - distNext);

                // Best answer for current index
                answer[curr] = min(distPrev, distNext);
            }
        }

        // Build final result for queries
        vector<int> result;
        for (int idx : queries)
        {
            result.push_back(answer[idx]);
        }

        return result;
    }
};