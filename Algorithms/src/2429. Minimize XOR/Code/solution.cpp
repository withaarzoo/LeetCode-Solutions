class Solution
{
public:
    int minimizeXor(int num1, int num2)
    {
        int count2 = __builtin_popcount(num2); // Number of 1s in num2
        int count1 = __builtin_popcount(num1); // Number of 1s in num1

        if (count1 == count2)
        {
            return num1; // Already satisfies the condition
        }

        int result = num1;
        if (count1 > count2)
        {
            // Remove extra 1s
            for (int i = 0; i < 32 && count1 > count2; ++i)
            {
                if (result & (1 << i))
                {
                    result &= ~(1 << i); // Clear bit
                    count1--;
                }
            }
        }
        else
        {
            // Add additional 1s
            for (int i = 0; i < 32 && count1 < count2; ++i)
            {
                if (!(result & (1 << i)))
                {
                    result |= (1 << i); // Set bit
                    count1++;
                }
            }
        }
        return result;
    }
};
