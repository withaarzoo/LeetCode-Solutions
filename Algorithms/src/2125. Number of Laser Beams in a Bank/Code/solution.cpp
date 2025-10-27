class Solution
{
public:
    int numberOfBeams(vector<string> &bank)
    {
        long long ans = 0;  // total beams (use long long to be safe)
        long long prev = 0; // device count in the previous non-empty row

        for (auto &row : bank)
        {
            long long cnt = 0;
            // count '1's in current row
            for (char ch : row)
                if (ch == '1')
                    ++cnt;

            if (cnt > 0)
            {
                // beams between this row and previous non-empty row
                ans += prev * cnt;
                prev = cnt; // update previous to current
            }
            // if cnt == 0, skip but prev remains unchanged
        }
        return (int)ans;
    }
};
