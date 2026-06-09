class Solution
{
public:
    long long maxTotalValue(vector<int> &nums, int k)
    {
        // Find the smallest element in the array
        long long mn = *min_element(nums.begin(), nums.end());

        // Find the largest element in the array
        long long mx = *max_element(nums.begin(), nums.end());

        // Best subarray value = global maximum - global minimum
        long long best = mx - mn;

        // We can choose the same best subarray k times
        return best * k;
    }
};