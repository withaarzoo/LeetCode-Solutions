#include <vector>
using namespace std;

class Solution {
public:
    int findSmallestInteger(vector<int>& nums, int value) {
        // freq[r] = count of elements with remainder r (0 <= r < value)
        vector<int> freq(value, 0);
        for (int a : nums) {
            int r = ((a % value) + value) % value; // normalize negative values
            freq[r]++;
        }
        // Try to form 0,1,2,... greedily using residues
        int x = 0;
        while (true) {
            int r = x % value;
            if (freq[r] == 0) return x; // can't form x
            freq[r]--; // use one element with residue r to form x
            x++;
        }
        // unreachable
        return -1;
    }
};
