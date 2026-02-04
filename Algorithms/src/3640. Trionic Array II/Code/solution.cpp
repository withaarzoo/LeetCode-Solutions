class Solution
{
public:
    vector<tuple<int, int, long long>> buildDecreasingParts(vector<int> &nums)
    {
        int n = nums.size();
        vector<tuple<int, int, long long>> parts;

        int l = 0;
        long long sum = nums[0];

        for (int i = 1; i < n; i++)
        {
            if (nums[i - 1] <= nums[i])
            {
                parts.push_back({l, i - 1, sum});
                l = i;
                sum = 0;
            }
            sum += nums[i];
        }
        parts.push_back({l, n - 1, sum});
        return parts;
    }

    long long maxSumTrionic(vector<int> &nums)
    {
        int n = nums.size();

        vector<long long> maxEndingAt(n), maxStartingAt(n);

        for (int i = 0; i < n; i++)
        {
            maxEndingAt[i] = nums[i];
            if (i > 0 && nums[i - 1] < nums[i] && maxEndingAt[i - 1] > 0)
            {
                maxEndingAt[i] += maxEndingAt[i - 1];
            }
        }

        for (int i = n - 1; i >= 0; i--)
        {
            maxStartingAt[i] = nums[i];
            if (i + 1 < n && nums[i] < nums[i + 1] && maxStartingAt[i + 1] > 0)
            {
                maxStartingAt[i] += maxStartingAt[i + 1];
            }
        }

        auto parts = buildDecreasingParts(nums);
        long long ans = LLONG_MIN;

        for (auto &[p, q, s] : parts)
        {
            if (p > 0 && q < n - 1 &&
                nums[p - 1] < nums[p] &&
                nums[q] < nums[q + 1] &&
                p < q)
            {
                ans = max(ans, maxEndingAt[p - 1] + s + maxStartingAt[q + 1]);
            }
        }
        return ans;
    }
};
