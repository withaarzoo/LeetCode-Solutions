class Solution
{
public:
    vector<long long> distance(vector<int> &nums)
    {
        int n = nums.size();
        unordered_map<int, vector<int>> mp;

        // Step 1: Group indices by value
        for (int i = 0; i < n; i++)
        {
            mp[nums[i]].push_back(i);
        }

        vector<long long> res(n, 0);

        // Step 2: Process each group
        for (auto &it : mp)
        {
            vector<int> &idx = it.second;
            int k = idx.size();

            long long prefixSum = 0;
            long long totalSum = 0;

            // Total sum of indices
            for (int x : idx)
                totalSum += x;

            for (int i = 0; i < k; i++)
            {
                long long curr = idx[i];

                long long left = curr * i - prefixSum;
                long long right = (totalSum - prefixSum - curr) - curr * (k - i - 1);

                res[curr] = left + right;

                prefixSum += curr;
            }
        }

        return res;
    }
};