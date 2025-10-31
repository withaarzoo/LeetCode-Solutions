#include <vector>
using namespace std;

class Solution {
public:
    vector<int> getSneakyNumbers(vector<int>& nums) {
        int n = (int)nums.size() - 2;          // original range size
        vector<char> seen(n, 0);              // boolean marks (char to save a bit)
        vector<int> res;
        res.reserve(2);
        for (int x : nums) {
            if (seen[x]) {
                res.push_back(x);             // found a repeated number
            } else {
                seen[x] = 1;                 // mark seen
            }
            if ((int)res.size() == 2) break; // early exit once we have both
        }
        return res;
    }
};
