class Solution
{
public:
    vector<int> constructTransformedArray(vector<int> &nums)
    {
        int n = nums.size();
        vector<int> result(n);

        for (int i = 0; i < n; i++)
        {
            if (nums[i] == 0)
            {
                result[i] = nums[i];
            }
            else
            {
                int move = nums[i];
                int target = (i + move) % n;

                // handle negative index
                if (target < 0)
                    target += n;

                result[i] = nums[target];
            }
        }
        return result;
    }
};
