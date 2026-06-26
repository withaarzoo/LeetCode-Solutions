class Solution
{
public:
    // Fenwick Tree for prefix frequency counting
    class Fenwick
    {
    public:
        vector<int> bit;

        Fenwick(int n)
        {
            bit.assign(n + 1, 0);
        }

        // Add one occurrence at index
        void update(int idx)
        {
            while (idx < bit.size())
            {
                bit[idx]++;
                idx += idx & -idx;
            }
        }

        // Count frequencies from 1...idx
        long long query(int idx)
        {
            long long sum = 0;
            while (idx > 0)
            {
                sum += bit[idx];
                idx -= idx & -idx;
            }
            return sum;
        }
    };

    long long countMajoritySubarrays(vector<int> &nums, int target)
    {

        int n = nums.size();

        // Build prefix sums after converting target -> +1, others -> -1
        vector<int> pref(n + 1, 0);
        for (int i = 0; i < n; i++)
        {
            pref[i + 1] = pref[i] + (nums[i] == target ? 1 : -1);
        }

        // Coordinate compression because prefix sums may be negative
        vector<int> values = pref;
        sort(values.begin(), values.end());
        values.erase(unique(values.begin(), values.end()), values.end());

        Fenwick ft(values.size());

        long long ans = 0;

        // Process every prefix sum
        for (int x : pref)
        {

            // Compressed index (1-based)
            int idx = lower_bound(values.begin(), values.end(), x) - values.begin() + 1;

            // Count previous prefix sums strictly smaller than current
            ans += ft.query(idx - 1);

            // Insert current prefix sum
            ft.update(idx);
        }

        return ans;
    }
};