class Solution
{
public:
    int minPairSum(vector<int> &nums)
    {
        // Step 1: Sort the array
        sort(nums.begin(), nums.end());

        int left = 0;
        int right = nums.size() - 1;
        int maxPairSum = 0;

        // Step 2: Pair smallest with largest
        while (left < right)
        {
            int pairSum = nums[left] + nums[right];
            maxPairSum = max(maxPairSum, pairSum);
            left++;
            right--;
        }

        return maxPairSum;
    }
};
