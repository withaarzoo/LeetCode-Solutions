class Solution
{
public:
    int maxRotateFunction(vector<int> &nums)
    {
        int n = nums.size();

        long sum = 0; // Total sum of elements
        long F = 0;   // Initial rotation value F(0)

        // Step 1: Calculate total sum and F(0)
        for (int i = 0; i < n; i++)
        {
            sum += nums[i];         // accumulate total sum
            F += (long)i * nums[i]; // compute F(0)
        }

        long result = F; // store max result

        // Step 2: Use recurrence relation to compute next values
        for (int k = 1; k < n; k++)
        {
            // Transition from F(k-1) to F(k)
            F = F + sum - (long)n * nums[n - k];

            result = max(result, F); // update maximum
        }

        return (int)result;
    }
};