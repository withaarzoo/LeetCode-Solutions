class Solution
{
public:
    static const int MOD = 1e9 + 7;

    // Fast power: calculates (base ^ exp) % MOD
    long long modPow(long long base, long long exp)
    {
        long long result = 1;

        while (exp > 0)
        {
            if (exp & 1)
            {
                result = (result * base) % MOD;
            }

            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    // Modular inverse using Fermat's theorem
    long long modInverse(long long x)
    {
        return modPow(x, MOD - 2);
    }

    int xorAfterQueries(vector<int> &nums, vector<vector<int>> &queries)
    {
        int n = nums.size();
        int limit = sqrt(n) + 1;

        // Store small-k queries grouped by k
        unordered_map<int, vector<vector<int>>> smallQueries;

        for (auto &q : queries)
        {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            // Large k -> process directly
            if (k >= limit)
            {
                for (int i = l; i <= r; i += k)
                {
                    nums[i] = (1LL * nums[i] * v) % MOD;
                }
            }
            else
            {
                smallQueries[k].push_back(q);
            }
        }

        // Process all grouped small-k queries
        for (auto &[k, group] : smallQueries)
        {
            vector<long long> diff(n, 1);

            for (auto &q : group)
            {
                int l = q[0];
                int r = q[1];
                int v = q[3];

                // Start multiplying from l
                diff[l] = (diff[l] * v) % MOD;

                // Find first position after affected range
                int steps = (r - l) / k;
                int nextPos = l + (steps + 1) * k;

                // Stop multiplication effect after nextPos
                if (nextPos < n)
                {
                    diff[nextPos] = (diff[nextPos] * modInverse(v)) % MOD;
                }
            }

            // Propagate values with jump size k
            for (int i = 0; i < n; i++)
            {
                if (i >= k)
                {
                    diff[i] = (diff[i] * diff[i - k]) % MOD;
                }

                nums[i] = (1LL * nums[i] * diff[i]) % MOD;
            }
        }

        int answer = 0;

        for (int num : nums)
        {
            answer ^= num;
        }

        return answer;
    }
};