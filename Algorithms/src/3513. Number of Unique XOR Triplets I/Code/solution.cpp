class Solution
{
public:
    int uniqueXorTriplets(vector<int> &nums)
    {
        // Get the size of the permutation
        int n = nums.size();

        // Handle the small cases separately
        if (n <= 2)
            return n;

        // Find how many bits are needed to represent n
        int bits = 0;
        int x = n;
        while (x)
        {
            bits++;
            x >>= 1;
        }

        // Total values in the range [0, 2^bits - 1]
        return 1 << bits;
    }
};