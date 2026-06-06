class Solution
{
public:
    vector<int> leftRightDifference(vector<int> &nums)
    {

        int n = nums.size();

        // Store the sum of all elements initially
        int rightSum = 0;
        for (int num : nums)
        {
            rightSum += num;
        }

        // Sum of elements on the left side
        int leftSum = 0;

        // Result array
        vector<int> ans(n);

        for (int i = 0; i < n; i++)
        {

            // Remove current element so rightSum becomes
            // the sum of elements strictly to the right
            rightSum -= nums[i];

            // Store absolute difference
            ans[i] = abs(leftSum - rightSum);

            // Add current element to leftSum for next index
            leftSum += nums[i];
        }

        return ans;
    }
};