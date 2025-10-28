#include <vector>
using namespace std;

class Solution
{
public:
    // simulate from start index with direction dir (-1 left, +1 right)
    bool simulate(const vector<int> &nums, int start, int dir)
    {
        int n = nums.size();
        vector<int> a = nums; // copy to mutate
        int curr = start;
        while (curr >= 0 && curr < n)
        {
            if (a[curr] == 0)
            {
                curr += dir; // move in same direction
            }
            else
            {
                a[curr]--;   // decrement
                dir = -dir;  // reverse direction
                curr += dir; // step in new direction
            }
        }
        // check if all zero
        for (int v : a)
            if (v != 0)
                return false;
        return true;
    }

    int countValidSelections(vector<int> &nums)
    {
        int n = nums.size();
        int ans = 0;
        for (int i = 0; i < n; ++i)
        {
            if (nums[i] != 0)
                continue; // start must be a zero
            if (simulate(nums, i, -1))
                ++ans; // left
            if (simulate(nums, i, +1))
                ++ans; // right
        }
        return ans;
    }
};
