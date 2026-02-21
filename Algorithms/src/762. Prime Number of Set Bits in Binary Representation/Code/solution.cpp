class Solution
{
public:
    int countPrimeSetBits(int left, int right)
    {
        // Precompute prime counts possible up to 20
        unordered_set<int> primes = {2, 3, 5, 7, 11, 13, 17, 19};

        int ans = 0;

        for (int num = left; num <= right; num++)
        {
            // Count set bits using built-in function
            int setBits = __builtin_popcount(num);

            // If setBits is prime, increment answer
            if (primes.count(setBits))
            {
                ans++;
            }
        }

        return ans;
    }
};