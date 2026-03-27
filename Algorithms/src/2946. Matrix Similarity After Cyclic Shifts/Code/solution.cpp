class Solution
{
public:
    bool areSimilar(vector<vector<int>> &mat, int k)
    {
        int m = mat.size();
        int n = mat[0].size();

        k %= n; // reduce unnecessary shifts

        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                int newCol;

                if (i % 2 == 0)
                {
                    // even row → left shift
                    newCol = (j + k) % n;
                }
                else
                {
                    // odd row → right shift
                    newCol = (j - k + n) % n;
                }

                if (mat[i][j] != mat[i][newCol])
                {
                    return false;
                }
            }
        }
        return true;
    }
};