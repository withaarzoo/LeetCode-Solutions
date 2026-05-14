class Solution
{
public:
    bool isGood(vector<int> &nums)
    {
        // Sort the array so numbers come in order
        sort(nums.begin(), nums.end());

        // Total number of elements
        int n = nums.size();

        // Maximum element after sorting
        int mx = nums[n - 1];

        // A good array must have size = mx + 1
        if (n != mx + 1)
            return false;

        // Check first n-1 elements
        // They should be 1, 2, 3, ..., mx
        for (int i = 0; i < n - 1; i++)
        {

            // Expected value at current position
            if (nums[i] != i + 1)
                return false;
        }

        // Last element must also be mx
        return nums[n - 1] == mx;
    }
};