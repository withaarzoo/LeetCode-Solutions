class Solution
{
public:
    int minimumDistance(vector<int> &nums)
    {
        unordered_map<int, vector<int>> pos;

        // Store all indices for each value
        for (int i = 0; i < nums.size(); i++)
        {
            pos[nums[i]].push_back(i);
        }

        int ans = INT_MAX;

        // Process each value's index list
        for (auto &entry : pos)
        {
            vector<int> &indices = entry.second;

            // Need at least 3 occurrences
            if (indices.size() < 3)
                continue;

            // Check every consecutive group of 3 indices
            for (int i = 0; i + 2 < indices.size(); i++)
            {
                int distance = 2 * (indices[i + 2] - indices[i]);
                ans = min(ans, distance);
            }
        }

        return ans == INT_MAX ? -1 : ans;
    }
};