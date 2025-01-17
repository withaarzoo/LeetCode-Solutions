#include <vector>
using namespace std;

class Solution
{
public:
    bool doesValidArrayExist(vector<int> &derived)
    {
        int n = derived.size();
        // Case 1: Assume original[0] = 0
        bool validCase1 = true;
        int current = 0; // original[0]
        for (int i = 0; i < n; ++i)
        {
            current = derived[i] ^ current; // Compute original[i+1]
        }
        validCase1 = (current == 0); // Wrap-around condition

        // Case 2: Assume original[0] = 1
        bool validCase2 = true;
        current = 1; // original[0]
        for (int i = 0; i < n; ++i)
        {
            current = derived[i] ^ current; // Compute original[i+1]
        }
        validCase2 = (current == 1); // Wrap-around condition

        return validCase1 || validCase2;
    }
};
