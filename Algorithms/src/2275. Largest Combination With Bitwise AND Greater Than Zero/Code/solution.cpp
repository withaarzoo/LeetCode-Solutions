class Solution
{
public:
    int largestCombination(vector<int> &candidates)
    {
        int bitCount[31] = {0}; // Array to count '1's at each bit position

        // Count '1's in each bit position across all numbers
        for (int num : candidates)
        {
            for (int i = 0; i < 31; ++i)
            {
                if (num & (1 << i))
                {
                    bitCount[i]++;
                }
            }
        }

        // Find the maximum count in any bit position
        int maxCombinationSize = 0;
        for (int i = 0; i < 31; ++i)
        {
            maxCombinationSize = max(maxCombinationSize, bitCount[i]);
        }

        return maxCombinationSize;
    }
};
