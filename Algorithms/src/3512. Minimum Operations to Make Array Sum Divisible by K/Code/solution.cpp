class Solution
{
public:
    int minOperations(vector<int> &nums, int k)
    {
        long long sum = 0;

        // Calculate total sum of the array
        for (int x : nums)
        {
            sum += x;
        }

        // The minimum operations required is the remainder when sum is divided by k
        // Because each operation reduces the sum by exactly 1
        int remainder = sum % k;

        return remainder; // If remainder is 0, answer is 0. Otherwise, it's exactly 'remainder'.
    }
};
