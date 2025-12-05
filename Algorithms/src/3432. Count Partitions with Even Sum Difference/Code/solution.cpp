class Solution
{
public:
    int countPartitions(vector<int> &nums)
    {
        long long total = 0;
        // Compute the total sum of the array
        for (int x : nums)
        {
            total += x;
        }

        // If total sum is odd, no valid partition
        if (total % 2 != 0)
            return 0;

        // If total is even, every position between elements is a valid partition
        int n = (int)nums.size();
        return n - 1;
    }
};
