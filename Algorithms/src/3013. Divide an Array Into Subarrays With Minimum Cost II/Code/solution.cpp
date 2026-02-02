class Solution
{
public:
    long long minimumCost(vector<int> &nums, int k, int dist)
    {
        // target is the number of additional elements we need to pick from the window
        int target = k - 2;
        long long current_window_sum = 0;

        // L stores the smallest 'target' elements
        // R stores the rest
        multiset<int> L, R;

        // Initialize the window for the first pivot i = 1
        // The window covers nums[2] ... nums[dist + 1]
        for (int j = 2; j <= min((int)nums.size() - 1, dist + 1); ++j)
        {
            L.insert(nums[j]);
            current_window_sum += nums[j];
        }

        // Balance function to maintain L size == target and order property
        auto balance = [&]()
        {
            // If L is too large, move largest to R
            while (L.size() > target)
            {
                int x = *L.rbegin();
                L.erase(prev(L.end()));
                current_window_sum -= x;
                R.insert(x);
            }
            // If L is too small and R has elements, move smallest from R to L
            while (L.size() < target && !R.empty())
            {
                int x = *R.begin();
                R.erase(R.begin());
                L.insert(x);
                current_window_sum += x;
            }
            // Ensure max(L) <= min(R)
            while (!L.empty() && !R.empty() && *L.rbegin() > *R.begin())
            {
                int l_val = *L.rbegin();
                int r_val = *R.begin();
                L.erase(prev(L.end()));
                R.erase(R.begin());
                L.insert(r_val);
                R.insert(l_val);
                current_window_sum = current_window_sum - l_val + r_val;
            }
        };

        balance();

        long long min_cost = (long long)nums[0] + nums[1] + current_window_sum;

        // Slide the window
        // i is the index of the 2nd subarray's start
        for (int i = 2; i <= nums.size() - (k - 1); ++i)
        {
            // Element leaving the window: nums[i]
            // It was a candidate, now it becomes the fixed pivot
            int out_val = nums[i];
            auto it = L.find(out_val);
            if (it != L.end())
            {
                L.erase(it);
                current_window_sum -= out_val;
            }
            else
            {
                R.erase(R.find(out_val));
            }

            // Element entering the window: nums[i + dist]
            if (i + dist < nums.size())
            {
                int in_val = nums[i + dist];
                L.insert(in_val);
                current_window_sum += in_val;
            }

            balance();

            long long current_cost = (long long)nums[0] + nums[i] + current_window_sum;
            min_cost = min(min_cost, current_cost);
        }

        return min_cost;
    }
};