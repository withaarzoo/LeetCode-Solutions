class Solution
{
public:
    int waysToSplitArray(vector<int> &nums)
    {
        long long total_sum = 0; // Total sum of the array
        for (int num : nums)
        {
            total_sum += num;
        }

        long long prefix_sum = 0; // Prefix sum
        int count = 0;            // Count of valid splits

        for (int i = 0; i < nums.size() - 1; i++)
        {
            prefix_sum += nums[i];
            long long right_sum = total_sum - prefix_sum;
            if (prefix_sum >= right_sum)
            {
                count++;
            }
        }

        return count;
    }
};
