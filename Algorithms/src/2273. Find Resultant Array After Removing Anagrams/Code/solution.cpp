#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<string> removeAnagrams(vector<string>& words) {
        vector<string> result;
        string prevSig = ""; // signature of last kept word

        for (auto &w : words) {
            string sig = w;
            sort(sig.begin(), sig.end()); // signature: sorted characters
            if (sig != prevSig) {
                result.push_back(w);      // keep this word
                prevSig = move(sig);      // update prevSig
            }
        }
        return result;
    }
};
