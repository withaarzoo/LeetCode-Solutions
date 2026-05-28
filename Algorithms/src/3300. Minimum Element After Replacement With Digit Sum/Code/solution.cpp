class Solution
{
public:
    // Helper function to calculate digit sum of a number
    int digitSum(int num)
    {
        int sum = 0;

        // Process every digit
        while (num > 0)
        {
            sum += num % 10; // Add last digit
            num /= 10;       // Remove last digit
        }

        return sum;
    }

    int minElement(vector<int> &nums)
    {
        int ans = INT_MAX;

        // Calculate digit sum for every element
        for (int num : nums)
        {
            ans = min(ans, digitSum(num));
        }

        return ans;
    }
};