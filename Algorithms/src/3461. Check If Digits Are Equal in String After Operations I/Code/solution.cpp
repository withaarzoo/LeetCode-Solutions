#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    bool hasSameDigits(string s)
    {
        // convert string to vector<int> digits
        vector<int> digits;
        digits.reserve(s.size());
        for (char c : s)
            digits.push_back(c - '0');

        // reduce until we have exactly two digits
        while (digits.size() > 2)
        {
            vector<int> next;
            next.reserve(digits.size() - 1);
            for (size_t i = 0; i + 1 < digits.size(); ++i)
            {
                next.push_back((digits[i] + digits[i + 1]) % 10);
            }
            digits.swap(next); // use swap to avoid copy
        }

        // check final two digits equality
        return digits.size() == 2 && digits[0] == digits[1];
    }
};
