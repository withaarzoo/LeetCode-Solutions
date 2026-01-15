class Solution
{
public:
    int getMaxGap(vector<int> &bars)
    {
        sort(bars.begin(), bars.end());

        int maxLen = 1;
        int curLen = 1;

        for (int i = 1; i < bars.size(); i++)
        {
            if (bars[i] == bars[i - 1] + 1)
            {
                curLen++;
            }
            else
            {
                curLen = 1;
            }
            maxLen = max(maxLen, curLen);
        }

        return maxLen;
    }

    int maximizeSquareHoleArea(int n, int m, vector<int> &hBars, vector<int> &vBars)
    {
        int hGap = getMaxGap(hBars) + 1;
        int vGap = getMaxGap(vBars) + 1;

        int side = min(hGap, vGap);
        return side * side;
    }
};
