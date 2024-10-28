#include <vector>
#include <unordered_set>
#include <algorithm>
using namespace std;

class Solution
{
public:
    int longestSquareStreak(vector<int> &nums)
    {
        sort(nums.begin(), nums.end());
        unordered_set<int> numSet(nums.begin(), nums.end());
        int maxLength = -1;

        for (int num : nums)
        {
            int length = 0;
            long long current = num;
            while (numSet.count(current))
            {
                length++;
                current *= current;
                if (current > 1e9)
                    break;
            }
            if (length >= 2)
            {
                maxLength = max(maxLength, length);
            }
        }
        return maxLength;
    }
};