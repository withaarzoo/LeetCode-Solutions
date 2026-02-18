class Solution
{
public:
    bool hasAlternatingBits(int n)
    {
        // Get the last bit
        int prev = n & 1;
        n >>= 1;

        while (n > 0)
        {
            int curr = n & 1;

            // If two adjacent bits are same
            if (curr == prev)
            {
                return false;
            }

            prev = curr;
            n >>= 1;
        }

        return true;
    }
};
