class Solution
{
public:
    int minimumDifference(vector<int> &nums, int k)
    {
        // If we need only one score, difference is always 0
        if (k == 1)
            return 0;

        // Step 1: Sort the array
        sort(nums.begin(), nums.end());

        int minDiff = INT_MAX;

        // Step 2: Sliding window of size k
        for (int i = 0; i + k - 1 < nums.size(); i++)
        {
            int diff = nums[i + k - 1] - nums[i];
            minDiff = min(minDiff, diff);
        }

        return minDiff;
    }
};
