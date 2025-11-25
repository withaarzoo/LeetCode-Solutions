class Solution
{
public:
    int smallestRepunitDivByK(int k)
    {
        // If k is divisible by 2 or 5, no number consisting only of 1's
        // can be divisible by k (because such numbers always end with 1).
        if (k % 2 == 0 || k % 5 == 0)
            return -1;

        int rem = 0; // current remainder
        // Try at most k times due to pigeonhole principle on remainders
        for (int length = 1; length <= k; ++length)
        {
            // Build next remainder: (previous_number * 10 + 1) % k
            rem = (rem * 10 + 1) % k;
            if (rem == 0)
            {
                return length; // found smallest length
            }
        }
        // If we didn't hit remainder 0 in k steps, it's impossible
        return -1;
    }
};
