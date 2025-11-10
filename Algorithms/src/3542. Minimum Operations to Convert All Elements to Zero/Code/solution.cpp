class Solution
{
public:
    int minOperations(vector<int> &nums)
    {
        vector<int> stk; // non-decreasing stack of heights
        int ans = 0;
        for (int x : nums)
        {
            // Drop to the new lower height by popping taller ones
            while (!stk.empty() && stk.back() > x)
                stk.pop_back();
            if (x == 0)
                continue; // zero adds no new layer
            if (stk.empty() || stk.back() < x)
            {
                // New rise → one more operation needed
                ans++;
                stk.push_back(x);
            }
            // if stk.back() == x, nothing to do
        }
        return ans;
    }
};
