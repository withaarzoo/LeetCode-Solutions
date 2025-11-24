class Solution
{
public:
    vector<bool> prefixesDivBy5(vector<int> &nums)
    {
        vector<bool> ans;
        ans.reserve(nums.size()); // avoid reallocation, small optimization

        int rem = 0; // remainder of current prefix modulo 5

        for (int bit : nums)
        {
            // Shift left in binary: multiply by 2 and add current bit
            rem = (rem * 2 + bit) % 5; // keep only remainder to avoid large numbers

            // If remainder is 0, the current prefix is divisible by 5
            ans.push_back(rem == 0);
        }

        return ans;
    }
};
