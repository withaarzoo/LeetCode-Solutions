class Solution
{
public:
    int minimumDistance(vector<int> &nums)
    {
        unordered_map<int, vector<int>> positions;

        // Store all indices for each value
        for (int i = 0; i < nums.size(); i++)
        {
            positions[nums[i]].push_back(i);
        }

        int ans = INT_MAX;

        // Check every value's index list
        for (auto &entry : positions)
        {
            vector<int> &idx = entry.second;

            // Need at least 3 occurrences
            if (idx.size() < 3)
                continue;

            // Check every consecutive group of 3 indices
            for (int i = 0; i + 2 < idx.size(); i++)
            {
                int distance = 2 * (idx[i + 2] - idx[i]);
                ans = min(ans, distance);
            }
        }

        return (ans == INT_MAX) ? -1 : ans;
    }
};