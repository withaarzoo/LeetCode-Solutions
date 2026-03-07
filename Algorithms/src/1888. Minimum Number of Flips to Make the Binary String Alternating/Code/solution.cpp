class Solution
{
public:
    int minFlips(string s)
    {
        int n = s.size();
        string ss = s + s; // simulate all rotations

        int diff1 = 0, diff2 = 0;
        int ans = INT_MAX;

        for (int i = 0; i < ss.size(); i++)
        {
            char expected1 = (i % 2) ? '1' : '0'; // pattern 0101...
            char expected2 = (i % 2) ? '0' : '1'; // pattern 1010...

            if (ss[i] != expected1)
                diff1++;
            if (ss[i] != expected2)
                diff2++;

            // maintain window size
            if (i >= n)
            {
                char prev = ss[i - n];

                char prevExp1 = ((i - n) % 2) ? '1' : '0';
                char prevExp2 = ((i - n) % 2) ? '0' : '1';

                if (prev != prevExp1)
                    diff1--;
                if (prev != prevExp2)
                    diff2--;
            }

            if (i >= n - 1)
            {
                ans = min(ans, min(diff1, diff2));
            }
        }

        return ans;
    }
};