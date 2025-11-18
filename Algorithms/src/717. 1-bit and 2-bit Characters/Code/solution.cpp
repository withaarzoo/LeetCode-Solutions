#include <vector>
using namespace std;

class Solution
{
public:
    bool isOneBitCharacter(vector<int> &bits)
    {
        int n = bits.size();
        int i = 0;
        // iterate until we reach or pass the second-to-last index
        // because we don't need to analyze beyond last bit
        while (i < n - 1)
        {
            if (bits[i] == 1)
            {
                // 1 starts a two-bit character -> skip two bits
                i += 2;
            }
            else
            {
                // 0 is a one-bit character -> skip one bit
                i += 1;
            }
        }
        // if we land exactly on last index, last char is one-bit
        return i == n - 1;
    }
};
