class Solution
{
public:
    int minimumCost(vector<int> &nums)
    {
        // nums[0] must be included
        int first = nums[0];

        // Sort the remaining elements
        sort(nums.begin() + 1, nums.end());

        // Pick the two smallest from the rest
        return first + nums[1] + nums[2];
    }
};
