class Solution
{
public:
    long long gcdSum(vector<int> &nums)
    {
        int n = nums.size();

        // Store the gcd values for every prefix
        vector<int> prefixGcd(n);

        // Running maximum of the prefix
        int prefixMax = 0;

        // Build the prefixGcd array
        for (int i = 0; i < n; i++)
        {
            prefixMax = max(prefixMax, nums[i]);
            prefixGcd[i] = gcd(nums[i], prefixMax);
        }

        // Sort so that smallest and largest can be paired
        sort(prefixGcd.begin(), prefixGcd.end());

        long long ans = 0;

        // Pair smallest with largest
        int left = 0;
        int right = n - 1;

        while (left < right)
        {
            // Add gcd of the current pair
            ans += gcd(prefixGcd[left], prefixGcd[right]);

            left++;
            right--;
        }

        return ans;
    }
};