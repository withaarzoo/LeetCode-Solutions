class Solution
{
public:
    int minRemoval(vector<int> &nums, int k)
    {
        sort(nums.begin(), nums.end());

        int n = nums.size();
        int left = 0;
        int maxKeep = 1;

        for (int right = 0; right < n; right++)
        {
            // Shrink window until condition becomes valid
            while ((long long)nums[right] > (long long)nums[left] * k)
            {
                left++;
            }
            maxKeep = max(maxKeep, right - left + 1);
        }

        return n - maxKeep;
    }
};
