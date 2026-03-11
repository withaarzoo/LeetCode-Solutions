class Solution
{
public:
    int bitwiseComplement(int n)
    {

        // Edge case: complement of 0 is 1
        if (n == 0)
            return 1;

        int mask = 0;

        // Build a mask with all bits = 1
        // Example: if n = 5 (101), mask becomes 111
        while (mask < n)
        {
            mask = (mask << 1) | 1;
        }

        // XOR flips the bits
        return mask ^ n;
    }
};