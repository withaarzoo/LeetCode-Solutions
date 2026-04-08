class Solution
{
public:
    int xorAfterQueries(vector<int> &nums, vector<vector<int>> &queries)
    {
        const long long MOD = 1e9 + 7;

        // Process each query
        for (auto &q : queries)
        {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            // Visit indices: l, l+k, l+2k, ... <= r
            for (int i = l; i <= r; i += k)
            {
                nums[i] = (1LL * nums[i] * v) % MOD;
            }
        }

        // Compute XOR of all final values
        int ans = 0;
        for (int num : nums)
        {
            ans ^= num;
        }

        return ans;
    }
};