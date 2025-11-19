#include <vector>
#include <unordered_set>
using namespace std;

class Solution
{
public:
    int findFinalValue(vector<int> &nums, int original)
    {
        // Put all numbers in an unordered_set for O(1) average lookup
        unordered_set<int> s(nums.begin(), nums.end());
        // Keep doubling original while it exists in the set
        while (s.count(original))
        {
            original *= 2;
        }
        return original;
    }
};
