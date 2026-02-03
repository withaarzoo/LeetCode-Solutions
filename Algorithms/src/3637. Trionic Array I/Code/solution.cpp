class Solution
{
public:
    bool isTrionic(vector<int> &nums)
    {
        int n = nums.size();
        int i = 0;

        // 1) strictly increasing
        while (i + 1 < n && nums[i] < nums[i + 1])
        {
            i++;
        }
        if (i == 0 || i == n - 1)
            return false;

        // 2) strictly decreasing
        int mid = i;
        while (i + 1 < n && nums[i] > nums[i + 1])
        {
            i++;
        }
        if (i == mid || i == n - 1)
            return false;

        // 3) strictly increasing again
        while (i + 1 < n && nums[i] < nums[i + 1])
        {
            i++;
        }

        return i == n - 1;
    }
};
