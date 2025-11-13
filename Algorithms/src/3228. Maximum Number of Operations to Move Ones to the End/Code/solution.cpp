class Solution
{
public:
    int maxOperations(string s)
    {
        long long ans = 0;  // result can be up to O(n^2) in value -> use 64-bit
        long long ones = 0; // count of '1's seen so far
        int n = s.size();
        for (int i = 0; i < n; ++i)
        {
            if (s[i] == '1')
            {
                ++ones; // saw another '1'
            }
            else
            { // s[i] == '0'
                // if this zero sits immediately after a '1', then all previous '1's
                // contribute one operation for this zero over the whole process.
                if (i > 0 && s[i - 1] == '1')
                    ans += ones;
            }
        }
        return (int)ans;
    }
};
