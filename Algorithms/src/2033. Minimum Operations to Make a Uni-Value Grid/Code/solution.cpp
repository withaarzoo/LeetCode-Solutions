class Solution
{
public:
    int minOperations(vector<vector<int>> &grid, int x)
    {
        vector<int> nums;

        // Step 1: Flatten grid
        for (auto &row : grid)
        {
            for (int val : row)
            {
                nums.push_back(val);
            }
        }

        // Step 2: Check feasibility
        int rem = nums[0] % x;
        for (int num : nums)
        {
            if (num % x != rem)
                return -1;
        }

        // Step 3: Sort
        sort(nums.begin(), nums.end());

        // Step 4: Median
        int median = nums[nums.size() / 2];

        // Step 5: Calculate operations
        int ops = 0;
        for (int num : nums)
        {
            ops += abs(num - median) / x;
        }

        return ops;
    }
};