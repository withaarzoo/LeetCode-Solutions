class Solution
{
public:
    int maximumLength(vector<int> &nums)
    {
        // Store frequency of every number
        unordered_map<long long, int> freq;
        for (int x : nums)
            freq[x]++;

        int ans = 1;

        // Handle value 1 separately because 1^2 = 1
        if (freq.count(1))
        {
            int cnt = freq[1];

            // We can only take an odd number of ones
            ans = max(ans, (cnt % 2 == 1) ? cnt : cnt - 1);
        }

        // Try every distinct starting value (>1)
        for (auto &[start, cnt] : freq)
        {
            if (start == 1)
                continue;

            long long cur = start;
            int len = 0;

            while (freq.count(cur))
            {
                // If at least two copies exist, use both
                if (freq[cur] >= 2)
                {
                    len += 2;

                    // Move to the next squared value
                    cur = cur * cur;
                }
                // Only one copy exists, so it becomes the center
                else
                {
                    len++;
                    break;
                }
            }

            // If the length is even, we never found a center
            if (len % 2 == 0)
                len--;

            ans = max(ans, len);
        }

        return ans;
    }
};