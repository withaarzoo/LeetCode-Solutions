class Solution
{
public:
    int binaryGap(int n)
    {
        int lastPosition = -1;   // stores index of last seen 1
        int maxDistance = 0;     // stores maximum gap
        int currentPosition = 0; // current bit index

        while (n > 0)
        {
            // Check if current bit is 1
            if (n & 1)
            {
                // If we have seen a 1 before
                if (lastPosition != -1)
                {
                    maxDistance = max(maxDistance, currentPosition - lastPosition);
                }
                // Update last seen position
                lastPosition = currentPosition;
            }

            // Move to next bit
            n >>= 1;
            currentPosition++;
        }

        return maxDistance;
    }
};