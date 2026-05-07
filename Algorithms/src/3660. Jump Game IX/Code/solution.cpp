class Solution
{
public:
    vector<int> maxValue(vector<int> &nums)
    {
        int n = (int)nums.size();

        // suffixMin[i] = smallest value in nums[i...n-1]
        // I keep one extra cell at the end so the last segment can stop cleanly.
        vector<int> suffixMin(n + 1, INT_MAX);
        for (int i = n - 1; i >= 0; --i)
        {
            suffixMin[i] = min(nums[i], suffixMin[i + 1]);
        }

        vector<int> ans(n);
        int l = 0;

        // I build one connected component at a time.
        while (l < n)
        {
            int r = l;
            int componentMax = nums[l];

            // I keep extending this segment while some value on the left
            // is bigger than some value on the right, which means an inversion
            // crosses the cut and the two parts are still connected.
            while (r + 1 < n && componentMax > suffixMin[r + 1])
            {
                ++r;
                componentMax = max(componentMax, nums[r]);
            }

            // Every index in this connected component can reach the same maximum.
            for (int i = l; i <= r; ++i)
            {
                ans[i] = componentMax;
            }

            // Move to the next segment.
            l = r + 1;
        }

        return ans;
    }
};