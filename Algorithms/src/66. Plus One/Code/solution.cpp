class Solution
{
public:
    vector<int> plusOne(vector<int> &digits)
    {
        // Start from the last digit
        for (int i = digits.size() - 1; i >= 0; i--)
        {
            digits[i]++; // Add 1 to current digit

            if (digits[i] < 10)
            { // No carry needed
                return digits;
            }

            digits[i] = 0; // Carry generated, set digit to 0
        }

        // If all digits were 9, we need one extra digit at the front
        digits.insert(digits.begin(), 1);
        return digits;
    }
};
