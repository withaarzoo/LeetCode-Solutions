class Solution
{
public:
    int numSteps(string s)
    {
        int steps = 0;
        int carry = 0;

        // Traverse from right to left (ignore first bit for now)
        for (int i = s.size() - 1; i > 0; i--)
        {
            int bit = (s[i] - '0') + carry;

            if (bit == 1)
            {
                // Odd case: add 1 (carry becomes 1), then divide
                steps += 2;
                carry = 1;
            }
            else
            {
                // Even case: just divide by 2
                steps += 1;
                // carry remains same (if bit was 2, carry stays 1)
            }
        }

        // If carry remains, we need one extra step
        return steps + carry;
    }
};