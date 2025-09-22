#include <vector>
#include <unordered_map>
#include <algorithm>
using namespace std;

class Solution {
public:
    int maxFrequencyElements(vector<int>& nums) {
        unordered_map<int,int> freq;
        // Count frequency of each number
        for (int x : nums) freq[x]++;

        // Find maximum frequency
        int maxFreq = 0;
        for (auto &p : freq) maxFreq = max(maxFreq, p.second);

        // Sum frequencies of elements that have frequency == maxFreq
        int result = 0;
        for (auto &p : freq) if (p.second == maxFreq) result += p.second;
        return result;
    }
};
