class Solution
{
public:
    bool kLengthApart(vector<int> &nums, int k)
    {
        int prev = -1; // index of previous 1 (none seen yet)
        for (int i = 0; i < (int)nums.size(); ++i)
        {
            if (nums[i] == 1)
            {
                if (prev != -1)
                {
                    // zeros between current and previous 1 = i - prev - 1
                    if (i - prev - 1 < k)
                        return false;
                }
                prev = i;
            }
        }
        return true;
    }
};
