class Solution
{
public:
    int smallestNumber(int n)
    {
        // We search for smallest k such that (1 << k) - 1 >= n
        long long k = 1;
        while (true)
        {
            long long val = (1LL << k) - 1; // 2^k - 1
            if (val >= n)
                return (int)val;
            k++;
        }
    }
};
